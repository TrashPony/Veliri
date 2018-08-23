package get

import (
	"../../../gameObjects/unit"
	"../../../localGame"
)

func AllUnits(game *localGame.Game) (map[int]map[int]*unit.Unit, []*unit.Unit ){
	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)

	for _, gamePlayer := range game.GetPlayers() {
		// todo обработка MS
		for _, playerUnit := range gamePlayer.GetSquad().MatherShip.Units {
			if playerUnit.Unit != nil {
				if playerUnit.Unit.OnMap {

					UnitEffects(playerUnit.Unit) // берем эфекты юнита
					addUnitToMap(units, playerUnit.Unit) // и кладем на карту

				} else {
					unitStorage = append(unitStorage, playerUnit.Unit)
					gamePlayer.AddUnitStorage(playerUnit.Unit)
				}
			}
		}
	}

	return units, unitStorage
}

func addUnitToMap(units map[int]map[int]*unit.Unit, gameUnit *unit.Unit)  {
	if units[gameUnit.X] != nil { // кладем юнита в матрицу
		units[gameUnit.X][gameUnit.Y] = gameUnit
	} else {
		units[gameUnit.X] = make(map[int]*unit.Unit)
		units[gameUnit.X][gameUnit.Y] = gameUnit
	}
}
