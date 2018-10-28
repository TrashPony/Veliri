package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveMSAmmo(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}

func RemoveUnitAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveUnitAmmo(user, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
}
