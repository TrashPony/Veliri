package inventory

import (
	"../player"
	"../db/get"
	"../db/updateSquad"
)

func SetWeapon(user *player.Player, idWeapon, inventorySlot, numEquipSlot int) {
	weapon := user.GetSquad().Inventory[inventorySlot]

	if weapon.ItemID == idWeapon {
		newWeapon := get.Weapon(idWeapon)

		weaponSlot, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]
		if ok {

			if weaponSlot.Weapon != nil {
				AddItem(user.GetSquad().Inventory,  weaponSlot.Weapon, "weapon",  weaponSlot.Weapon.ID, 1)
			} else {
				weaponSlot.InsertToDB = true
			}

			if weaponSlot.Ammo != nil {
				AddItem(user.GetSquad().Inventory,  weaponSlot.Ammo, "ammo", weaponSlot.Ammo.ID, weaponSlot.AmmoQuantity)
				weaponSlot.Ammo = nil
			}

			RemoveInventoryItem(1, user.GetSquad().Inventory[inventorySlot])
			weaponSlot.Weapon = newWeapon

			updateSquad.Squad(user.GetSquad())
		}
	}
}
