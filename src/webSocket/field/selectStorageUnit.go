package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/unit"
	"../../mechanics/coordinate"
	"../../mechanics"

)

func selectStorageUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]

	if !ok {
		delete(usersFieldWs, ws)
	} else {
		if client.GetCreateZone() == nil {
			client.SetCreateZone(mechanics.GetGameZone(client.GetMatherShip().X, client.GetMatherShip().Y, client.GetMatherShip().RangeView, Games[client.GetGameID()]))
		}

		storageUnit, find := client.GetUnitStorage(msg.UnitID)

		if find {
			resp := SelectStorageUnit{Event: msg.Event, Unit: storageUnit, PlaceCoordinate: client.GetCreateZone()}
			ws.WriteJSON(resp)
		}
	}
}

type SelectStorageUnit struct {
	Event           string                                       `json:"event"`
	Unit            *unit.Unit                                   `json:"unit"`
	PlaceCoordinate map[string]map[string]*coordinate.Coordinate `json:"place_coordinate"`
}
