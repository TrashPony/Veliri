package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func AddAnomaly(x, y, radius, mapID, power int, typeAnomaly string) {
	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_danger_anomalies (x, y, radius, id_map, power, type) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		x, y, radius, mapID, power, typeAnomaly)
	if err != nil {
		log.Fatal("add new anomaly " + err.Error())
	}
}

func RemoveAnomaly(id int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_danger_anomalies WHERE id=$1 ", id)
	if err != nil {
		log.Fatal("remove anomaly " + err.Error())
	}
}
