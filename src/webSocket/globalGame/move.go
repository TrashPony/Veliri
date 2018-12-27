package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func move(ws *websocket.Conn, msg Message) {
	mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
	user := usersGlobalWs[ws]

	if find && user.InBaseID == 0 {
		if user.GetSquad().MoveChecker {
			user.GetSquad().GetMove() <- true // останавливаем прошлое движение
		}

		path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)

		if len(path) > 1 {
			user.GetSquad().ToX = float64(path[len(path)-1].X)
			user.GetSquad().ToY = float64(path[len(path)-1].Y)
		} else {
			user.GetSquad().ToX = float64(user.GetSquad().GlobalX)
			user.GetSquad().ToY = float64(user.GetSquad().GlobalY)
		}

		if err != nil && len(path) == 0 {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
		}

		globalPipe <- Message{Event: "PreviewPath", Path: path, idUserSend: user.GetID()}
		if err != nil {
			DisconnectUser(usersGlobalWs[ws])
		}

		go MoveUserMS(ws, msg, user, path)
		user.GetSquad().MoveChecker = true
	}
}

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit) {
	for i, pathUnit := range path {
		select {
		case exitNow := <-user.GetSquad().GetMove():
			if exitNow {
				user.GetSquad().MoveChecker = false
				return
			}
		default:
			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner)

			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			for { // ожидаем пока другой игрок уйдет с пути или первый не изменил путь
				if !globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, usersGlobalWs) {
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

			// если клиент отключился то останавливаем его
			if ws == nil || usersGlobalWs[ws] == nil {
				user.GetSquad().MoveChecker = false
				return
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit}

			if i+1 != len(path) { // бeз этого ифа канал будет ловить деад лок
				time.Sleep(100 * time.Millisecond)
				user.GetSquad().CurrentSpeed = pathUnit.Speed
			} else {
				user.GetSquad().CurrentSpeed = 0
			}

			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)

			if (pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != user.GetSquad().Q && pathUnit.R != user.GetSquad().R) {
				user.GetSquad().Q = pathUnit.Q
				user.GetSquad().R = pathUnit.R

				go update.Squad(user.GetSquad(), false)
			}
		}
	}
	user.GetSquad().MoveChecker = false
	return
}
