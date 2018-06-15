package useEquip

import (
	"../game"
	"../equip"
	"../player"
	"../coordinate"
	"../../mechanics"
	"../db"
)

func ToMap(useCoordinate *coordinate.Coordinate, activeGame *game.Game, useEquip *equip.Equip, client *player.Player) {

	useEquip.Used = false //TODO делаем эквип использованым но сейчас нет для тестов надо исправитьв будущем

	zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate.X, useCoordinate.Y, useEquip.Region)

	for _, zoneCoordinate := range zoneCoordinates {
		gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.X, zoneCoordinate.Y)
		if find {
			for _, effect := range useEquip.Effects { // переносим все эфекты из эквипа выбраной координате
				mechanics.AddNewCoordinateEffect(gameCoordinate, effect)
			}
			db.UpdateCoordinateEffects(gameCoordinate)
		}
	}

	db.UpdatePlayer(client)
}
