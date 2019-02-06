package squadInventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func RemoveWeapon(user *player.Player, numEquipSlot int, unit *unit.Unit, dst string, updateDB bool) error {
	if user.InBaseID > 0 {
		slot, ok := unit.Body.Weapons[numEquipSlot]

		if ok && slot != nil && slot.Weapon != nil {
			if slot.Ammo != nil {
				RemoveAmmo(user, numEquipSlot, unit, dst, updateDB)
			}

			if dst == "squadInventory" {
				// TODO
				return nil
			}

			if dst == "storage" {
				okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Weapon, "weapon", slot.Weapon.ID,
					1, slot.HP, slot.Weapon.Size, slot.Weapon.MaxHP)

				if okAddItem {
					slot.Weapon = nil
					unit.CalculateParams()

					if updateDB {
						go update.Squad(user.GetSquad(), true)
					}

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
		return RemoveWeapon(user, numEquipSlot, unitSlot.Unit, dst, true)
	} else {
		return errors.New("no unit")
	}
}
