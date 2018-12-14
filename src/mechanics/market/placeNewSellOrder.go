package market

import (
	"../db/market"
	"../db/squad/update"
	"../factories/storages"
	"../gameObjects/order"
	"../player"
	"errors"
	"time"
)

func (o *OrdersPool) PlaceNewSellOrder(storageSlot, price, quantity, minBuyOut, expires int, user *player.Player) error {
	// todo имя базы захарджкожено
	// todo если expires 0 то ставим 1 месяц
	// todo expires в днях, надо брать текущее время  + expires и это значение класть в бд

	if user.InBaseID > 0 {

		baseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]
		if slot == nil {
			return errors.New("no find slot")
		}

		if slot.MaxHP == slot.HP && quantity <= slot.Quantity {

			if minBuyOut == 0 {
				minBuyOut = 1
			}

			// смотрит есть ли на рынке итемы которые уже подходят под условия цены, и если есть покупаем
			for _, marketOrder := range o.orders {
				if marketOrder.IdItem == slot.ItemID && marketOrder.TypeItem == slot.Type && marketOrder.Price >= price {
					if marketOrder.Count > quantity {
						err := o.Sell(marketOrder.Id, quantity, user)
						if err == nil {
							return nil // т.к. мы продали все итемы то ордер создавать не ненадо
						}
					} else {
						err := o.Sell(marketOrder.Id, marketOrder.Count, user)
						if err == nil {
							quantity -= marketOrder.Count
						}
					}
				}
			}

			newOrder := order.Order{IdUser: user.GetID(), Price: price, Count: quantity, Type: "sell",
				MinBuyOut: minBuyOut, TypeItem: slot.Type, IdItem: slot.ItemID, Expires: time.Now(),
				PlaceName: "База 1", PlaceID: user.InBaseID, Item: slot.Item}

			// добавляем их в магазин
			id := market.PlaceNewOrder(&newOrder)

			if id > 0 {
				newOrder.Id = id
				// добавляем его на фабрику
				o.AddNewOrder(newOrder)
				// удаляем итем со склада
				storages.Storages.RemoveItem(user.GetID(), user.InBaseID, storageSlot, quantity)
			} else {
				return errors.New("unknown error")
			}
		} else {
			if slot.MaxHP > slot.HP {
				return errors.New("equip is damage")
			}

			if slot.MaxHP < slot.HP {
				return errors.New("wrong hp")
			}

			if quantity > slot.Quantity {
				return errors.New("low amount")
			}
		}

		update.Squad(user.GetSquad())
		return nil
	} else {
		return errors.New("user not in base")
	}
}
