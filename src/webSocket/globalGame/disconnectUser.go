package globalGame

import "../../mechanics/player"

func DisconnectUser(user *player.Player) {
	globalPipe <- Message{Event: "DisconnectUser", OtherUser: GetShortUserInfo(user),
		idSender: user.GetID(), idMap: user.GetSquad().MapID}
}
