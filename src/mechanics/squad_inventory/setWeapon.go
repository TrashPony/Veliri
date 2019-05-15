package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func SetWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int, unit *unit.Unit, source string) error {
	if user.InBaseID > 0 {

		slot := getSlotBySource(user, inventorySlot, source)

		body := unit.Body

		if slot != nil && body != nil && slot.ItemID == idWeapon && slot.Type == "weapon" {

			if slot.HP <= 0 {
				return errors.New("equip broken")
			}

			newWeapon, _ := gameTypes.Weapons.GetByID(idWeapon)
			weaponSlot, ok := body.Weapons[numEquipSlot]

			if ok {
				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (weaponSlot.Weapon != nil && body.MaxPower-body.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
					(weaponSlot.Weapon == nil && body.MaxPower-body.GetUsePower() >= newWeapon.Power) {

					if newWeapon.StandardSize == 1 && body.StandardSizeSmall {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, slot.HP, unit, source)
						return nil
					}
					if newWeapon.StandardSize == 2 && body.StandardSizeMedium {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, slot.HP, unit, source)
						return nil
					}
					if newWeapon.StandardSize == 3 && body.StandardSizeBig {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, slot.HP, unit, source)
						return nil
					}

					return errors.New("wrong standard size")
				} else {
					return errors.New("lacking power")
				}
			} else {
				return errors.New("wrong weapon slot")
			}
		} else {
			return errors.New("wrong inventory slot")
		}
	} else {
		return errors.New("not in base")
	}
}

func SetUnitWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot, numberUnitSlot int, source string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return SetWeapon(user, idWeapon, inventorySlot, numEquipSlot, unitSlot.Unit, source)
	} else {
		return errors.New("no unit")
	}
}

func setWeapon(weaponSlot *detail.BodyWeaponSlot, user *player.Player, newWeapon *detail.Weapon, inventorySlot, hp int, unit *unit.Unit, source string) {

	defer unit.CalculateParams()

	if weaponSlot.Weapon != nil {
		RemoveWeapon(user, weaponSlot.Number, unit, "storage")
	}

	if weaponSlot.Ammo != nil {
		RemoveAmmo(user, weaponSlot.Number, unit, "storage")
	}

	update.Squad(user.GetSquad(), true) // без этого если в слоте есть снаряжение то оно не заменяется, а добавляется в бд

	weaponSlot.HP = hp
	RemoveSlotBySource(user, inventorySlot, source, 1)

	weaponSlot.Weapon = newWeapon
	weaponSlot.InsertToDB = true // говорим что бы обновилась в бд инфа о вепоне
}
