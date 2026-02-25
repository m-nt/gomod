package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-nt/gomod/auth/application"
)

type Handler struct {
	svc *application.Service
}

func New(s *application.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Login(c *gin.Context) {

	var req struct {
		UserID int    `json:"user_id"`
		Email  string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Login(c, req.UserID, req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
