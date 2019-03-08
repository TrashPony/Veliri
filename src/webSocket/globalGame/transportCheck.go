package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"time"
)

func CheckTransportCoordinate(q, r, seconds, distCheck, mapID int) bool { // заставляет игроков эвакуироватся с точки респауна базы
	x, y := globalGame.GetXYCenterHex(q, r)

	lock := false
	for ws, user := range globalGame.Clients.GetAll() {
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if int(dist) < distCheck && mapID == user.GetSquad().MapID {
			if !user.GetSquad().ForceEvacuation {
				globalPipe <- Message{Event: "setFreeCoordinate", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Seconds: seconds, Bot: user.Bot}
				go ForceEvacuation(ws, user, x, y, seconds, distCheck)
			}
			lock = true
			user.GetSquad().ForceEvacuation = true
		}
	}

	return lock
}

func ForceEvacuation(ws *websocket.Conn, user *player.Player, x, y, seconds, distCheck int) {
	timeCount := 0
	for {
		timeCount++
		time.Sleep(100 * time.Millisecond)

		if ws == nil {
			break
		}

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if int(dist) > distCheck {
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
