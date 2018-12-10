package squadInventory

import (
	"../db/updateSquad"
	"../player"
	"../storage"
	"errors"
)

func ItemToStorage(user *player.Player, inventorySlot int) error {
	if user.InBaseID > 0 {
		slot := user.GetSquad().Inventory.Slots[inventorySlot]

		if slot == nil {
			return errors.New("no find slot")
		}

		ok := storage.Storages.AddItem(user.GetID(), user.InBaseID, slot.Item, slot.Type, slot.ItemID, slot.Quantity,
			slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP)

		if ok {
			user.GetSquad().Inventory.Slots[inventorySlot].RemoveItem(slot.Quantity)
		}

		updateSquad.Squad(user.GetSquad())
		return nil
	} else {
		return errors.New("user not in base")
	}
}
