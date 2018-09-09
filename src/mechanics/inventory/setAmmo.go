package inventory

import (
	"../db/get"
	"../db/updateSquad"
	"../gameObjects/ammo"
	"../gameObjects/detail"
	"../player"
)

func SetMSAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot int) {
	ammoItem := user.GetSquad().Inventory[inventorySlot]

	if ammoItem.ItemID == idAmmo {
		newAmmo := get.Ammo(idAmmo)

		ammoSlot, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]
		if ok {
			SetAmmo(ammoSlot, user, newAmmo, inventorySlot)
		}
	}
}

func SetUnitAmmo(user *player.Player, idAmmo, inventorySlot, numEquipSlot, numberUnitSlot int) {
	ammoItem := user.GetSquad().Inventory[inventorySlot]

	if ammoItem.ItemID == idAmmo {
		newAmmo := get.Ammo(idAmmo)

		unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
		if ok && unitSlot.Unit != nil {
			ammoSlot, ok := unitSlot.Unit.Body.Weapons[numEquipSlot]
			if ok {
				SetAmmo(ammoSlot, user, newAmmo, inventorySlot)
			}
		}
	}
}

func SetAmmo(ammoSlot *detail.BodyWeaponSlot, user *player.Player, newAmmo *ammo.Ammo, inventorySlot int) {
	if ammoSlot.Weapon == nil {
		return // если нет оружия ему нельзя поставить боеприпас
	}

	if ammoSlot.Ammo != nil {
		AddItem(user.GetSquad().Inventory, ammoSlot.Ammo, "ammo", ammoSlot.Ammo.ID, ammoSlot.AmmoQuantity, 1)
	}

	ammoSlot.Ammo = newAmmo
	ammoSlot.AmmoQuantity = RemoveInventoryItem(ammoSlot.Weapon.AmmoCapacity, user.GetSquad().Inventory[inventorySlot])

	updateSquad.Squad(user.GetSquad())
}
