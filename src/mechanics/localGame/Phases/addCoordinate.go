package Phases

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"strconv"
)

func AddCoordinate(res map[string]map[string]*coordinate.Coordinate, gameCoordinate *coordinate.Coordinate) {
	if res[strconv.Itoa(gameCoordinate.Q)] != nil {
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*coordinate.Coordinate)
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	}
}
