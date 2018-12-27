package globalGame

import "github.com/gorilla/websocket"

func afterburnerToggle(ws *websocket.Conn, msg Message) {
	user, ok := usersGlobalWs[ws]
	if ok {

		if user.GetSquad().Afterburner {
			user.GetSquad().Afterburner = false
		} else {
			user.GetSquad().Afterburner = true
		}

		msg.ToX = user.GetSquad().ToX
		msg.ToY = user.GetSquad().ToY

		move(ws, msg) // пересчитываем путь т.к. эффективность двиготеля изменилась
		globalPipe <- Message{Event: "AfterburnerToggle", Afterburner: user.GetSquad().Afterburner, idUserSend: user.GetID()}
	}
}
