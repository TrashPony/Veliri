package attackPhase

import (
	"../../../db/localGame/update"
	"../../../gameObjects/detail"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../localGame/map/watchZone"
	"../../useEquip"
)

func ToUnit(useUnit, toUseUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot) *ResultBattle {
	if !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {

		targetsUnit := make([]TargetUnit, 0)

		useUnit.Power -= useEquipSlot.Equip.UsePower
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		useEquipSlot.Used = true

		for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту
			useEquip.AddNewUnitEffect(toUseUnit, effect, useEquipSlot.Equip.StepsTime)

			if effect.Type == "replenishes" && effect.Parameter == "hp" {
				toUseUnit.HP += effect.Quantity
				targetsUnit = append(targetsUnit, TargetUnit{Unit: *toUseUnit, Heal: effect.Quantity})
			}

			if effect.Type == "takes_away" && effect.Parameter == "hp" {
				toUseUnit.HP -= effect.Quantity
				targetsUnit = append(targetsUnit, TargetUnit{Unit: *toUseUnit, Damage: effect.Quantity})
			}

			if effect.Type == "replenishes" && effect.Parameter == "power" {
				toUseUnit.Power += effect.Quantity
				targetsUnit = append(targetsUnit, TargetUnit{Unit: *toUseUnit, Power: effect.Quantity})
			}

			if effect.Type == "takes_away" && effect.Parameter == "power" {
				toUseUnit.Power -= effect.Quantity
				targetsUnit = append(targetsUnit, TargetUnit{Unit: *toUseUnit, Power: -effect.Quantity})
			}
		}

		toUseUnit.CalculateParams() // обновляем параметры юнита
		update.UnitEffects(toUseUnit)

		// добавляет того кто использует т.к. у него отнимается энергия на использование
		targetsUnit = append(targetsUnit, TargetUnit{Unit: *useUnit, Power: useEquipSlot.Equip.UsePower})

		return &ResultBattle{AttackUnit: *useUnit, TargetUnits: targetsUnit, EquipSlot: useEquipSlot}
	} else {
		return &ResultBattle{Error: "no power"}
	}
}

func MoveEquip(useUnit *unit.Unit, game *localGame.Game, useEquipSlot *detail.BodyEquipSlot) *ResultBattle {
	resultAction := &ResultBattle{AttackUnit: *useUnit, EquipSlot: useEquipSlot, Target: useEquipSlot.Target}
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
			resultAction.WatchNode = watchNode
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
