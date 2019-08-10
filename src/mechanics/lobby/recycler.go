package lobby

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
)

type RecycleItem struct {
	Slot       *inventory.Slot `json:"slot"`
	Recycled   bool            `json:"recycled"` // если false то итем не плавится
	TaxPercent int             `json:"tax_percent"`
	Source     string          `json:"source"`
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
	GetWires() int
	GetGear() int
	GetTitaniumPlate() int
	GetBatteries() int
	GetArmorItems() int
}

func Recycle(user *player.Player, recycleItems *map[string]map[int]*RecycleItem, gameBase *base.Base) error {

	storage, ok := storages.Storages.Get(user.GetID(), user.InBaseID)
	if !ok {
		return errors.New("user no base")
	}

	// удаляем все слоты которые смогли переработать, с последующием добавление новых итемов на склад

	for sourceKey, source := range *recycleItems {
		for i, item := range source { // source: storage, squadInventory

			// проверка перед переработкой на наличие всех айтемов
			var slot *inventory.Slot
			ok := false

			if sourceKey == "storage" {
				slot, ok = storage.Slots[i]
			}

			if sourceKey == "squadInventory" {
				slot, ok = user.GetSquad().Inventory.Slots[i]
			}

			if ok && item.Recycled && item.Slot.Quantity == slot.Quantity {

				// небольшой костыль связаный с принимаемым типом )
				recycleItem := make(map[string]map[int]*RecycleItem)
				recycleItem[sourceKey] = make(map[int]*RecycleItem)
				recycleItem[sourceKey][i] = &RecycleItem{Slot: slot}

				possibleRecycleItems, baseTaxItems := GetRecycleItems(&recycleItem, user, gameBase)

				// отдаем игроку его итемы, обязательно на склад
				for _, slot := range possibleRecycleItems {
					storages.Storages.AddSlot(user.GetID(), user.InBaseID, slot)
				}

				// отдаем налог базе
				for _, slot := range baseTaxItems {
					if gameBase.CurrentResources[slot.ItemID] != nil && slot.Type == "recycle" && slot.ItemID == gameBase.CurrentResources[slot.ItemID].ItemID {
						gameBase.CurrentResources[slot.ItemID].Quantity += slot.Quantity
					}
				}

				delete(source, i)

				// удаляем переработаные итемы из инвентаря источника
				if sourceKey == "storage" {
					storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, i, slot.Quantity)
				}
				if sourceKey == "squadInventory" {
					// мы уже сделали проверку на наичие и в теории можем не беспокоится об ошибке
					user.GetSquad().Inventory.RemoveItem(slot.ItemID, slot.Type, slot.Quantity)
				}
			}
		}
	}

	return nil
}

func GetRecycleItems(recycleItems *map[string]map[int]*RecycleItem, client *player.Player, gameBase *base.Base) ([]*inventory.Slot, []*inventory.Slot) {
	recyclerItems := make([]*inventory.Slot, 0)
	inBaseItems := make([]*inventory.Slot, 0)

	for _, source := range *recycleItems {
		for _, item := range source { // source: storage, squadInventory

			// обычные ресурсы имеют прямое преобразование в recycle ресурсы
			if item.Slot.Type == "resource" {

				// потери при переработки из за недостатка скила
				percentUserSkills := 25 - client.CurrentSkills["processing"].Level*5

				res, ok := gameTypes.Resource.GetBaseByID(item.Slot.ItemID)
				if ok {
					item.Recycled = true
					item.Slot.Tax = (item.Slot.Quantity / 100) * gameBase.GetRecyclePercent(item.Slot.ItemID)
					item.TaxPercent = gameBase.GetRecyclePercent(item.Slot.ItemID)

					ParseItems(&recyclerItems, 100-(percentUserSkills+gameBase.GetRecyclePercent(item.Slot.ItemID)), res, item.Slot.Quantity)
					ParseItems(&inBaseItems, gameBase.GetRecyclePercent(item.Slot.ItemID), res, item.Slot.Quantity)
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

				percentUserSkills := 25 - client.CurrentSkills["processing"].Level*5

				bluePrint := gameTypes.BluePrints.GetByItemTypeAndID(item.Slot.ItemID, item.Slot.Type)
				if bluePrint != nil {

					item.Recycled = true
					//item.ProductionLossPercent = percentUserSkills + gameBase.GetRecyclePercent(item.Slot.ItemID)

					ParseItems(&recyclerItems, 100-percentUserSkills, bluePrint, item.Slot.Quantity)
				}
			}
		}
	}
	return recyclerItems, inBaseItems
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

	if resourcer.GetWires() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("wires"), resourcer.GetWires())
	}

	if resourcer.GetGear() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("gear"), resourcer.GetGear())
	}

	if resourcer.GetTitaniumPlate() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("titanium_plate"), resourcer.GetTitaniumPlate())
	}

	if resourcer.GetBatteries() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("batteries"), resourcer.GetBatteries())
	}

	if resourcer.GetArmorItems() > 0 {
		appendDetail(gameTypes.Resource.GetDetailByName("armor_items"), resourcer.GetArmorItems())
	}
}
