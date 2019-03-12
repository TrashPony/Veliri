package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func ChangeType(typeID int, move, watch, attack bool) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE coordinate_type "+
		"SET move = $3, view = $4, attack = $5 "+
		"WHERE id = $6",
		move, watch, attack, typeID)
	if err != nil {
		log.Fatal("update type coordinate" + err.Error())
	}
}
