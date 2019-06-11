package market

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
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
		//todo одинаковый код с заполнением инвентаря, опасный код
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

		if mOrder.TypeItem == "resource" {
			resource, _ := gameTypes.Resource.GetBaseByID(mOrder.IdItem)

			mOrder.Item = resource
			mOrder.ItemSize = resource.Size
			mOrder.ItemHP = 1 // у ресов нет хп
		}

		if mOrder.TypeItem == "recycle" {
			resource, _ := gameTypes.Resource.GetRecycledByID(mOrder.IdItem)

			mOrder.Item = resource
			mOrder.ItemSize = resource.Size
			mOrder.ItemHP = 1 // у ресов нет хп
		}

		if mOrder.TypeItem == "detail" {
			detail, _ := gameTypes.Resource.GetDetailByID(mOrder.IdItem)

			mOrder.Item = detail
			mOrder.ItemSize = detail.Size
			mOrder.ItemHP = 1 // у ресов нет хп
		}

		if mOrder.TypeItem == "boxes" {
			box, _ := gameTypes.Boxes.GetByID(mOrder.IdItem)

			mOrder.Item = box
			mOrder.ItemSize = box.FoldSize
			mOrder.ItemHP = 1 // у ящиков тож нет хп
		}

		if mOrder.TypeItem == "blueprints" {
			blueprint, _ := gameTypes.BluePrints.GetByID(mOrder.IdItem)

			mOrder.Item = blueprint
			mOrder.ItemSize = 0 // чертежи не занимают места
			mOrder.ItemHP = 1   // у чертежов нет хп
		}

		if mOrder.TypeItem == "trash" {
			trashItem, _ := gameTypes.TrashItems.GetByID(mOrder.IdItem)

			mOrder.Item = trashItem
			mOrder.ItemSize = trashItem.Size
			mOrder.ItemHP = 1 // у мусора нет хп
		}

		orders[mOrder.Id] = &mOrder
	}

	return orders
}
