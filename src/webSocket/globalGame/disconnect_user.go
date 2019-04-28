package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func DisconnectUser(user *player.Player, ws *websocket.Conn, onlyMessage bool) {
	if !onlyMessage {
		globalGame.Clients.DelClientByWS(ws)
	}

	if user != nil && user.GetSquad() != nil {
		go SendMessage(Message{Event: "DisconnectUser", OtherUser: user.GetShortUserInfo(true, false),
			IDSender: user.GetID(), IDMap: user.GetSquad().MapID})
	}
}
