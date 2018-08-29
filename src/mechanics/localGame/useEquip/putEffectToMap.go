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

	AddAnchor(useCoordinate, useEquip) // добавим эфект с якорем в центральную ячекй что бы знать куда ставить спрайт и анимацию

	zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate.X, useCoordinate.Y, useEquip.Region)

	effectCoordinates := make(map[string]map[string]*coordinate.Coordinate)

	for _, zoneCoordinate := range zoneCoordinates {
		gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.X, zoneCoordinate.Y)
		if find {
			for _, effect := range useEquip.Effects { // переносим все эфекты из эквипа выбраной координате
				AddNewCoordinateEffect(gameCoordinate, effect, useEquip.StepsTime)
			}
			AddCoordinate(effectCoordinates, gameCoordinate)
			update.CoordinateEffects(gameCoordinate)
		}
	}

	update.Player(client)

	return effectCoordinates
}

func AddAnchor(useCoordinate *coordinate.Coordinate, useEquip *equip.Equip)  {
	addAnchor := true

	for _, effect := range useEquip.Effects {
		for _, coordinateEffect := range useCoordinate.Effects {
			if coordinateEffect.Type == "anchor" && effect.Name == coordinateEffect.Name {
				addAnchor = false
			}
		}
	}

	if addAnchor {
		for _, effect := range useEquip.Effects {
			if effect.Type == "anchor" {
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