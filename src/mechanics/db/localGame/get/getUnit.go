package get

import (
	"../../../gameObjects/unit"
	"../../../localGame"
)

func AllUnits(game *localGame.Game) (map[int]map[int]*unit.Unit, []*unit.Unit) {
	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)

	for _, gamePlayer := range game.GetPlayers() {

		gamePlayer.GetSquad().MatherShip.Owner = gamePlayer.GetLogin()
		UnitEffects(gamePlayer.GetSquad().MatherShip)         // берем эфекты ms
		addUnitToMap(units, gamePlayer.GetSquad().MatherShip) // и кладем на карту, ms на карте с начала игры
		gamePlayer.GetSquad().MatherShip.CalculateParams()    // пересчитываем статы

		for _, playerUnit := range gamePlayer.GetSquad().MatherShip.Units {
			if playerUnit.Unit != nil {
				if playerUnit.Unit.OnMap {

					playerUnit.Unit.Owner = gamePlayer.GetLogin()
					UnitEffects(playerUnit.Unit)         // берем эфекты юнита
					addUnitToMap(units, playerUnit.Unit) // и кладем на карту

				} else {
					playerUnit.Unit.Owner = gamePlayer.GetLogin()
					unitStorage = append(unitStorage, playerUnit.Unit)
					gamePlayer.AddUnitStorage(playerUnit.Unit)
				}

				playerUnit.Unit.CalculateParams()
			}
		}
	}
	return units, unitStorage
}

func addUnitToMap(units map[int]map[int]*unit.Unit, gameUnit *unit.Unit) {
	if units[gameUnit.Q] != nil { // кладем юнита в матрицу
		units[gameUnit.Q][gameUnit.R] = gameUnit
	} else {
		units[gameUnit.Q] = make(map[int]*unit.Unit)
		units[gameUnit.Q][gameUnit.R] = gameUnit
	}
}
