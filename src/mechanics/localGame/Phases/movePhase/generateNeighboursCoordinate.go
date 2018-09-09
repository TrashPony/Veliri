package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../gameObjects/unit"
	"../../../player"
	"../../Phases"
)

// TODO возможно есть способ это все упаковать в минимальное количества кода т.к. он тут ппц повтяющиеся
func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit) (res map[string]map[string]*coordinate.Coordinate) {
	/*
	   соседи гексов беруться по разному в зависимости от четности строки
	   // even {Q,R}

	      {0,-1}  {+1,-1}
	   {-1,0} {0,0} {+1,0}
	      {0,+1}  {+1,+1}

	   // odd
	     {-1,-1}  {0,-1}
	   {-1,0} {0,0} {+1,0}
	     {-1,+1}  {0,+1}
	*/
	curr, find := gameMap.GetCoordinate(curr.Q, curr.R) // из алгоритмов иногда приходять координаты без высоты
	if !find {
		return
	}

	res = make(map[string]map[string]*coordinate.Coordinate)

	//left
	checkNeighbour(curr.Q-1, curr.R, client, curr, gameMap, gameUnit, res)
	//right
	checkNeighbour(curr.Q+1, curr.R, client, curr, gameMap, gameUnit, res)

	if curr.R%2 != 0 {
		// topLeft
		checkNeighbour(curr.Q, curr.R-1, client, curr, gameMap, gameUnit, res)
		// topRight
		checkNeighbour(curr.Q+1, curr.R-1, client, curr, gameMap, gameUnit, res)
		// botLeft
		checkNeighbour(curr.Q, curr.R+1, client, curr, gameMap, gameUnit, res)
		// botRight
		checkNeighbour(curr.Q+1, curr.R+1, client, curr, gameMap, gameUnit, res)
	} else {
		// topLeft
		checkNeighbour(curr.Q-1, curr.R-1, client, curr, gameMap, gameUnit, res)
		// topRight
		checkNeighbour(curr.Q, curr.R-1, client, curr, gameMap, gameUnit, res)
		// botLeft
		checkNeighbour(curr.Q-1, curr.R+1, client, curr, gameMap, gameUnit, res)
		// botRight
		checkNeighbour(curr.Q, curr.R+1, client, curr, gameMap, gameUnit, res)
	}

	return
}

func checkNeighbour(q, r int, client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map, gameUnit *unit.Unit,
	res map[string]map[string]*coordinate.Coordinate) {

	neighbour, find := checkValidForMoveCoordinate(client, gameMap, q, r)
	if find && checkLevelCoordinate(curr, neighbour) && checkMSPlace(client, neighbour) &&
		checkMSPatency(neighbour, gameUnit, client, gameMap) {
		Phases.AddCoordinate(res, neighbour)
	}
}

func checkMSPlace(client *player.Player, neighbour *coordinate.Coordinate) bool {
	for _, q := range client.GetUnits() {
		for _, myUnit := range q {
			if myUnit.Body.MotherShip {
				if checkMSCoordinate(myUnit, neighbour) {
					return false
				}
			}
		}
	}

	for _, q := range client.GetHostileUnits() {
		for _, hostileUnit := range q {
			if hostileUnit.Body.MotherShip {
				if checkMSCoordinate(hostileUnit, neighbour) {
					return false
				}
			}
		}
	}

	return true
}

func checkMSCoordinate(gameUnit *unit.Unit, neighbour *coordinate.Coordinate) bool {
	if gameUnit.Q-1 == neighbour.Q && gameUnit.R == neighbour.R {
		return true
	} // left
	if gameUnit.Q+1 == neighbour.Q && gameUnit.R == neighbour.R {
		return true
	} // right

	if gameUnit.R%2 != 0 {
		if gameUnit.Q == neighbour.Q && gameUnit.R-1 == neighbour.R {
			return true
		} // topLeft
		if gameUnit.Q+1 == neighbour.Q && gameUnit.R-1 == neighbour.R {
			return true
		} // topRight
		if gameUnit.Q == neighbour.Q && gameUnit.R+1 == neighbour.R {
			return true
		} // botLeft
		if gameUnit.Q+1 == neighbour.Q && gameUnit.R+1 == neighbour.R {
			return true
		} // botRight
	} else {
		if gameUnit.Q-1 == neighbour.Q && gameUnit.R-1 == neighbour.R {
			return true
		} // topLeft
		if gameUnit.Q == neighbour.Q && gameUnit.R-1 == neighbour.R {
			return true
		} // topRight
		if gameUnit.Q-1 == neighbour.Q && gameUnit.R+1 == neighbour.R {
			return true
		} // botLeft
		if gameUnit.Q == neighbour.Q && gameUnit.R+1 == neighbour.R {
			return true
		} // botRight
	}

	return false
}

func checkMSPatency(curr *coordinate.Coordinate, gameUnit *unit.Unit, client *player.Player, gameMap *_map.Map) bool {
	if gameUnit.Body.MotherShip {

		var left, right, topLeft, topRight, botLeft, botRight bool

		left = checkMsCoordinate(curr, curr.Q-1, curr.R, gameUnit, client, gameMap)
		right = checkMsCoordinate(curr, curr.Q+1, curr.R, gameUnit, client, gameMap)

		if curr.R%2 != 0 {
			topLeft = checkMsCoordinate(curr, curr.Q, curr.R-1, gameUnit, client, gameMap)
			topRight = checkMsCoordinate(curr, curr.Q+1, curr.R-1, gameUnit, client, gameMap)
			botLeft = checkMsCoordinate(curr, curr.Q, curr.R+1, gameUnit, client, gameMap)
			botRight = checkMsCoordinate(curr, curr.Q+1, curr.R+1, gameUnit, client, gameMap)
		} else {
			topLeft = checkMsCoordinate(curr, curr.Q-1, curr.R-1, gameUnit, client, gameMap)
			topRight = checkMsCoordinate(curr, curr.Q, curr.R-1, gameUnit, client, gameMap)
			botLeft = checkMsCoordinate(curr, curr.Q-1, curr.R+1, gameUnit, client, gameMap)
			botRight = checkMsCoordinate(curr, curr.Q, curr.R+1, gameUnit, client, gameMap)
		}

		return left && right && topLeft && topRight && botLeft && botRight
	} else {
		return true
	}
}

func checkMsCoordinate(curr *coordinate.Coordinate, q, r int, gameUnit *unit.Unit, client *player.Player, gameMap *_map.Map) bool {
	if !(gameUnit.Q == q && gameUnit.R == r) {
		neighbours, pass := checkValidForMoveCoordinate(client, gameMap, q, r)
		if pass {
			return checkLevelCoordinate(curr, neighbours)
		} else {
			return false
		}
	} else {
		return true
	}
}

func checkLevelCoordinate(one, two *coordinate.Coordinate) bool {
	if one.Level > two.Level {
		diffLevel := one.Level - two.Level
		if diffLevel < 2 {
			return true
		} else {
			return false
		}
	} else {
		diffLevel := two.Level - one.Level
		if diffLevel < 2 {
			return true
		} else {
			return false
		}
	}
}
