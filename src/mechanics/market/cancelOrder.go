package market

import (
	"../db/market"
	dbPlayer "../db/player"
	"../factories/storages"
	"../player"
	"errors"
)

func (o *OrdersPool) Cancel(orderID int, user *player.Player) error {
	find, userOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	if find {
		if userOrder.IdUser == user.GetID() {
			if userOrder.Type == "buy" {
				// возвращаем деньги ща те итемы что остались
				user.SetCredits(user.GetCredits() + userOrder.Price*userOrder.Count)
				dbPlayer.UpdateUser(user)
				// удаляем заказ
				market.RemoveOrder(userOrder)
				delete(o.orders, userOrder.Id) // удаляем из фабрики т.к. мьютекс тут работает, это безопасно
				return nil
			} else {
				// возвращаем итемы что не куплены
				storages.Storages.AddItem(userOrder.IdUser, userOrder.PlaceID, userOrder.Item, userOrder.TypeItem,
					userOrder.IdItem, userOrder.Count, userOrder.ItemHP, userOrder.ItemSize*float32(userOrder.Count),
					userOrder.ItemHP)
				// удаляем заказ
				market.RemoveOrder(userOrder)
				delete(o.orders, userOrder.Id) // удаляем из фабрики т.к. мьютекс тут работает, это безопасно
				return nil
			}
		} else {
			return errors.New("not allow")
		}
	} else {
		return errors.New("error not find")
	}
}
