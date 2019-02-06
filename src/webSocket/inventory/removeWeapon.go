package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "RemoveMotherShipWeapon" {
		err = squadInventory.RemoveWeapon(user, msg.EquipSlot, user.GetSquad().MatherShip, msg.Destination, true)
	}

	if msg.Event == "RemoveUnitWeapon" {
		err = squadInventory.RemoveUnitWeapon(user, msg.EquipSlot, msg.UnitSlot, msg.Destination)
	}

	UpdateSquad(user, err, ws, msg)
}
