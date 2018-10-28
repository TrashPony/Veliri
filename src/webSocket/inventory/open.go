package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
	"log"
)

func Open(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		inventory.GetInventory(user)
	}

	err := ws.WriteJSON(Response{Event: msg.Event, Squad: user.GetSquad(), InventorySize: user.GetSquad().GetUseAllInventorySize()})
	if err != nil {
		log.Fatal(err.Error())
	}
}
