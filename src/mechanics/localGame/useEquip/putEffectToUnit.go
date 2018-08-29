package useEquip

import (
	"../../gameObjects/unit"
	"../../gameObjects/detail"
	"../../db/localGame/update"
	"../../db/updateSquad"
	"../../player"
	"errors"
)

func ToUnit(gameUnit *unit.Unit, useEquipSlot *detail.BodyEquipSlot, client *player.Player) error {

	if !gameUnit.UseEquip && !useEquipSlot.Used && gameUnit.Power >= useEquipSlot.Equip.UsePower {

		gameUnit.Power = gameUnit.Power - useEquipSlot.Equip.UsePower
		gameUnit.UseEquip = true

		useEquipSlot.Used = true
		useEquipSlot.StepsForReload = useEquipSlot.Equip.Reload

		for _, effect := range useEquipSlot.Equip.Effects { // переносим все эфекты из него выбраному юниту
			AddNewUnitEffect(gameUnit, effect, useEquipSlot.Equip.StepsTime)
		} // TODO разрулить время на которое должен работать эфект

		update.UnitEffects(gameUnit)
		updateSquad.Squad(client.GetSquad())

		return nil
	} else {
		return errors.New("not allow")
	}
}
