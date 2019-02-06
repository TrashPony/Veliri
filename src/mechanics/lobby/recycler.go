package lobby

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

type RecycleItem struct {
	Slot     *inventory.Slot `json:"slot"`
	Recycled bool            `json:"recycled"` // если false то итем не плавится
}

// я хз как еще это нормально назвать)
type Resourcer interface {
	GetEnrichedThorium() int
	GetIron() int
	GetCopper() int
	GetTitanium() int
	GetSilicon() int
	GetPlastic() int
	GetSteel() int
	GetWire() int
	GetElectronics() int
}

func Recycle(user *player.Player, recycleItems *map[int]*RecycleItem) error {

	storage, ok := storages.Storages.Get(user.GetID(), user.InBaseID)
	if !ok {
		return errors.New("user no base")
	}

	// удаляем все слоты которые смогли переработать, с последующием добавление новых итемов на склад
	for i, item := range *recycleItems {
		slot, ok := storage.Slots[i]
		if ok && item.Recycled && item.Slot.Quantity == slot.Quantity {

			// небольшой костыль связаный с принимаемым типом )
			recycleItem := make(map[int]*RecycleItem)
			recycleItem[i] = &RecycleItem{Slot: slot}

			possibleRecycleItems := GetRecycleItems(&recycleItem)
			for _, item := range possibleRecycleItems {
				storages.Storages.AddSlot(user.GetID(), user.InBaseID, item)
			}

			delete(*recycleItems, i)
			storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, i, slot.Quantity)
		}
	}

	return nil
}

func GetRecycleItems(recycleItems *map[int]*RecycleItem) []*inventory.Slot {

	// TODO вычеслять процент переработки, уровень базы + скил игрока

	recyclerItems := make([]*inventory.Slot, 0)

	for _, item := range *recycleItems {
		// обычные ресурсы имеют прямое преобразование в recycle ресурсы
		if item.Slot.Type == "resource" {
			res, ok := gameTypes.Resource.GetBaseByID(item.Slot.ItemID)
			if ok {
				item.Recycled = true
				ParseItems(&recyclerItems, 50, res, item.Slot.Quantity)
			}
		}

		// recycle ресурсы не могут быть обработаны
		if item.Slot.Type == "recycle" {
			continue
		}

		// если итем поврежден то его нельзя расплавить
		if item.Slot.HP < item.Slot.MaxHP {
			continue
		}

		// для всех остальных сущьностей, надо брать рецепт и смотреть ресурсы в нем
		if item.Slot.Type == "body" || item.Slot.Type == "ammo" || item.Slot.Type == "equip" ||
			item.Slot.Type == "weapon" || item.Slot.Type == "boxes" || item.Slot.Type == "detail" {
			bluePrint := gameTypes.BluePrints.GetByItemTypeAndID(item.Slot.ItemID, item.Slot.Type)
			if bluePrint != nil {
				item.Recycled = true
				ParseItems(&recyclerItems, 50, bluePrint, item.Slot.Quantity)
			}
		}
	}
	return recyclerItems
}

func ParseItems(itemsPool *[]*inventory.Slot, percent int, resourcer Resourcer, countItems int) {

	var appendPrimitiveRes = func(res *resource.RecycledResource, count int) {

		realCount := ((count * percent) / 100) * countItems

		for _, item := range *itemsPool {
			if item.ItemID == res.TypeID && item.Type == "recycle" {
				item.Quantity += realCount
				return
			}
		}

		*itemsPool = append(*itemsPool, &inventory.Slot{
			Item:     res,
			Quantity: realCount,
			Type:     "recycle",
			ItemID:   res.TypeID,
			HP:       1,
			MaxHP:    1,
			Size:     res.Size * float32(realCount),
		})
	}

	var appendDetail = func(detail *resource.CraftDetail, count int) {

		realCount := ((count * percent) / 100) * countItems

		for _, item := range *itemsPool {
			if item.ItemID == detail.TypeID && item.Type == "detail" {
				item.Quantity += realCount
				return
			}
		}

		*itemsPool = append(*itemsPool, &inventory.Slot{
			Item:     detail,
			Quantity: realCount,
			Type:     "detail",
			ItemID:   detail.TypeID,
			HP:       1,
			MaxHP:    1,
			Size:     detail.Size * float32(realCount),
		})
	}

	//-- ПРИМИТИВЫ
	if resourcer.GetIron() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("iron"), resourcer.GetIron())
	}

	if resourcer.GetCopper() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("copper"), resourcer.GetCopper())
	}

	if resourcer.GetTitanium() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("titanium"), resourcer.GetTitanium())
	}

	if resourcer.GetSilicon() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("silicon"), resourcer.GetSilicon())
	}

	if resourcer.GetPlastic() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("plastic"), resourcer.GetPlastic())
	}

	if resourcer.GetEnrichedThorium() > 0 {
		appendPrimitiveRes(gameTypes.Resource.GetRecycledByName("enriched_thorium"), resourcer.GetEnrichedThorium())
	}

	//-- ДЕТАЛИ
	if resourcer.GetSteel() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("steel"), resourcer.GetSteel())
	}

	if resourcer.GetWire() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("wire"), resourcer.GetWire())
	}

	if resourcer.GetElectronics() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("electronics"), resourcer.GetElectronics())
	}
}
