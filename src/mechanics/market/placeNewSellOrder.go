package market

import (
	"../db/market"
	"../db/updateSquad"
	"../gameObjects/order"
	"../player"
	"../storage"
	"errors"
	"time"
)

func (o *OrdersPool) PlaceNewSellOrder(storageSlot, price, quantity, minBuyOut, expires int, user *player.Player) error {

	if user.InBaseID > 0 {

		baseStorage, _ := storage.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]
		if slot == nil {
			return errors.New("no find slot")
		}

		if slot.MaxHP == slot.HP && quantity <= slot.Quantity {

			newOrder := order.Order{IdUser: user.GetID(), Price: price, Count: quantity, Type: "sell",
				MinBuyOut: minBuyOut, TypeItem: slot.Type, IdItem: slot.ItemID, Expires: time.Now(),
				PlaceName: "База 1", PlaceID: user.InBaseID, Item: slot.Item}

			// todo имя базы захарджкожено
			// todo если expires 0 то ставим 1 месяц
			// todo если minBuyOut 0 то ставим 1
			// todo expires в днях, надо брать текущее время  + expires и это значение класть в бд

			// добавляем их в магазин
			id := market.PlaceNewOrder(&newOrder)

			if id > 0 {
				newOrder.Id = id
				// добавляем его на фабрику
				o.AddNewOrder(newOrder)
				// удаляем итем со склада
				storage.Storages.RemoveItem(user.GetID(), user.InBaseID, storageSlot, quantity)
			} else {
				return errors.New("unknown error")
			}
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
