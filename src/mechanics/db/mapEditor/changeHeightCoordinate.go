package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func ChangeHeightCoordinate(idMap, q, r, change int) {

	changeCoordinate := getMapCoordinateInMC(idMap, q, r)

	if changeCoordinate != nil {

		level := changeCoordinate.Level + change

		if level > 5 {
			level = 5
		}

		_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET level = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
			level, idMap, q, r)
		if err != nil {
			log.Fatal("update storage item" + err.Error())
		}
	} else {

		defaultLevel, defaultType := getDefaultMap(idMap)

		defaultLevel += change

		if defaultLevel > 5 {
			defaultLevel = 5
		}

		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES ($1, $2, $3, $4, $5, '')",
			idMap, defaultType, q, r, defaultLevel)
		if err != nil {
			log.Fatal("add new level in map editor" + err.Error())
		}
	}
}
