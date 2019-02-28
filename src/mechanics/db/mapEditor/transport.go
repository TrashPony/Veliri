package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func PlaceTransport(idMap, q, r int) {
	changeCoordinate := getMapCoordinateInMC(idMap, q, r)
	if changeCoordinate != nil {
		_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET transport = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
			true, idMap, q, r)
		if err != nil {
			log.Fatal("update transport in map editor" + err.Error())
		}
	} else {
		defaultLevel, defaultType := getDefaultMap(idMap)

		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed, "+
			"x_offset, y_offset, transport, texture_over_flore, handler, to_q, to_r, to_base_id, to_map_id) "+
			"VALUES ($1, $2, $3, $4, $5, '', $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)",
			idMap, defaultType, q, r, defaultLevel, 0, 60, 0, 0, true, "", "", 0, 0, 1, 2)
		if err != nil {
			log.Fatal("add new transport in map editor " + err.Error())
		}
	}
}

func RemoveTransport(idMap, q, r int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET transport = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
		false, idMap, q, r)
	if err != nil {
		log.Fatal("delete transport in map editor" + err.Error())
	}
}
