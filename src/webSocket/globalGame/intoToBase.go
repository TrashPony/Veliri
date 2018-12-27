package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func intoToBase(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	intoBase, find := bases.Bases.Get(msg.BaseID)
	if find {
		x, y := globalGame.GetXYCenterHex(intoBase.Q, intoBase.R)

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
		if dist < 270 { // 250 пикселей, выбрано рандомно
			if user.GetSquad().MoveChecker {
				user.GetSquad().GetMove() <- true // останавливаем движение
			}

			user.InBaseID = intoBase.ID
			bases.UserIntoBase(user.GetID(), intoBase.ID)
			user.GetSquad().GlobalX = 0
			user.GetSquad().GlobalY = 0

			globalPipe <- Message{Event: "IntoToBase", idUserSend: user.GetID()}
			DisconnectUser(user)
		}
	}
}
