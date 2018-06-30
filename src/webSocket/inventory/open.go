package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func Open(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.Open(user)
}
