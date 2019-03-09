package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
)

func DisconnectUser(user *player.Player, ws *websocket.Conn, onlyMessage bool) {
	if !onlyMessage {
		globalGame.Clients.DelClientByWS(ws)
	}

	if user != nil && user.GetSquad() != nil {
		go sendMessage(Message{Event: "DisconnectUser", OtherUser: GetShortUserInfo(user),
			idSender: user.GetID(), idMap: user.GetSquad().MapID})
	}
}
