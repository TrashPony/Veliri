package coordinate

import (
	"strconv"
)

func AddCoordinate(res map[string]map[string]*Coordinate, gameCoordinate *Coordinate) {
	if res[strconv.Itoa(gameCoordinate.Q)] != nil {
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*Coordinate)
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	}
}

func AddXYCoordinate(res map[string]map[string]*Coordinate, gameCoordinate *Coordinate) {
	if res[strconv.Itoa(gameCoordinate.X)] != nil {
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*Coordinate)
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	}
}
