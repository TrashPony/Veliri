package useEquip

import (
	"../../../mechanics/gameObjects/coordinate"
	"../../../mechanics/localGame"
	"../../../mechanics/gameObjects/equip"
	"../../../mechanics/player"
	"../../../mechanics/db/localGame/update"
	"strconv"
)

func ToMap(useCoordinate *coordinate.Coordinate, activeGame *localGame.Game, useEquip *equip.Equip, client *player.Player) map[string]map[string]*coordinate.Coordinate {

	AddAnchor(useCoordinate, useEquip, "anchor") // добавим эфект с якорем в центральную ячекй что бы знать куда ставить спрайт и анимацию
	AddAnchor(useCoordinate, useEquip, "animate") // добавим эфект с анимацией что бы проиграть анимация взрыва при фазе атаки

	zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate.X, useCoordinate.Y, useEquip.Region)

	effectCoordinates := make(map[string]map[string]*coordinate.Coordinate)

	for _, zoneCoordinate := range zoneCoordinates {
		gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.X, zoneCoordinate.Y)
		if find {
			for _, effect := range useEquip.Effects { // переносим все эфекты из эквипа выбраной координате
				if effect.Type != "anchor" && effect.Type != "animate" {
					newEffect := *effect // создаем копию эфекта что бы обнулить ид и добавить в бд как новую
					newEffect.ID = 0
					AddNewCoordinateEffect(gameCoordinate, &newEffect, useEquip.StepsTime)
				}
			}
			AddCoordinate(effectCoordinates, gameCoordinate)
			update.CoordinateEffects(gameCoordinate)
		}
	}

	update.Player(client)

	return effectCoordinates
}

func AddAnchor(useCoordinate *coordinate.Coordinate, useEquip *equip.Equip, typeEffect string)  {
	addAEffect := true

	for _, effect := range useEquip.Effects {
		for _, coordinateEffect := range useCoordinate.Effects {
			if coordinateEffect.Type == typeEffect && effect.Name == coordinateEffect.Name {
				addAEffect = false
			}
		}
	}

	if addAEffect {
		for _, effect := range useEquip.Effects {
			if effect.Type == typeEffect {
				effect.StepsTime = useEquip.StepsTime
				useCoordinate.Effects = append(useCoordinate.Effects, effect)
			}
		}
	}
}

func AddCoordinate(res map[string]map[string]*coordinate.Coordinate, gameCoordinate *coordinate.Coordinate)  {
	if res[strconv.Itoa(gameCoordinate.X)] != nil {
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.X)] = make(map[string]*coordinate.Coordinate)
		res[strconv.Itoa(gameCoordinate.X)][strconv.Itoa(gameCoordinate.Y)] = gameCoordinate
	}
}