package watchZone

import (
	"strconv"
	"errors"
	"../../../localGame"
	"../../../gameObjects/unit"
	"../../../gameObjects/coordinate"
)

type Watcher interface {
	GetX() int
	GetY() int
	GetWatchZone() int
	GetOwnerUser() string
}

func watch(gameObject Watcher, login string, game *localGame.Game) (allCoordinate map[string]*coordinate.Coordinate, unitsCoordinate map[int]map[int]*unit.Unit, Err error) {

	allCoordinate = make(map[string]*coordinate.Coordinate)
	unitsCoordinate = make(map[int]map[int]*unit.Unit)

	if login == gameObject.GetOwnerUser() {

		RadiusCoordinates := coordinate.GetCoordinatesRadius(gameObject.GetX(), gameObject.GetY(), gameObject.GetWatchZone())
		PermCoordinates   := filter(gameObject, RadiusCoordinates, game)

		for _, gameCoordinate := range PermCoordinates{
			unitInMap, ok := game.GetUnit(gameCoordinate.X,gameCoordinate.Y)

			newCoordinate, find := game.Map.GetCoordinate(gameCoordinate.X, gameCoordinate.Y)
			if find { // TODO костыль // TODO проеб сылок координата gameCoordinate не так что у игры >_<
				allCoordinate[strconv.Itoa(gameCoordinate.X)+":"+strconv.Itoa(gameCoordinate.Y)] = newCoordinate
			}

			if ok {
				if unitsCoordinate[gameCoordinate.X] != nil {
					unitsCoordinate[gameCoordinate.X][gameCoordinate.Y] = unitInMap
				} else {
					unitsCoordinate[gameCoordinate.X] = make(map[int]*unit.Unit)
					unitsCoordinate[gameCoordinate.X][gameCoordinate.Y] = unitInMap
				}
			} else {

			}
		}
	} else {
		return allCoordinate, unitsCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, nil
}


