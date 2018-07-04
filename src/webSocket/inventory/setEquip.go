package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func SetMotherShipEquip(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.SetEquip(user, msg.EquipID, msg.InventorySlot, msg.EquipSlot, msg.EquipSlotType)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
