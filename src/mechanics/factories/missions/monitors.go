package missions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

func toSector(gameMission *mission.Mission, action *mission.Action, client *player.Player) {
	// проверяем что бы все предыдущие экшоны были выполнены
	for {
		if gameMission.CheckAvailableActionByIndex(action.Number) {
			// этот воркер проверяет что бы игрок находился в нужном секторе
			if client.GetSquad() != nil && client.GetSquad().MatherShip != nil && client.GetSquad().MatherShip.MapID == action.MapID {
				action.Complete = true
			} else {
				action.Complete = false
			}
		} // поидее ситуации иначе произойти не должно)

		time.Sleep(100 * time.Millisecond)
	}
}

func toXY(gameMission *mission.Mission, action *mission.Action, client *player.Player) {
	// проверяет что игрок находится в Q,R радиусе Radius в нужном секторе
	for {
		if gameMission.CheckAvailableActionByIndex(action.Number) {
			if client.GetSquad() != nil && client.GetSquad().MatherShip != nil && client.GetSquad().MatherShip.MapID == action.MapID {

				dist := game_math.GetBetweenDist(client.GetSquad().MatherShip.X, client.GetSquad().MatherShip.Y, action.X, action.Y)
				if int(dist) < action.Radius {
					action.Complete = true
				} else {
					action.Complete = false
				}
			} else {
				action.Complete = false
			}
		} // поидее ситуации иначе произойти не должно)
		time.Sleep(100 * time.Millisecond)
	}
}
