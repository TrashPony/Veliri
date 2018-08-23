package get

import (
	"../../../gameObjects/unit"
	"../../../gameObjects/matherShip"
	"../../../localGame"
)

func AllUnits(game *localGame.Game) (map[int]map[int]*unit.Unit, []*unit.Unit, map[int]map[int]*matherShip.MatherShip){
	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)
	var motherShips = make(map[int]map[int]*matherShip.MatherShip)

	for _, gamePlayer := range game.GetPlayers() {

		// todo обработка MSEffects
		gamePlayer.GetMatherShip().Owner = gamePlayer.GetLogin()
		addMStToMap(motherShips, gamePlayer.GetMatherShip())

		for _, playerUnit := range gamePlayer.GetSquad().MatherShip.Units {
			if playerUnit.Unit != nil {
				if playerUnit.Unit.OnMap {

					playerUnit.Unit.Owner = gamePlayer.GetLogin()
					UnitEffects(playerUnit.Unit) // берем эфекты юнита
					addUnitToMap(units, playerUnit.Unit) // и кладем на карту

				} else {
					playerUnit.Unit.Owner = gamePlayer.GetLogin()
					unitStorage = append(unitStorage, playerUnit.Unit)
					gamePlayer.AddUnitStorage(playerUnit.Unit)
				}
			}
		}
	}

	return units, unitStorage, motherShips
}

func addUnitToMap(units map[int]map[int]*unit.Unit, gameUnit *unit.Unit)  {
	if units[gameUnit.X] != nil { // кладем юнита в матрицу
		units[gameUnit.X][gameUnit.Y] = gameUnit
	} else {
		units[gameUnit.X] = make(map[int]*unit.Unit)
		units[gameUnit.X][gameUnit.Y] = gameUnit
	}

}

func addMStToMap(motherShips map[int]map[int]*matherShip.MatherShip, ms *matherShip.MatherShip)  {
	if motherShips[ms.X] != nil {
		motherShips[ms.X][ms.Y] = ms
	} else {
		motherShips[ms.X] = make(map[int]*matherShip.MatherShip)
		motherShips[ms.X][ms.Y] = ms
	}
}
