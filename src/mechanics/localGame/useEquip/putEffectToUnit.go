package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/detail"
	"../../db/localGame/update"
	"../../player"
)

func ToUnit(gameUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot, client *player.Player) {

	useEquipSlot.Used = true

	for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту
		AddNewUnitEffect(gameUnit, effect)
	}

	update.Player(client)
	update.UnitEffects(gameUnit)
}
