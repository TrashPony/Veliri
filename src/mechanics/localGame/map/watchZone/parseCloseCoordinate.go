package watchZone

import (
	"../coordinate"
	"../../../unit"
	"../../../matherShip"
	"../../"
)

func parseCloseCoordinate(closeCoordinates []*coordinate.Coordinate, closeUnit []*unit.Unit, closeMatherShip []*matherShip.MatherShip, game *localGame.Game) ([]*coordinate.Coordinate) {

	for _, closeUnit := range closeUnit {
		closeCoordinate, find := game.GetMap().GetCoordinate(closeUnit.X, closeUnit.Y)
		if find {
			closeCoordinates = append(closeCoordinates, closeCoordinate)
		}
	}

	for _, closeMatherShip := range closeMatherShip {
		closeCoordinate, find := game.GetMap().GetCoordinate(closeMatherShip.X, closeMatherShip.Y)
		if find {
			closeCoordinates = append(closeCoordinates, closeCoordinate)
		}
	}

	return closeCoordinates
}
