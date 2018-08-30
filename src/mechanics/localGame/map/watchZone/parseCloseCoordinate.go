package watchZone

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
)

func parseCloseCoordinate(closeCoordinates []*coordinate.Coordinate, closeUnit []*unit.Unit, game *localGame.Game) ([]*coordinate.Coordinate) {

	for _, closeUnit := range closeUnit {
		closeCoordinate, find := game.GetMap().GetCoordinate(closeUnit.X, closeUnit.Y)
		if find {
			closeCoordinates = append(closeCoordinates, closeCoordinate)
		}
	}

	return closeCoordinates
}
