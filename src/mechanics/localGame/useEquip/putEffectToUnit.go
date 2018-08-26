package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/detail"
	"../../db/localGame/update"
	"../../db/updateSquad"
	"../../player"
)

func ToUnit(gameUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot, client *player.Player) {

	useEquipSlot.Used = true
	useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

	for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту
		AddNewUnitEffect(gameUnit, effect)
	}

	update.UnitEffects(gameUnit)
	updateSquad.Squad(client.GetSquad())
}
