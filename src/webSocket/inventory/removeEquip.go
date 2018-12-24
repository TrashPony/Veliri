package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveMSEquip(user, msg.EquipSlot, msg.EquipSlotType)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}

func RemoveUnitEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveUnitEquip(user, msg.EquipSlot, msg.EquipSlotType, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
