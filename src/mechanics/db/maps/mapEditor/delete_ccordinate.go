package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func DeleteCooordinateByQR(q, r, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE q=$1 AND r=$2 AND id_map=$3", q, r, mapID)
	if err != nil {
		log.Fatal(err)
	}
}
