package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func Rotate(idMap, q, r, rotate, speed, XOffset, YOffset int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET rotate = $1, animate_speed = $5, x_offset = $6, y_offset = $7 "+
		"WHERE id_map = $2 AND q=$3 AND r = $4",
		rotate, idMap, q, r, speed, XOffset, YOffset)
	if err != nil {
		log.Fatal("update rotate coordinate" + err.Error())
	}
}
