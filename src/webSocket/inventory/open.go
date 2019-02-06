package inventory

import (
	"github.com/TrashPony/Veliri/src/mechanics/squadInventory"
	"github.com/gorilla/websocket"
	"log"
)

func Open(ws *websocket.Conn, msg Message) {
	user := usersInventoryWs[ws]

	if user.GetSquad() == nil {
		squadInventory.GetInventory(user)
	}

	var err error

	if user.GetSquad() != nil {
		err = ws.WriteJSON(Response{Event: msg.Event, Squad: user.GetSquad(), BaseSquads: user.GetSquadsByBaseID(user.InBaseID),
			InventorySize: user.GetSquad().Inventory.GetSize(), InBase: user.InBaseID > 0})
	} else {
		err = ws.WriteJSON(Response{Event: msg.Event, Squad: user.GetSquad(), BaseSquads: user.GetSquadsByBaseID(user.InBaseID),
			InventorySize: 0, InBase: user.InBaseID > 0})
	}

	if err != nil {
		log.Fatal(err.Error())
	}
}
