package inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func Open(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		squad_inventory.GetInventory(user)
	}
	UpdateSquad("openInventory", user, nil, ws, msg)
}
