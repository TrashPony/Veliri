package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamicMapObject"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"math/rand"
	"time"
)

func selectDigger(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	squadCoordinate := globalGame.GetQRfromXY(user.GetSquad().GlobalX, user.GetSquad().GlobalY, mp)

	if user != nil && squadCoordinate != nil {

		diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if diggerSlot == nil || diggerSlot.Equip == nil && diggerSlot.Equip.Applicable == "digger" {
			go SendMessage(Message{Event: "Error", Error: "no equip", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
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

		go SendMessage(Message{Event: "SelectDigger", Coordinates: result, IDUserSend: user.GetID(),
			TypeSlot: msg.TypeSlot, Slot: msg.Slot, IDMap: user.GetSquad().MapID})
	}
}

func useDigger(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	squadCoordinate := globalGame.GetQRfromXY(user.GetSquad().GlobalX, user.GetSquad().GlobalY, mp)

	if user != nil && squadCoordinate != nil {

		stopMove(user, true)

		diggerSlot := user.GetSquad().MatherShip.Body.GetEquip(msg.TypeSlot, msg.Slot)
		if diggerSlot == nil || diggerSlot.Equip == nil && diggerSlot.Equip.Applicable == "digger" {
			go SendMessage(Message{Event: "Error", Error: "no equip", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
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

		users, rLock := globalGame.Clients.GetAll()
		defer rLock.Unlock()

		for _, resultCoordinate := range result {
			if resultCoordinate != nil && msg.Q == resultCoordinate.Q && msg.R == resultCoordinate.R {
				diggerCoordinate, ok := mp.GetCoordinate(msg.Q, msg.R)
				if ok && diggerCoordinate.Move {

					mpCoordinate, _ := mp.GetCoordinate(msg.Q, msg.R)
					var dynamicObject dynamicMapObject.DynamicObject
					dynamicObject.TextureBackground = "crater_2"
					dynamicObject.BackgroundScale = 75
					dynamicObject.BackgroundRotate = rand.Intn(360)

					// todo проверить что координата свободна от игрока
					anomaly := maps.Maps.GetMapAnomaly(mp.Id, msg.Q, msg.R)
					if anomaly != nil {
						maps.Maps.RemoveMapAnomaly(mp.Id, msg.Q, msg.R)

						box, res, AnomalyText := anomaly.GetLoot()
						// TODO запуст горутины на уничтожение

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

								mp.AddResourceInMap(res)
							}
						}

						if AnomalyText != nil {
							if mpCoordinate != nil {
								dynamicObject.TextureObject = "infoAnomaly"
								dynamicObject.Dialog = AnomalyText
								dynamicObject.Destroyed = true
								dynamicObject.DestroyTime = time.Now()
								dynamicObject.ObjectScale = 20
								dynamicObject.ObjectRotate = rand.Intn(360)
								dynamicObject.Shadow = 50
								dynamicObject.Move = true
								dynamicObject.View = true
								dynamicObject.Attack = true
							}
						}

						mpCoordinate.DynamicObject = &dynamicObject

						go SendMessage(Message{Event: msg.Event, OtherUser: user.GetShortUserInfo(true, false), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: box, Reservoir: res,
							DynamicObject: &dynamicObject, Name: diggerSlot.Equip.Name, IDMap: user.GetSquad().MapID})

						for _, otherUser := range users {
							equipSlot := otherUser.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
							anomalies, err := globalGame.GetVisibleAnomaly(otherUser, equipSlot)
							if err == nil {
								go SendMessage(Message{Event: "AnomalySignal", IDUserSend: otherUser.GetID(),
									Anomalies: anomalies, IDMap: user.GetSquad().MapID})
							} else {
								go SendMessage(Message{Event: "RemoveAnomalies", IDUserSend: otherUser.GetID(),
									IDMap: user.GetSquad().MapID})
							}
						}
					} else {
						mpCoordinate.DynamicObject = &dynamicObject
						go SendMessage(Message{Event: msg.Event, OtherUser: user.GetShortUserInfo(true, false), Q: msg.Q, R: msg.R,
							TypeSlot: msg.TypeSlot, Slot: msg.Slot, Box: nil, Reservoir: nil, DynamicObject: &dynamicObject,
							Name: diggerSlot.Equip.Name, IDMap: user.GetSquad().MapID})
					}
				}
			}
		}
	}
}
