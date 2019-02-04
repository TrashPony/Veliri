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

		if user.GetSquad().MatherShip.Body != nil && user.GetSquad().MatherShip.Body.CapacitySize >= user.GetSquad().Inventory.GetSize()+slot.Size {
			ok := user.GetSquad().Inventory.AddItem(slot.Item, slot.Type, slot.ItemID, slot.Quantity, slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP)
			if ok {
				storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, storageSlot, slot.Quantity)
				go update.Squad(user.GetSquad(), true)
				return nil
			} else {
				return errors.New("no free slots")
			}
		} else {
			return errors.New("weight exceeded")
		}
	} else {
		return errors.New("user not in base")
	}
}
