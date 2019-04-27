package market

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/market"
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"time"
)

func (o *OrdersPool) PlaceNewSellOrder(storageSlot, price, quantity, minBuyOut, expires int, user *player.Player) error {

	// todo если expires 0 то ставим 14 дней
	// todo expires в днях, надо брать текущее время  + expires и это значение класть в бд

	if user.InBaseID > 0 {

		baseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)

		slot := baseStorage.Slots[storageSlot]
		if slot == nil {
			return errors.New("no find slot")
		}

		base, findBase := bases.Bases.Get(user.InBaseID)

		if slot.MaxHP == slot.HP && quantity <= slot.Quantity && findBase {

			if minBuyOut == 0 {
				minBuyOut = 1
			}

			// смотрит есть ли на рынке итемы которые уже подходят под условия цены, и если есть покупаем
			for _, marketOrder := range o.orders {
				if marketOrder.IdItem == slot.ItemID && marketOrder.TypeItem == slot.Type &&
					marketOrder.Price >= price && base.ID == marketOrder.PlaceID {
					if marketOrder.Count > quantity {
						err := o.Sell(marketOrder.Id, quantity, user)
						if err == nil {
							return nil // т.к. мы продали все итемы то ордер создавать не ненадо
						}
					} else {
						countOrder := marketOrder.Count // т.к. при продаже удалиться count у ордера, сохраним его
						err := o.Sell(marketOrder.Id, marketOrder.Count, user)
						if err == nil {
							quantity -= countOrder
						}
					}
				}
			}

			newOrder := order.Order{IdUser: user.GetID(), Price: price, Count: quantity, Type: "sell",
				MinBuyOut: minBuyOut, TypeItem: slot.Type, IdItem: slot.ItemID, Expires: time.Now(),
				PlaceName: base.Name, PlaceID: base.ID, Item: slot.Item}

			// добавляем их в магазин
			id := market.PlaceNewOrder(&newOrder)

			if id > 0 {
				newOrder.Id = id
				// добавляем его на фабрику
				o.AddNewOrder(newOrder)
				// удаляем итем со склада
				storages.Storages.RemoveItemBySlot(user.GetID(), user.InBaseID, storageSlot, quantity)
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

			if !findBase {
				return errors.New("wrong base")
			}
		}

		update.Squad(user.GetSquad(), true)
		return nil
	} else {
		return errors.New("user not in base")
	}
}
