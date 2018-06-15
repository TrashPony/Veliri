package useEquip

import (
	"../game"
	"../equip"
	"../player"
	"../coordinate"
	"../../mechanics"
	"../db"
	"strconv"
)

func ToMap(useCoordinate *coordinate.Coordinate, activeGame *game.Game, useEquip *equip.Equip, client *player.Player) map[string]map[string]*coordinate.Coordinate {

	useEquip.Used = false //TODO делаем эквип использованым но сейчас нет для тестов надо исправитьв будущем

	zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate.X, useCoordinate.Y, useEquip.Region)

	effectCoordinates := make(map[string]map[string]*coordinate.Coordinate)

	for _, zoneCoordinate := range zoneCoordinates {
		gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.X, zoneCoordinate.Y)
		if find {
			for _, effect := range useEquip.Effects { // переносим все эфекты из эквипа выбраной координате
				mechanics.AddNewCoordinateEffect(gameCoordinate, *effect)
			}

			AddCoordinate(effectCoordinates, gameCoordinate)

			db.UpdateCoordinateEffects(gameCoordinate)
		}
	}

	db.UpdatePlayer(client)

	return effectCoordinates
}

func AddCoordinate(res map[string]map[string]*coordinate.Coordinate, gameCoordinate *coordinate.Coordinate)  {
	if res[strconv.Itoa(gameCoordinate.X)] != nil {
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	}
}