package inventory

import (
	"../../mechanics/inventory"
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit/detailUnit"
)

func GetDetailSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "GetDetailOfUnits" {

		weapons := detailUnit.GetWeapons()
		bodies := detailUnit.GetBodies()

		var resp = Response{Event: msg.Event, Weapons: weapons, Bodies: bodies}
		ws.WriteJSON(resp)
	}

	if msg.Event == "GetEquipping" {
		var equipping = inventory.GetTypeEquipping()
		var resp = Response{Event: msg.Event, Equipping: equipping}
		ws.WriteJSON(resp)
	}

	if msg.Event == "GetListSquad" {
		squads, err := inventory.GetUserSquads(usersInventoryWs[ws].GetID())

		var resp Response

		if err != nil {
			resp = Response{Event: "GetListSquad", Error: err.Error()}
			ws.WriteJSON(resp)
		} else {
			usersInventoryWs[ws].Squads = squads
			resp = Response{Event: "GetListSquad", Error: "none", Squads: squads}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "GetMatherShips" {
		var matherShips = inventory.GetTypeMatherShips()
		var resp = Response{Event: msg.Event, MatherShips: matherShips}
		ws.WriteJSON(resp)
	}
}
