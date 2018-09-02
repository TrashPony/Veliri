package placePhase

import (
	"../../../gameObjects/coordinate"
	"../../../localGame"
	"strconv"
)

func GetPlaceCoordinate(qCenter, rCenter, watchRadius int, actionGame *localGame.Game) (zone map[string]map[string]*coordinate.Coordinate) {

	placeCoordinate, find := actionGame.Map.GetCoordinate(qCenter, rCenter)
	if find {
		tmpCoordinates := coordinate.GetCoordinatesRadius(placeCoordinate, watchRadius)

		for _, zoneCoordinate := range tmpCoordinates {
			gameCoordinate, find := actionGame.GetMap().GetCoordinate(zoneCoordinate.Q, zoneCoordinate.R)

			if find && gameCoordinate.Type == "" && gameCoordinate.Move {
				_, find = actionGame.GetUnit(zoneCoordinate.Q, zoneCoordinate.R)
				if !find {
					if zone != nil {
						if zone[strconv.Itoa(gameCoordinate.Q)] != nil {
							zone[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
						} else {
							zone[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*coordinate.Coordinate)
							zone[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
						}
					} else {
						zone = make(map[string]map[string]*coordinate.Coordinate)
						zone[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*coordinate.Coordinate)
						zone[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
					}
				}
			}
		}
	}
	return zone
}
