package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func RemoveMotherShipEquip(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]

	inventory.RemoveEquip(user, msg.EquipSlot, msg.EquipSlotType)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
