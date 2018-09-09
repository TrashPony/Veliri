package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipWeapon(ws *websocket.Conn, msg Message) {

	user := usersInventoryWs[ws]

	err := inventory.SetMSWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot)

	if err != nil {
		ws.WriteJSON(Error{Event: "error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}

func SetUnitWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := inventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)

	if err != nil {
		ws.WriteJSON(Error{Event: "error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
