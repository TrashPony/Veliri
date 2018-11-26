package mapEditor

func PlaceAnimate(idMap, idType, q, r int) {

	newObject := getTypeByID(idType)             // отсюда берем анимация
	oldType := getMapCoordinateInMC(idMap, q, r) // отсюда терейн

	if oldType == nil {
		// если там нет значащей координаты берем терейн и лвл по умолчанию на карте
		defaultLevel, defaultType := getDefaultMap(idMap)
		oldType = getTypeByID(defaultType)
		oldType.Level = defaultLevel
	}
	// т.к. мы ставим именно анимацию, то можем игнорировать обьект
	newType := getTypeByTerrainAndObject(oldType.TextureFlore, "", newObject.AnimateSpriteSheets)

	if newType != nil {
		PlaceCoordinate(idMap, newType.ID, q, r)
	} else {
		// создаем новый тип
		newId := AddNewTypeCoordinate("", oldType.TextureFlore, "",
			newObject.AnimateSpriteSheets, true, newObject.Move, newObject.View, newObject.Attack, newObject.ImpactRadius)

		PlaceCoordinate(idMap, newId, q, r)
	}
}
