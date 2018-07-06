package inventory

import (
	"../player"
	"../db/updateSquad"
)

func RemoveWeapon(user *player.Player, numEquipSlot int) {
	slotWeapon, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotWeapon != nil && slotWeapon.Weapon != nil {

		if slotWeapon.Ammo != nil {
			RemoveAmmo(user, numEquipSlot)
		}

		AddItem(user.GetSquad().Inventory, slotWeapon.Weapon, "weapon", slotWeapon.Weapon.ID, 1)
		slotWeapon.Weapon = nil

		updateSquad.Squad(user.GetSquad())
	}
}
