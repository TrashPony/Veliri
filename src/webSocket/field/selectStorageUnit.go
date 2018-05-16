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

		/*coordinates := client.GetCreateZone()
		respawn :=	client.GetRespawn()
		activeGame := Games[client.GetGameID()]

		_, ok := coordinates[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
		if ok && !(msg.X == respawn.X && msg.Y == respawn.Y) {
			var unit game.Unit
			unit, price, createError := game.CreateUnit(msg.IdGame, strconv.Itoa(client.GetID()), msg.TypeUnit, msg.X, msg.Y)

			if createError == nil {
				activeGame.SetUnit(&unit)

				UpdateWatchZone(client, activeGame, nil)

				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), PlayerPrice: price, X: unit.X, Y: unit.Y}
				fieldPipe <- resp

				var unitsParameter InitUnit
				unitsParameter.initUnit("InitUnit", &unit, client.GetLogin())
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, Error: createError.Error()}
				fieldPipe <- resp
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, Error: "not allow"}
			fieldPipe <- resp
		}*/
	}
}

type SelectStorageUnit struct {
	Event           string                   `json:"event"`
	Unit            *unit.Unit               `json:"unit"`
	PlaceCoordinate []*coordinate.Coordinate `json:"place_coordinate"`
}
