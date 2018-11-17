package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveMSWeapon(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}

func RemoveUnitWeapon(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveUnitWeapon(user, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}
