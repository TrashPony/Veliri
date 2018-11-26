package mapEditor

func PlaceObject(idMap, idType, q, r int) {

	newObject := getTypeByID(idType)             // отсюда берем обьект
	oldType := getMapCoordinateInMC(idMap, q, r) // отсюда терейн

	if oldType == nil {
		// если там нет значащей координаты берем терейн и лвл по умолчанию на карте
		defaultLevel, defaultType := getDefaultMap(idMap)
		oldType = getTypeByID(defaultType)
		oldType.Level = defaultLevel
	}
	// т.к. мы ставим именно обьект, то можем игнорировать анимацию
	newType := getTypeByTerrainAndObject(oldType.TextureFlore, newObject.TextureObject, "")

	if newType != nil {
		PlaceCoordinate(idMap, newType.ID, q, r)
	} else {
		// создаем новый тип
		newId := AddNewTypeCoordinate("", oldType.TextureFlore, newObject.TextureObject,
			"", false, newObject.Move, newObject.View, newObject.Attack, newObject.ImpactRadius)

		PlaceCoordinate(idMap, newId, q, r)
	}
}
