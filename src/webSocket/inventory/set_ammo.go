package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func SetAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "SetMotherShipAmmo" {
		err = squad_inventory.SetAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitAmmo" {
		err = squad_inventory.SetUnitAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot, msg.Source)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
