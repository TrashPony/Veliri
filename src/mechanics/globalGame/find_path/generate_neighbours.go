package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func generateNeighboursCoordinate(curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit, scaleMap int,
	units map[int]*unit.ShortUnitInfo, xSize, ySize int, regions []*_map.Region) (res map[string]map[string]*coordinate.Coordinate) {

	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)

	//строго лево
	leftCoordinate, left := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y, xSize, ySize, gameUnit, scaleMap, regions, units)
	if left {
		coordinate.AddCoordinate(res, leftCoordinate)
	}

	//строго право
	rightCoordinate, right := checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y, xSize, ySize, gameUnit, scaleMap, regions, units)
	if right {
		coordinate.AddCoordinate(res, rightCoordinate)
	}

	//верх центр
	topCoordinate, top := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if top {
		coordinate.AddCoordinate(res, topCoordinate)
	}

	//низ центр
	bottomCoordinate, bottom := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if bottom {
		coordinate.AddCoordinate(res, bottomCoordinate)
	}

	//верх лево
	gameCoordinate, find := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}

	//верх право
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}

	//низ лево
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}

	//низ право
	gameCoordinate, find = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}

	return
}
