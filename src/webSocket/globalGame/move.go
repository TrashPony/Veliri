package globalGame

import (
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
	"../../mechanics/db/squad/update"
)

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit, exit chan bool, move *bool) {
	oldQ := 0
	oldR := 0

	for _, pathUnit := range path {
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

			time.Sleep(100 * time.Millisecond)
			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)

			if (pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != oldQ && pathUnit.R != oldR) {
				user.GetSquad().Q = pathUnit.Q
				user.GetSquad().R = pathUnit.R

				oldQ = pathUnit.Q
				oldR = pathUnit.R

				go update.Squad(user.GetSquad(), false)
			}
		}
	}
	*move = false
	return
}
