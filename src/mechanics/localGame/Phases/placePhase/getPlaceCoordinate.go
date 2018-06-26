package placePhase

import (
	"../../../gameObjects/coordinate"
	"strconv"
	"../../../localGame"
)

func GetPlaceCoordinate(xCenter, yCenter, watchRadius int, actionGame *localGame.Game) (zone map[string]map[string]*coordinate.Coordinate) {

	tmpCoordinates := coordinate.GetCoordinatesRadius(xCenter, yCenter, watchRadius)

	for _, zoneCoordinate := range tmpCoordinates {
		gameCoordinate, find := actionGame.GetMap().GetCoordinate(zoneCoordinate.X, zoneCoordinate.Y)

		if find {
			if gameCoordinate.Type == "" {
				_, find = actionGame.GetUnit(zoneCoordinate.X, zoneCoordinate.Y)
				if !find {
					_, find = actionGame.GetMatherShip(zoneCoordinate.X, zoneCoordinate.Y)
					if !find {
						if zone != nil {
							if zone[strconv.Itoa(gameCoordinate.X)] != nil {
								zone[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
							} else {
								zone[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
								zone[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
							}
						} else {
							zone = make(map[string]map[string]*coordinate.Coordinate)
							zone[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
							zone[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
						}
					}
				}
			}
		}
	}
	return zone
}
