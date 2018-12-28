package globalGame

import (
	"../../mechanics/factories/bases"
	"../../mechanics/factories/boxes"
	"../../mechanics/factories/maps"
	"../../mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func loadGame(ws *websocket.Conn, msg Message) {
	user := Clients.GetByWs(ws)

	mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

	usersGlobalWs, mx := Clients.GetAll()
	globalGame.GetPlaceCoordinate(user, usersGlobalWs, mp)
	mx.Unlock()

	user.GetSquad().Afterburner = false
	user.GetSquad().MoveChecker = false
	user.GetSquad().CreateMove()

	otherUsers := make([]*hostileMS, 0)

	mx.Lock()
	for _, otherUser := range usersGlobalWs {
		if user.GetID() != otherUser.GetID() {
			otherUsers = append(otherUsers, GetShortUserInfo(otherUser))
		}
	}
	mx.Unlock()

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
