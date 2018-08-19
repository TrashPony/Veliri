package inventory

import (
	"../player"
	"../db/updateSquad"
)

func RemoveMSWeapon(user *player.Player, numEquipSlot int) {
	slotWeapon, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotWeapon != nil && slotWeapon.Weapon != nil {
		if slotWeapon.Ammo != nil {
			RemoveMSAmmo(user, numEquipSlot)
		}

		AddItem(user.GetSquad().Inventory, slotWeapon.Weapon, "weapon", slotWeapon.Weapon.ID, 1)
		slotWeapon.Weapon = nil

		updateSquad.Squad(user.GetSquad())
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

			AddItem(user.GetSquad().Inventory, slotWeapon.Weapon, "weapon", slotWeapon.Weapon.ID, 1)
			slotWeapon.Weapon = nil

			updateSquad.Squad(user.GetSquad())
		}
	}
}