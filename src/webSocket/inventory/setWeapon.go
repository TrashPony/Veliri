package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipWeapon(ws *websocket.Conn, msg Message) {

	user := usersInventoryWs[ws]

	err := inventory.SetMSWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot)

	if err != nil {
		ws.WriteJSON(Response{Event: "ms error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}

func SetUnitWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)

	if err != nil {
		ws.WriteJSON(Response{Event: "unit error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}
