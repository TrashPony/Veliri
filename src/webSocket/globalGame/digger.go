package globalGame

import (
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/maps"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/dynamicMapObject"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
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
					// todo проверить что координата свободна от игрока
					anomaly := maps.Maps.GetMapAnomaly(mp.Id, msg.Q, msg.R)
					if anomaly != nil {
						maps.Maps.RemoveMapAnomaly(mp.Id, msg.Q, msg.R)

						box, res, AnomalyText := anomaly.GetLoot()
						mpCoordinate, _ := mp.GetCoordinate(msg.Q, msg.R)

						if box != nil {

							box.MapID = mp.Id
							box.Q = anomaly.GetQ()
							box.R = anomaly.GetR()
							box.Rotate = rand.Intn(360)

							boxes.Boxes.InsertNewBox(box)
						}

						if res != nil {
							if mpCoordinate != nil {

								res.Q = anomaly.GetQ()
								res.R = anomaly.GetR()
								res.Rotate = rand.Intn(360)
								res.MapID = mp.Id
								mpCoordinate.Move = false

								maps.AddResourceInMap(mp, res)
							}
						}

						var dynamicObject dynamicMapObject.DynamicObject

						if AnomalyText != nil {
							if mpCoordinate != nil {

								dynamicObject = dynamicMapObject.DynamicObject{
									TextureObject: "infoAnomaly",
									Dialog:        AnomalyText,
									Destroyed:     true,
									DestroyTime:   time.Now(),
									// TODO запуст горутины на уничтожение
								}

								mpCoordinate.DynamicObject = &dynamicObject
							}
						}

						globalPipe <- Message{Event: msg.Event, OtherUser: GetShortUserInfo(user), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: box, Reservoir: res,
							DynamicObject: &dynamicObject, Name: diggerSlot.Equip.Name}

						usersGlobalWs, mx := Clients.GetAll()
						mx.Unlock()

						for _, otherUser := range usersGlobalWs {
							equipSlot := otherUser.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
							anomalies, err := globalGame.GetVisibleAnomaly(otherUser, equipSlot)
							if err == nil {
								globalPipe <- Message{Event: "AnomalySignal", idUserSend: otherUser.GetID(), Anomalies: anomalies}
							} else {
								globalPipe <- Message{Event: "RemoveAnomalies", idUserSend: otherUser.GetID()}
							}
						}
					} else {
						globalPipe <- Message{Event: msg.Event, OtherUser: GetShortUserInfo(user), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: nil, Reservoir: nil, DynamicObject: nil,
							Name: diggerSlot.Equip.Name}
					}
				}
			}
		}
	}
}
