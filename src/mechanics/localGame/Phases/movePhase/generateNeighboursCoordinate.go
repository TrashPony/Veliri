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
			// even
			   {-1,0}  {-1,+1}
			{0,-1} {0,0} {0,+1}
			   {+1,0}  {+1,+1}

			// odd
			  {-1,-1}  {-1,0}
			{0,-1} {0,0} {0,+1}
			  {-1,+1}  {+1,0}
	*/
	res = make(map[string]map[string]*coordinate.Coordinate)

	NeighboursOne, oneOk := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R)
	if oneOk && checkLevelCoordinate(curr, NeighboursOne) {
		Phases.AddCoordinate(res, NeighboursOne)
	}

	NeighboursTwo, twoOk := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R - 1)
	if twoOk && checkLevelCoordinate(curr, NeighboursTwo) {
		Phases.AddCoordinate(res, NeighboursTwo)
	}

	NeighboursThree, threeOk := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R)
	if threeOk && checkLevelCoordinate(curr, NeighboursThree) {
		Phases.AddCoordinate(res, NeighboursThree)
	}

	NeighboursFour, fourOk := checkValidForMoveCoordinate(client, gameMap, curr.Q, curr.R + 1)
	if fourOk && checkLevelCoordinate(curr, NeighboursFour) {
		Phases.AddCoordinate(res, NeighboursFour)
	}

	if curr.R % 2 != 0 {
		NeighboursFive, fiveOk := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R - 1)
		if fiveOk && checkLevelCoordinate(curr, NeighboursFive) {
			Phases.AddCoordinate(res, NeighboursFive)
		}

		NeighboursSix, sixOk := checkValidForMoveCoordinate(client, gameMap, curr.Q + 1, curr.R + 1)
		if sixOk && checkLevelCoordinate(curr, NeighboursSix) {
			Phases.AddCoordinate(res, NeighboursSix)
		}
	} else {
		NeighboursFive, fiveOk := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R - 1)
		if fiveOk && checkLevelCoordinate(curr, NeighboursFive) {
			Phases.AddCoordinate(res, NeighboursFive)
		}

		NeighboursSix, sixOk := checkValidForMoveCoordinate(client, gameMap, curr.Q - 1, curr.R + 1)
		if sixOk && checkLevelCoordinate(curr, NeighboursSix) {
			Phases.AddCoordinate(res, NeighboursSix)
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
