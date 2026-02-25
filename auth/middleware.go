package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const ctxUserIDKey = "auth_user_id"

// RequireAuth blocks unauthenticated requests
func RequireAuth(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		uid, err := svc.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set(ctxUserIDKey, uid)
		c.Next()
	}
}

// GetUserID safely extracts the ID for other modules to use
func GetUserID(c *gin.Context) int {
	id, exists := c.Get(ctxUserIDKey)
	if !exists {
		return 0
	}
	return id.(int)
}
