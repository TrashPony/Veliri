package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../player"
	"../../Phases"
)

func generateNeighboursCoordinate(client *player.Player, curr *coordinate.Coordinate, gameMap *_map.Map) (res map[string]map[string]*coordinate.Coordinate) {
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
	curr, find := gameMap.GetCoordinate(curr.Q, curr.R)// из алгоритмов иногда приходять координаты без высоты
	if !find {
		return
	}

	res = make(map[string]map[string]*coordinate.Coordinate)

	neighboursLeft, left := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R)
	if left && checkLevelCoordinate(curr, neighboursLeft) {
		Phases.AddCoordinate(res, neighboursLeft)
	}

	NeighboursRight, right := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R)
	if right && checkLevelCoordinate(curr, NeighboursRight) {
		Phases.AddCoordinate(res, NeighboursRight)
	}

	if curr.R % 2 != 0 {
		NeighboursTopLeft, topLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R - 1)
		if topLeft && checkLevelCoordinate(curr, NeighboursTopLeft) {
			Phases.AddCoordinate(res, NeighboursTopLeft)
		}

		NeighboursTopRight, topRight := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R - 1)
		if topRight && checkLevelCoordinate(curr, NeighboursTopRight) {
			Phases.AddCoordinate(res, NeighboursTopRight)
		}

		NeighboursBotLeft, botLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R + 1)
		if botLeft && checkLevelCoordinate(curr, NeighboursBotLeft) {
			Phases.AddCoordinate(res, NeighboursBotLeft)
		}

		NeighboursBotRight, botRight := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R + 1)
		if botRight && checkLevelCoordinate(curr, NeighboursBotRight) {
			Phases.AddCoordinate(res, NeighboursBotRight)
		}
	} else {
		NeighboursTopLeft, topLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R - 1)
		if topLeft && checkLevelCoordinate(curr, NeighboursTopLeft) {
			Phases.AddCoordinate(res, NeighboursTopLeft)
		}

		NeighboursTopRight, topRight := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R - 1)
		if topRight && checkLevelCoordinate(curr, NeighboursTopRight) {
			Phases.AddCoordinate(res, NeighboursTopRight)
		}

		NeighboursBotLeft, botLeft := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R + 1)
		if botLeft && checkLevelCoordinate(curr, NeighboursBotLeft) {
			Phases.AddCoordinate(res, NeighboursBotLeft)
		}

		NeighboursBotRight, botRight := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R + 1)
		if botRight && checkLevelCoordinate(curr, NeighboursBotRight) {
			Phases.AddCoordinate(res, NeighboursBotRight)
		}
	}

	return
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
