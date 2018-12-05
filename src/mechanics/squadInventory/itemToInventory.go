package squadInventory

import (
	"../player"
	"../storage"
	"errors"
	"../db/updateSquad"
)

func ItemToInventory(user *player.Player, storageSlot int) error {
	if user.InBaseID > 0 {

		baseStorage, _ := storage.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]

		if slot == nil {
			return errors.New("no find slot")
		}

		ok := user.GetSquad().Inventory.AddItem(slot.Item, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.Size/float32(slot.Quantity))

		if ok {
			storage.Storages.RemoveItem(user.GetID(), user.InBaseID, storageSlot, slot.Quantity)
		}

		updateSquad.Squad(user.GetSquad())
		return nil
	} else {
		return errors.New("user not in base")
	}
}
