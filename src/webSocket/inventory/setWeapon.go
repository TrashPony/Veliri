package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func SetMotherShipWeapon(ws *websocket.Conn, msg Message) {

	user := usersInventoryWs[ws]

	err := squadInventory.SetMSWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot)

	if err != nil {
		ws.WriteJSON(Response{Event: "ms error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}

func SetUnitWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	err := squadInventory.SetUnitWeapon(user, msg.WeaponID, msg.InventorySlot, msg.EquipSlot, msg.UnitSlot)

	if err != nil {
		ws.WriteJSON(Response{Event: "unit error", Error: err.Error()})
	}

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
