package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func RemoveBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipBody" {
		squad_inventory.RemoveMSBody(user)
	}

	if msg.Event == "RemoveUnitBody" {
		squad_inventory.RemoveUnitBody(user, msg.UnitSlot)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
