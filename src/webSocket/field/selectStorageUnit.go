package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/unit"
	"../../mechanics/coordinate"
)

func selectStorageUnit(msg Message, ws *websocket.Conn) {
	client, ok := usersFieldWs[ws]

	if !ok {
		delete(usersFieldWs, ws)
	} else {
		storageUnit, find := client.GetUnitStorage(msg.UnitID)

		if find {
			resp := SelectStorageUnit{Event: msg.Event, Unit: storageUnit, PlaceCoordinate: client.GetCreateZone()}
			ws.WriteJSON(resp)
		}
	}
}

type SelectStorageUnit struct {
	Event           string                   `json:"event"`
	Unit            *unit.Unit               `json:"unit"`
	PlaceCoordinate []*coordinate.Coordinate `json:"place_coordinate"`
}