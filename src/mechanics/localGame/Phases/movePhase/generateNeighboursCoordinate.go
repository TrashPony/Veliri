package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../player"
	"../../Phases"
)

func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map) (res map[string]map[string]*coordinate.Coordinate) {
	// берет все соседние клетки от текущей
	res = make(map[string]map[string]*coordinate.Coordinate)
	curr, ok := gameMap.GetCoordinate(curr.X, curr.Y)
	if !ok {
		return
	}
	//строго лево
	leftCoordinate, left := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y)
	if left && checkLevelCoordinate(curr, leftCoordinate) {
		Phases.AddCoordinate(res, leftCoordinate)
	}

	//строго право
	rightCoordinate, right := checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y)
	if right && checkLevelCoordinate(curr, rightCoordinate)  {
		Phases.AddCoordinate(res, rightCoordinate)
	}

	//верх центр
	topCoordinate, top := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y-1)
	if top && checkLevelCoordinate(curr, topCoordinate) {
		Phases.AddCoordinate(res, topCoordinate)
	}

	//низ центр
	bottomCoordinate, bottom := checkValidForMoveCoordinate(client, gameMap, curr.X, curr.Y+1)
	if bottom && checkLevelCoordinate(curr, bottomCoordinate) {
		Phases.AddCoordinate(res, bottomCoordinate)
	}

	//верх лево
	gameCoordinate, find := checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y-1)
	if find {
		checkEdgesCoordinate(curr, leftCoordinate, topCoordinate, gameCoordinate, res)
	}

	//верх право
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y-1)
	if find {
		checkEdgesCoordinate(curr, rightCoordinate, topCoordinate, gameCoordinate, res)
	}

	//низ лево
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X-1, curr.Y+1)
	if find {
		checkEdgesCoordinate(curr, leftCoordinate, bottomCoordinate, gameCoordinate, res)
	}

	//низ право
	gameCoordinate, find = checkValidForMoveCoordinate(client, gameMap, curr.X+1, curr.Y+1)
	if find {
		checkEdgesCoordinate(curr, rightCoordinate, bottomCoordinate, gameCoordinate, res)
	}

	return
}

func checkEdgesCoordinate(curr, edgCorOne, edgCorTwo, gameCoordinate *coordinate.Coordinate, res map[string]map[string]*coordinate.Coordinate) {
	if (edgCorOne != nil && edgCorOne.PassableEdges) && (edgCorTwo != nil && edgCorTwo.PassableEdges) {
		// смотрим можно ли пройти прилежашие координаты по углам
		if checkLevelCoordinate(curr, edgCorOne) && checkLevelCoordinate(curr, edgCorTwo){
			// сравниваем прилежащие координаты на предмет проходимости по высоте с текущей
			if checkLevelCoordinate(gameCoordinate, edgCorOne) && checkLevelCoordinate(gameCoordinate, edgCorTwo) {
				// сравниваем будующую координату на предмет проходимости по высоте с прилежащими координатами
				if checkLevelCoordinate(curr, gameCoordinate) {
					// сравниваем текущую координату на предмет проходимости по высоте с будущей координатами
					  // немного упоротый метод но работает отлично
					Phases.AddCoordinate(res, gameCoordinate)
				}
			}
		}
	}
}

func checkLevelCoordinate(one, two *coordinate.Coordinate) bool  {
	if one.Level > two.Level {
		diffLevel :=  one.Level - two.Level
		if diffLevel < 2 {
			return true
		} else {
			return false
		}
	} else {
		diffLevel :=  two.Level - one.Level
		if diffLevel < 2 {
			return true
		} else {
			return false
		}
	}
}
