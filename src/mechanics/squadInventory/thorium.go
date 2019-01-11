package squadInventory

import (
	"../db/squad/update"
	"../factories/gameTypes"
	"../player"
	"errors"
)

func SetThorium(user *player.Player, nubInventorySlot, numThoriumSlot int) error {

	thoriumSlot, _ := user.GetSquad().MatherShip.Body.ThoriumSlots[numThoriumSlot]
	inventorySlot, _ := user.GetSquad().Inventory.Slots[nubInventorySlot]

	// торий это ресурс с ид 1 и типом "recycle"
	if thoriumSlot != nil && inventorySlot != nil && inventorySlot.ItemID == 1 && inventorySlot.Item != nil && inventorySlot.Type == "recycle" {
		needThorium := thoriumSlot.MaxCount - thoriumSlot.Count

		if needThorium <= inventorySlot.Quantity {
			thoriumSlot.Count += needThorium
			inventorySlot.RemoveItemBySlot(needThorium)
		} else {
			thoriumSlot.Count += inventorySlot.Quantity
			inventorySlot.RemoveItemBySlot(inventorySlot.Quantity)
		}

		go update.Squad(user.GetSquad(), true)
	} else {
		return errors.New("no find slot")
	}

	return errors.New("unknown error")
}

func RemoveThorium(user *player.Player, numThoriumSlot int) error {

	thoriumSlot, _ := user.GetSquad().MatherShip.Body.ThoriumSlots[numThoriumSlot]

	// торий это ресурс с ид 1 и типом "recycle"
	if thoriumSlot != nil && thoriumSlot.Count > 0 {
		item, _ := gameTypes.Resource.GetRecycledByID(1)

		user.GetSquad().Inventory.AddItem(item, "recycle", 1, thoriumSlot.Count, 1,
			item.Size/float32(thoriumSlot.Count), 1)

		thoriumSlot.Count = 0
		go update.Squad(user.GetSquad(), true)
	} else {
		return errors.New("no thorium")
	}

	return errors.New("unknown error")
}
