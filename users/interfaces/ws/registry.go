package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Registry struct {
	m sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Set(id string, c *websocket.Conn) {
	r.m.Store(id, c)
}

func (r *Registry) Delete(id string) {
	r.m.Delete(id)
}

func (r *Registry) Send(id string, msg []byte) error {
	if v, ok := r.m.Load(id); ok {
		return v.(*websocket.Conn).WriteMessage(websocket.TextMessage, msg)
	}
	return nil
}
