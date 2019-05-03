package market

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/market"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func (o *OrdersPool) Sell(orderID, count int, user *player.Player) error {

	find, sellOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	base, findBase := bases.Bases.Get(user.InBaseID)

	if find && sellOrder.Type == "buy" && findBase {
		if sellOrder.Count >= count && count%sellOrder.MinBuyOut == 0 && base.ID == sellOrder.PlaceID {

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
				sellOrder.IdItem, count, sellOrder.ItemHP, sellOrder.ItemSize*float32(count), sellOrder.ItemHP, false)
		} else {
			if sellOrder.Count < count || count%sellOrder.MinBuyOut != 0 {
				return errors.New("wrong count")
			}
			if base.ID != sellOrder.PlaceID {
				return errors.New("wrong base")
			}
		}
	} else {
		if !find {
			return errors.New("no find order")
		}

		if sellOrder.Type != "buy" {
			return errors.New("wrong order type")
		}
	}
	return nil
}
