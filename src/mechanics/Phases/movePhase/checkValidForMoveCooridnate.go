package movePhase

import (
	"../../coordinate"
	"../../gameMap"
	"../../player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *gameMap.Map, x int, y int) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(x, y)

	_, findUnit := client.GetUnit(x, y)
	_, findHostileUnit := client.GetHostileUnit(x, y)
	_, findMSHostile := client.GetHostileMatherShip(x, y)

	if ok && !findUnit && !findHostileUnit && !findMSHostile{
		if !(x == client.GetMatherShip().X && y == client.GetMatherShip().Y) {
			if gameCoordinate.Type != "obstacle" && gameCoordinate.Type != "terrain" {
				return gameCoordinate, true
			}
		}
	}

	return nil, false
}