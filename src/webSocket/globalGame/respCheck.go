package globalGame

import (
	"../../mechanics/gameObjects/base"
	"../../mechanics/globalGame"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func RespCheck(respBase *base.Base) bool { // заставляет игроков эвакуироватся с точки респауна базы
	x, y := globalGame.GetXYCenterHex(respBase.RespQ, respBase.RespR)

	lock := false
	usersGlobalWs, mx := Clients.GetAll()
	for ws, user := range usersGlobalWs {
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 150 {
			if !user.GetSquad().ForceEvacuation {
				globalPipe <- Message{Event: "setFreeResp", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
				go ForceEvacuation(ws, user, x, y)
			}
			lock = true
			user.GetSquad().ForceEvacuation = true
		}
	}
	mx.Unlock()

	return lock
}

func ForceEvacuation(ws *websocket.Conn, user *player.Player, x, y int) {
	timeCount := 0
	for {
		timeCount++
		time.Sleep(100 * time.Millisecond)

		if ws == nil {
			break
		}

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if dist > 150 {
			globalPipe <- Message{Event: "removeNoticeFreeResp", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
			user.GetSquad().ForceEvacuation = false
			break
		} else {
			if timeCount > 100 && !user.GetSquad().Evacuation {
				go evacuationSquad(ws)
			}
		}
	}
}
