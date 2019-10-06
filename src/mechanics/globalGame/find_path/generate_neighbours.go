package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
)

func generateNeighboursCoordinate(curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit, scaleMap int,
	units map[int]*unit.ShortUnitInfo, xSize, ySize int, regions []*_map.Region, unitsID []int) (res map[string]map[string]*coordinate.Coordinate) {

	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)
	unitCollisionFind := false //todo ,jkmit ujdyjrjle ,jue ujdyjrjlf

	//строго лево
	leftCoordinate, left, unitCollision := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if left {
		coordinate.AddCoordinate(res, leftCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//строго право
	rightCoordinate, right, unitCollision := checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if right {
		coordinate.AddCoordinate(res, rightCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//верх центр
	topCoordinate, top, unitCollision := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if top {
		coordinate.AddCoordinate(res, topCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//низ центр
	bottomCoordinate, bottom, unitCollision := checkValidForMoveCoordinate(gameMap, curr.X, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if bottom {
		coordinate.AddCoordinate(res, bottomCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//верх лево
	gameCoordinate, find, unitCollision := checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//верх право
	gameCoordinate, find, unitCollision = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y-1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//низ лево
	gameCoordinate, find, unitCollision = checkValidForMoveCoordinate(gameMap, curr.X-1, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	//низ право
	gameCoordinate, find, unitCollision = checkValidForMoveCoordinate(gameMap, curr.X+1, curr.Y+1, xSize, ySize, gameUnit, scaleMap, regions, units, unitsID)
	if find {
		coordinate.AddCoordinate(res, gameCoordinate)
	}
	if unitCollision {
		unitCollisionFind = true
	}

	// если вокруг карента есть юнит то делаем дополнительную проверку на достежимость точек
	if unitCollisionFind && curr.State != 1 {
		for _, xLine := range res {
			for y, resCoordinate := range xLine {

				if debug.Store.UnitUnitCollision {
					debug.Store.AddMessage("CreateLine", "red", resCoordinate.X*scaleMap+scaleMap/2,
						resCoordinate.Y*scaleMap+scaleMap/2, curr.X*scaleMap+scaleMap/2, curr.Y*scaleMap+scaleMap/2,
						0, gameMap.Id, 20)
				}

				collision := collisions.SearchCollisionInLine(
					float64(resCoordinate.X*scaleMap+scaleMap/2),
					float64(resCoordinate.Y*scaleMap+scaleMap/2),
					float64(curr.X*scaleMap+scaleMap/2),
					float64(curr.Y*scaleMap+scaleMap/2),
					gameMap, gameUnit, 5, units, true, false, false, unitsID)

				if collision {
					delete(xLine, y)
				}
			}
		}
	}

	return
}
