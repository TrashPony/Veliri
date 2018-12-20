package globalGame

import (
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path []globalGame.PathUnit, exit chan bool, move *bool) {

	for _, pathUnit := range path {

		select {
		case exit := <-exit:
			if exit {
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
		}
	}
	*move = false
	return
}
