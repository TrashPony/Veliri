package inventory

import (
	"database/sql"
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
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

	PlaceUserID int  `json:"place_user_id"`
	Tax         int  `json:"tax"`  // поле для налогов
	Find        bool `json:"find"` // поле для верстака, обозначающие естли такое количество итемов на складе или нет
}

func (inv *Inventory) SetSlotsSize(size int) {
	inv.size = size
}

func (inv *Inventory) AddItemFromSlot(slot *Slot, userID int) bool {

	if slot.Quantity <= 0 { // slot.Size/float32(slot.Quantity) деление на ноль все сломает
		return false
	}

	return inv.AddItem(slot.Item, slot.Type, slot.ItemID, slot.Quantity,
		slot.HP, slot.Size/float32(slot.Quantity), slot.MaxHP, false, userID)
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

func (inv *Inventory) AddItem(item interface{}, itemType string, itemID, quantity, hp int, itemSize float32, maxHP int, newSlot bool, userID int) bool {

	//newSlot говорит о том что этот айтем надо положить в строго новый слот
	if !newSlot {
		for _, slot := range inv.Slots { // ищем стопку с такими же элементами
			if slot.ItemID == itemID && slot.Type == itemType && slot.HP == hp && slot.Item != nil {
				slot.Quantity = slot.Quantity + quantity
				slot.Size = slot.Size + (itemSize * float32(quantity))
				slot.PlaceUserID = userID
				return true
			}
		}
	}

	newNumberSlot := inv.GetEmptySlot()
	if newNumberSlot > 0 {
		newItem := Slot{Item: item, Type: itemType, ItemID: itemID, InsertToDB: true,
			Quantity: quantity, HP: hp, MaxHP: maxHP, Size: itemSize * float32(quantity), PlaceUserID: userID}
		inv.Slots[newNumberSlot] = &newItem
		return true
	}

	return false
}

func (inv *Inventory) GetEmptySlot() int {
	for i := 1; i <= inv.size; i++ { // ищем пустой слот
		_, ok := inv.Slots[i]
		if !ok {
			return i
		}
	}
	return 0
}

func (inv *Inventory) RemoveItem(itemID int, itemType string, quantityRemove int) error {
	// надо убедиться что необходимое количество есть и его можно удалить
	if inv.ViewItems(itemID, itemType, quantityRemove) {

		for _, slot := range inv.Slots {
			if slot.ItemID == itemID && slot.Type == itemType {
				if slot.Quantity >= quantityRemove {
					slot.RemoveItemBySlot(quantityRemove)
					return nil
				} else {
					// если в слоте не чего либо для полного удаления,
					// то удаляем все из слота, и уменьшаем количество итемов которые еще надо удалить
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

// метод делает сравнение инвентарей слот к слоту
func (inv *Inventory) ViewItemsBySlots(slots map[int]*Slot) bool {
	checkItems := true
	for number, slot := range slots {
		realSlot, findSlot := inv.Slots[number]
		if !findSlot || slot == nil || slot.Quantity > realSlot.Quantity {
			checkItems = false
		}
	}
	return checkItems
}

// метод смотрит все предметы inv2 что бы они были в inv на наличие
func (inv *Inventory) SearchItemsByOtherInventory(inv2 *Inventory) bool {
	for _, slot := range inv2.Slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.Quantity) {
			return false
		}
	}
	return true
}

// метод удаляем все итемы из inv которые есть в inv2 если они все в наличие
func (inv *Inventory) RemoveItemsByOtherInventory(inv2 *Inventory) bool {
	for _, slot := range inv2.Slots {
		if !inv.ViewItems(slot.ItemID, slot.Type, slot.Quantity) {
			return false
		}
	}

	for _, slot := range inv2.Slots {
		inv.RemoveItem(slot.ItemID, slot.Type, slot.Quantity)
	}
	return true
}

// метод считает все итемы в инвентаре
func (inv *Inventory) ViewQuantityItems(itemID int, itemType string) int {
	countRealItems := 0
	for _, slot := range inv.Slots {
		if slot.ItemID == itemID && slot.Type == itemType {
			countRealItems += slot.Quantity
		}
	}

	return countRealItems
}

// метод смотрим естли необходимое количество предметов в инвентаре
func (inv *Inventory) ViewItems(itemID int, itemType string, quantityFind int) bool {
	if inv.ViewQuantityItems(itemID, itemType) >= quantityFind {
		return true
	} else {
		return false
	}
}

func (slot *Slot) AddItemBySlot(quantity, userID int) {
	// определяем вес 1 вещи
	sizeOneItem := slot.Size / float32(slot.Quantity)
	slot.Quantity += quantity
	slot.PlaceUserID = userID
	// находим новый вес для всей стопки
	slot.Size = sizeOneItem * float32(slot.Quantity)
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
		slot.Quantity = 0
		slot.Size = 0
		return slot.Quantity
	}
}

func (inv *Inventory) FillInventory(rows *sql.Rows) {
	inv.SetSlotsSize(9999) // у всех инвентарей 9999 ячеек
	for rows.Next() {

		var inventorySlot = Slot{}
		var slot int

		err := rows.Scan(&slot, &inventorySlot.Type, &inventorySlot.ItemID, &inventorySlot.Quantity, &inventorySlot.HP, &inventorySlot.PlaceUserID)
		if err != nil {
			log.Fatal("scan inventory slots " + err.Error())
		}

		FillSlot(&inventorySlot, slot, inv, false)
	}
}

func FillSlot(inventorySlot *Slot, slot int, inv *Inventory, maxHP bool) {
	if inventorySlot.Type == "weapon" {
		weapon, _ := gameTypes.Weapons.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = weapon
		inventorySlot.Size = weapon.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = weapon.MaxHP

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "ammo" {
		ammo, _ := gameTypes.Ammo.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = ammo
		inventorySlot.Size = ammo.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у аммо нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "equip" {
		equip, _ := gameTypes.Equips.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = equip
		inventorySlot.Size = equip.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = equip.MaxHP

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "body" {
		body, _ := gameTypes.Bodies.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = body
		inventorySlot.Size = body.CapacitySize * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = body.MaxHP

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "resource" {
		resource, _ := gameTypes.Resource.GetBaseByID(inventorySlot.ItemID)

		inventorySlot.Item = resource
		inventorySlot.Size = resource.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у ресов нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "recycle" {
		resource, _ := gameTypes.Resource.GetRecycledByID(inventorySlot.ItemID)

		inventorySlot.Item = resource
		inventorySlot.Size = resource.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у ресов нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "detail" {
		detail, _ := gameTypes.Resource.GetDetailByID(inventorySlot.ItemID)

		inventorySlot.Item = detail
		inventorySlot.Size = detail.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у ящиков тож нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "boxes" {
		box, _ := gameTypes.Boxes.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = box
		inventorySlot.Size = box.FoldSize * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у ящиков тож нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "blueprints" {
		blueprint, _ := gameTypes.BluePrints.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = blueprint
		inventorySlot.Size = 0  // чертежи не занимают места
		inventorySlot.MaxHP = 1 // у чертежов нет хп

		inv.Slots[slot] = inventorySlot
	}

	if inventorySlot.Type == "trash" {
		trashItem, _ := gameTypes.TrashItems.GetByID(inventorySlot.ItemID)

		inventorySlot.Item = trashItem
		inventorySlot.Size = trashItem.Size * float32(inventorySlot.Quantity)
		inventorySlot.MaxHP = 1 // у мусора нет хп

		inv.Slots[slot] = inventorySlot
	}

	if maxHP {
		inventorySlot.HP = inventorySlot.MaxHP
	}
}
