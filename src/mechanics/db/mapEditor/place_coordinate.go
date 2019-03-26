package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceCoordinate(coordinate *coordinate.Coordinate, mp *_map.Map, newIDType int) {

	oldType := getMapCoordinateInMC(mp.Id, coordinate.Q, coordinate.R)

	newType := getTypeByID(newIDType)

	coordinate.ID = newIDType
	coordinate.Type = newType.Type
	coordinate.TextureFlore = newType.TextureFlore
	coordinate.TextureObject = newType.TextureObject
	coordinate.AnimateSpriteSheets = newType.AnimateSpriteSheets
	coordinate.AnimateLoop = newType.AnimateLoop
	coordinate.UnitOverlap = newType.UnitOverlap

	coordinate.Move = newType.Move
	coordinate.View = newType.View
	coordinate.Attack = newType.Attack

	coordinate.Scale = 100
	coordinate.Shadow = true
	coordinate.XShadowOffset = 10
	coordinate.YShadowOffset = 10

	if oldType != nil {
		UpdateMapCoordinate(coordinate, mp)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}