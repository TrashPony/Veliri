package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func ChangeHeightCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.Q, coordinate.R)

	if changeCoordinate != nil {
		updateChangeHeight(coordinate, mp, change)
	} else {
		insertNewHeightCoordinate(coordinate, mp, change)
	}
}

func insertNewHeightCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {
	newLvl := mp.DefaultLevel + change
	if newLvl > 5 {
		newLvl = 5
	}
	if newLvl < 1 {
		newLvl = 1
	}
	coordinate.Level = newLvl
	InsertMapCoordinate(coordinate, mp)
}

func updateChangeHeight(coordinate *coordinate.Coordinate, mp *_map.Map, change int) {
	newLvl := coordinate.Level + change

	if newLvl > 5 {
		newLvl = 5
	}

	if newLvl < 1 {
		newLvl = 1
	}
	coordinate.Level = newLvl
	UpdateMapCoordinate(coordinate, mp)
}
