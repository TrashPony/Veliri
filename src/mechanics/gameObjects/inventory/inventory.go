package inventory

import (
	"../../factories/gameTypes"
	"database/sql"
	"errors"
	"log"
	"math"
)

type Inventory struct {
	Slots map[int]*Slot `json:"slots"`
	size  int
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

func (inv *Inventory) SetSlotsSize(size int) {
	inv.size = size
}

func (inv *Inventory) AddItemFromSlot(slot *Slot) bool {
	ok := inv.AddItem(slot.Item, slot.Type, slot.ItemID, slot.Quantity,
		slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP)

	return ok
}

func (inv *Inventory) GetSize() float32 {
	var inventorySquadSize float32
	for _, slot := range inv.Slots {
		if slot.Item != nil {
			inventorySquadSize = inventorySquadSize + slot.Size
		}
	}

	return float32(round(float64(inventorySquadSize), 1))
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func (inv *Inventory) AddItem(item interface{}, itemType string, itemID int, quantity int, hp int, itemSize float32, maxHP int) bool {

	for _, slot := range inv.Slots { // ищем стопку с такими же элементами
		if slot.ItemID == itemID && slot.Type == itemType && slot.HP == hp && slot.Item != nil {
			slot.Quantity = slot.Quantity + quantity
			slot.Size = slot.Size + (itemSize * float32(quantity))
			return true
		}
	}

	for i := 1; i <= inv.size; i++ { // ищем пустой слот
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

func (inv *Inventory) RemoveItem(itemID int, itemType string, quantityRemove int) error {
	// надо убедиться что необходимое количество есть и его можно удалить
	countRealItems := 0

	for _, slot := range inv.Slots {
		if slot.ItemID == itemID && slot.Type == itemType {
			countRealItems += slot.Quantity
		}
	}

	if countRealItems >= quantityRemove {
		for _, slot := range inv.Slots {
			if slot.ItemID == itemID && slot.Type == itemType {
				if slot.Quantity > quantityRemove {
					slot.RemoveItemBySlot(quantityRemove)
					return nil
				} else {
					quantityRemove -= slot.Quantity
					slot.RemoveItemBySlot(slot.Quantity)
				}
			}
		}
		return nil
	} else {
		return errors.New("few items")
	}
}

/* когда slot.Item = nil он удалиться из бд при обновление данных */
func (slot *Slot) RemoveItemBySlot(quantityRemove int) (CountRemove int) {
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

		if inventorySlot.Type == "resource" {
			resource, _ := gameTypes.Resource.GetBaseByID(inventorySlot.ItemID)

			inventorySlot.Item = resource
			inventorySlot.Size = resource.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = 1 // у ресов нет хп

			inv.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "recycle" {
			resource, _ := gameTypes.Resource.GetRecycledByID(inventorySlot.ItemID)

			inventorySlot.Item = resource
			inventorySlot.Size = resource.Size * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = 1 // у ресов нет хп

			inv.Slots[slot] = &inventorySlot
		}

		if inventorySlot.Type == "boxes" {
			resource, _ := gameTypes.Boxes.GetByID(inventorySlot.ItemID)

			inventorySlot.Item = resource
			inventorySlot.Size = resource.FoldSize * float32(inventorySlot.Quantity)
			inventorySlot.MaxHP = 1 // у ящиков тож нет хп

			inv.Slots[slot] = &inventorySlot
		}
	}
}
