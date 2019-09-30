package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceHandler(coordinate *coordinate.Coordinate, mp *_map.Map, toX, toY, toBaseId, toMapId int, typeHandler string) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.X, coordinate.Y)

	coordinate.Handler = typeHandler
	// TODO
	//coordinate.ToQ = toQ
	//coordinate.ToR = toR
	coordinate.ToBaseID = toBaseId
	coordinate.ToMapID = toMapId

	if changeCoordinate != nil {
		UpdateMapCoordinate(coordinate, mp)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}

func RemoveHandler(coordinate *coordinate.Coordinate, mp *_map.Map) {
	coordinate.Handler = ""
	UpdateMapCoordinate(coordinate, mp)
}
