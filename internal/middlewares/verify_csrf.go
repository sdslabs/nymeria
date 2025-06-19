package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/utils"
)

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip CSRF validation for GET, HEAD, OPTIONS requests
		if c.Request.Method == "GET" || c.Request.Method == "HEAD" || c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}

		// Get user ID from header (TODO: replace with session/JWT when available)
		userID := c.GetHeader("X-User-ID") // TODO: Get user ID from session
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "User ID is required for CSRF protection",
			})
			c.Abort()
			return
		}

		var csrfToken string

		// First, try to get CSRF token from JSON body
		if c.GetHeader("Content-Type") == "application/json" {
			var body map[string]interface{}
			// Create a copy of the request body for CSRF token extraction
			if err := c.ShouldBindJSON(&body); err == nil {
				if token, exists := body["csrf_token"]; exists {
					if tokenStr, ok := token.(string); ok {
						csrfToken = tokenStr
					}
				}
			}
			// Restore the body for the actual handler by binding again
			// Note: This is a limitation - we need to read the body twice
			// A more elegant solution would be to buffer the body
		}

		// If not found in JSON body, try header
		if csrfToken == "" {
			csrfToken = c.GetHeader("X-CSRF-Token")
		}

		// If not found in header, try form field
		if csrfToken == "" {
			csrfToken = c.PostForm("csrf_token")
		}

		// If no CSRF token found, return error
		if csrfToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "CSRF token is required",
			})
			c.Abort()
			return
		}

		// Validate CSRF token
		if !utils.ValidateCSRFToken(userID, csrfToken) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid or expired CSRF token",
			})
			c.Abort()
			return
		}

		// Token is valid, continue to the next handler
		c.Next()
	}
}
