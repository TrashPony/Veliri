package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveBody(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipBody" {
		squadInventory.RemoveMSBody(user)
	}

	if msg.Event == "RemoveUnitBody" {
		squadInventory.RemoveUnitBody(user, msg.UnitSlot, true)
	}

	UpdateSquad(user, err, ws, msg)
}
