package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
)

func SetMotherShipAmmo(ws *websocket.Conn, msg Message)  {
	user := usersInventoryWs[ws]
	inventory.SetAmmo(user, msg.AmmoID, msg.InventorySlot, msg.EquipSlot)

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
