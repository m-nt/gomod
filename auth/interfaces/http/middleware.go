package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/m-nt/gomod/auth/application"
)

func AuthMiddleware(svc *application.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		claims, err := svc.Validate(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
