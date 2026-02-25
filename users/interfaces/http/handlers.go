package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/m-nt/gomod/users/application"
)

type Handler struct {
	svc *application.Service
}

func New(svc *application.Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.Create(c, req.Email, req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	u, err := h.svc.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
