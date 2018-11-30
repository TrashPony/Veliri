package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func ChangeType(typeID, scale int, shadow bool) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE coordinate_type SET scale = $1, shadow = $2 WHERE id = $3",
		scale, shadow, typeID)
	if err != nil {
		log.Fatal("update type coordinate" + err.Error())
	}
}
