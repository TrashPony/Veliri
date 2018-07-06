package inventory

import (
	"../player"
	"../db/updateSquad"
)

func RemoveAmmo(user *player.Player, numEquipSlot int) {
	slotAmmo, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotAmmo != nil && slotAmmo.Ammo != nil {

		AddItem(user.GetSquad().Inventory, slotAmmo.Ammo, "ammo", slotAmmo.Ammo.ID, slotAmmo.AmmoQuantity)
		slotAmmo.Ammo = nil

		updateSquad.Squad(user.GetSquad())
	}
}
