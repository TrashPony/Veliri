package market

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/market"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/players"
	"github.com/TrashPony/Veliri/src/mechanics/factories/storages"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/satori/go.uuid"
)

func (o *OrdersPool) Sell(orderID, count int, user *player.Player) error {

	find, sellOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	base, findBase := bases.Bases.Get(user.InBaseID)

	if find && sellOrder.Type == "buy" && findBase && count > 0 && user.GetID() != sellOrder.IdUser {
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

			// создаем покупателю нотификацию
			// из за этой махинации с нотификация метод AddCash почти не имеет смысла, но трогать не буду)
			ownerDeal, _ := players.Users.Get(sellOrder.IdUser)
			notifyUUID := uuid.Must(uuid.NewV4(), nil).String()
			base, _ := bases.Bases.Get(sellOrder.PlaceID)
			mp, _ := maps.Maps.GetByID(base.MapID)

			// для пользователя который разместил ордер это покупка, поэтому нотификация о покупке
			ownerDeal.NotifyQueue[notifyUUID] = &player.Notify{
				Name:  "buy",
				UUID:  notifyUUID,
				Event: "complete",
				Item:  &inventory.Slot{Item: sellOrder.Item, Quantity: count, Type: sellOrder.TypeItem},
				Base:  base,
				Map:   mp.GetShortInfoMap(),
				Price: sellOrder.Price,
			}
			dbPlayer.UpdateUser(ownerDeal)
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

		if count < 1 {
			return errors.New("wrong count items")
		}

		if user.GetID() == sellOrder.IdUser {
			return errors.New("it's you order")
		}
	}
	return nil
}
