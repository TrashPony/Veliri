package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func ChangeType(typeID, scale int, shadow, move, watch, attack bool) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE coordinate_type SET scale = $1, shadow = $2, move = $3, view = $4, attack = $5 WHERE id = $6",
		scale, shadow, move, watch, attack, typeID)
	if err != nil {
		log.Fatal("update type coordinate" + err.Error())
	}
}
