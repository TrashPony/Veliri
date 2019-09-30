package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func DeleteCooordinateByQR(x, y, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE x=$1 AND y=$2 AND id_map=$3", x, y, mapID)
	if err != nil {
		log.Fatal(err)
	}
}
