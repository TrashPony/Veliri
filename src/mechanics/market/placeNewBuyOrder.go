package market

import (
	"../db/market"
	dbPlayer "../db/player"
	"../factories/gameTypes"
	"../gameObjects/order"
	"../player"
	"errors"
	"time"
)

func (o *OrdersPool) PlaceNewBuyOrder(itemId, price, quantity, minBuyOut, expires int, itemType string, user *player.Player) error {

	// todo имя базы захарджкожено
	// todo если expires 0 то ставим 1 месяц
	// todo expires в днях, надо брать текущее время  + expires и это значение класть в бд

	if user.GetCredits() >= price*quantity {

		var item interface{}
		var ok bool

		if itemType == "weapon" {
			item, ok = gameTypes.Weapons.GetByID(itemId)
		}

		if itemType == "ammo" {
			item, ok = gameTypes.Ammo.GetByID(itemId)
		}

		if itemType == "equip" {
			item, ok = gameTypes.Equips.GetByID(itemId)
		}

		if itemType == "body" {
			item, ok = gameTypes.Bodies.GetByID(itemId)
		}

		if ok {

			if minBuyOut == 0 {
				minBuyOut = 1
			}

			// смотрит есть ли на рынке итемы которые уже подходят под условия цены, и если есть покупаем
			for _, marketOrder := range o.orders {
				if marketOrder.IdItem == itemId && marketOrder.TypeItem == itemType && marketOrder.Price <= price {
					if marketOrder.Count > quantity {
						err := o.Buy(marketOrder.Id, quantity, user)
						if err == nil {
							return nil // т.к. мы купили все итемы то ордер создавать не ненадо
						}
					} else {
						if marketOrder.Count%minBuyOut == 0 {
							err := o.Buy(marketOrder.Id, marketOrder.Count, user)
							if err == nil {
								quantity -= marketOrder.Count
							}
						} else {
							err := o.Buy(marketOrder.Id, marketOrder.Count-marketOrder.Count%minBuyOut, user)
							if err == nil {
								quantity -= marketOrder.Count - marketOrder.Count%minBuyOut
							}
						}
					}
				}
			}

			newOrder := order.Order{IdUser: user.GetID(), Price: price, Count: quantity, Type: "buy",
				MinBuyOut: minBuyOut, TypeItem: itemType, IdItem: itemId, Expires: time.Now(),
				PlaceName: "База 1", PlaceID: user.InBaseID, Item: item}

			// добавляем ордер в магазин
			id := market.PlaceNewOrder(&newOrder)

			if id > 0 {
				newOrder.Id = id

				// добавляем ордер на фабрику
				o.AddNewOrder(newOrder)
				// отнимаем деньги :)
				user.SetCredits(user.GetCredits() - price*quantity)
				dbPlayer.UpdateUser(user)
			}
		} else {
			return errors.New("no item")
		}
	} else {
		return errors.New("no credits")
	}

	return nil
}
