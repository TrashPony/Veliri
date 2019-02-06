package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func ChangeType(typeID, scale int, shadow, move, watch, attack bool, radius int) {
	//TODO т.к. мы меняем тип координаты то он отаражется на ВСЕХ карты в игре значит надо перерасчитывать занятое пространство на всех картах
	_, err := dbConnect.GetDBConnect().Exec("UPDATE coordinate_type "+
		"SET scale = $1, shadow = $2, move = $3, view = $4, attack = $5, impact_radius = $7 "+
		"WHERE id = $6",
		scale, shadow, move, watch, attack, typeID, radius)
	if err != nil {
		log.Fatal("update type coordinate" + err.Error())
	}
}
