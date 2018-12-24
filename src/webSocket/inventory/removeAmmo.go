package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
)

func RemoveMotherShipAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveMSAmmo(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}

func RemoveUnitAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	squadInventory.RemoveUnitAmmo(user, msg.EquipSlot, msg.UnitSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad(), InventorySize: user.GetSquad().Inventory.GetSize()})
}
