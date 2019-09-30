package mapEditor

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

func PlaceTextures(coordinate *coordinate.Coordinate, mp *_map.Map, textureName string) {
	changeCoordinate := getMapCoordinateInMC(mp.Id, coordinate.X, coordinate.Y)

	coordinate.TextureOverFlore = textureName

	coordinate.TexturePriority = mp.GetMaxPriorityTexture()
	coordinate.TexturePriority++

	if changeCoordinate != nil {
		UpdateMapCoordinate(coordinate, mp)
	} else {
		InsertMapCoordinate(coordinate, mp)
	}
}

func RemoveTextures(coordinate *coordinate.Coordinate, mp *_map.Map) {
	coordinate.TextureOverFlore = ""
	UpdateMapCoordinate(coordinate, mp)
}
