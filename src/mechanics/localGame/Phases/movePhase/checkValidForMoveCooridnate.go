package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *_map.Map, q int, r int) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(q, r)

	_, findUnit := client.GetUnit(q, r)
	_, findHostileUnit := client.GetHostileUnit(q, r)

	if ok && !findUnit && !findHostileUnit {
		if gameCoordinate.Move {
			return gameCoordinate, true
		} else {
			return gameCoordinate, false
		}
	}

	return nil, false
}
