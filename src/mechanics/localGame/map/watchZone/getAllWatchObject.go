package watchZone

import (
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func getAllWatchObject(activeGame *localGame.Game, client *player.Player) {

	for _, xLine := range activeGame.GetUnits() {
		for _, gameUnit := range xLine {

			if gameUnit.HP <= 0 { // если у юнита нет хп то он не может смотреть ¯\_(ツ)_/¯
				continue
			}

			watchCoordinate, watchUnit, err := Watch(gameUnit, client.GetLogin(), activeGame)

			// если юнит не игрока и не его союзника то пропускаем следующие действия
			owner := activeGame.GetUserByName(gameUnit.GetOwnerUser())

			if owner != nil {
				if err != nil && !activeGame.CheckPacts(client.GetID(), owner.GetID()) {
					continue
				}

				// если юнит игрока то добавляем его в пул юнитов игрока
				if gameUnit.GetOwnerUser() == client.GetLogin() {
					client.AddUnit(gameUnit)
				}

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, gameCoordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[gameCoordinate.Q][gameCoordinate.R]
					if ok {
						client.AddCoordinate(gameCoordinate)
					}
				}
			}
		}
	}
}
