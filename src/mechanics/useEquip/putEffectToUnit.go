package useEquip

import (
	"../unit"
	"../equip"
	"../db"
	"../player"
	"../../mechanics"
)

func ToUnit(gameUnit *unit.Unit, useEquip *equip.Equip, client *player.Player) {

	useEquip.Used = false //TODO делаем эквип использованым но сейчас нет для тестов надо исправитьв будущем

	for _, effect := range useEquip.Effects { // переносим все эфекты из него выбраному юниту
		mechanics.AddNewEffect(gameUnit, effect)
	}

	db.UpdateUnit(gameUnit)
	db.UpdatePlayer(client)
}
