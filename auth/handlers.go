package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r gin.IRouter, svc *Service, verifier UserVerifier) {
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Ask the outside world (Users module) to verify
		uid, err := verifier.VerifyCredentials(c, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token, _ := svc.GenerateToken(uid)
		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
