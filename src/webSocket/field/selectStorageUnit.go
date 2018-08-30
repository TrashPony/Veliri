package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/localGame/Phases/placePhase"
)

func selectStorageUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	if client.GetReady() == false && findGame {
		if !ok {
			delete(usersFieldWs, ws)
		} else {

			storageUnit, find := client.GetUnitStorage(msg.UnitID)

			if find {
				resp := SelectStorageUnit{Event: msg.Event, Unit: storageUnit,
					PlaceCoordinate: placePhase.GetPlaceCoordinate(client.GetSquad().MatherShip.Q, client.GetSquad().MatherShip.R,
						client.GetSquad().MatherShip.Body.RangeView, activeGame)}
				ws.WriteJSON(resp)
			}
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type SelectStorageUnit struct {
	Event           string                                       `json:"event"`
	Unit            *unit.Unit                                   `json:"unit"`
	PlaceCoordinate map[string]map[string]*coordinate.Coordinate `json:"place_coordinate"`
}
