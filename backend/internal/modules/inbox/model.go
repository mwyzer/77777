package inbox

import "time"

// Customer represents a customer from WhatsApp or Telegram.
type Customer struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone,omitempty"`
	Email      string    `json:"email,omitempty"`
	Provider   string    `json:"provider"`
	ProviderID string    `json:"provider_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Conversation represents a conversation thread with a customer.
type Conversation struct {
	ID             string     `json:"id"`
	CustomerID     string     `json:"customer_id"`
	Customer       *Customer  `json:"customer,omitempty"`
	Channel        string     `json:"channel"`
	Status         string     `json:"status"`
	LastMessageAt  *time.Time `json:"last_message_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// Message represents a single message in a conversation.
type Message struct {
	ID                string    `json:"id"`
	ConversationID    string    `json:"conversation_id"`
	SenderType        string    `json:"sender_type"`
	Content           string    `json:"content"`
	Status            string    `json:"status"`
	ProviderMessageID string    `json:"provider_message_id,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}
