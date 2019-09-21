package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func generateNeighboursCoordinate(curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit, scaleMap int,
	allUnits map[int]*unit.ShortUnitInfo) (res map[string]map[string]*coordinate.Coordinate) {

	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)

	//строго лево
	leftCoordinate, left := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if left {
		coordinate.AddXYCoordinate(res, leftCoordinate)
	}

	//строго право
	rightCoordinate, right := checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if right {
		coordinate.AddXYCoordinate(res, rightCoordinate)
	}

	//верх центр
	topCoordinate, top := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y-1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if top {
		coordinate.AddXYCoordinate(res, topCoordinate)
	}

	//низ центр
	bottomCoordinate, bottom := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y+1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if bottom {
		coordinate.AddXYCoordinate(res, bottomCoordinate)
	}

	//верх лево
	gameCoordinate, find := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y-1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//верх право
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y-1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//низ лево
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y+1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	//низ право
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y+1, curr.X, curr.Y, curr.Rotate, gameUnit, scaleMap, allUnits)
	if find {
		coordinate.AddXYCoordinate(res, gameCoordinate)
	}

	return
}
