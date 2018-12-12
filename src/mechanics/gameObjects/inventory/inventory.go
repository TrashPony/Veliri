package inventory

import (
	"../../factories/gameTypes"
	"database/sql"
	"log"
)

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
	MaxHP      int         `json:"max_hp"`
	Size       float32     `json:"size"`
}

func (inv *Inventory) AddItem(item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32, maxHP int) bool {

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
				Quantity: quantity, HP: hp, MaxHP: maxHP, Size: itemSize * float32(quantity)}
			inv.Slots[i] = &newItem
			return true
		}
	}

	return false
}

/* когда slot.Item = nil он удалиться из бд при обновление данных */
// TODO переделать метод удаления на инвентарь целиком что бы была возможность удалять итемы из много столов сразу
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

func (inv *Inventory) FillInventory(rows *sql.Rows) {
	for rows.Next() {

		var inventorySlot = Slot{}
		var slot int

		err := rows.Scan(&slot, &inventorySlot.Type, &inventorySlot.ItemID, &inventorySlot.Quantity, &inventorySlot.HP)
		if err != nil {
			log.Fatal("scan inventory slots " + err.Error())
		}

		if inventorySlot.Type == "weapon" {
			weapon, _ := gameTypes.Weapons.GetByID(inventorySlot.ItemID)

			inventorySlot.Item = weapon
			inventorySlot.Size = weapon.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = weapon.MaxHP

			inv.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "ammo" {
			ammo, _ := gameTypes.Ammo.GetByID(inventorySlot.ItemID)

			inventorySlot.Item = ammo
			inventorySlot.Size = ammo.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = 1 // у аммо нет хп

			inv.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "equip" {
			equip, _ := gameTypes.Equips.GetByID(inventorySlot.ItemID)

			inventorySlot.Item = equip
			inventorySlot.Size = equip.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = equip.MaxHP

			inv.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "body" {
			body, _ := gameTypes.Bodies.GetByID(inventorySlot.ItemID)

			inventorySlot.Item = body
			inventorySlot.Size = body.CapacitySize * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = body.MaxHP

			inv.Slots[slot] = &inventorySlot
		}
	}
}
