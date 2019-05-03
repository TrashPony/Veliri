package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squad_inventory"
	"github.com/gorilla/websocket"
)

func RemoveEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad("UpdateSquad", user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipEquip" {
		err = squad_inventory.RemoveEquip(user, msg.EquipSlot, msg.EquipSlotType, user.GetSquad().MatherShip, msg.Destination, true)
	}

	if msg.Event == "RemoveUnitEquip" {
		err = squad_inventory.RemoveUnitEquip(user, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot, msg.Destination)
	}

	UpdateSquad("UpdateSquad", user, err, ws, msg)
}
