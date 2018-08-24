package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/equip"
	"../../db/localGame/update"
	"../../player"
)

func ToUnit(gameUnit *unit.Unit, useEquip *equip.Equip, client *player.Player) {

	useEquip.Used = false //TODO делаем эквип использованым но сейчас нет для тестов надо исправитьв будущем

	for _, effect := range useEquip.Effects { // переносим все эфекты из него выбраному юниту
		AddNewUnitEffect(gameUnit, effect)
	}

	update.Player(client)
	update.UnitEffects(gameUnit)
}
