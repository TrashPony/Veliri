package squadInventory

import (
	"../db/squad/update"
	"../factories/storages"
	"../gameObjects/unit"
	"../player"
	"errors"
)

func RemoveAmmo(user *player.Player, numEquipSlot int, unit *unit.Unit) error {
	if user.InBaseID > 0 {
		slot, ok := unit.Body.Weapons[numEquipSlot]

		if ok && slot != nil && slot.Ammo != nil {
			okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Ammo, "ammo", slot.Ammo.ID,
				slot.AmmoQuantity, 1, slot.Ammo.Size, 1)
			if okAddItem {
				slot.Ammo = nil
				go update.Squad(user.GetSquad(), true)
				return nil
			} else {
				return errors.New("add item error")
			}
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("not in base")
	}
}

func RemoveUnitAmmo(user *player.Player, numEquipSlot, numberUnitSlot int) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return RemoveAmmo(user, numEquipSlot, unitSlot.Unit)
	} else {
		return errors.New("no unit")
	}
}
