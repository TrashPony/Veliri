package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func move(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
	user := usersGlobalWs[ws]

	if find && user.InBaseID == 0 {
		if *moveChecker {
			stopMove <- true // останавливаем прошлое движение
		}

		path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp, usersGlobalWs)

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

		go MoveUserMS(ws, user, path, stopMove, moveChecker)
		*moveChecker = true
	}
}

func MoveUserMS(ws *websocket.Conn, user *player.Player, path []globalGame.PathUnit, exit chan bool, moveChecker *bool) {
	for i, pathUnit := range path {
		select {
		case exitNow := <-exit:
			if exitNow {
				*moveChecker = false
				return
			}
		default:
			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner)

			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			if !globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, usersGlobalWs) {
				// на пути появился непроходимых игрок
				*moveChecker = false
				return
			}

			// если клиент отключился то останавливаем его
			if ws == nil || usersGlobalWs[ws] == nil {
				*moveChecker = false
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
	*moveChecker = false
	return
}
