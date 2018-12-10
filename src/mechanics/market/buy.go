package market

import (
	"../db/dbPlayer"
	"../db/market"
	"../player"
	"../players"
	"../storage"
	"errors"
)

func (o *OrdersPool) Buy(orderID, count int, user *player.Player) error {
	// 3тий ретурн это мх, его надо вызрывать только после всех изменений с ордером
	find, buyOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	if find {
		if user.GetCredits() >= buyOrder.Price*count && buyOrder.Count >= count {

			user.SetCredits(user.GetCredits() - buyOrder.Price*count) // отнимаем деньги :)
			buyOrder.Count -= count                                   // отнимаем количество покупаемых итемов у ордера

			if buyOrder.Count > 0 {
				market.UpdateOrder(buyOrder)
			} else {
				market.RemoveOrder(buyOrder)
				delete(o.orders, buyOrder.Id) // удаляем из фабрики т.к. мьютекс тут работает, это безопасно
			}

			storage.Storages.AddItem(user.GetID(), buyOrder.PlaceID, buyOrder.Item, buyOrder.TypeItem,
				buyOrder.IdItem, count, buyOrder.ItemHP, buyOrder.ItemSize*float32(count), buyOrder.ItemHP)

			dbPlayer.UpdateUser(user)

			players.Users.AddCash(buyOrder.IdUser, buyOrder.Price*count) // пополням баланс продавца
		} else {
			if user.GetCredits() < buyOrder.Price*count {
				return errors.New("no credits")
			}
			if buyOrder.Count < count {
				return errors.New("wrong count")
			}
		}
	} else {
		return errors.New("no find order")
	}

	return nil
}
