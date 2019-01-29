package lobby

import (
	"../gameObjects/inventory"
	"../factories/gameTypes"
)

type RecycleItem struct {
	Slot     *inventory.Slot `json:"slot"`
	Recycled bool            `json:"recycled"` // если false то итем не плавится
}

type Resourcer interface {
	// я хз как еще это нормально назвать)
	GetEnrichedThorium() int
	GetIron() int
	GetCopper() int
	GetTitanium() int
	GetSilicon() int
	GetPlastic() int
	GetSteel() int
	GetWire() int
}

func GetRecycleItems(recycleItems *map[int]*RecycleItem) []*inventory.Slot {

	recyclerItems := make([]*inventory.Slot, 0)

	for _, item := range *recycleItems {
		// обычные ресурсы имеют прямое преобразование в recycle ресурсы
		if item.Slot.Type == "resource" {
			item.Recycled = true
			// TODO parsing
		}
		// recycle ресурсы не могут быть обработаны
		if item.Slot.Type == "recycle" {
			item.Recycled = false
			continue
		}

		// для всех остальных сущьностей, надо брать рецепт и смотреть ресурсы в нем
		if item.Slot.Type == "body" || item.Slot.Type == "ammo" || item.Slot.Type == "equip" ||
			item.Slot.Type == "weapon" || item.Slot.Type == "boxes" || item.Slot.Type == "detail" {
			bluePrint := gameTypes.BluePrints.GetByItemTypeAndID(item.Slot.ItemID, item.Slot.Type)
			if bluePrint != nil {
				// TODO parsing
			} else {
				item.Recycled = false
			}
		}
	}
	return recyclerItems
}

func ParseItems(itemsPool *[]*inventory.Slot, percent int, resourcer Resourcer) {

	if resourcer.GetEnrichedThorium() > 0 {
		// todo добавление в itemsPool новый итем
	}

	if resourcer.GetIron() > 0 {

	}

	if resourcer.GetCopper() > 0 {

	}

	if resourcer.GetTitanium() > 0 {

	}

	if resourcer.GetSilicon() > 0 {

	}

	if resourcer.GetPlastic() > 0 {

	}

	if resourcer.GetSteel() > 0 {

	}

	if resourcer.GetWire() > 0 {

	}
}
