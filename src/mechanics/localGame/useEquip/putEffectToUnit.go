package useEquip

import (
	"../../db/localGame/update"
	"../../db/updateSquad"
	"../../gameObjects/detail"
	"../../gameObjects/unit"
	"../../player"
	"errors"
)

func ToUnit(useUnit, toUseUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot, client *player.Player) error {

	if !useUnit.UseEquip && !useEquipSlot.Used && useUnit.Power >= useEquipSlot.Equip.UsePower {

		useUnit.Power -= useEquipSlot.Equip.UsePower
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		useUnit.UseEquip = true
		useEquipSlot.Used = true

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
