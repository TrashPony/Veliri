package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit, scaleMap int) (res map[string]map[string]*coordinate.Coordinate) {
	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)

	//строго лево
	leftCoordinate, left := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y, gameUnit, scaleMap)
	if left {
		coordinate.AddXYCoordinate(res, leftCoordinate)
	}

	//строго право
	rightCoordinate, right := checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y, gameUnit, scaleMap)
	if right {
		coordinate.AddXYCoordinate(res, rightCoordinate)
	}

	//верх центр
	topCoordinate, top := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y-1, gameUnit, scaleMap)
	if top {
		coordinate.AddXYCoordinate(res, topCoordinate)
	}

	//низ центр
	bottomCoordinate, bottom := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y+1, gameUnit, scaleMap)
	if bottom {
		coordinate.AddXYCoordinate(res, bottomCoordinate)
	}

	//верх лево
	gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y-1, gameUnit, scaleMap)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//верх право
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y-1, gameUnit, scaleMap)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//низ лево
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y+1, gameUnit, scaleMap)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//низ право
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y+1, gameUnit, scaleMap)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	return
}
