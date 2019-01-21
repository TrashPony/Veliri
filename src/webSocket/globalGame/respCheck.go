package globalGame

import (
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func CheckTransportCoordinate(q, r, seconds int) bool { // заставляет игроков эвакуироватся с точки респауна базы
	x, y := globalGame.GetXYCenterHex(q, r)

	lock := false
	usersGlobalWs, mx := Clients.GetAll()
	for ws, user := range usersGlobalWs {
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 150 {
			if !user.GetSquad().ForceEvacuation {
				globalPipe <- Message{Event: "setFreeCoordinate", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Seconds: seconds}

				go ForceEvacuation(ws, user, x, y, seconds)
			}
			lock = true
			user.GetSquad().ForceEvacuation = true
		}
	}
	mx.Unlock()

	return lock
}

func ForceEvacuation(ws *websocket.Conn, user *player.Player, x, y, seconds int) {
	timeCount := 0
	for {
		timeCount++
		time.Sleep(100 * time.Millisecond)

		if ws == nil {
			break
		}

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if dist > 75 {
			globalPipe <- Message{Event: "removeNoticeFreeCoordinate", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			user.GetSquad().ForceEvacuation = false
			break
		} else {
			if timeCount > seconds*10 && !user.GetSquad().Evacuation {
				go evacuationSquad(ws)
			}
		}
	}
}
