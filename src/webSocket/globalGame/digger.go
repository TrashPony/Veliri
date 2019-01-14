package globalGame

import (
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func selectDigger(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	squadCoordinate := globalGame.GetQRfromXY(user.GetSquad().GlobalX, user.GetSquad().GlobalY, mp)

	if user != nil && squadCoordinate != nil {

		diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if diggerSlot == nil || diggerSlot.Equip == nil && diggerSlot.Equip.Applicable == "digger" {
			globalPipe <- Message{Event: "Error", Error: "no equip", idUserSend: user.GetID()}
			return
		}

		result := make([]*coordinate.Coordinate, 0)

		radius := coordinate.GetCoordinatesRadius(squadCoordinate, diggerSlot.Equip.Radius+1)
		deadZone := coordinate.GetCoordinatesRadius(squadCoordinate, 1)

		for _, radiusCoordinate := range radius {
			mapCoordinate, find := mp.GetCoordinate(radiusCoordinate.Q, radiusCoordinate.R)
			if find && mapCoordinate.Move && !(radiusCoordinate.Q == squadCoordinate.Q && radiusCoordinate.R == squadCoordinate.R) {
				result = append(result, mapCoordinate)
			}
		}

		// т.к. мс стоит на 7 клетках то копать он может только на следующих
		for i, resultCoordinate := range result {
			for _, deadCoordinate := range deadZone {
				if resultCoordinate.Q == deadCoordinate.Q && resultCoordinate.R == deadCoordinate.R {
					result[i] = nil
				}
			}
		}

		globalPipe <- Message{Event: "SelectDigger", Coordinates: result, idUserSend: user.GetID(), TypeSlot: msg.TypeSlot, Slot: msg.Slot}
	}
}

func useDigger(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)
	if user.GetSquad().MoveChecker {
		user.GetSquad().GetMove() <- true // останавливаем движение
	}

	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	squadCoordinate := globalGame.GetQRfromXY(user.GetSquad().GlobalX, user.GetSquad().GlobalY, mp)

	if user != nil && squadCoordinate != nil {
		diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if diggerSlot == nil || diggerSlot.Equip == nil && diggerSlot.Equip.Applicable == "digger" {
			globalPipe <- Message{Event: "Error", Error: "no equip", idUserSend: user.GetID()}
			return
		}

		result := make([]*coordinate.Coordinate, 0)

		radius := coordinate.GetCoordinatesRadius(squadCoordinate, diggerSlot.Equip.Radius+1)
		deadZone := coordinate.GetCoordinatesRadius(squadCoordinate, 1)

		for _, radiusCoordinate := range radius {
			mapCoordinate, find := mp.GetCoordinate(radiusCoordinate.Q, radiusCoordinate.R)
			if find && mapCoordinate.Move && !(radiusCoordinate.Q == squadCoordinate.Q && radiusCoordinate.R == squadCoordinate.R) {
				result = append(result, mapCoordinate)
			}
		}

		// т.к. мс стоит на 7 клетках то копать он может только на следующих
		for i, resultCoordinate := range result {
			for _, deadCoordinate := range deadZone {
				if resultCoordinate.Q == deadCoordinate.Q && resultCoordinate.R == deadCoordinate.R {
					result[i] = nil
				}
			}
		}

		for _, resultCoordinate := range result {
			if resultCoordinate != nil && msg.Q == resultCoordinate.Q && msg.R == resultCoordinate.R {
				diggerCoordinate, ok := mp.GetCoordinate(msg.Q, msg.R)
				if ok && diggerCoordinate.Move {
					// todo проверить что координата свободна
					anomaly := maps.Maps.GetMapAnomaly(mp.Id, msg.Q, msg.R)
					if anomaly != nil {
						// TODO хранить аномалии как обьекты на бекенде
						box, res, AnomalyText := anomaly.GetLoot()
						globalPipe <- Message{Event: msg.Event, OtherUser: GetShortUserInfo(user), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: box, Reservoir: res, AnomalyText: AnomalyText}
					} else {
						globalPipe <- Message{Event: msg.Event, OtherUser: GetShortUserInfo(user), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: nil, Reservoir: nil, AnomalyText: nil}
					}
				}
			}
		}
	}
}
