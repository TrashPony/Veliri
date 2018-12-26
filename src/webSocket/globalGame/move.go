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
	// TODO попробовать использовать контекст
	if find && user.InBaseID == 0 {
		if *moveChecker {
			stopMove <- true // останавливаем прошлое движение
		}

		path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)
		if err != nil && len(path) == 0 {
			err = ws.WriteJSON(Message{Event: "Error", Error: err.Error()})
		}

		err = ws.WriteJSON(Message{Event: "PreviewPath", Path: path})
		if err != nil {
			DisconnectUser(usersGlobalWs[ws])
		}

		go MoveUserMS(ws, msg, user, path, stopMove, moveChecker)
		*moveChecker = true
	}
}

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit, exit chan bool, moveChecker *bool) {
	for i, pathUnit := range path {
		select {
		case exitNow := <-exit:
			if exitNow {
				*moveChecker = false
				return
			}
		default:
			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots)

			// если клиент отключился то останавливаем его
			if ws == nil || usersGlobalWs[ws] == nil {
				*moveChecker = false
				return
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: msg.Event, OtherUser: GetShortUserInfo(user), PathUnit: pathUnit}

			if i+1 != len(path) { // бeз этого ифа канал будет ловить деад лок
				time.Sleep(100 * time.Millisecond)
			}

			// TODO проверка колизий игрок - игрок

			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)
			user.GetSquad().CurrentSpeed = pathUnit.Speed

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
