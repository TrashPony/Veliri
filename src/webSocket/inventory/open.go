package inventory

import (
	"../../mechanics/gameObjects/squad"
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
	"log"
)

func Open(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		inventory.GetInventory(user)
	}

	err := ws.WriteJSON(RespSquad{Event: msg.Event, Squad: user.GetSquad()})
	if err != nil {
		log.Fatal(err.Error())
	}
}

type RespSquad struct {
	Event string       `json:"event"`
	Squad *squad.Squad `json:"squad"`
}
