package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../gameObjects/detail"
	"../player"
	"errors"
)

func SetMSWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int) error {
	weapon := user.GetSquad().Inventory.Slots[inventorySlot]
	msBody := user.GetSquad().MatherShip.Body

	if weapon != nil && msBody != nil && weapon.ItemID == idWeapon && weapon.Type == "weapon"{
		newWeapon, _ := gameTypes.Weapons.GetByID(idWeapon)

		weaponSlot, ok := msBody.Weapons[numEquipSlot]
		if ok {
			// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
			if (weaponSlot.Weapon != nil && msBody.MaxPower-msBody.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
				(weaponSlot.Weapon == nil && msBody.MaxPower-msBody.GetUsePower() >= newWeapon.Power) {
				if newWeapon.StandardSize == 1 && msBody.StandardSizeSmall {
					SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
					user.GetSquad().MatherShip.CalculateParams()
					return nil
				}
				if newWeapon.StandardSize == 2 && msBody.StandardSizeMedium {
					SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
					user.GetSquad().MatherShip.CalculateParams()
					return nil
				}
				if newWeapon.StandardSize == 3 && msBody.StandardSizeBig {
					SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
					user.GetSquad().MatherShip.CalculateParams()
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
}

func SetUnitWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot, numberUnitSlot int) error {
	weapon := user.GetSquad().Inventory.Slots[inventorySlot]

	if weapon.Item != nil && weapon.ItemID == idWeapon && weapon.Type == "weapon"{
		newWeapon, _ := gameTypes.Weapons.GetByID(idWeapon)

		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {

			weaponSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot].Unit.Body.Weapons[numEquipSlot]
			unitBody := user.GetSquad().MatherShip.Units[numberUnitSlot].Unit.Body

			if ok {
				// писос, но тут смотрить можно ли поставить из расчета свободной энергии, или в замену текущему эквипу
				if (weaponSlot.Weapon != nil && unitBody.MaxPower-unitBody.GetUsePower()+weaponSlot.Weapon.Power >= newWeapon.Power) ||
					(weaponSlot.Weapon == nil && unitBody.MaxPower-unitBody.GetUsePower() >= newWeapon.Power) {
					if unitBody.GetUseCapacitySize()+newWeapon.Size <= unitBody.CapacitySize {
						if newWeapon.StandardSize == 1 && unitBody.StandardSizeSmall {
							SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
							unitSlot.Unit.CalculateParams()
							return nil
						}
						if newWeapon.StandardSize == 2 && unitBody.StandardSizeMedium {
							SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
							unitSlot.Unit.CalculateParams()
							return nil
						}
						if newWeapon.StandardSize == 3 && unitBody.StandardSizeBig {
							SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
							unitSlot.Unit.CalculateParams()
							return nil
						}
						return errors.New("wrong standard size")
					} else {
						return errors.New("lacking size")
					}
				} else {
					return errors.New("lacking power")
				}
			} else {
				return errors.New("wrong weapon slot")
			}
		} else {
			return errors.New("wrong unit")
		}
	} else {
		return errors.New("wrong inventory slot")
	}
}

func SetWeapon(weaponSlot *detail.BodyWeaponSlot, user *player.Player, newWeapon *detail.Weapon, inventorySlot int, hp int) {
	if weaponSlot.Weapon != nil {
		user.GetSquad().Inventory.AddItem(weaponSlot.Weapon, "weapon", weaponSlot.Weapon.ID, 1,
			weaponSlot.HP, weaponSlot.Weapon.Size, weaponSlot.Weapon.MaxHP)
		weaponSlot.Weapon = nil
	}

	if weaponSlot.Ammo != nil {
		user.GetSquad().Inventory.AddItem(weaponSlot.Ammo, "ammo", weaponSlot.Ammo.ID, weaponSlot.AmmoQuantity,
			1, weaponSlot.Ammo.Size, 1)
		weaponSlot.Ammo = nil
	}

	update.Squad(user.GetSquad(), true)

	weaponSlot.HP = hp

	user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(1)
	weaponSlot.Weapon = newWeapon
	weaponSlot.InsertToDB = true // говорим что бы обновилась в бд инфа о вепоне

	update.Squad(user.GetSquad(), true)
}
