package objects

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

func Watch(gameObject Watcher, login string, units map[int]map[int]*Unit, allStructures map[int]map[int]*Structure) (allCoordinate map[string]*Coordinate, unitsCoordinate map[int]map[int]*Unit, structureCoordinate map[int]map[int]*Structure, Err error) {
	allCoordinate = make(map[string]*Coordinate)
	unitsCoordinate = make(map[int]map[int]*Unit)
	structureCoordinate = make(map[int]map[int]*Structure)

	if login == gameObject.getNameUser() {
		PermCoordinates := GetCoordinates(gameObject.getX(), gameObject.getY(), gameObject.getWatchZone())
		for i := 0; i < len(PermCoordinates); i++ {
			unitInMap, ok := units[PermCoordinates[i].X][PermCoordinates[i].Y]
			if ok {

				if unitsCoordinate[PermCoordinates[i].X] != nil {
					unitsCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				} else {
					unitsCoordinate[PermCoordinates[i].X] = make(map[int]*Unit)
					unitsCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = unitInMap
				}
			} else {
				var structureInMap *Structure
				structureInMap, ok = allStructures[PermCoordinates[i].X][PermCoordinates[i].Y]
				if ok {
					if structureCoordinate[PermCoordinates[i].X] != nil {
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					} else {
						structureCoordinate[PermCoordinates[i].X] = make(map[int]*Structure)
						structureCoordinate[PermCoordinates[i].X][PermCoordinates[i].Y] = structureInMap
					}
				} else {
					allCoordinate[strconv.Itoa(PermCoordinates[i].X)+":"+strconv.Itoa(PermCoordinates[i].Y)] = PermCoordinates[i]
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, structureCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, structureCoordinate, nil
}


