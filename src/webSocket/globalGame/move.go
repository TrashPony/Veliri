package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func move(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	if find && user.InBaseID == 0 && !user.GetSquad().Evacuation {

		stopMove(ws, false)

		path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)
		user.GetSquad().ActualPath = &path

		if len(path) > 1 {
			user.GetSquad().ToX = float64(path[len(path)-1].X)
			user.GetSquad().ToY = float64(path[len(path)-1].Y)
		} else {
			user.GetSquad().ToX = float64(user.GetSquad().GlobalX)
			user.GetSquad().ToY = float64(user.GetSquad().GlobalY)
		}

		if err != nil && len(path) == 0 {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
		}

		globalPipe <- Message{Event: "PreviewPath", Path: path, idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
		if err != nil {
			DisconnectUser(user)
		}

		go MoveUserMS(ws, msg, user, &path)
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

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path *[]squad.PathUnit) {

	moveRepeat := false

	defer func() {
		user.GetSquad().MoveChecker = false
		stopMove(ws, false)
		if moveRepeat {
			move(ws, msg)
		}
	}()

	for i, pathUnit := range *path {
		select {
		case exitNow := <-user.GetSquad().GetMove():
			if exitNow {
				return
			}
		default:

			if user.GetSquad().ActualPath != path {
				return
			}

			newGravity := globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)
			if user.GetSquad().HighGravity != newGravity {
				user.GetSquad().HighGravity = newGravity
				globalPipe <- Message{Event: "ChangeGravity", idUserSend: user.GetID(), Squad: user.GetSquad(), idMap: user.GetSquad().MapID, Bot: user.Bot}
				moveRepeat = true
				return
			}

			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)
			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			// колизии игрок-игрок // TODO столкновения,  урон, оталкивание
			//for !globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, Clients.GetAll()) {
			//	time.Sleep(1 * time.Second)
			//	return
			//}

			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				globalPipe <- Message{Event: "AnomalySignal", idUserSend: user.GetID(), Anomalies: anomalies, idMap: user.GetSquad().MapID, Bot: user.Bot}
			}

			// если на пути встречается ящик то мы его давим и падает скорость
			mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, user.GetSquad().MapID, user.GetSquad().MatherShip.Body)
			if mapBox != nil {
				globalPipe <- Message{Event: "DestroyBox", BoxID: mapBox.ID, idMap: user.GetSquad().MapID}
				boxes.Boxes.DestroyBox(mapBox)
				user.GetSquad().CurrentSpeed -= float64(user.GetSquad().MatherShip.Body.Speed)
				moveRepeat = true
				return
			}

			// если клиент отключился то останавливаем его
			if ws == nil || Clients.GetByWs(ws) == nil {
				return
			}

			coor := globalGame.HandlerDetect(user)
			if coor != nil && coor.HandlerOpen {
				HandlerParse(user, coor)
				return
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID, Bot: user.Bot}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit, idMap: user.GetSquad().MapID}

			if i+1 != len(*path) { // бeз этого ифа канал будет ловить деад лок
				time.Sleep(100 * time.Millisecond)
				user.GetSquad().CurrentSpeed = pathUnit.Speed
			} else {
				user.GetSquad().CurrentSpeed = 0
			}

			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)

			if ((pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != user.GetSquad().Q && pathUnit.R != user.GetSquad().R)) ||
				i+1 == len(*path) {
				user.GetSquad().Q = pathUnit.Q
				user.GetSquad().R = pathUnit.R

				if !user.Bot {
					go update.Squad(user.GetSquad(), false)
				}
			}
		}
	}
}
