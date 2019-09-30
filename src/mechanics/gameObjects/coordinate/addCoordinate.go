package coordinate

import (
	"strconv"
)

func AddCoordinate(res map[string]map[string]*Coordinate, gameCoordinate *Coordinate) {
	if res[strconv.Itoa(gameCoordinate.X)] != nil {
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*Coordinate)
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	}
}

func AddIntCoordinate(res map[int]map[int]*Coordinate, gameCoordinate *Coordinate) {
	if res[gameCoordinate.X] != nil {
		res[gameCoordinate.X][gameCoordinate.Y] = gameCoordinate
	} else {
		res[gameCoordinate.X] = make(map[int]*Coordinate)
		res[gameCoordinate.X][gameCoordinate.Y] = gameCoordinate
	}
}
