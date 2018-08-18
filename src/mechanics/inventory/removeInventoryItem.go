package inventory

import "../gameObjects/squad"

/* когда slot.Item = nil он удалиться из бд при обновление отряда */
func RemoveInventoryItem(quantityRemove int, slot *squad.InventorySlot) (CountRemove int) {
	if quantityRemove < slot.Quantity {
		slot.Quantity = slot.Quantity - quantityRemove
		return quantityRemove
	} else {
		slot.Item = nil
		return slot.Quantity
	}
}
