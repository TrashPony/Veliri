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
