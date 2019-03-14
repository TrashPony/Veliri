package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func AnomaliesLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		for _, anomaly := range mp.Anomalies {
			launchAnomaly(anomaly)
		}
	}
}

func launchAnomaly(anomaly *_map.Anomalies) {
	if anomaly.Type == "mortality" {
		go mortalityAnomaly(anomaly)
	}
	if anomaly.Type == "unknown" {

	}
}

func mortalityAnomaly(anomaly *_map.Anomalies) {
	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()
	for ws, user := range users {
		// TODO провекра на попадание в зону, отправка сообщения что "опасность! вы находитесь рядом с аномалией!"
		println(ws, user)
		SquadDamage(user, anomaly.Power) // todo damage anomaly.Power * (radius/radius * dist))
	}
}

func SquadDamage(user *player.Player, damage int) {
	// 1 наносим урон корпусу
	user.GetSquad().MatherShip.HP -= damage

	//todo 2 наносим урон 2м рандомным эквипам

	if user.GetSquad().MatherShip.HP <= 0 {
		// игрок умирает 
	}
}
