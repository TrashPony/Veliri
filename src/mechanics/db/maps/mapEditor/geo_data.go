package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func AddGeoData(x, y, radius, mapdID int) {
	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO global_geo_data (x, y, radius, id_map) "+
		"VALUES ($1, $2, $3, $4)",
		x, y, radius, mapdID)
	if err != nil {
		log.Fatal("add new geo data " + err.Error())
	}
}

func RemoveGeoData(id int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM global_geo_data WHERE id=$1 ", id)
	if err != nil {
		log.Fatal("add new geo data " + err.Error())
	}
}
