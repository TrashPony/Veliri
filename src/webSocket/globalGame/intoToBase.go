package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func intoToBase(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	user := usersGlobalWs[ws]

	intoBase, _ := bases.Bases.Get(msg.BaseID)
	x, y := globalGame.GetXYCenterHex(intoBase.Q, intoBase.R)

	dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)
	if dist < 220 { // 220 пикселей, выбрано рандомно
		if *moveChecker {
			stopMove <- true // останавливаем движение
		}

		user.InBaseID = intoBase.ID
		bases.UserIntoBase(user.GetID(), intoBase.ID)
		user.GetSquad().GlobalX = 0
		user.GetSquad().GlobalY = 0

		ws.WriteJSON(Message{Event: "IntoToBase"})
		DisconnectUser(user)
	}
}
