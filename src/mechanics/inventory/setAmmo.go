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

			capacity := ammoSlot.Weapon.AmmoCapacity
			ammoCount := user.GetSquad().Inventory[inventorySlot].Quantity

			if capacity < ammoCount {
				user.GetSquad().Inventory[inventorySlot].Quantity = ammoCount - capacity
				ammoSlot.Ammo = newAmmo
				ammoSlot.AmmoQuantity = capacity
			} else {
				user.GetSquad().Inventory[inventorySlot].Item = nil // ставим итему nil что бы при обновление удалился слот из бд
				ammoSlot.Ammo = newAmmo
				ammoSlot.AmmoQuantity = ammoCount
			}

			updateSquad.Squad(user.GetSquad())
		}
	}
}
