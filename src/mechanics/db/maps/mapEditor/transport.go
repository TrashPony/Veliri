package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceTransport(coordinate *coordinate.Coordinate, mp *_map.Map) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.X, coordinate.Y)

	coordinate.Transport = true

	if changeCoordinate != nil {
		UpdateMapCoordinate(coordinate, mp, coordinate.X, coordinate.Y)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}

func RemoveTransport(coordinate *coordinate.Coordinate, mp *_map.Map) {
	coordinate.Transport = false
	UpdateMapCoordinate(coordinate, mp, coordinate.X, coordinate.Y)
}
