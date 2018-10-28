package inventory

import "../gameObjects/squad"

/* когда slot.Item = nil он удалиться из бд при обновление отряда */
func RemoveInventoryItem(quantityRemove int, slot *squad.InventorySlot) (CountRemove int) {
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
