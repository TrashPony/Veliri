package inventory

import (
	"../player"
	"../db/get"
	"../db/updateSquad"
)

func SetAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot int) {
	ammo := user.GetSquad().Inventory[inventorySlot]

	if ammo.ItemID == idAmmo {
		newAmmo := get.Ammo(idAmmo)

		ammoSlot, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]
		if ok {

			if ammoSlot.Weapon == nil {
				return // если нет оружия ему нельзя поставить боеприпас
			}

			if ammoSlot.Ammo != nil {
				AddItem(user.GetSquad().Inventory, ammoSlot.Ammo, "ammo", ammoSlot.Ammo.ID, ammoSlot.AmmoQuantity)
			}

			ammoSlot.Ammo = newAmmo
			ammoSlot.AmmoQuantity = RemoveInventoryItem(ammoSlot.Weapon.AmmoCapacity, user.GetSquad().Inventory[inventorySlot])

			updateSquad.Squad(user.GetSquad())
		}
	}
}
