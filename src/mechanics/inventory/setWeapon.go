package inventory

import (
	"../player"
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/detail"
)

func SetMSWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int) {
	weapon := user.GetSquad().Inventory[inventorySlot]

	if weapon.ItemID == idWeapon {
		newWeapon := get.Weapon(idWeapon)

		weaponSlot, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]
		if ok {
			SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
		}
	}
}

func SetUnitWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot, numberUnitSlot int) {
	weapon := user.GetSquad().Inventory[inventorySlot]

	if weapon.ItemID == idWeapon {
		newWeapon := get.Weapon(idWeapon)
		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {
			weaponSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot].Unit.Body.Weapons[numEquipSlot]
			if ok {
				SetWeapon(weaponSlot, user, newWeapon, inventorySlot, weapon.HP)
			}
		}
	}
}

func SetWeapon(weaponSlot *detail.BodyWeaponSlot, user *player.Player, newWeapon *detail.Weapon, inventorySlot int, hp int)  {
	if weaponSlot.Weapon != nil {
		AddItem(user.GetSquad().Inventory,  weaponSlot.Weapon, "weapon",  weaponSlot.Weapon.ID, 1, weaponSlot.HP)
		weaponSlot.Weapon = nil
	}

	if weaponSlot.Ammo != nil {
		AddItem(user.GetSquad().Inventory,  weaponSlot.Ammo, "ammo", weaponSlot.Ammo.ID, weaponSlot.AmmoQuantity, 1)
		weaponSlot.Ammo = nil
	}

	updateSquad.Squad(user.GetSquad())

	weaponSlot.HP = hp

	RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
	weaponSlot.Weapon = newWeapon
	weaponSlot.InsertToDB = true // говорим что бы обновилась в бд инфа о вепоне

	updateSquad.Squad(user.GetSquad())
}