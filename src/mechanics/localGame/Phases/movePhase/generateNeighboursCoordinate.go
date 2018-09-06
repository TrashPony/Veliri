package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../player"
	"../../Phases"
	"../../../gameObjects/unit"
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

	neighboursLeft, left := checkValidForMoveCoordinate(client, gameMap, curr.Q-1, curr.R)
	if left && checkLevelCoordinate(curr, neighboursLeft) && checkMSPatency(neighboursLeft, gameUnit, client, gameMap) {
		Phases.AddCoordinate(res, neighboursLeft)
	}

	NeighboursRight, right := checkValidForMoveCoordinate(client, gameMap, curr.Q+1, curr.R)
	if right && checkLevelCoordinate(curr, NeighboursRight) && checkMSPatency(NeighboursRight, gameUnit, client, gameMap) {
		Phases.AddCoordinate(res, NeighboursRight)
	}

	if curr.R%2 != 0 {
		NeighboursTopLeft, topLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R-1)
		if topLeft && checkLevelCoordinate(curr, NeighboursTopLeft) && checkMSPatency(NeighboursTopLeft, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursTopLeft)
		}

		NeighboursTopRight, topRight := checkValidForMoveCoordinate(client, gameMap, curr.Q+1, curr.R-1)
		if topRight && checkLevelCoordinate(curr, NeighboursTopRight) && checkMSPatency(NeighboursTopRight, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursTopRight)
		}

		NeighboursBotLeft, botLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R+1)
		if botLeft && checkLevelCoordinate(curr, NeighboursBotLeft) && checkMSPatency(NeighboursBotLeft, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursBotLeft)
		}

		NeighboursBotRight, botRight := checkValidForMoveCoordinate(client, gameMap, curr.Q+1, curr.R+1)
		if botRight && checkLevelCoordinate(curr, NeighboursBotRight) && checkMSPatency(NeighboursBotRight, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursBotRight)
		}
	} else {
		NeighboursTopLeft, topLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q-1, curr.R-1)
		if topLeft && checkLevelCoordinate(curr, NeighboursTopLeft) && checkMSPatency(NeighboursTopLeft, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursTopLeft)
		}

		NeighboursTopRight, topRight := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R-1)
		if topRight && checkLevelCoordinate(curr, NeighboursTopRight) && checkMSPatency(NeighboursTopRight, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursTopRight)
		}

		NeighboursBotLeft, botLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q-1, curr.R+1)
		if botLeft && checkLevelCoordinate(curr, NeighboursBotLeft) && checkMSPatency(NeighboursBotLeft, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursBotLeft)
		}
		NeighboursBotRight, botRight := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R+1)
		if botRight && checkLevelCoordinate(curr, NeighboursBotRight) && checkMSPatency(NeighboursBotRight, gameUnit, client, gameMap) {
			Phases.AddCoordinate(res, NeighboursBotRight)
		}
	}

	return
}

// TODO возможно есть способ это все упаковать в минимальное количества кода т.к. он тут ппц повтяющиеся
func checkMSPatency(curr *coordinate.Coordinate, gameUnit *unit.Unit, client *player.Player, gameMap *_map.Map) bool {
	if gameUnit.Body.MotherShip && curr.Q == 9 && curr.R == 7{

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
	if gameUnit.Q != q && gameUnit.R != r {
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
