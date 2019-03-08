package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"sync"
)

type wsUsers struct {
	users map[*websocket.Conn]*player.Player
	mx    sync.RWMutex
}

var Clients = NewClientsStore()

func NewClientsStore() *wsUsers {
	return &wsUsers{
		users: make(map[*websocket.Conn]*player.Player),
	}
}

func (c *wsUsers) AddNewClient(ws *websocket.Conn, client *player.Player) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.users[ws] = client
}

func (c *wsUsers) GetByWs(ws *websocket.Conn) *player.Player {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.users[ws]
}

func (c *wsUsers) GetAll() map[*websocket.Conn]*player.Player {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.users
}
