package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func RemoveMotherShipAmmo(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	inventory.RemoveAmmo(user, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
