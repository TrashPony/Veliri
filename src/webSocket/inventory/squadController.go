package inventory

import (
	"../../mechanics/db/removeInDB"
	"github.com/gorilla/websocket"
)

func SquadSettings(ws *websocket.Conn, msg Message)  {

	/*if msg.Event == "AddNewSquad" {
		err, squad := insert.AddNewSquad(msg.SquadName, usersInventoryWs[ws].GetID())

		var resp Response

		if err != nil {
			resp = Response{Event: "AddNewSquad", Error: err.Error()}
			ws.WriteJSON(resp)
		} else {
			usersInventoryWs[ws].Squads = append(usersInventoryWs[ws].Squads, squad)
			usersInventoryWs[ws].GetSquad() = squad
			resp = Response{Event: "AddNewSquad", Error: "none", Squad: squad}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "SelectSquad" {
		for _, squad := range  usersInventoryWs[ws].Squads {
			if squad.ID == msg.SquadID {
				usersInventoryWs[ws].GetSquad() = squad
				resp := Response{Event: "SelectSquad", Error: "none", Squad: squad}
				ws.WriteJSON(resp)
			}
		}
	}*/

	if msg.Event == "DeleteSquad" {
		if usersInventoryWs[ws].GetSquad() != nil {

			id := usersInventoryWs[ws].GetSquad().ID
			removeInDB.DeleteSquad(id)
			usersInventoryWs[ws].SetSquad(nil)

			resp := Response{Event: "RemoveSquad", SquadID: id}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}

	if msg.Event == "SelectMatherShip" {
		if usersInventoryWs[ws].GetSquad() != nil {
			if usersInventoryWs[ws].GetSquad().MatherShip != nil {
				//usersInventoryWs[ws].Squad.ReplaceMatherShip(msg.MatherShipID)
			} else {
				//usersInventoryWs[ws].Squad.AddMatherShip(msg.MatherShipID)
			}
			resp := Response{Event: "UpdateSquad", Squad: usersInventoryWs[ws].GetSquad()}
			ws.WriteJSON(resp)
		} else {
			resp := Response{Event: msg.Event, Error: "No select squad"}
			ws.WriteJSON(resp)
		}
	}
}