package globalGame

import (
	"../../mechanics/factories/boxes"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func openBox(ws *websocket.Conn, msg Message, stopMove chan bool, moveChecker *bool) {
	user := usersGlobalWs[ws]

	box := boxes.Boxes.Get(msg.BoxID)

	if box != nil {
		x, y := globalGame.GetXYCenterHex(box.Q, box.R)

		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, x, y)

		if dist < 150 {
			if *moveChecker {
				stopMove <- true // останавливаем движение
			}
			ws.WriteJSON(Message{Event: msg.Event, BoxID: box.ID, Inventory: box.GetStorage()})
		}
	}
}
