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
	if find && user != nil && user.InBaseID == 0 {

		otherUsers := make([]*hostileMS, 0)

		usersGlobalWs, mx := Clients.GetAll()
		globalGame.GetPlaceCoordinate(user, usersGlobalWs, mp)
		for _, otherUser := range usersGlobalWs {
			if user.GetID() != otherUser.GetID() {
				otherUsers = append(otherUsers, GetShortUserInfo(otherUser))
			}
		}
		mx.Unlock()

		user.GetSquad().Afterburner = false
		user.GetSquad().MoveChecker = false
		user.GetSquad().CreateMove()
		user.GetSquad().HighGravity = globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)

		globalPipe <- Message{Event: "ConnectNewUser", OtherUser: GetShortUserInfo(user), idSender: user.GetID(), idMap: user.GetSquad().MapID}
		globalPipe <- Message{
			Event:      msg.Event,
			Map:        mp,
			User:       user,
			Squad:      user.GetSquad(),
			Bases:      bases.Bases.GetBasesByMap(mp.Id),
			OtherUsers: otherUsers,
			Boxes:      boxes.Boxes.GetAllBoxByMapID(mp.Id),
			idUserSend: user.GetID(),
			Credits:    user.GetCredits(),
			Experience: user.GetExperiencePoint(),
			idMap:      user.GetSquad().MapID,
		}

		// находим аномалии
		equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
		anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
		if err == nil {
			globalPipe <- Message{Event: "AnomalySignal", idUserSend: user.GetID(), Anomalies: anomalies, idMap: user.GetSquad().MapID}
		}
	} else {
		globalPipe <- Message{Event: "Error", Error: "no allow", idUserSend: user.GetID(), idMap: user.GetSquad().MapID}
	}
}
