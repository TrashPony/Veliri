package squad_inventory

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func RemoveAmmo(user *player.Player, numEquipSlot int, unit *unit.Unit, dst string) error {
	if user.InBaseID == 0 {
		// если мы на улице то можем скинуть аммо только в инвентарь
		dst = "squadInventory"
	}

	defer unit.CalculateParams()

	slot, ok := unit.Body.Weapons[numEquipSlot]

	if ok && slot != nil && slot.Ammo != nil {

		if dst == "squadInventory" {
			okAddItem := user.GetSquad().Inventory.AddItem(slot.Ammo, "ammo", slot.Ammo.ID, slot.AmmoQuantity, 1, slot.Ammo.Size, 1, false)
			if okAddItem {
				slot.Ammo = nil
				return nil
			} else {
				return errors.New("add item in inventory error")
			}
		}

		if dst == "storage" {
			okAddItem := storages.Storages.AddItem(user.GetID(), user.InBaseID, slot.Ammo, "ammo", slot.Ammo.ID,
				slot.AmmoQuantity, 1, slot.Ammo.Size, 1, false)

			if okAddItem {
				slot.Ammo = nil
				return nil
			} else {
				return errors.New("add item error")
			}
		}
	} else {
		return errors.New("no item")
	}
	return errors.New("unknown error")
}

func RemoveUnitAmmo(user *player.Player, numEquipSlot, numberUnitSlot int, dst string) error {
	unitSlot, ok := user.GetSquad().MatherShip.Units[numberUnitSlot]
	if ok && unitSlot.Unit != nil {
		return RemoveAmmo(user, numEquipSlot, unitSlot.Unit, dst)
	} else {
		return errors.New("no unit")
	}
}
