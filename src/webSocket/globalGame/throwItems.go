package globalGame

import (
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func throwItems(ws *websocket.Conn, msg Message) {
	user := usersGlobalWs[ws]

	err, newBox := globalGame.ThrowItems(user, msg.ThrowItems)

	if err != nil {
		ws.WriteJSON(Message{Event: "Error", Error: err.Error()})
	} else {
		ws.WriteJSON(Message{Event: "UpdateInventory"})
		for ws := range usersGlobalWs {
			ws.WriteJSON(Message{Event: "NewBox", Box: newBox, X: user.GetSquad().GlobalX, Y: user.GetSquad().GlobalY})
		}
	}
}
