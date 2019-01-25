package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func move(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find && user.InBaseID == 0 {

		stopMove(ws, false)

		path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)

		if len(path) > 1 {
			user.GetSquad().ToX = float64(path[len(path)-1].X)
			user.GetSquad().ToY = float64(path[len(path)-1].Y)
		} else {
			user.GetSquad().ToX = float64(user.GetSquad().GlobalX)
			user.GetSquad().ToY = float64(user.GetSquad().GlobalY)
		}

		if err != nil && len(path) == 0 {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		}

		globalPipe <- Message{Event: "PreviewPath", Path: path, idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
		if err != nil {
			DisconnectUser(user)
		}

		go MoveUserMS(ws, msg, user, path)
		user.GetSquad().MoveChecker = true
	}
}

func stopMove(ws *websocket.Conn, reserSpeed bool) {
	user := Clients.GetByWs(ws)
	if user.GetSquad().MoveChecker && user.GetSquad().GetMove() != nil {
		user.GetSquad().GetMove() <- true // останавливаем прошлое движение
		if reserSpeed {
			user.GetSquad().CurrentSpeed = 0
		}
	}
}

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit) {
	// TODO убрать много селектов но я не знаю как инче %(
	for i, pathUnit := range path {
		select {
		case exitNow := <-user.GetSquad().GetMove():
			if exitNow {
				user.GetSquad().MoveChecker = false
				return
			}
		default:

			newGravity := globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)
			if user.GetSquad().HighGravity != newGravity {
				user.GetSquad().HighGravity = newGravity
				globalPipe <- Message{Event: "ChangeGravity", idUserSend: user.GetID(), Squad: user.GetSquad(), idMap: user.GetSquad().MapID}
				select {
				case exitNow := <-user.GetSquad().GetMove():
					if exitNow {
						user.GetSquad().MoveChecker = false
						move(ws, msg)
						return
					}
				default:
					user.GetSquad().MoveChecker = false
					move(ws, msg)
					return
				}
			}

			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)
			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			for { // ожидаем пока другой игрок уйдет с пути или первый не изменил путь

				obstacle := !globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, Clients.GetAll())

				if obstacle {
					select {
					case exitNow := <-user.GetSquad().GetMove():
						if exitNow {
							user.GetSquad().MoveChecker = false
							return
						}
					default:
						time.Sleep(100 * time.Millisecond)
					}
				} else {
					break
				}
			}

			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				globalPipe <- Message{Event: "AnomalySignal", idUserSend: user.GetID(), Anomalies: anomalies, idMap: user.GetSquad().MapID}
			}

			// если на пути встречается ящик то мы его давим и падает скорость
			mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, user.GetSquad().MapID)
			if mapBox != nil {
				globalPipe <- Message{Event: "DestroyBox", BoxID: mapBox.ID, idMap: user.GetSquad().MapID}
				boxes.Boxes.DestroyBox(mapBox)
				user.GetSquad().CurrentSpeed -= float64(user.GetSquad().MatherShip.Body.Speed)
				select {
				case exitNow := <-user.GetSquad().GetMove():
					if exitNow {
						user.GetSquad().MoveChecker = false
						move(ws, msg)
						return
					}
				default:
					user.GetSquad().MoveChecker = false
					move(ws, msg)
					return
				}
			}

			// если клиент отключился то останавливаем его
			if ws == nil || Clients.GetByWs(ws) == nil {
				user.GetSquad().MoveChecker = false
				return
			}

			coor := globalGame.HandlerDetect(user)
			if coor != nil && coor.HandlerOpen {
				user.GetSquad().MoveChecker = false
				HandlerParse(user, coor)
				return
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit, idMap: user.GetSquad().MapID}

			if i+1 != len(path) { // бeз этого ифа канал будет ловить деад лок
				time.Sleep(100 * time.Millisecond)
				user.GetSquad().CurrentSpeed = pathUnit.Speed
			} else {
				user.GetSquad().CurrentSpeed = 0
			}

			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)

			if ((pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != user.GetSquad().Q && pathUnit.R != user.GetSquad().R)) ||
				i+1 == len(path) {
				user.GetSquad().Q = pathUnit.Q
				user.GetSquad().R = pathUnit.R

				go update.Squad(user.GetSquad(), false)
			}
		}
	}
	user.GetSquad().MoveChecker = false
	return
}
