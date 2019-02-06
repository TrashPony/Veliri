package inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetWeapon(ws *websocket.Conn, msg Message) {

	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		UpdateSquad(user, errors.New("no select squad"), ws, msg)
		return
	}

	var err error

	if msg.Event == "SetMotherShipWeapon" {
		err = squadInventory.SetWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, user.GetSquad().MatherShip, msg.Source)
	}

	if msg.Event == "SetUnitWeapon" {
		err = squadInventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot, msg.Source)
	}

	UpdateSquad(user, err, ws, msg)
}
