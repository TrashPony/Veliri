package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "SetMotherShipEquip" {
		err = squadInventory.SetEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitEquip" {
		err = squadInventory.SetUnitEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot, msg.Source)
	}

	UpdateSquad(user, err, ws, msg)
}
