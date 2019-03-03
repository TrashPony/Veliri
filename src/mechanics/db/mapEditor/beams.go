package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func AddBeam(x, y, toX, toY, mapID int, color string) {
	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_beams (x_start, y_start, x_end, y_end, color, id_map) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		x, y, toX, toY, color, mapID)
	if err != nil {
		log.Fatal("add new beam " + err.Error())
	}
}

func RemoveBeam(id int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM global_geo_data WHERE id=$1 ", id)
	if err != nil {
		log.Fatal("add new geo data " + err.Error())
	}
}
