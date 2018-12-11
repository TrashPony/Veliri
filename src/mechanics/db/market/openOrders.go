package market

import (
	"../../../dbConnect"
	"../../factories/gameTypes"
	"../../gameObjects/order"
	"log"
)

func OpenOrders() map[int]*order.Order {

	var orders = make(map[int]*order.Order)

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, id_user, price, count, type, min_buy_out, type_item," +
		" id_item, expires, place_name, place FROM orders")
	if err != nil {
		log.Fatal("get orders" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var mOrder order.Order

		rows.Scan(&mOrder.Id, &mOrder.IdUser, &mOrder.Price, &mOrder.Count, &mOrder.Type, &mOrder.MinBuyOut, &mOrder.TypeItem,
			&mOrder.IdItem, &mOrder.Expires, &mOrder.PlaceName, &mOrder.PlaceID)

		if mOrder.TypeItem == "weapon" {
			weapon, _ := gameTypes.Weapons.GetByID(mOrder.IdItem)

			mOrder.Item = weapon
			mOrder.ItemSize = weapon.Size
			mOrder.ItemHP = weapon.MaxHP
		}

		if mOrder.TypeItem == "body" {
			body, _ := gameTypes.Bodies.GetByID(mOrder.IdItem)

			mOrder.Item = body
			mOrder.ItemSize = body.CapacitySize
			mOrder.ItemHP = body.MaxHP
		}

		if mOrder.TypeItem == "ammo" {
			ammo, _ := gameTypes.Ammo.GetByID(mOrder.IdItem)

			mOrder.Item = ammo
			mOrder.ItemSize = ammo.Size
			mOrder.ItemHP = 1
		}

		if mOrder.TypeItem == "equip" {
			equip, _ := gameTypes.Equips.GetByID(mOrder.IdItem)

			mOrder.Item = equip
			mOrder.ItemSize = equip.Size
			mOrder.ItemHP = equip.MaxHP
		}

		orders[mOrder.Id] = &mOrder
	}

	return orders
}
