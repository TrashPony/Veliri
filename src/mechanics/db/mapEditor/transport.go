package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceTransport(coordinate *coordinate.Coordinate, mp *_map.Map) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.Q, coordinate.R)

	coordinate.Transport = true

	if changeCoordinate != nil {
		UpdateMapCoordinate(coordinate, mp)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}

func RemoveTransport(coordinate *coordinate.Coordinate, mp *_map.Map) {
	coordinate.Transport = false
	UpdateMapCoordinate(coordinate, mp)
}
