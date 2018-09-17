package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../localGame/Phases"
	"../../../player"
	"strconv"
)

func GetMoveCoordinate(gameUnit *unit.Unit, client *player.Player, activeGame *localGame.Game, event string) map[string]map[string]*coordinate.Coordinate {

	start, _ := activeGame.Map.GetCoordinate(gameUnit.Q, gameUnit.R)

	openCoordinate := make(map[string]map[string]*coordinate.Coordinate)
	closeCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	startMatrix := generateNeighboursCoordinate(client, start, activeGame.Map, gameUnit, event) // берет все соседние клетки от старта

	for _, xLine := range startMatrix {
		for _, gameCoordinate := range xLine {
			_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R)
			if find {
				Phases.AddCoordinate(openCoordinate, gameCoordinate)
			}
		}
	}

	for i := 0; i < gameUnit.ActionPoints-1; i++ {
		for _, xLine := range openCoordinate {
			for _, gameCoordinate := range xLine {
				matrix := generateNeighboursCoordinate(client, gameCoordinate, activeGame.Map, gameUnit, event)
				for _, xLine := range matrix {
					for _, gameCoordinate := range xLine {
						_, ok := openCoordinate[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)]
						if !ok {
							_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R)
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
				_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R)
				if find {
					Phases.AddCoordinate(openCoordinate, gameCoordinate)
				}
			}
		}
	}

	return openCoordinate
}
