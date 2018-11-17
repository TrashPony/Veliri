package inventory

type Inventory struct {
	Slots map[int]*Slot `json:"slots"`
}

type Slot struct {
	Item       interface{} `json:"item"`
	Quantity   int         `json:"quantity"`
	Type       string      `json:"type"`
	ItemID     int         `json:"item_id"`
	InsertToDB bool        `json:"insert_to_db"`
	HP         int         `json:"hp"`
	Size       float32     `json:"size"`
}

func (inv *Inventory) AddItem(item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32) bool {

	for _, slot := range inv.Slots { // ищем стопку с такими же элементами
		if slot.ItemID == itemID && slot.Type == itemType && slot.HP == hp && slot.Item != nil {
			slot.Quantity = slot.Quantity + quantity
			slot.Size = slot.Size + (itemSize * float32(quantity))
			return true
		}
	}

	for i := 1; i <= 40; i++ { // ищем пустой слот
		_, ok := inv.Slots[i]
		if !ok {
			newItem := Slot{Item: item, Type: itemType, ItemID: itemID, InsertToDB: true,
				Quantity: quantity, HP: hp, Size: itemSize * float32(quantity)}
			inv.Slots[i] = &newItem
			return true
		}
	}

	return false
}

/* когда slot.Item = nil он удалиться из бд при обновление данных */
func (slot *Slot) RemoveItem(quantityRemove int) (CountRemove int) {
	if quantityRemove < slot.Quantity {
		// определяем вес 1 вещи
		itemSize := slot.Size / float32(slot.Quantity)
		// отнимает вес по количеству предметов
		slot.Size = slot.Size - (itemSize * float32(quantityRemove))
		// отнимаем количество итемов
		slot.Quantity = slot.Quantity - quantityRemove
		return quantityRemove
	} else {
		slot.Item = nil
		return slot.Quantity
	}
}
