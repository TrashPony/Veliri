package lobby

import (
	"../../mechanics/factories/storages"
	"github.com/gorilla/websocket"
)

func openWorkbench(ws *websocket.Conn, msg Message) {
	user := usersLobbyWs[ws]

	// GET STORAGE INVENTORY
	// todo GET current works
	baseStorage, find := storages.Storages.Get(user.GetID(), user.InBaseID)
	if user != nil && find {
		lobbyPipe <- Message{Event: "WorkbenchStorage", UserID: user.GetID(), Storage: baseStorage}
	}
}
