package market

import (
	"../db/market"
	"../db/updateSquad"
	"../player"
	"../storage"
	"errors"
)

func PlaceNewSellOrder(storageSlot, price, quantity, minBuyOut, expires int, user *player.Player) error {
	if user.InBaseID > 0 {

		baseStorage, _ := storage.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]
		if slot == nil {
			return errors.New("no find slot")
		}

		if slot.MaxHP == slot.HP && quantity <= slot.Quantity {

			// добавляем их в магазин
			market.PlaceNewOrder(user.GetID(), price, quantity, minBuyOut, slot.ItemID, expires, user.InBaseID,
				"sell", slot.Type, "База 1")
			// todo имя базы захарджкожено
			// todo если expires 0 то ставим месяц
			// todo если minBuyOut 0 то ставим 1

			// удаляем из инвентаря
			storage.Storages.RemoveItem(user.GetID(), user.InBaseID, storageSlot, quantity)

		} else {
			if slot.MaxHP > slot.HP {
				return errors.New("equip is damage")
			}
			if quantity > slot.Quantity {
				return errors.New("low amount")
			}
		}

		updateSquad.Squad(user.GetSquad())
		return nil
	} else {
		return errors.New("user not in base")
	}
}
