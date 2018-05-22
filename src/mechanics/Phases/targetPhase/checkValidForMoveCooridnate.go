package targetPhase

import (
	"../../coordinate"
	"../../gameMap"
	"../../player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *gameMap.Map, x int, y int) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(x, y)

	if ok {
		if gameCoordinate.Type != "obstacle" {
			return gameCoordinate, true
		}
	}

	return nil, false
}
