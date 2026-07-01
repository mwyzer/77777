package inbox

import (
	"net/http"

	"github.com/customer-comm-dashboard/backend/internal/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// ── Conversations ──

// ListConversations handles GET /api/inbox/conversations
func (h *Handler) ListConversations(c *gin.Context) {
	var params ConversationFilterParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.BadRequest(c, "Invalid query parameters")
		return
	}

	result, err := h.service.ListConversations(c.Request.Context(), params)
	if err != nil {
		response.InternalError(c, "Failed to list conversations")
		return
	}

	response.OK(c, result)
}

// GetConversation handles GET /api/inbox/conversations/:id
func (h *Handler) GetConversation(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Conversation ID is required")
		return
	}

	result, err := h.service.GetConversation(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "Conversation not found")
		return
	}

	response.OK(c, result)
}

// ── Messages ──

// ListMessages handles GET /api/inbox/conversations/:id/messages
func (h *Handler) ListMessages(c *gin.Context) {
	conversationID := c.Param("id")
	if conversationID == "" {
		response.BadRequest(c, "Conversation ID is required")
		return
	}

	var params PaginationParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.BadRequest(c, "Invalid query parameters")
		return
	}

	result, err := h.service.GetMessages(c.Request.Context(), conversationID, params.Page, params.Limit)
	if err != nil {
		response.InternalError(c, "Failed to list messages")
		return
	}

	response.OK(c, result)
}

// SendMessage handles POST /api/inbox/conversations/:id/messages
func (h *Handler) SendMessage(c *gin.Context) {
	conversationID := c.Param("id")
	if conversationID == "" {
		response.BadRequest(c, "Conversation ID is required")
		return
	}

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Content is required")
		return
	}

	msg, err := h.service.SendMessage(c.Request.Context(), conversationID, req.Content)
	if err != nil {
		c.JSON(http.StatusNotFound, response.APIResponse{
			Success: false,
			Error:   "Conversation not found",
		})
		return
	}

	c.JSON(http.StatusCreated, response.APIResponse{
		Success: true,
		Message: "Message sent",
		Data:    msg,
	})
}

// ── Customers ──

// ListCustomers handles GET /api/inbox/customers
func (h *Handler) ListCustomers(c *gin.Context) {
	var params CustomerFilterParams
	if err := c.ShouldBindQuery(&params); err != nil {
		response.BadRequest(c, "Invalid query parameters")
		return
	}

	result, err := h.service.ListCustomers(c.Request.Context(), params)
	if err != nil {
		response.InternalError(c, "Failed to list customers")
		return
	}

	response.OK(c, result)
}

// GetCustomer handles GET /api/inbox/customers/:id
func (h *Handler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Customer ID is required")
		return
	}

	result, err := h.service.GetCustomer(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "Customer not found")
		return
	}

	response.OK(c, result)
}
