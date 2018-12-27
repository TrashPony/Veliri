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

	globalGame.GetPlaceCoordinate(user, usersGlobalWs, mp)

	user.GetSquad().Afterburner = false
	user.GetSquad().MoveChecker = false

	user.GetSquad().CreateMove()

	otherUsers := make([]*hostileMS, 0)

	for _, otherUser := range usersGlobalWs {
		if user.GetID() != otherUser.GetID() {
			otherUsers = append(otherUsers, GetShortUserInfo(otherUser))
		}
	}

	globalPipe <- Message{Event: "ConnectNewUser", OtherUser: GetShortUserInfo(user), idSender: user.GetID()}

	if find && user != nil && user.InBaseID == 0 {
		globalPipe <- Message{
			Event:      msg.Event,
			Map:        mp,
			User:       user,
			Squad:      user.GetSquad(),
			Bases:      bases.Bases.GetBasesByMap(mp.Id),
			OtherUsers: otherUsers,
			Boxes:      boxes.Boxes.GetAllBoxByMapID(mp.Id),
			idUserSend: user.GetID(),
		}
	} else {
		globalPipe <- Message{Event: "Error", Error: "no allow", idUserSend: user.GetID()}
	}
}
