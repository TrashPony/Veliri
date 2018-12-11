package squadInventory

import (
	"../db/squad/update"
	"../factories/storages"
	"../player"
	"errors"
)

func ItemToInventory(user *player.Player, storageSlot int) error {
	if user.InBaseID > 0 {

		baseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]

		if slot == nil {
			return errors.New("no find slot")
		}

		ok := user.GetSquad().Inventory.AddItem(slot.Item, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP)

		if ok {
			storages.Storages.RemoveItem(user.GetID(), user.InBaseID, storageSlot, slot.Quantity)
		}

		update.Squad(user.GetSquad())
		return nil
	} else {
		return errors.New("user not in base")
	}
}
