package watchZone

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
)

func parseCloseCoordinate(closeCoordinates []*coordinate.Coordinate, closeUnit []*unit.Unit, game *localGame.Game) []*coordinate.Coordinate {

	for _, closeUnit := range closeUnit {
		closeCoordinate, find := game.GetMap().GetCoordinate(closeUnit.Q, closeUnit.R)
		if find {
			closeCoordinates = append(closeCoordinates, closeCoordinate)
		}
	}

	return closeCoordinates
}
