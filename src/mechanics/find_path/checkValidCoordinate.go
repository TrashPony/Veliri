package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *_map.Map, q int, r int) (*coordinate.Coordinate, bool) {

	gameCoordinate, ok := gameMap.GetCoordinate(q, r)

	if ok && gameCoordinate.Move {
		return gameCoordinate, true
	}
	return nil, false
}
