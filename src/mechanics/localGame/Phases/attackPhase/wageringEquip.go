package attackPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/Phases"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/useEquip"
)

func ToUnit(useUnit, toUseUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot) *ResultBattle {
	if !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {

		targetsUnit := make([]TargetUnit, 0)

		useUnit.Power -= useEquipSlot.Equip.UsePower
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		useEquipSlot.Used = true

		for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту

			if (useEquipSlot.Equip.StepsTime > 1) || (useEquipSlot.Equip.StepsTime == 1 && effect.Type != "replenishes" && effect.Type != "takes_away") {
				useEquip.AddNewUnitEffect(toUseUnit, effect, useEquipSlot.Equip.StepsTime)
			}

			powEnEffect(effect, toUseUnit, &targetsUnit)
		}

		toUseUnit.CalculateParams() // обновляем параметры юнита
		update.UnitEffects(toUseUnit)

		// добавляет того кто использует т.к. у него отнимается энергия на использование
		targetsUnit = append(targetsUnit, TargetUnit{Unit: *useUnit, Power: useEquipSlot.Equip.UsePower})

		return &ResultBattle{AttackUnit: *useUnit, TargetUnits: targetsUnit, EquipSlot: *useEquipSlot}
	} else {
		return &ResultBattle{Error: "no power"}
	}
}

func MoveEquip(useUnit *unit.Unit, game *localGame.Game, useEquipSlot *detail.BodyEquipSlot) *ResultBattle {
	resultAction := &ResultBattle{AttackUnit: *useUnit, EquipSlot: *useEquipSlot, Target: *useEquipSlot.Target}
	// тут мы обновляем позицию юнита для всех пользователей игры, + обновляем зону видимости владельца юнита
	for _, gamePlayer := range game.GetPlayers() {
		if gamePlayer.GetLogin() == useUnit.Owner {
			gamePlayer.DelUnit(useUnit, false)
		} else {
			_, find := gamePlayer.GetHostileUnit(useUnit.Q, useUnit.R)
			if find {
				gamePlayer.DelHostileUnit(useUnit.ID)
			}
		}
	}
	game.DelUnit(useUnit)

	useUnit.Q = useEquipSlot.Target.Q
	useUnit.R = useEquipSlot.Target.R

	for _, gamePlayer := range game.GetPlayers() {
		if gamePlayer.GetLogin() == useUnit.Owner {
			gamePlayer.AddUnit(useUnit)
			watchNode := watchZone.UpdateWatchZone(game, gamePlayer)
			if len(watchNode.OpenUnit) > 0 {
				for _, openHostileUnit := range watchNode.OpenUnit {
					gamePlayer.AddNewMemoryHostileUnit(*openHostileUnit)
				}
			}
		} else {
			_, find := gamePlayer.GetWatchCoordinate(useUnit.Q, useUnit.R)
			if find {
				gamePlayer.AddHostileUnit(useUnit)
			}
		}
	}
	game.SetUnit(useUnit)

	return resultAction
}

func ToMap(useUnit *unit.Unit, useCoordinate *coordinate.Coordinate, activeGame *localGame.Game, useEquipSlot *detail.BodyEquipSlot) *ResultBattle {
	if useUnit.Power >= useEquipSlot.Equip.UsePower {

		targetsUnit := make([]TargetUnit, 0)

		useUnit.Power -= useEquipSlot.Equip.UsePower
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload // устанавливает время перезарядки

		useEquipSlot.Used = true

		if !(useEquipSlot.Equip.StepsTime == 1) { // если время действия эквипа 1 ход то вся анимация проигрывается сразу
			AddAnchor(useCoordinate, useEquipSlot.Equip, "anchor")  // добавим эфект с якорем в центральную ячекй что бы знать куда ставить спрайт и анимацию
			AddAnchor(useCoordinate, useEquipSlot.Equip, "animate") // добавим эфект с анимацией что бы проиграть анимация взрыва при фазе атаки
		}

		zoneCoordinates := coordinate.GetCoordinatesRadius(useCoordinate, useEquipSlot.Equip.Region)

		effectCoordinates := make(map[string]map[string]*coordinate.Coordinate)

		for _, zoneCoordinate := range zoneCoordinates {
			gameCoordinate, find := activeGame.Map.GetCoordinate(zoneCoordinate.Q, zoneCoordinate.R)
			gameUnit, findUnit := activeGame.GetUnit(zoneCoordinate.Q, zoneCoordinate.R)

			if find {
				for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из эквипа выбраной координате
					if effect.Type != "anchor" && effect.Type != "animate" {

						if (useEquipSlot.Equip.StepsTime > 1) || (useEquipSlot.Equip.StepsTime == 1 && effect.Type != "replenishes" && effect.Type != "takes_away") {
							newEffect := *effect // создаем копию эфекта что бы обнулить ид и добавить в бд как новую
							newEffect.ID = 0
							useEquip.AddNewCoordinateEffect(gameCoordinate, &newEffect, useEquipSlot.Equip.StepsTime)

							if findUnit {
								// накладываем эфекты на всех юнитов которые стоят на зоне поражения
								useEquip.AddNewUnitEffect(gameUnit, effect, useEquipSlot.Equip.StepsTime)
							}
						}

						if findUnit {
							powEnEffect(effect, gameUnit, &targetsUnit) // сразу отнимаем/даем хп или энергию
						}
					}
				}
				Phases.AddCoordinate(effectCoordinates, gameCoordinate)
				update.CoordinateEffects(gameCoordinate)
			}

			if findUnit {
				gameUnit.CalculateParams() // применяем все новые наложеные эффекты
			}
		}

		// добавляет того кто использует т.к. у него отнимается энергия на использование
		targetsUnit = append(targetsUnit, TargetUnit{Unit: *useUnit, Power: useEquipSlot.Equip.UsePower})

		return &ResultBattle{AttackUnit: *useUnit, Target: *useCoordinate, TargetUnits: targetsUnit, EquipSlot: *useEquipSlot}
	} else {
		return &ResultBattle{Error: "no power"}
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
