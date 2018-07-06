package inventory

import (
	"github.com/gorilla/websocket"
	"../../mechanics/inventory"
	"../../mechanics/db/updateSquad"
)

func RemoveMotherShipBody(ws *websocket.Conn) {
	user := usersInventoryWs[ws]

	if user.GetSquad().MatherShip.Body != nil {
		inventory.BodyRemove(user.GetSquad().Inventory, user.GetSquad().MatherShip.Body)
		user.GetSquad().MatherShip.Body = nil
	}

	updateSquad.Squad(user.GetSquad())

	ws.WriteJSON(Response{Event: "UpdateSquad", Squad: user.GetSquad()})
}
