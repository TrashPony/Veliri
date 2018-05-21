package movePhase

import (
	"../../coordinate"
	"../../unit"
	"../../player"
	"../../game"
	"strconv"
)

func GetMoveCoordinate(gameUnit *unit.Unit, client *player.Player, activeGame *game.Game) (res map[string]map[string]*coordinate.Coordinate) {

	start, _ := activeGame.Map.GetCoordinate(gameUnit.X, gameUnit.Y)
	obstacles := GetObstacles(client, activeGame)

	openCoordinate := make(map[int]map[int]*coordinate.Coordinate)
	closeCoordinate := make(map[int]map[int]*coordinate.Coordinate)

	startMatrix := generateNeighboursCoordinate(start, obstacles) // берет все соседние клетки от старта

	for _, xLine := range startMatrix {
		for _, gameCoordinate := range xLine {
			addCoordinateIfValid(openCoordinate, obstacles, gameCoordinate.X, gameCoordinate.Y)
		}
	}

	for i := 0; i < gameUnit.MoveSpeed-1; i++ {
		for _, xLine := range openCoordinate {
			for _, gameCoordinate := range xLine {
				matrix := generateNeighboursCoordinate(gameCoordinate, obstacles)
				for _, xLine := range matrix {
					for _, gameCoordinate := range xLine {
						_, ok := openCoordinate[gameCoordinate.X][gameCoordinate.Y]
						if !ok {
							addCoordinateIfValid(closeCoordinate, obstacles, gameCoordinate.X, gameCoordinate.Y)
						}
					}
				}
			}
		}

		for _, xLine := range closeCoordinate {
			for _, gameCoordinate := range xLine {
				addCoordinateIfValid(openCoordinate, obstacles, gameCoordinate.X, gameCoordinate.Y)
			}
		}
	}

	for _, xLine := range openCoordinate{
		for _, gameCoordinate := range xLine {
			if res != nil {
				if res[strconv.Itoa(gameCoordinate.X)] != nil {
					res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
				} else {
					res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
					res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
				}
			} else {
				res = make(map[string]map[string]*coordinate.Coordinate)
				res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
				res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
			}
		}
	}

	return res
}
