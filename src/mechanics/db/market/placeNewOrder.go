package market

import (
	"../../../dbConnect"
	"../../gameObjects/order"
	"log"
)

func PlaceNewOrder(newOrder *order.Order) int {
	var id int

	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO orders "+
		"(id_user, price, count, type, min_buy_out, type_item, id_item, expires, place_name, place) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		newOrder.IdUser, newOrder.Price, newOrder.Count, newOrder.Type, newOrder.MinBuyOut, newOrder.TypeItem,
		newOrder.IdItem, newOrder.Expires, newOrder.PlaceName, newOrder.PlaceID).Scan(&id)
	if err != nil {
		log.Fatal("add new order" + err.Error())
	}

	return id
}
