package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "SetMotherShipAmmo" {
		err = squadInventory.SetAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitAmmo" {
		err = squadInventory.SetUnitAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot, msg.Source)
	}

	UpdateSquad(user, err, ws, msg)
}
