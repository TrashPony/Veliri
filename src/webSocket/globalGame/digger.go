package globalGame

import (
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/coordinate"
	"github.com/gorilla/websocket"
)

func selectDigger(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	squadCoordinate, _ := mp.GetCoordinate(user.GetSquad().Q, user.GetSquad().R)

	if user != nil && squadCoordinate != nil {

		diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if diggerSlot == nil || diggerSlot.Equip == nil && diggerSlot.Equip.Applicable == "digger" {
			globalPipe <- Message{Event: "Error", Error: "no equip", idUserSend: user.GetID()}
			return
		}

		result := make([]*coordinate.Coordinate, 0)

		radius := coordinate.GetCoordinatesRadius(squadCoordinate, diggerSlot.Equip.Radius)
		for _, radiusCoordinate := range radius {
			mapCoordinate, find := mp.GetCoordinate(radiusCoordinate.Q, radiusCoordinate.R)
			if find && mapCoordinate.Move && !(radiusCoordinate.Q == squadCoordinate.Q && radiusCoordinate.R == squadCoordinate.R) {
				result = append(result, mapCoordinate)
			}
		}

		globalPipe <- Message{Event: "SelectDigger", Coordinates: result, idUserSend: user.GetID()}
	}

	// coordinate.GetCoordinatesRadius()
	// отдавать координаты которые можно копать
	// проверять что там не кто не стоит, делать координату мове=фелс, проверять наличие анмалии, отдавать лут аномалии
}
