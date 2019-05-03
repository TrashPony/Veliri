package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func SetEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "SetMotherShipEquip" {
		err = squad_inventory.SetEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitEquip" {
		err = squad_inventory.SetUnitEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot, msg.Source)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
