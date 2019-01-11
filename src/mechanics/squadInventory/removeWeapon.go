package squadInventory

import (
	"../db/squad/update"
	"../player"
)

func RemoveMSWeapon(user *player.Player, numEquipSlot int) {
	slotWeapon, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotWeapon != nil && slotWeapon.Weapon != nil {
		if slotWeapon.Ammo != nil {
			RemoveMSAmmo(user, numEquipSlot)
		}

		user.GetSquad().Inventory.AddItem(slotWeapon.Weapon, "weapon", slotWeapon.Weapon.ID, 1,
			slotWeapon.HP, slotWeapon.Weapon.Size, slotWeapon.Weapon.MaxHP)
		slotWeapon.Weapon = nil

		user.GetSquad().MatherShip.CalculateParams()

		go update.Squad(user.GetSquad(), true)
	}
}

func RemoveUnitWeapon(user *player.Player, numEquipSlot, numberUnitSlot int) {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

	if ok && unitSlot.Unit != nil {
		slotWeapon, ok := unitSlot.Unit.Body.Weapons[numEquipSlot]

		if ok && slotWeapon != nil && slotWeapon.Weapon != nil {
			if slotWeapon.Ammo != nil {
				RemoveUnitAmmo(user, numEquipSlot, numberUnitSlot)
			}

			user.GetSquad().Inventory.AddItem(slotWeapon.Weapon, "weapon", slotWeapon.Weapon.ID, 1,
				slotWeapon.HP, slotWeapon.Weapon.Size, slotWeapon.Weapon.MaxHP)
			slotWeapon.Weapon = nil

			unitSlot.Unit.CalculateParams()

			go update.Squad(user.GetSquad(), true)
		}
	}
}
