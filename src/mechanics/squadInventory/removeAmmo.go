package squadInventory

import (
	"../db/updateSquad"
	"../player"
)

func RemoveMSAmmo(user *player.Player, numEquipSlot int) {
	slotAmmo, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotAmmo != nil && slotAmmo.Ammo != nil {
		user.GetSquad().Inventory.AddItem(slotAmmo.Ammo, "ammo", slotAmmo.Ammo.ID, slotAmmo.AmmoQuantity, 1, slotAmmo.Ammo.Size)
		slotAmmo.Ammo = nil

		updateSquad.Squad(user.GetSquad())
	}
}

func RemoveUnitAmmo(user *player.Player, numEquipSlot, numberUnitSlot int) {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

	if ok && unitSlot.Unit != nil {

		slotAmmo, ok := unitSlot.Unit.Body.Weapons[numEquipSlot]

		if ok && slotAmmo != nil && slotAmmo.Ammo != nil {
			user.GetSquad().Inventory.AddItem(slotAmmo.Ammo, "ammo", slotAmmo.Ammo.ID, slotAmmo.AmmoQuantity, 1, slotAmmo.Ammo.Size)
			slotAmmo.Ammo = nil

			updateSquad.Squad(user.GetSquad())
		}
	}
}