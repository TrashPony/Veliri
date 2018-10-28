package inventory

import "../gameObjects/squad"

func AddItem(inventory map[int]*squad.InventorySlot, item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32) bool {

	for _, slot := range inventory { // ищем стопку с такими же элементами
		if slot.ItemID == itemID && slot.Type == itemType && slot.HP == hp && slot.Item != nil {
			slot.Quantity = slot.Quantity + quantity
			slot.Size = slot.Size + (itemSize * float32(quantity))
			return true
		}
	}

	for i := 1; i <= 40; i++ { // ищем пустой слот
		_, ok := inventory[i]
		if !ok {
			newItem := squad.InventorySlot{Item: item, Type: itemType, ItemID: itemID, InsertToDB: true,
				Quantity: quantity, HP: hp, Size: itemSize * float32(quantity)}
			inventory[i] = &newItem
			return true
		}
	}

	return false
}
