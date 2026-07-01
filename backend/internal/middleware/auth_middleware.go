package middleware

import (
	"net/http"
	"strings"

	"github.com/customer-comm-dashboard/backend/internal/modules/auth"
	"github.com/customer-comm-dashboard/backend/internal/response"
	"github.com/gin-gonic/gin"
)

// AuthRequired returns a Gin middleware that validates JWT tokens.
func AuthRequired(authService *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
			response.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.APIResponse{
				Success: false,
				Error:   "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user ID in context for downstream handlers
		c.Set("userID", claims.Subject)

		c.Next()
	}
}
