package market

import (
	"../db/market"
	"../gameObjects/order"
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

func (o *OrdersPool) GetOrders() map[int]*order.Order {
	o.mx.Lock()
	defer o.mx.Unlock()
	return o.orders
}

func (o *OrdersPool) GetOrder(id int) (bool, *order.Order, *sync.Mutex) {
	o.mx.Lock()
	openOrder, find := o.orders[id]
	// 3тий ретурн это мх, его надо вызрывать только после всех изменений с ордером
	return find, openOrder, &o.mx
}

func (o *OrdersPool) AddNewOrder(newOrder order.Order) {
	o.mx.Lock()
	defer o.mx.Unlock()
	o.orders[newOrder.Id] = &newOrder
}
