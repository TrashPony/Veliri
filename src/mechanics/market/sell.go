package market

import (
	"../db/market"
	"../factories/players"
	"../factories/storages"
	"../player"
	"errors"
)

func (o *OrdersPool) Sell(orderID, count int, user *player.Player) error {

	find, sellOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	if find {
		if sellOrder.Count >= count && count%sellOrder.MinBuyOut == 0 {

			// пытаемся удалить итемы у продовца
			sellUserBaseStorage, _ := storages.Storages.Get(user.GetID(), user.InBaseID)
			err := sellUserBaseStorage.RemoveItem(sellOrder.IdItem, sellOrder.TypeItem, count)
			if err != nil {
				return errors.New("not items")
			}

			sellOrder.Count -= count
			if sellOrder.Count > 0 {
				market.UpdateOrder(sellOrder)
			} else {
				market.RemoveOrder(sellOrder)
				delete(o.orders, sellOrder.Id) // удаляем из фабрики т.к. мьютекс тут работает, это безопасно
			}

			// пополням баланс продавца
			players.Users.AddCash(user.GetID(), sellOrder.Price*count)

			// добавляем покупателю итемы в склад базы
			storages.Storages.AddItem(sellOrder.IdUser, sellOrder.PlaceID, sellOrder.Item, sellOrder.TypeItem,
				sellOrder.IdItem, count, sellOrder.ItemHP, sellOrder.ItemSize*float32(count), sellOrder.ItemHP)
		} else {
			return errors.New("wrong count")
		}
	}
	return nil
}
