package inventory

import (
	"../../mechanics/db/get"
	"github.com/gorilla/websocket"
)

func GetDetailSquad(ws *websocket.Conn, msg Message)  {
	if msg.Event == "GetDetailOfUnits" {

		//weapons := detailUnit.GetWeapons()
		//bodies := detailUnit.GetBodies()

		//var resp = Response{Event: msg.Event, Weapons: weapons, Bodies: bodies}
		//ws.WriteJSON(resp)
	}

	if msg.Event == "GetEquipping" {
		var equipping = get.TypeEquipping()
		var resp = Response{Event: msg.Event, Equipping: equipping}
		ws.WriteJSON(resp)
	}

	if msg.Event == "GetListSquad" {
		squads, err := get.UserSquads(usersInventoryWs[ws].GetID())

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
		var matherShips = get.MatherShips()
		var resp = Response{Event: msg.Event, MatherShips: matherShips}
		ws.WriteJSON(resp)
	}
}
