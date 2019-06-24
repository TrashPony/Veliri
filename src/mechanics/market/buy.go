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

func (o *OrdersPool) Buy(orderID, count int, user *player.Player) error {

	find, buyOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	if find && buyOrder.Type == "sell" && count > 0 && user.GetID() != buyOrder.IdUser {
		if user.GetCredits() >= buyOrder.Price*count && buyOrder.Count >= count {

			user.SetCredits(user.GetCredits() - buyOrder.Price*count) // отнимаем деньги :)
			dbPlayer.UpdateUser(user)

			buyOrder.Count -= count // отнимаем количество покупаемых итемов у ордера

			if buyOrder.Count > 0 {
				market.UpdateOrder(buyOrder)
			} else {
				market.RemoveOrder(buyOrder)
				delete(o.orders, buyOrder.Id) // удаляем из фабрики т.к. мьютекс тут работает, это безопасно
			}

			storages.Storages.AddItem(user.GetID(), buyOrder.PlaceID, buyOrder.Item, buyOrder.TypeItem,
				buyOrder.IdItem, count, buyOrder.ItemHP, buyOrder.ItemSize, buyOrder.ItemHP, false)

			players.Users.AddCash(buyOrder.IdUser, buyOrder.Price*count) // пополням баланс продавца

			// создаем продавцу нотификацию
			// из за этой махинации с нотификация метод AddCash почти не имеет смысла, но трогать не буду)
			ownerDeal, _ := players.Users.Get(buyOrder.IdUser)
			notifyUUID := uuid.Must(uuid.NewV4(), nil).String()
			base, _ := bases.Bases.Get(buyOrder.PlaceID)
			mp, _ := maps.Maps.GetByID(base.MapID)

			// для пользователя который разместил ордер это продажа, поэтому нотификация о продаже
			ownerDeal.NotifyQueue[notifyUUID] = &player.Notify{
				Name:  "sell",
				UUID:  notifyUUID,
				Event: "complete",
				Item:  &inventory.Slot{Item: buyOrder.Item, Quantity: count, Type: buyOrder.TypeItem},
				Base:  base,
				Map:   mp.GetShortInfoMap(),
				Price: buyOrder.Price,
			}
			dbPlayer.UpdateUser(ownerDeal)
		} else {
			if user.GetCredits() < buyOrder.Price*count {
				return errors.New("no credits")
			}
			if buyOrder.Count < count {
				return errors.New("wrong count")
			}
		}
	} else {
		if !find {
			return errors.New("no find order")
		}

		if buyOrder.Type != "sell" {
			return errors.New("wrong order type")
		}

		if count < 1 {
			return errors.New("wrong count items")
		}

		if user.GetID() == buyOrder.IdUser {
			return errors.New("it's you order")
		}
	}

	return nil
}
