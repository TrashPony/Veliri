package globalGame

import (
	"../../mechanics/db/squad/update"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit, exit chan bool, move *bool) {
	for i, pathUnit := range path {
		select {
		case exitNow := <-exit:
			if exitNow {
				*move = false
				return
			}
		default:
			err := ws.WriteJSON(Message{Event: msg.Event, PathUnit: pathUnit})
			if err != nil {
				*move = false
				return
			}

			if i+1 != len(path) { // бкз этого ифа канал будет ловить дед лок
				time.Sleep(100 * time.Millisecond)
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
	*move = false
	return
}
