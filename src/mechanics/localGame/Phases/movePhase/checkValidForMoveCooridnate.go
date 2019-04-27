package movePhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *_map.Map, q int, r int, event string) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(q, r)

	myUnit, findUnit := client.GetUnit(q, r)
	_, findHostileUnit := client.GetHostileUnit(q, r)

	if ok && gameCoordinate.Move && !findUnit && !findHostileUnit {
		return gameCoordinate, true
	} else {
		// если грок направляется в трюм мса то немного переопределим фильтр что бы координата мс попала в матрицу
		if ok && gameCoordinate.Move && !findHostileUnit && event == "ToMC" && myUnit.Body.MotherShip {
			return gameCoordinate, true
		}
	}

	return nil, false
}
