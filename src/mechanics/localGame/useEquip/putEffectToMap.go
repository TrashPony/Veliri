package useEquip

import (
	"../../db/localGame/update"
	"../../db/updateSquad"
	"../../gameObjects/coordinate"
	"../../gameObjects/detail"
	"../../gameObjects/equip"
	"../../gameObjects/unit"
	"../../localGame"
	"../../player"
	"errors"
	"strconv"
)

func ToMap(useUnit *unit.Unit, useCoordinate *coordinate.Coordinate, activeGame *localGame.Game, useEquipSlot *detail.BodyEquipSlot, client *player.Player) (map[string]map[string]*coordinate.Coordinate, error) {
	if !useUnit.UseEquip && !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {

		useUnit.Power = useUnit.Power - useEquipSlot.Equip.UsePower
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		useUnit.UseEquip = false  // todo для тестов false, для игры true
		useEquipSlot.Used = false // todo для тестов false, для игры true

		AddAnchor(useCoordinate, useEquipSlot.Equip, "anchor")  // добавим эфект с якорем в центральную ячекй что бы знать куда ставить спрайт и анимацию
		AddAnchor(useCoordinate, useEquipSlot.Equip, "animate") // добавим эфект с анимацией что бы проиграть анимация взрыва при фазе атаки

		zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate, useEquipSlot.Equip.Region)

		effectCoordinates := make(map[string]map[string]*coordinate.Coordinate)

		for _, zoneCoordinate := range zoneCoordinates {
			gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.Q, zoneCoordinate.R)
			if find {
				for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из эквипа выбраной координате
					if effect.Type != "anchor" && effect.Type != "animate" {
						newEffect := *effect // создаем копию эфекта что бы обнулить ид и добавить в бд как новую
						newEffect.ID = 0
						AddNewCoordinateEffect(gameCoordinate, &newEffect, useEquipSlot.Equip.StepsTime)
					}
				}
				AddCoordinate(effectCoordinates, gameCoordinate)
				update.CoordinateEffects(gameCoordinate)
			}
		}

		update.Player(client)
		updateSquad.Squad(client.GetSquad())

		return effectCoordinates, nil
	} else {
		if useUnit.Power < useEquipSlot.Equip.UsePower {
			return nil, errors.New("no power")
		}

		if useUnit.UseEquip || useEquipSlot.Used {
			return nil, errors.New("you not ready")
		}

		return nil, errors.New("unknown error")
	}
}

func AddAnchor(useCoordinate *coordinate.Coordinate, useEquip *equip.Equip, typeEffect string) {
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

func AddCoordinate(res map[string]map[string]*coordinate.Coordinate, gameCoordinate *coordinate.Coordinate) {
	if res[strconv.Itoa(gameCoordinate.Q)] != nil {
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	} else {
		res[strconv.Itoa(gameCoordinate.Q)] = make(map[string]*coordinate.Coordinate)
		res[strconv.Itoa(gameCoordinate.Q)][strconv.Itoa(gameCoordinate.R)] = gameCoordinate
	}
}
