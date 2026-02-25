package http

import "github.com/gin-gonic/gin"

func Register(r gin.IRouter, h *Handler, m ...gin.HandlerFunc) {
	g := r.Group("/users")
	g.Use(m...)
	g.POST("", h.Create)
	g.GET("/:id", h.Get)
}
