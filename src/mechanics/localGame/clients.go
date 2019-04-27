package localGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/gorilla/websocket"
	"sync"
)

type wsUsers struct {
	users map[*websocket.Conn]*player.Player // карта для проверок игроков (колизии, хедлеры и тд)
	mx    sync.RWMutex
}

var Clients = NewClientsStore()

func NewClientsStore() *wsUsers {
	return &wsUsers{
		users: make(map[*websocket.Conn]*player.Player),
	}
}

func (c *wsUsers) AddNewClient(newWS *websocket.Conn, newClient *player.Player) {
	c.mx.Lock()
	defer c.mx.Unlock()

	for ws, client := range c.users {
		if !client.Bot && client.GetLogin() == newClient.GetLogin() {
			delete(c.users, ws) // удаляем его из активных подключений
			ws.Close()
		}
	}
	c.users[newWS] = newClient
}

func (c *wsUsers) GetByWs(ws *websocket.Conn) *player.Player {
	user := c.users[ws]
	return user
}

func (c *wsUsers) GetAllConnects() (map[*websocket.Conn]*player.Player, *sync.RWMutex) {
	c.mx.Lock()
	return c.users, &c.mx
}

func (c *wsUsers) DelClientByWS(ws *websocket.Conn) {
	c.mx.Lock()
	if c.users[ws] != nil && !c.users[ws].Bot {
		delete(c.users, ws)
		if ws != nil {
			ws.Close()
		}
	}
	c.mx.Unlock()
}
