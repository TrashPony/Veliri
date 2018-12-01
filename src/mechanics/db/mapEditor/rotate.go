package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func Rotate(idMap, q, r, rotate int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET rotate = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
		rotate, idMap, q, r)
	if err != nil {
		log.Fatal("update rotate coordinate" + err.Error())
	}
}
