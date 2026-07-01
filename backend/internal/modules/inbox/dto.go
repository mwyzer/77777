package inbox

// ── Request DTOs ──

// SendMessageRequest is the payload for POST /api/inbox/conversations/:id/messages
type SendMessageRequest struct {
	Content string `json:"content" binding:"required,min=1"`
}

// ── Response DTOs ──

// ConversationListResponse is the response for GET /api/inbox/conversations
type ConversationListResponse struct {
	Conversations []Conversation `json:"conversations"`
	Total         int            `json:"total"`
	Page          int            `json:"page"`
	Limit         int            `json:"limit"`
}

// ConversationDetailResponse includes the customer info
type ConversationDetailResponse struct {
	Conversation Conversation `json:"conversation"`
	Customer     Customer     `json:"customer"`
}

// MessageListResponse is the response for GET /api/inbox/conversations/:id/messages
type MessageListResponse struct {
	Messages []Message `json:"messages"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
}

// CustomerListResponse is the response for GET /api/inbox/customers
type CustomerListResponse struct {
	Customers []Customer `json:"customers"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
}

// CustomerDetailResponse includes the customer and their conversations
type CustomerDetailResponse struct {
	Customer      Customer       `json:"customer"`
	Conversations []Conversation `json:"conversations"`
}

// ── Query / Filter Params ──

// PaginationParams holds common pagination query parameters.
type PaginationParams struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

// ConversationFilterParams holds filter params for listing conversations.
type ConversationFilterParams struct {
	Status  string `form:"status"`
	Channel string `form:"channel"`
	Search  string `form:"search"`
	Page    int    `form:"page"`
	Limit   int    `form:"limit"`
}

// CustomerFilterParams holds filter params for listing customers.
type CustomerFilterParams struct {
	Search   string `form:"search"`
	Provider string `form:"provider"`
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
}

// ── Helpers ──

// defaultPagination applies default values for page and limit.
func defaultPagination(page, limit int) (int, int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	return page, limit
}

// offset calculates SQL offset from page and limit.
func offset(page, limit int) int {
	return (page - 1) * limit
}
