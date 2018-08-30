package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *_map.Map, x int, y int) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(x, y)

	_, findUnit := client.GetUnit(x, y)
	_, findHostileUnit := client.GetHostileUnit(x, y)

	if ok && !findUnit && !findHostileUnit {
		if gameCoordinate.Move {
			return gameCoordinate, true
		} else {
			return gameCoordinate, false
		}
	}

	return nil, false
}
