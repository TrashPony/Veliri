package game

import (
	"strconv"
	"errors"
)

type Watcher interface {
	getX() int
	getY() int
	getWatchZone() int
	getOwnerUser() string
}

func Watch(gameObject Watcher, login string, game *Game) (allCoordinate map[string]*Coordinate, unitsCoordinate map[int]map[int]*Unit, matherShipCoordinate map[int]map[int]*MatherShip, Err error) {

	allCoordinate = make(map[string]*Coordinate)
	unitsCoordinate = make(map[int]map[int]*Unit)
	matherShipCoordinate = make(map[int]map[int]*MatherShip)

	if login == gameObject.getOwnerUser() {

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
				var matherShipInMap *MatherShip
				matherShipInMap, ok = game.GetMatherShip(coordinate.X, coordinate.Y)
				if ok {
					if matherShipCoordinate[coordinate.X] != nil {
						matherShipCoordinate[coordinate.X][coordinate.Y] = matherShipInMap
					} else {
						matherShipCoordinate[coordinate.X] = make(map[int]*MatherShip)
						matherShipCoordinate[coordinate.X][coordinate.Y] = matherShipInMap
					}
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, matherShipCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, matherShipCoordinate, nil
}


