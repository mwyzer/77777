package webhook

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/customer-comm-dashboard/backend/internal/database"
	"github.com/customer-comm-dashboard/backend/internal/modules/inbox"
	"github.com/customer-comm-dashboard/backend/internal/providers/telegram"
)

// Service handles webhook processing logic.
type Service struct {
	inboxRepo  *inbox.Repository
	tgClient   *telegram.Client
}

// NewService creates a new webhook service.
func NewService(inboxRepo *inbox.Repository, tgClient *telegram.Client) *Service {
	return &Service{
		inboxRepo: inboxRepo,
		tgClient:  tgClient,
	}
}

// ProcessTelegramUpdate handles an incoming Telegram webhook update.
func (s *Service) ProcessTelegramUpdate(ctx context.Context, update *telegram.Update) error {
	// ── 1. Idempotency check ──
	if database.RedisClient != nil {
		key := fmt.Sprintf("idempotency:telegram:%d", update.UpdateID)
		ok, err := database.RedisClient.SetNX(ctx, key, "1", 24*time.Hour).Result()
		if err != nil {
			log.Printf("WARNING: Redis idempotency check failed: %v", err)
		} else if !ok {
			log.Printf("Duplicate Telegram update_id=%d — skipped", update.UpdateID)
			return nil
		}
	}

	if update.Message == nil {
		return nil // non-message updates (e.g., status) — ignore silently
	}

	msg := update.Message
	chat := msg.Chat

	// ── 2. Find or create customer ──
	providerID := strconv.FormatInt(chat.ID, 10)
	customer, err := s.inboxRepo.FindCustomerByProvider(ctx, "telegram", providerID)
	if err != nil {
		return fmt.Errorf("find customer: %w", err)
	}

	if customer == nil {
		name := chat.FirstName
		if chat.LastName != "" {
			name = chat.FirstName + " " + chat.LastName
		}
		customer = &inbox.Customer{
			Name:       name,
			Provider:   "telegram",
			ProviderID: providerID,
		}
		if chat.Username != "" {
			customer.Email = chat.Username + "@telegram.local"
		}
		if err := s.inboxRepo.CreateCustomer(ctx, customer); err != nil {
			return fmt.Errorf("create customer: %w", err)
		}
		log.Printf("Created customer: %s (id=%s)", customer.Name, customer.ID)
	}

	// ── 3. Find or create conversation ──
	conv, err := s.inboxRepo.FindConversationByCustomerAndChannel(ctx, customer.ID, "telegram")
	if err != nil {
		return fmt.Errorf("find conversation: %w", err)
	}

	if conv == nil {
		conv = &inbox.Conversation{
			CustomerID: customer.ID,
			Channel:    "telegram",
			Status:     "open",
		}
		if err := s.inboxRepo.CreateConversation(ctx, conv); err != nil {
			return fmt.Errorf("create conversation: %w", err)
		}
		log.Printf("Created conversation: %s", conv.ID)
	}

	// ── 4. Save message ──
	message := &inbox.Message{
		ConversationID:    conv.ID,
		SenderType:        "customer",
		Content:           msg.Text,
		Status:            "delivered",
		ProviderMessageID: strconv.FormatInt(msg.MessageID, 10),
	}

	if err := s.inboxRepo.CreateMessage(ctx, message); err != nil {
		return fmt.Errorf("create message: %w", err)
	}
	log.Printf("Saved message: %s from %s", message.ID, customer.Name)

	// ── 5. Update conversation last_message_at ──
	if err := s.inboxRepo.UpdateConversationLastMessage(ctx, conv.ID, message.CreatedAt); err != nil {
		log.Printf("WARNING: update last_message_at: %v", err)
	}

	// ── 6. Publish realtime event (Phase 11 placeholder) ──
	if database.RedisClient != nil {
		database.RedisClient.Publish(ctx, "pubsub:inbox:new-message", message.ID)
	}

	return nil
}

// SendReply sends a reply message back to the customer via Telegram.
func (s *Service) SendReply(ctx context.Context, conversationID, customerProviderID, text string) (int64, error) {
	chatID, err := strconv.ParseInt(customerProviderID, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid provider ID: %w", err)
	}

	return s.tgClient.SendMessage(chatID, text)
}
