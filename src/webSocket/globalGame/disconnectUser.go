package globalGame

import "../../mechanics/player"

func DisconnectUser(user *player.Player) {
	for ws, otherUser := range usersGlobalWs {
		if user != nil && otherUser != nil {
			if otherUser.GetID() != user.GetID() {
				ws.WriteJSON(Message{Event: "DisconnectUser", OtherUser: GetShortUserInfo(user)})
			}
		}
	}
}
