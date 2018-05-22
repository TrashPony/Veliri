package targetPhase

import (
	"../../coordinate"
	"../../player"
	"../../unit"
	"../../game"
	"../../Phases"
	"strconv"
)

func GetTargetCoordinate(gameUnit *unit.Unit, client *player.Player, activeGame *game.Game) map[string]map[string]*coordinate.Coordinate {

	start, _ := activeGame.Map.GetCoordinate(gameUnit.X, gameUnit.Y)

	openCoordinate := make(map[string]map[string]*coordinate.Coordinate)
	closeCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	startMatrix := generateNeighboursCoordinate(client, start, activeGame.Map) // берет все соседние клетки от старта

	for _, xLine := range startMatrix {
		for _, gameCoordinate := range xLine {
			_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.X, gameCoordinate.Y)
			if find {
				Phases.AddCoordinate(openCoordinate, gameCoordinate)
			}
		}
	}

	for i := 0; i < gameUnit.MoveSpeed-1; i++ {
		for _, xLine := range openCoordinate {
			for _, gameCoordinate := range xLine {
				matrix := generateNeighboursCoordinate(client, gameCoordinate, activeGame.Map)
				for _, xLine := range matrix {
					for _, gameCoordinate := range xLine {
						_, ok := openCoordinate[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)]
						if !ok {
							_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.X, gameCoordinate.Y)
							if find {
								Phases.AddCoordinate(closeCoordinate, gameCoordinate)
							}
						}
					}
				}
			}
		}

		for _, xLine := range closeCoordinate {
			for _, gameCoordinate := range xLine {
				_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.X, gameCoordinate.Y)
				if find {
					Phases.AddCoordinate(openCoordinate, gameCoordinate)
				}
			}
		}
	}

	return openCoordinate
}
