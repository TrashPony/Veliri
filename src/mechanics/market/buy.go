package market

import (
	"../player"
	"errors"
)

func (o *OrdersPool) Buy(orderID, count int, user *player.Player) error {
	// 3тий ретурн это мх, его надо вызрывать только после всех изменений с ордером
	find, buyOrder, mx := o.GetOrder(orderID)
	defer mx.Unlock()

	if find {
		if user.GetCredits() >= buyOrder.Price*count && buyOrder.Count >= count {
			user.SetCredits(user.GetCredits() - buyOrder.Price*count) // отнимаем деньги :)

			// 3 обновляем ордер в бд, в фабрике
			// 4 кладем новый итем в количестве купленых штук в склад той базы где был ордер

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
