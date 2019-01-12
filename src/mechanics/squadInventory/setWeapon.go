package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../gameObjects/detail"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func SetWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int, unit *unit.Unit) error {
	if user.InBaseID > 0 {
		weapon := user.GetSquad().Inventory.Slots[inventorySlot]
		body := unit.Body

		if weapon != nil && body != nil && weapon.ItemID == idWeapon && weapon.Type == "weapon" {

			newWeapon, _ := gameTypes.Weapons.GetByID(idWeapon)
			weaponSlot, ok := body.Weapons[numEquipSlot]

			if ok {
				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (weaponSlot.Weapon != nil && body.MaxPower-body.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
					(weaponSlot.Weapon == nil && body.MaxPower-body.GetUsePower() >= newWeapon.Power) {

					if newWeapon.StandardSize == 1 && body.StandardSizeSmall {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP, unit)
						unit.CalculateParams()
						return nil
					}
					if newWeapon.StandardSize == 2 && body.StandardSizeMedium {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP, unit)
						unit.CalculateParams()
						return nil
					}
					if newWeapon.StandardSize == 3 && body.StandardSizeBig {
						setWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP, unit)
						unit.CalculateParams()
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

func SetUnitWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot, numberUnitSlot int) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return SetWeapon(user, idWeapon, inventorySlot, numEquipSlot, unitSlot.Unit)
	} else {
		return errors.New("no unit")
	}
}

func setWeapon(weaponSlot *detail.BodyWeaponSlot, user *player.Player, newWeapon *detail.Weapon, inventorySlot int, hp int, unit *unit.Unit) {
	if weaponSlot.Weapon != nil {
		RemoveWeapon(user, weaponSlot.Number, unit)
	}

	if weaponSlot.Ammo != nil {
		RemoveAmmo(user, weaponSlot.Number, unit)
	}

	update.Squad(user.GetSquad(), true) // без этого если в слоте есть снаряжение то оно не заменяется, а добавляется в бд

	weaponSlot.HP = hp

	user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(1)
	weaponSlot.Weapon = newWeapon
	weaponSlot.InsertToDB = true // говорим что бы обновилась в бд инфа о вепоне

	go update.Squad(user.GetSquad(), true)
}
