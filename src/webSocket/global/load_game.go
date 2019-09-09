package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
)

func LoadGame(user *player.Player, msg Message) {
	// TODO при загрузке нового игрока весь мир замерает на некоторые секунды, возможно актуально только для ботов
	mp, find := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)

	if find {

		// обнуляем все параметры глобальной игры
		user.GetSquad().MatherShip.Afterburner = false
		user.GetSquad().MatherShip.MoveChecker = false
		user.GetSquad().MatherShip.ActualPath = nil
		user.GetSquad().MatherShip.HighGravity = globalGame.GetGravity(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, user.GetSquad().MatherShip.MapID)

		//TODO globalGame.GetPlaceCoordinate(user)
		user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y = globalGame.GetXYCenterHex(user.GetSquad().MatherShip.Q, user.GetSquad().MatherShip.R)

		go SendMessage(Message{Event: "ConnectNewUser", ShortUnit: user.GetSquad().MatherShip.GetShortInfo(), IDSender: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID})
		go SendMessage(Message{
			Event:       msg.Event,
			Map:         mp,
			User:        user,
			Squad:       user.GetSquad(),
			Bases:       bases.Bases.GetBasesByMap(mp.Id),
			Boxes:       boxes.Boxes.GetAllBoxByMapID(mp.Id),
			IDUserSend:  user.GetID(),
			Credits:     user.GetCredits(),
			IDMap:       user.GetSquad().MatherShip.MapID,
			ShortUnits:  globalGame.Clients.GetAllShortUnits(user.GetSquad().MatherShip.MapID, true),
			Bot:         user.Bot,
			HighGravity: globalGame.GetGravity(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, user.GetSquad().MatherShip.MapID),
		})

		// находим аномалии
		equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
		anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
		if err == nil {
			go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
		}
	} else {
		go SendMessage(Message{Event: "Error", Error: "no allow", IDUserSend: user.GetID(), IDMap: user.GetSquad().MatherShip.MapID, Bot: user.Bot})
	}
}
