package market

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/market"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"sync"
)

type OrdersPool struct {
	mx     sync.Mutex
	orders map[int]*order.Order
}

var Orders = NewOrdersPool()

func NewOrdersPool() *OrdersPool {
	return &OrdersPool{
		orders: market.OpenOrders(),
	}
}

func (o *OrdersPool) GetUserOrders(userID int) map[int]*order.Order {
	userOrders := make(map[int]*order.Order)

	for _, poolOrder := range o.orders {
		if poolOrder.IdUser == userID {
			userOrders[poolOrder.Id] = poolOrder
		}
	}

	return userOrders
}

func (o *OrdersPool) GetOrders() map[int]*order.Order {
	o.mx.Lock()
	defer o.mx.Unlock()
	return o.orders
}

func (o *OrdersPool) GetOrder(id int) (bool, *order.Order, *sync.Mutex) {
	o.mx.Lock()
	openOrder, find := o.orders[id]
	// 3тий ретурн это мх, его надо закрывать только после всех изменений с ордером
	return find, openOrder, &o.mx
}

func (o *OrdersPool) AddNewOrder(newOrder order.Order) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.orders[newOrder.Id] = &newOrder
}
