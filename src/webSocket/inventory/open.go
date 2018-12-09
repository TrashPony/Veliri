package inventory

import (
	"../../mechanics/squadInventory"
	"github.com/gorilla/websocket"
	"log"
)

func Open(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		squadInventory.GetInventory(user)
	}

	err := ws.WriteJSON(Response{Event: msg.Event, Squad: user.GetSquad(),
		InventorySize: user.GetSquad().GetUseAllInventorySize(), InBase: user.InBaseID > 0})
	if err != nil {
		log.Fatal(err.Error())
	}
}
