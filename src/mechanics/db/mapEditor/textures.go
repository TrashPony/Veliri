package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

func PlaceTextures(idMap, q, r int, textureName string) {
	changeCoordinate := getMapCoordinateInMC(idMap, q, r)
	if changeCoordinate != nil {
		_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET texture_over_flore = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
			textureName, idMap, q, r)
		if err != nil {
			log.Fatal("update texture in map editor" + err.Error())
		}
	} else {
		defaultLevel, defaultType := getDefaultMap(idMap)

		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed, "+
			"x_offset, y_offset, texture_over_flore) "+
			"VALUES ($1, $2, $3, $4, $5, '', $6, $7, $8, $9, $10)",
			idMap, defaultType, q, r, defaultLevel, 0, 60, 0, 0, textureName)
		if err != nil {
			log.Fatal("add new texture in map editor " + err.Error())
		}
	}
}

func RemoveTextures(idMap, q, r int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET texture_over_flore = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
		"", idMap, q, r)
	if err != nil {
		log.Fatal("delete texture in map editor" + err.Error())
	}
}
