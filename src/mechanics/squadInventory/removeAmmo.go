package squadInventory

import (
	"../db/squad/update"
	"../player"
)

func RemoveMSAmmo(user *player.Player, numEquipSlot int) {
	slotAmmo, ok := user.GetSquad().MatherShip.Body.Weapons[numEquipSlot]

	if ok && slotAmmo != nil && slotAmmo.Ammo != nil {
		user.GetSquad().Inventory.AddItem(slotAmmo.Ammo, "ammo", slotAmmo.Ammo.ID, slotAmmo.AmmoQuantity, 1, slotAmmo.Ammo.Size, 1)
		slotAmmo.Ammo = nil

		go update.Squad(user.GetSquad(), true)
	}
}

func RemoveUnitAmmo(user *player.Player, numEquipSlot, numberUnitSlot int) {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]

	if ok && unitSlot.Unit != nil {

		slotAmmo, ok := unitSlot.Unit.Body.Weapons[numEquipSlot]

		if ok && slotAmmo != nil && slotAmmo.Ammo != nil {
			user.GetSquad().Inventory.AddItem(slotAmmo.Ammo, "ammo", slotAmmo.Ammo.ID, slotAmmo.AmmoQuantity, 1, slotAmmo.Ammo.Size, 1)
			slotAmmo.Ammo = nil

			go update.Squad(user.GetSquad(), true)
		}
	}
}
