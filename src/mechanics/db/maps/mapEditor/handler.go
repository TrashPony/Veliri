package mapEditor

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceHandler(coordinate *coordinate.Coordinate, mp *_map.Map, pos string, toBaseId, toMapId int, typeHandler string) {

	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.X, coordinate.Y)

	coordinate.Handler = typeHandler

	err := json.Unmarshal([]byte(pos), &coordinate.Positions)
	if err != nil {
		println(err.Error())
	}

	coordinate.ToBaseID = toBaseId
	coordinate.ToMapID = toMapId

	if changeCoordinate != nil {
		UpdateMapCoordinate(coordinate, mp, coordinate.X, coordinate.Y)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}

func RemoveHandler(coordinate *coordinate.Coordinate, mp *_map.Map) {
	coordinate.Handler = ""
	UpdateMapCoordinate(coordinate, mp, coordinate.X, coordinate.Y)
}
