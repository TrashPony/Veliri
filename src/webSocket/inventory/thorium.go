package inventory

import (
	"../../mechanics/squadInventory"
	"errors"
	"github.com/gorilla/websocket"
)

func setThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]
	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}
	err := squadInventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot, msg.Source)
	UpdateSquad(user, err, ws, msg)
}

func removeThoriumThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]
	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}
	err := squadInventory.RemoveThorium(user, msg.ThoriumSlot, true)
	UpdateSquad(user, err, ws, msg)
}
