package movePhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"strconv"
)

func GetMoveCoordinate(gameUnit *unit.Unit, client *player.Player, activeGame *localGame.Game, event string) map[string]map[string]*coordinate.Coordinate {

	start, _ := activeGame.Map.GetCoordinate(gameUnit.Q, gameUnit.R)

	openCoordinate := make(map[string]map[string]*coordinate.Coordinate)
	closeCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	startMatrix := generateNeighboursCoordinate(client, start, activeGame.Map, gameUnit, event) // берет все соседние клетки от старта

	for _, xLine := range startMatrix {
		for _, gameCoordinate := range xLine {
			_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R, event)
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
							_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R, event)
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
				_, find := checkValidForMoveCoordinate(client, activeGame.Map, gameCoordinate.Q, gameCoordinate.R, event)
				if find {
					Phases.AddCoordinate(openCoordinate, gameCoordinate)
				}
			}
		}
	}

	// удаляем все координаты вокруг МС т.к. нам доступно только войти в него а не стоять рядом
	if event == "ToMC" {
		msQ := client.GetSquad().MatherShip.Q
		msR := client.GetSquad().MatherShip.R

		//left
		delete(openCoordinate[strconv.Itoa(msQ-1)], strconv.Itoa(msR))
		//right
		delete(openCoordinate[strconv.Itoa(msQ+1)], strconv.Itoa(msR))

		if msR%2 != 0 {
			// topLeft
			delete(openCoordinate[strconv.Itoa(msQ)], strconv.Itoa(msR-1))
			// topRight
			delete(openCoordinate[strconv.Itoa(msQ+1)], strconv.Itoa(msR-1))
			// botLeft
			delete(openCoordinate[strconv.Itoa(msQ)], strconv.Itoa(msR+1))
			// botRight
			delete(openCoordinate[strconv.Itoa(msQ+1)], strconv.Itoa(msR+1))
		} else {
			// topLeft
			delete(openCoordinate[strconv.Itoa(msQ-1)], strconv.Itoa(msR-1))
			// topRight
			delete(openCoordinate[strconv.Itoa(msQ)], strconv.Itoa(msR-1))
			// botLeft
			delete(openCoordinate[strconv.Itoa(msQ-1)], strconv.Itoa(msR+1))
			// botRight
			delete(openCoordinate[strconv.Itoa(msQ)], strconv.Itoa(msR+1))
		}
	}

	return openCoordinate
}
