package get

import (
	"../../../dbConnect"
	"../../gameObjects/order"
	"log"
)

func OpenOrders() []*order.Order {
	orders := make([]*order.Order, 0)

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, id_user, price, count, type, min_buy_out, type_item," +
		" id_item, expires, place_name, place FROM orders")
	if err != nil {
		log.Fatal("get orders" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var mOrder order.Order

		rows.Scan(mOrder.Id, mOrder.IdUser, mOrder.Price, mOrder.Count, mOrder.Type, mOrder.MinBuyOut, mOrder.TypeItem,
			mOrder.IdItem, mOrder.Expires, mOrder.PlaceName, mOrder.PlaceID)

		if mOrder.TypeItem == "weapon" {
			mOrder.Item = Weapon(mOrder.IdItem)
		}

		if mOrder.TypeItem == "body" {
			mOrder.Item = Body(mOrder.IdItem)
		}

		if mOrder.TypeItem == "ammo" {
			mOrder.Item = Ammo(mOrder.IdItem)
		}

		if mOrder.TypeItem == "equip" {
			mOrder.Item = TypeEquip(mOrder.IdItem)
		}

		orders = append(orders, &mOrder)
	}

	return orders
}
