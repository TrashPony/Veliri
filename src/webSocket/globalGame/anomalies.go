package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"time"
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
		dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, anomaly.X, anomaly.Y)

		if int(dist) < anomaly.Radius {
			// чем ближе к центру тем полнее урон
			SquadDamage(user, anomaly.Power*(anomaly.Radius/anomaly.Radius-int(dist)), ws)

			// отправка сообщения что "опасность! вы находитесь рядом с аномалией!"
			go sendMessage(Message{Event: "AnomalyCatch", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot, Anomaly: anomaly})
		}
	}
}

func SquadDamage(user *player.Player, damage int, ws *websocket.Conn) {
	// 1 наносим урон корпусу
	user.GetSquad().MatherShip.HP -= damage

	//todo 2 наносим урон рандомным эквипам

	go sendMessage(Message{Event: "DamageSquad", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot, Squad: user.GetSquad()})
	if user.GetSquad().MatherShip.HP <= 0 {
		go sendMessage(Message{Event: "DeadSquad", OtherUser: GetShortUserInfo(user), idMap: user.GetSquad().MapID})
		user.SetSquad(nil) // отнимание всего отряда и инвентаря в трюме

		time.Sleep(3 * time.Second)// время для проигрыша анимации например
		intoToBase(user, user.LastBaseID, ws)
	}
}
