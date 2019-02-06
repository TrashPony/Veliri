package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipAmmo" {
		err = squadInventory.RemoveAmmo(user, msg.EquipSlot, user.GetSquad().MatherShip, msg.Destination, true)
	}

	if msg.Event == "RemoveUnitAmmo" {
		err = squadInventory.RemoveUnitAmmo(user, msg.EquipSlot, msg.UnitSlot, msg.Destination)
	}

	UpdateSquad(user, err, ws, msg)
}
