package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
	"time"
)

func OutBase(base *base.Base) *coordinate.Coordinate {

	base.RespawnLock.Lock()
	defer func() {
		time.Sleep(time.Second)
		base.RespawnLock.Unlock()
	}()

	findRespawn := false
	var respCoordinate *coordinate.Coordinate

	for !findRespawn && respCoordinate == nil {
		findRespawn, respCoordinate = CheckBaseRespawn(base)
		// запускаем механизм проверки и эвакуации игрока с респауна))))
		time.Sleep(time.Millisecond * 100)
	}

	return respCoordinate
}

func CheckBaseRespawn(base *base.Base) (bool, *coordinate.Coordinate) {
	for _, resp := range base.Respawns {
		if !CheckTransportCoordinate(resp.Q, resp.R, 10, 100, base.MapID) {
			return true, resp
		}
	}

	return false, nil
}

func CheckTransportCoordinate(q, r, seconds, distCheck, mapID int) bool { // заставляет игроков эвакуироватся с точки респауна базы

	x, y := globalGame.GetXYCenterHex(q, r)

	lock := false
	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()

	for ws, user := range users {
		if user.GetSquad() != nil {
			dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

			if int(dist) < distCheck && mapID == user.GetSquad().MapID {
				if !user.GetSquad().ForceEvacuation {
					go SendMessage(Message{Event: "setFreeCoordinate", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Seconds: seconds, Bot: user.Bot})
					go ForceEvacuation(ws, user, x, y, seconds, distCheck)
				}
				lock = true
				user.GetSquad().ForceEvacuation = true
			}
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
			go SendMessage(Message{Event: "removeNoticeFreeCoordinate", IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID})
			user.GetSquad().ForceEvacuation = false
			break
		} else {
			if timeCount > seconds*10 && !user.GetSquad().Evacuation {
				go evacuationSquad(ws)
			}
		}
	}
}
