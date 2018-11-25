package mapEditor

import "../../gameObjects/coordinate"

func PlaceTerrain(idMap, idType, q, r int) {

	newTerrain := getTypeByID(idType)            // отсюда берем терейн
	oldType := getMapCoordinateInMC(idMap, q, r) // отсюда обьект если он есть

	var newType *coordinate.Coordinate

	if oldType != nil {
		newType = getTypeByTerrainAndObject(newTerrain.TextureFlore, oldType.TextureObject)
	} else {
		newType = getTypeByTerrainAndObject(newTerrain.TextureFlore, "")
	}

	if newType != nil {
		PlaceCoordinate(idMap, newType.ID, q, r)
	} else {
		var newId int
		if oldType != nil {
			newId = AddNewTypeCoordinate("", newTerrain.TextureFlore, oldType.TextureObject,
				"", false, oldType.Move, oldType.View, oldType.Attack, oldType.ImpactRadius)
		} else {
			// т.к. все настройки координаты зависят от обьекта делаем координату полностью открытой
			newId = AddNewTypeCoordinate("", newTerrain.TextureFlore, "", "",
				false, true, true, true, 0)
		}

		PlaceCoordinate(idMap, newId, q, r)
	}
}
