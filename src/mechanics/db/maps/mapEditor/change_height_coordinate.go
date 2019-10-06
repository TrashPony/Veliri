package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func ChangeHeightCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.X, coordinate.Y)

	if changeCoordinate != nil {
		updateChangeHeight(coordinate, mp, change)
	} else {
		insertNewHeightCoordinate(coordinate, mp, change)
	}
}

func insertNewHeightCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {
	InsertMapCoordinate(coordinate, mp)
}

func updateChangeHeight(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {
	UpdateMapCoordinate(coordinate, mp)
}
