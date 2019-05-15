package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func RemoveWeapon(user *player.Player, numEquipSlot int, unit *unit.Unit, dst string) error {
	if user.InBaseID > 0 {

		defer unit.CalculateParams()

		slot, ok := unit.Body.Weapons[numEquipSlot]

		if ok && slot != nil && slot.Weapon != nil {
			if slot.Ammo != nil {
				RemoveAmmo(user, numEquipSlot, unit, dst)
			}

			if dst == "squadInventory" {
				// TODO
				return nil
			}

			if dst == "storage" {
				okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Weapon, "weapon", slot.Weapon.ID,
					1, slot.HP, slot.Weapon.Size, slot.Weapon.MaxHP, false)

				if okAddItem {
					slot.Weapon = nil
					return nil
				} else {
					return errors.New("add item error")
				}
			}
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}

	return errors.New("unknown error")
}

func RemoveUnitWeapon(user *player.Player, numEquipSlot, numberUnitSlot int, dst string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return RemoveWeapon(user, numEquipSlot, unitSlot.Unit, dst)
	} else {
		return errors.New("no unit")
	}
}
