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

	dist := ((user.GetSquad().GlobalX - x) * (user.GetSquad().GlobalX - x)) +
		((user.GetSquad().GlobalY - y) * (user.GetSquad().GlobalX - y))

	if dist < 320*320 { // 320*320 это 320 пикселей, выбрано рандомно

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
