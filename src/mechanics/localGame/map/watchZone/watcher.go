package watchZone

import (
	"strconv"
	"errors"
	"../../../localGame"
	"../../../gameObjects/unit"
	"../../../gameObjects/coordinate"
)

type Watcher interface {
	GetQ() int
	GetR() int
	GetY() int
	GetWatchZone() int
	GetOwnerUser() string
}

func watch(gameObject Watcher, login string, game *localGame.Game) (allCoordinate map[string]*coordinate.Coordinate, unitsCoordinate map[int]map[int]*unit.Unit, Err error) {

	allCoordinate = make(map[string]*coordinate.Coordinate)
	unitsCoordinate = make(map[int]map[int]*unit.Unit)

	if login == gameObject.GetOwnerUser() {

		centerCoordinate, _ := game.Map.GetCoordinate(gameObject.GetQ(), gameObject.GetR())

		RadiusCoordinates := coordinate.GetCoordinatesRadius(centerCoordinate, gameObject.GetWatchZone())
		//PermCoordinates   := filter(gameObject, RadiusCoordinates, game)

		for _, gameCoordinate := range RadiusCoordinates{
			unitInMap, ok := game.GetUnit(gameCoordinate.Q,gameCoordinate.R)

			newCoordinate, find := game.Map.GetCoordinate(gameCoordinate.Q, gameCoordinate.R)
			if find {
				allCoordinate[strconv.Itoa(gameCoordinate.Q)+":"+strconv.Itoa(gameCoordinate.R)] = newCoordinate
			}

			if ok {
				if unitsCoordinate[gameCoordinate.Q] != nil {
					unitsCoordinate[gameCoordinate.Q][gameCoordinate.R] = unitInMap
				} else {
					unitsCoordinate[gameCoordinate.Q] = make(map[int]*unit.Unit)
					unitsCoordinate[gameCoordinate.Q][gameCoordinate.R] = unitInMap
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, nil
}


