package globalGame

import "github.com/TrashPony/Veliri/src/mechanics/player"

func DisconnectUser(user *player.Player) {
	globalPipe <- Message{Event: "DisconnectUser", OtherUser: GetShortUserInfo(user),
		idSender: user.GetID(), idMap: user.GetSquad().MapID}
}
