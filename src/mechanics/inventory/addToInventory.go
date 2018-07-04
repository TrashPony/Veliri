package inventory

import "../gameObjects/squad"

func AddItem(inventory map[int]*squad.InventorySlot, item interface{}, itemType string, itemID int, quantity int) bool {

	for _, slot := range inventory { // ищем стопку с такими же элементами
		if slot.ItemID == itemID && slot.Type == itemType {
			slot.Quantity++
			return true
		}
	}

	for i := 0; i < 40; i++ { // ищем пустой слот
		_, ok := inventory[i]
		if !ok {
			newItem := squad.InventorySlot{Item: item, Type: itemType, ItemID: itemID, InsertToDB:true, Quantity: quantity}
			inventory[i] = &newItem
		}
	}

	return false
}
