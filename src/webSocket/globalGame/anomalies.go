package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/remove"
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
		for ws, user := range users {
			if user.GetSquad() != nil && user.GetSquad().MapID == mp.Id {
				dist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, anomaly.X, anomaly.Y)

				if int(dist) < anomaly.Radius {
					// чем ближе к центру тем полнее урон
					SquadDamage(user, int(float64(anomaly.Power)*((float64(anomaly.Radius)-dist)/float64(anomaly.Radius))), ws)

					// отправка сообщения что "опасность! вы находитесь рядом с аномалией!"
					go sendMessage(Message{Event: "AnomalyCatch", idUserSend: user.GetID(), Bot: user.Bot, Anomaly: anomaly})
				}
			}
		}
		rLock.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func SquadDamage(user *player.Player, damage int, ws *websocket.Conn) {
	// 1 наносим урон корпусу
	user.GetSquad().MatherShip.HP -= damage

	// todo 2 наносим урон рандомным эквипам
	// todo обновление в бд

	go sendMessage(Message{Event: "DamageSquad", idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot, Squad: user.GetSquad()})
	if user.GetSquad().MatherShip.HP <= 0 {
		go sendMessage(Message{Event: "DeadSquad", OtherUser: GetShortUserInfo(user), idMap: user.GetSquad().MapID})
		// время для проигрыша анимации например
		time.Sleep(3 * time.Second)
		// удаляем отряд из игры
		remove.Squad(user.GetSquad())
		// отнимание всего отряда и инвентаря в трюме
		user.SetSquad(nil)
		// тащим юзера в последнюю посещенную им базу
		intoToBase(user, user.LastBaseID, ws)
	}
}
