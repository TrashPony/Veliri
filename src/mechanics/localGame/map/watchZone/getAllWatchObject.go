package watchZone

import (
	"../../../localGame"
	"../../../player"
)

func getAllWatchObject(activeGame *localGame.Game, client *player.Player) {

	for _, xLine := range activeGame.GetUnits() {
		for _, gameUnit := range xLine {

			watchCoordinate, watchUnit, err := watch(gameUnit, client.GetLogin(), activeGame)

			if err != nil { // если крип не мой то пропускаем дальнейшее действие
				continue
			} else {
				client.AddUnit(gameUnit)

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
