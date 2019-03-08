package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/gorilla/websocket"
)

func loadGame(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)
	// TODO при загрузке нового игрока весь мир замерает на некоторые секунды, возможно актуально только для ботов
	if user != nil {

		mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

		if find && user.InBaseID == 0 {

			otherUsers := make([]*hostileMS, 0)

			globalGame.GetPlaceCoordinate(user, globalGame.Clients.GetAll(), mp)
			for _, otherUser := range globalGame.Clients.GetAll() {
				if user.GetID() != otherUser.GetID() && user.GetSquad().MapID == otherUser.GetSquad().MapID && otherUser.InBaseID == 0 {
					otherUsers = append(otherUsers, GetShortUserInfo(otherUser))
				}
			}

			user.GetSquad().Afterburner = false
			user.GetSquad().MoveChecker = false
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
}
