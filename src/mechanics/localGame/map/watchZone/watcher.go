package watchZone

import (
	"strconv"
	"errors"
	"../../"
	"../../../gameObjects/matherShip"
	"../../../gameObjects/unit"
	"../coordinate"
)

type Watcher interface {
	GetX() int
	GetY() int
	GetWatchZone() int
	GetOwnerUser() string
}

func watch(gameObject Watcher, login string, game *localGame.Game) (allCoordinate map[string]*coordinate.Coordinate, unitsCoordinate map[int]map[int]*unit.Unit, matherShipCoordinate map[int]map[int]*matherShip.MatherShip, Err error) {

	allCoordinate = make(map[string]*coordinate.Coordinate)
	unitsCoordinate = make(map[int]map[int]*unit.Unit)
	matherShipCoordinate = make(map[int]map[int]*matherShip.MatherShip)

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
				var matherShipInMap *matherShip.MatherShip
				matherShipInMap, ok = game.GetMatherShip(gameCoordinate.X, gameCoordinate.Y)
				if ok {
					if matherShipCoordinate[gameCoordinate.X] != nil {
						matherShipCoordinate[gameCoordinate.X][gameCoordinate.Y] = matherShipInMap
					} else {
						matherShipCoordinate[gameCoordinate.X] = make(map[int]*matherShip.MatherShip)
						matherShipCoordinate[gameCoordinate.X][gameCoordinate.Y] = matherShipInMap
					}
				}
			}
		}
	} else {
		return allCoordinate, unitsCoordinate, matherShipCoordinate, errors.New("no owned")
	}
	return allCoordinate, unitsCoordinate, matherShipCoordinate, nil
}


