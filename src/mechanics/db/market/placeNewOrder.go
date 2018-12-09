package market

import (
	"../../../dbConnect"
	"log"
	"time"
)

func PlaceNewOrder(userID, price, count, minBuyOut, itemID, expires, baseID int, typeOrder, typeItem, baseName string) {
	// todo expires в часах, надо брать текущее время  + expires и это значение класть в бд
	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO orders "+
		"(id_user, price, count, type, min_buy_out, type_item, id_item, expires, place_name, place) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		userID, price, count, typeOrder, minBuyOut, typeItem, itemID, time.Now(), baseName, baseID)
	if err != nil {
		log.Fatal("add new order" + err.Error())
	}
}
