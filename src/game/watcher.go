package game

import (
	"strconv"
	"errors"
)

type Watcher interface {
	getX() int
	getY() int
	getWatchZone() int
	getNameUser() string
}

func Watch(gameObject Watcher, login string, game *Game) (allCoordinate map[string]*Coordinate, unitsCoordinate map[int]map[int]*Unit, structureCoordinate map[int]map[int]*Structure, Err error) {

	allCoordinate = make(map[string]*Coordinate)
	unitsCoordinate = make(map[int]map[int]*Unit)
	structureCoordinate = make(map[int]map[int]*Structure)

	if login == gameObject.getNameUser() {

		RadiusCoordinates := GetCoordinates(gameObject.getX(), gameObject.getY(), gameObject.getWatchZone())
		PermCoordinates   := Filter(gameObject, RadiusCoordinates, game)

		for _, coordinate := range PermCoordinates{
			unitInMap, ok := game.GetUnit(coordinate.X,coordinate.Y)

			allCoordinate[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)] = coordinate

			if ok {
				if unitsCoordinate[coordinate.X] != nil {
					unitsCoordinate[coordinate.X][coordinate.Y] = unitInMap
				} else {
					unitsCoordinate[coordinate.X] = make(map[int]*Unit)
					unitsCoordinate[coordinate.X][coordinate.Y] = unitInMap
				}
			} else {
				var structureInMap *Structure
				structureInMap, ok = game.GetStructure(coordinate.X, coordinate.Y)
				if ok {
					if structureCoordinate[coordinate.X] != nil {
						structureCoordinate[coordinate.X][coordinate.Y] = structureInMap
					} else {
						structureCoordinate[coordinate.X] = make(map[int]*Structure)
						structureCoordinate[coordinate.X][coordinate.Y] = structureInMap
					}
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, structureCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, structureCoordinate, nil
}


