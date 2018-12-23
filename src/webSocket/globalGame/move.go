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

		path := globalGame.MoveTo(user, msg.ToX, msg.ToY, mp)

		err := ws.WriteJSON(Message{Event: "PreviewPath", Path: path})
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
			err := ws.WriteJSON(Message{Event: msg.Event, PathUnit: pathUnit})
			if err != nil {
				DisconnectUser(usersGlobalWs[ws])
				*moveChecker = false
				return
			}

			for ws, otherUser := range usersGlobalWs {
				if otherUser.GetID() != user.GetID() {
					ws.WriteJSON(Message{Event: "MoveOtherUser", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit})
				}
			}

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
