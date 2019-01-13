package squadInventory

import (
	"../factories/storages"
	"../gameObjects/inventory"
	"../player"
)

func getSlotBySource(user *player.Player, inventorySlot int, source string) *inventory.Slot {
	var slot *inventory.Slot

	if source == "squadInventory" {
		slot = user.GetSquad().Inventory.Slots[inventorySlot]
	}

	if source == "storage" {
		storage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
		slot = storage.Slots[inventorySlot]
	}

	return slot
}

func RemoveSlotBySource(user *player.Player, inventorySlot int, source string, quantity int) int {
	if source == "squadInventory" {
		return user.GetSquad().Inventory.Slots[inventorySlot].RemoveItemBySlot(quantity)
	}

	if source == "storage" {
		_, countRemove := storages.Storages.RemoveItem(user.GetID(), user.InBaseID, inventorySlot, quantity)
		return countRemove
	}

	return 0
}
