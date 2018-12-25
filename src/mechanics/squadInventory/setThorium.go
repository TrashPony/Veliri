package squadInventory

import (
	"../player"
	"errors"
)

func SetThorium(user *player.Player, nubInventorySlot, numThoriumSlot int) error {

	thoriumSlot, _ := user.GetSquad().MatherShip.Body.ThoriumSlots[numThoriumSlot]
	inventorySlot, _ := user.GetSquad().Inventory.Slots[nubInventorySlot]

	// торий это ресурс с ид 1 и типом "recycle"
	if thoriumSlot != nil && inventorySlot != nil && inventorySlot.ItemID == 1 && inventorySlot.Item != nil && inventorySlot.Type == "recycle" {
		// TODO вычслить сколько нехватает до максимума, установить и удалить из инвентаря
	} else {
		return errors.New("no find slot")
	}

	return errors.New("unknown error")
}
