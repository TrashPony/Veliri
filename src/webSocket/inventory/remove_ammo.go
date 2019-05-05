package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func RemoveAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipAmmo" {
		err = squad_inventory.RemoveAmmo(user, msg.EquipSlot, user.GetSquad().MatherShip, msg.Destination)
	}

	if msg.Event == "RemoveUnitAmmo" {
		err = squad_inventory.RemoveUnitAmmo(user, msg.EquipSlot, msg.UnitSlot, msg.Destination)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
