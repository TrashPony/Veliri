package watchZone

import (
	"../game"
	"../player"
)

func getAllWatchObject(activeGame *game.Game, client *player.Player) {

	for _, xLine := range activeGame.GetUnits() {
		for _, gameUnit := range xLine {

			watchCoordinate, watchUnit, watchMatherShip, err := watch(gameUnit, client.GetLogin(), activeGame) //PermissionCoordinates(client, unit, units)

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

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileMatherShip(hostile)
						}
					}
				}

				for _, gameCoordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[gameCoordinate.X][gameCoordinate.Y]
					if ok {
						client.AddCoordinate(gameCoordinate)
					}
				}
			}
		}
	}

	for _, xLine := range activeGame.GetMatherShips() {
		for _, gameMatherShip := range xLine {

			watchCoordinate, watchUnit, watchMatherShip, err := watch(gameMatherShip, client.GetLogin(), activeGame)

			if err != nil { // если структура не моя то пропускаем дальнейшее действие
				continue
			} else {
				client.AddMatherShips(gameMatherShip)

				for _, xLine := range watchUnit {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileUnit(hostile)
						}
					}
				}

				for _, xLine := range watchMatherShip {
					for _, hostile := range xLine {
						if hostile.Owner != client.GetLogin() {
							client.AddHostileMatherShip(hostile)
						}
					}
				}

				for _, gameCoordinate := range watchCoordinate {
					_, ok := activeGame.GetMap().OneLayerMap[gameCoordinate.X][gameCoordinate.Y]
					if ok {
						client.AddCoordinate(gameCoordinate)
					}
				}
			}
		}
	}
}
