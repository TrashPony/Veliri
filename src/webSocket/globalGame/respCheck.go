package globalGame

import (
	"../../mechanics/gameObjects/base"
	"../../mechanics/globalGame"
	"time"
)

func RespCheck(respBase *base.Base) { // заставляет игроков эвакуироватся с точки респауна базы
	x, y := globalGame.GetXYCenterHex(respBase.RespQ, respBase.RespR)

	for ws, user := range usersGlobalWs {
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 300 && !user.GetSquad().Evacuation {
			globalPipe <- Message{Event: "setFreeResp", idUserSend: user.GetID()}
			timeCount := 0
			for {

				timeCount++
				time.Sleep(100 * time.Millisecond)

				if ws == nil {
					break
				}

				dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

				if dist > 300 {
					break
				} else {
					if timeCount > 50 && !user.GetSquad().Evacuation{
						go evacuationSquad(ws)
					}
				}
			}
		}
	}
}
