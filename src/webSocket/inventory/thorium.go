package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func setThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]
	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}
	err := squad_inventory.SetThorium(user, msg.InventorySlot, msg.ThoriumSlot, msg.Source)
	UpdateSquad("UpdateSquad", user, err, ws, msg)
}

func removeThoriumThorium(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]
	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}
	err := squad_inventory.RemoveThorium(user, msg.ThoriumSlot, true)
	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
