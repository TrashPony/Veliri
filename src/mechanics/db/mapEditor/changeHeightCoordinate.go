package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"log"
)

func ChangeHeightCoordinate(idMap, q, r, change int) {

	changeCoordinate := getMapCoordinateInMC(idMap, q, r)

	if changeCoordinate != nil {
		if changeCoordinate.Impact != nil {
			return
		}

		if changeCoordinate.ImpactRadius == 0 {
			updateChangeHeight(idMap, q, r, change, changeCoordinate.Level)
		} else {

			changeCoordinate.Q = q
			changeCoordinate.R = r
			changeCoordinate.CalculateXYZ()

			radiusCoordinates := coordinate.GetCoordinatesRadius(changeCoordinate, changeCoordinate.ImpactRadius)

			for _, coor := range radiusCoordinates {
				mapCoor := getMapCoordinateInMC(idMap, coor.Q, coor.R)
				if mapCoor != nil {
					updateChangeHeight(idMap, coor.Q, coor.R, change, mapCoor.Level)
				} else {
					insertNewHeightCoordinate(idMap, coor.Q, coor.R, change)
				}
			}

			// обновляем центр
			updateChangeHeight(idMap, q, r, change, changeCoordinate.Level)
		}
	} else {
		insertNewHeightCoordinate(idMap, q, r, change)
	}
}

func insertNewHeightCoordinate(idMap, q, r, change int) {
	defaultLevel, defaultType := getDefaultMap(idMap)

	defaultLevel += change

	if defaultLevel > 5 {
		defaultLevel = 5
	}

	if defaultLevel < 1 {
		defaultLevel = 1
	}

	_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed, "+
		"x_offset, y_offset) "+
		"VALUES ($1, $2, $3, $4, $5, '', $6, $7, $8, $9)",
		idMap, defaultType, q, r, defaultLevel, 0, 60, 0, 0)
	if err != nil {
		log.Fatal("add new level in map editor " + err.Error())
	}
}

func updateChangeHeight(idMap, q, r, change, oldLvl int) {
	level := oldLvl + change

	if level > 5 {
		level = 5
	}

	if level < 1 {
		level = 1
	}

	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET level = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
		level, idMap, q, r)
	if err != nil {
		log.Fatal("update level in map editor" + err.Error())
	}
}
