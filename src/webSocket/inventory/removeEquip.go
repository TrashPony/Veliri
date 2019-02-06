package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipEquip" {
		err = squadInventory.RemoveEquip(user, msg.EquipSlot, msg.EquipSlotType, user.GetSquad().MatherShip, msg.Destination, true)
	}

	if msg.Event == "RemoveUnitEquip" {
		err = squadInventory.RemoveUnitEquip(user, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot, msg.Destination)
	}

	UpdateSquad(user, err, ws, msg)
}
