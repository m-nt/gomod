package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Register(r gin.IRouter, reg *Registry) {
	r.GET("/ws/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		reg.Set(id, conn)
		defer reg.Delete(id)

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				break
			}
		}
	})
}
