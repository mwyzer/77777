package webhook

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/customer-comm-dashboard/backend/internal/providers/telegram"
	"github.com/customer-comm-dashboard/backend/internal/response"
	"github.com/gin-gonic/gin"
)

// Handler handles HTTP webhook requests.
type Handler struct {
	service *Service
}

// NewHandler creates a new webhook handler.
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// TelegramWebhook handles POST /api/webhooks/telegram
func (h *Handler) TelegramWebhook(c *gin.Context) {
	// Read raw body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("ERROR: Failed to read webhook body: %v", err)
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Error:   "Failed to read request body",
		})
		return
	}

	// Parse Telegram update
	var update telegram.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Printf("ERROR: Failed to parse Telegram update: %v", err)
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Error:   "Invalid payload format",
		})
		return
	}

	// Process update (async-safe — Telegram expects 200 quickly)
	if err := h.service.ProcessTelegramUpdate(c.Request.Context(), &update); err != nil {
		log.Printf("ERROR: Failed to process Telegram update: %v", err)
		// Always return 200 to Telegram to prevent retry storms
		c.JSON(http.StatusOK, response.APIResponse{
			Success: true,
			Message: "Webhook received",
		})
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Webhook processed",
	})
}
