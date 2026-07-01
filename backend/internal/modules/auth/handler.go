package auth

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

// Login handles POST /api/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Email and password are required")
		return
	}

	res, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		response.Unauthorized(c, "Invalid email or password")
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Login success",
		Data:    res,
	})
}

// Me handles GET /api/auth/me
func (h *Handler) Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "Not authenticated")
		return
	}

	user, err := h.service.GetMe(c.Request.Context(), userID.(string))
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "User retrieved",
		Data: gin.H{
			"user": user,
		},
	})
}
