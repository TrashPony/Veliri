package squadInventory

import (
	"../player"
	"../storage"
	"errors"
)

func ItemToStorage(user *player.Player, inventorySlot int) error {
	if user.InBaseID > 0 {
		slot := user.GetSquad().Inventory.Slots[inventorySlot]

		storage.AddNewItem(user, user.InBaseID, slot)
		user.GetSquad().Inventory.Slots[inventorySlot].RemoveItem(slot.Quantity)

		return nil
	} else {
		return errors.New("user not in base")
	}
}
