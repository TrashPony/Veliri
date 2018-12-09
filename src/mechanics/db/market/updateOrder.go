package market

import (
	"../../../dbConnect"
	"../../gameObjects/order"
	"log"
)

func UpdateOrder(marketOrder *order.Order) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE orders "+
		"SET count = $2, price = $3, min_buy_out = $4, expires = $5 "+
		"WHERE id = $1",
		marketOrder.Id, marketOrder.Count, marketOrder.Price, marketOrder.MinBuyOut, marketOrder.Expires)
	if err != nil {
		log.Fatal("update order " + err.Error())
	}
}
