package globalGame

import (
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func throwItems(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	if user != nil {
		err, newBox := globalGame.ThrowItems(user, msg.ThrowItems)

		if err != nil {
			globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID()}
		} else {
			globalPipe <- Message{Event: "UpdateInventory", idUserSend: user.GetID()}
			globalPipe <- Message{Event: "NewBox", Box: newBox, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY}
		}
	}
}
