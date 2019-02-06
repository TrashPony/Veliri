package market

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/market"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"time"
)

func (o *OrdersPool) PlaceNewBuyOrder(itemId, price, quantity, minBuyOut, expires int, itemType string, user *player.Player) error {

	// todo если expires 0 то ставим 14 дней
	// todo expires в днях, надо брать текущее время  + expires и это значение класть в бд

	base, findBase := bases.Bases.Get(user.InBaseID)

	if user.GetCredits() >= price*quantity && (minBuyOut == 0 || quantity%minBuyOut == 0) && findBase {

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
				if marketOrder.IdItem == itemId && marketOrder.TypeItem == itemType &&
					marketOrder.Price <= price && base.ID == marketOrder.PlaceID {
					if marketOrder.Count > quantity {
						err := o.Buy(marketOrder.Id, quantity, user)
						if err == nil {
							return nil // т.к. мы купили все итемы то ордер создавать не ненадо
						}
					} else {
						if marketOrder.Count%minBuyOut == 0 {
							countOrder := marketOrder.Count // т.к. при продаже удалиться count у ордера, сохраним его
							err := o.Buy(marketOrder.Id, marketOrder.Count, user)
							if err == nil {
								quantity -= countOrder
							}
						} else {
							countOrder := marketOrder.Count - marketOrder.Count%minBuyOut
							err := o.Buy(marketOrder.Id, marketOrder.Count-marketOrder.Count%minBuyOut, user)
							if err == nil {
								quantity -= countOrder
							}
						}
					}
				}
			}

			newOrder := order.Order{IdUser: user.GetID(), Price: price, Count: quantity, Type: "buy",
				MinBuyOut: minBuyOut, TypeItem: itemType, IdItem: itemId, Expires: time.Now(),
				PlaceName: base.Name, PlaceID: base.ID, Item: item}

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
		if user.GetCredits() < price*quantity {
			return errors.New("no credits")
		}
		if minBuyOut != 0 && quantity%minBuyOut != 0 {
			return errors.New("wrong count")
		}
		if !findBase {
			return errors.New("wrong base")
		}
	}

	return nil
}
