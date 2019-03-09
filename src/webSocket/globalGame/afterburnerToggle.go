package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func afterburnerToggle(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	if user != nil {

		if user.GetSquad().Afterburner {
			user.GetSquad().Afterburner = false
		} else {
			user.GetSquad().Afterburner = true
		}

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		go sendMessage(Message{Event: "AfterburnerToggle", Afterburner: user.GetSquad().Afterburner, idUserSend: user.GetID(), idMap: user.GetSquad().MapID})
	}
}
