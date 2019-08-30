package ai

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	wsGlobal "github.com/TrashPony/Veliri/src/webSocket/global"
	"time"
)

func AnomaliesLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		for _, anomaly := range mp.Anomalies {
			launchAnomaly(anomaly, mp)
		}
	}
}

func launchAnomaly(anomaly *_map.Anomalies, mp *_map.Map) {
	if anomaly.Type == "mortality" {
		go mortalityAnomaly(anomaly, mp)
	}
	if anomaly.Type == "unknown" {

	}
}

func mortalityAnomaly(anomaly *_map.Anomalies, mp *_map.Map) {
	for {
		users, rLock := globalGame.Clients.GetAll()
		for _, user := range users {
			if user.GetSquad() != nil && user.GetSquad().MatherShip.MapID == mp.Id {
				dist := globalGame.GetBetweenDist(user.GetSquad().MatherShip.X, user.GetSquad().MatherShip.Y, anomaly.X, anomaly.Y)

				if int(dist) < anomaly.Radius {
					// чем ближе к центру тем полнее урон
					wsGlobal.SquadDamage(user, int(float64(anomaly.Power)*((float64(anomaly.Radius)-dist)/float64(anomaly.Radius))), user.GetSquad().MatherShip)

					// отправка сообщения что "опасность! вы находитесь рядом с аномалией!"
					go wsGlobal.SendMessage(wsGlobal.Message{Event: "AnomalyCatch", IDUserSend: user.GetID(), Bot: user.Bot, Anomaly: anomaly})
				}
			}
		}
		rLock.Unlock()
		time.Sleep(1 * time.Second)
	}
}
