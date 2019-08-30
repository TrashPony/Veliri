package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
)

func DisconnectUser(user *player.Player, onlyMessage bool) {
	if !onlyMessage && user != nil {
		globalGame.Clients.DelClientByID(user.ID)
	}

	if user != nil && user.GetSquad() != nil {
		go SendMessage(Message{Event: "DisconnectUser", OtherUser: user.GetShortUserInfo(true),
			IDSender: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
	}
}
