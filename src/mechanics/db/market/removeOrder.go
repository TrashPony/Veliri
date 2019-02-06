package market

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/order"
	"log"
)

func RemoveOrder(marketOrder *order.Order) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM orders WHERE id=$1",
		marketOrder.Id)
	if err != nil {
		log.Fatal("delete order" + err.Error())
	}
}
