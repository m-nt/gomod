package http

import "github.com/gin-gonic/gin"

func Register(r gin.IRouter, h *Handler) {
	g := r.Group("/auth")
	g.POST("/login", h.Login)
}
