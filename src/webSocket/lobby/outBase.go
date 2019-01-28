package lobby

import (
	"../../mechanics/lobby"
	"github.com/gorilla/websocket"
)

func outBase(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	if user != nil {
		err := lobby.OutBase(usersLobbyWs[ws])

		// todo запускать метод в отдельной горутине
		// todo флаг выхода с базы, т.к. пока освобождается респаун игрок может передумать

		if err != nil {
			lobbyPipe <- Message{Event: "Error", Error: err.Error(), UserID: user.GetID()}
		} else {
			lobbyPipe <- Message{Event: msg.Event, UserID: user.GetID()}
		}
	}
}
