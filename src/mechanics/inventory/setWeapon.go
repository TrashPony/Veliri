package inventory

import (
	"../player"
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/detail"
	"errors"
)

func SetMSWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int) error {
	weapon := user.GetSquad().Inventory[inventorySlot]
	msBody := user.GetSquad().MatherShip.Body

	if weapon.ItemID == idWeapon {
		newWeapon := get.Weapon(idWeapon)

		weaponSlot, ok := msBody.Weapons[numEquipSlot]
		if ok {
			// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
			if (weaponSlot.Weapon != nil && msBody.MaxPower-msBody.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
				(weaponSlot.Weapon == nil && msBody.MaxPower-msBody.GetUsePower() >= newWeapon.Power) {

				SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
			} else {
				return errors.New("lacking power")
			}
		}
	}

	return nil
}

func SetUnitWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot, numberUnitSlot int) error {
	weapon := user.GetSquad().Inventory[inventorySlot]

	if weapon.ItemID == idWeapon {
		newWeapon := get.Weapon(idWeapon)
		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {

			weaponSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot].Unit.Body.Weapons[numEquipSlot]
			unitBody := user.GetSquad().MatherShip.Units[numberUnitSlot].Unit.Body

			if ok {
				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (weaponSlot.Weapon != nil && unitBody.MaxPower-unitBody.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
					(weaponSlot.Weapon == nil && unitBody.MaxPower-unitBody.GetUsePower() >= newWeapon.Power) {
					SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
				} else {
					return errors.New("lacking power")
				}
			}
		}
	}

	return nil
}

func SetWeapon(weaponSlot *detail.BodyWeaponSlot, user *player.Player, newWeapon *detail.Weapon, inventorySlot int, hp int) {
	if weaponSlot.Weapon != nil {
		AddItem(user.GetSquad().Inventory, weaponSlot.Weapon, "weapon", weaponSlot.Weapon.ID, 1, weaponSlot.HP)
		weaponSlot.Weapon = nil
	}

	if weaponSlot.Ammo != nil {
		AddItem(user.GetSquad().Inventory, weaponSlot.Ammo, "ammo", weaponSlot.Ammo.ID, weaponSlot.AmmoQuantity, 1)
		weaponSlot.Ammo = nil
	}

	updateSquad.Squad(user.GetSquad())

	weaponSlot.HP = hp

	RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
	weaponSlot.Weapon = newWeapon
	weaponSlot.InsertToDB = true // говорим что бы обновилась в бд инфа о вепоне

	updateSquad.Squad(user.GetSquad())
}
