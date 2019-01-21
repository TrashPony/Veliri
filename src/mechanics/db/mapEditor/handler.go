package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func PlaceHandler(idMap, q, r, toQ, toR, toBaseId, toMapId int, typeHandler string) {
	if toMapId == 0 {
		// TODO костыль из за ограничение сылок на ключи
		toMapId = 1
	}

	if toBaseId == 0 {
		// TODO костыль из за ограничение сылок на ключи
		toBaseId = 1
	}

	changeCoordinate := getMapCoordinateInMC(idMap, q, r)
	if changeCoordinate != nil {
		_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET handler = $4, to_q = $5, to_r = $6, to_base_id = $7, to_map_id = $8"+
			"WHERE id_map = $1 AND q=$2 AND r = $3",
			idMap, q, r, typeHandler, toQ, toR, toBaseId, toMapId)
		if err != nil {
			log.Fatal("update handler in map editor" + err.Error())
		}
	} else {
		defaultLevel, defaultType := getDefaultMap(idMap)
		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed, "+
			"x_offset, y_offset, handler, to_q, to_r, to_base_id, to_map_id) "+
			"VALUES ($1, $2, $3, $4, $5, '', $6, $7, $8, $9, $10, $11, $12, $13, $14)",
			idMap, defaultType, q, r, defaultLevel, 0, 60, 0, 0, typeHandler, toQ, toR, toBaseId, toMapId)
		if err != nil {
			log.Fatal("add new handler in map editor " + err.Error())
		}
	}
}

func RemoveHandler(idMap, q, r int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET handler = $4, to_q = $5, to_r = $6,"+
		" to_base_id = $7, to_map_id = $8"+
		" WHERE id_map = $1 AND q=$2 AND r = $3",
		idMap, q, r, "", 0, 0, 0, 0)
	if err != nil {
		log.Fatal("delete handler in map editor" + err.Error())
	}
}
