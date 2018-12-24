package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func loadGame(ws *websocket.Conn, msg Message) {
	mp, find := maps.Maps.GetByID(usersGlobalWs[ws].GetSquad().MapID)
	user := usersGlobalWs[ws]

	if user.GetSquad().GlobalX == 0 && user.GetSquad().GlobalY == 0 {
		x, y := globalGame.GetXYCenterHex(user.GetSquad().Q, user.GetSquad().R)
		user.GetSquad().GlobalX = x
		user.GetSquad().GlobalY = y
		user.GetSquad().CurrentSpeed = 0
	}

	otherUsers := make([]*hostileMS, 0)

	for ws, otherUser := range usersGlobalWs {
		if otherUser.GetID() != user.GetID() {
			otherUsers = append(otherUsers, GetShortUserInfo(otherUser))
			ws.WriteJSON(Message{Event: "ConnectNewUser", OtherUser: GetShortUserInfo(user)})
		}
	}

	if find && user != nil && user.InBaseID == 0 {
		ws.WriteJSON(Message{
			Event:      msg.Event,
			Map:        mp,
			User:       user,
			Squad:      user.GetSquad(),
			Bases:      bases.Bases.GetBasesByMap(mp.Id),
			OtherUsers: otherUsers,
			Boxes:      boxes.Boxes.GetAllBoxByMapID(mp.Id),
		})
	} else {
		ws.WriteJSON(Message{Event: "Error", Error: "no allow"})
	}
}
