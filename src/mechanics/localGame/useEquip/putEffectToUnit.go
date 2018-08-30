package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/detail"
	"../../db/localGame/update"
	"../../db/updateSquad"
	"../../player"
	"errors"
)

func ToUnit(useUnit, toUseUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot, client *player.Player) error {

	if !useUnit.UseEquip && !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {

		useUnit.Power = useUnit.Power - useEquipSlot.Equip.UsePower

		useUnit.UseEquip = false // todo для тестов false, для игры true

		useEquipSlot.Used = false // todo для тестов false, для игры true
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту
			AddNewUnitEffect(toUseUnit, effect, useEquipSlot.Equip.StepsTime)
		}

		update.UnitEffects(toUseUnit)

		updateSquad.Squad(client.GetSquad())

		return nil
	} else {
		return errors.New("not allow")
	}
}
