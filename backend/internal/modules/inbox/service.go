package inbox

import (
	"context"
	"fmt"
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ── Conversation ──

// GetConversation returns a conversation by ID with its customer.
func (s *Service) GetConversation(ctx context.Context, id string) (*ConversationDetailResponse, error) {
	conv, err := s.repo.FindConversationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if conv == nil {
		return nil, fmt.Errorf("conversation not found")
	}

	customer, err := s.repo.FindCustomerByID(ctx, conv.CustomerID)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}

	return &ConversationDetailResponse{
		Conversation: *conv,
		Customer:     *customer,
	}, nil
}

// ListConversations returns a paginated list of conversations with customer info.
func (s *Service) ListConversations(ctx context.Context, params ConversationFilterParams) (*ConversationListResponse, error) {
	convs, total, err := s.repo.ListConversations(ctx, params)
	if err != nil {
		return nil, err
	}

	page, limit := defaultPagination(params.Page, params.Limit)

	return &ConversationListResponse{
		Conversations: convs,
		Total:         total,
		Page:          page,
		Limit:         limit,
	}, nil
}

// ── Message ──

// GetMessages returns paginated messages for a conversation.
func (s *Service) GetMessages(ctx context.Context, conversationID string, page, limit int) (*MessageListResponse, error) {
	messages, total, err := s.repo.ListMessages(ctx, conversationID, page, limit)
	if err != nil {
		return nil, err
	}

	page, limit = defaultPagination(page, limit)

	return &MessageListResponse{
		Messages: messages,
		Total:    total,
		Page:     page,
		Limit:    limit,
	}, nil
}

// SendMessage sends a reply from agent to customer.
func (s *Service) SendMessage(ctx context.Context, conversationID, content string) (*Message, error) {
	// Verify conversation exists
	conv, err := s.repo.FindConversationByID(ctx, conversationID)
	if err != nil {
		return nil, err
	}
	if conv == nil {
		return nil, fmt.Errorf("conversation not found")
	}

	msg := &Message{
		ConversationID: conversationID,
		SenderType:     "agent",
		Content:        content,
		Status:         "sent",
	}

	if err := s.repo.CreateMessage(ctx, msg); err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	// Update conversation last_message_at
	now := time.Now()
	if err := s.repo.UpdateConversationLastMessage(ctx, conversationID, now); err != nil {
		return nil, fmt.Errorf("update last message: %w", err)
	}

	return msg, nil
}

// ── Customer ──

// GetCustomer returns a customer by ID with their conversations.
func (s *Service) GetCustomer(ctx context.Context, id string) (*CustomerDetailResponse, error) {
	customer, err := s.repo.FindCustomerByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, fmt.Errorf("customer not found")
	}

	convs, err := s.repo.ListConversationsByCustomerID(ctx, id)
	if err != nil {
		return nil, err
	}
	if convs == nil {
		convs = []Conversation{}
	}

	return &CustomerDetailResponse{
		Customer:      *customer,
		Conversations: convs,
	}, nil
}

// ListCustomers returns a paginated, filtered list of customers.
func (s *Service) ListCustomers(ctx context.Context, params CustomerFilterParams) (*CustomerListResponse, error) {
	customers, total, err := s.repo.ListCustomers(ctx, params)
	if err != nil {
		return nil, err
	}
	if customers == nil {
		customers = []Customer{}
	}

	page, limit := defaultPagination(params.Page, params.Limit)

	return &CustomerListResponse{
		Customers: customers,
		Total:     total,
		Page:      page,
		Limit:     limit,
	}, nil
}
