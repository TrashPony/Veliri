package mapEditor

import (
	"../../../dbConnect"
	"../../gameObjects/coordinate"
	"log"
	"strconv"
)

func PlaceCoordinate(idMap, idType, q, r int) {

	defaultLevel, _ := getDefaultMap(idMap)

	oldType := getMapCoordinateInMC(idMap, q, r)
	newType := getTypeByID(idType)

	// присваиваем позицию что бы правильно расчитать радиус действия
	newType.Q = q
	newType.R = r
	newType.CalculateXYZ()

	if oldType != nil {
		// если старая координата чьято подчиненная то мы не можем ставить туда что либо
		if oldType.Impact != nil {
			return
		}

		// если есть старая координата то у новой будет такая же высота
		newType.Level = oldType.Level

		if oldType.ImpactRadius == 0 && newType.ImpactRadius == 0 {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET id_type = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
				idType, idMap, q, r)
			if err != nil {
				log.Fatal("update type in mc " + err.Error())
			}
		} else {
			if newType.ImpactRadius != 0 && oldType.ImpactRadius != 0 {
				if checkRadiusCoordinate(newType, idMap) {
					// т.к. q и r одиновые у старой и новой координаты то сначало удались старые ключи
					removeImpact(newType, idMap)
					// и уже добавим новые
					placeRadiusCoordinate(newType, idMap, true)
				}
			} else {
				if newType.ImpactRadius != 0 {
					if checkRadiusCoordinate(newType, idMap) {

						_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE id_map = $1 AND q = $2 AND r = $3",
							idMap, q, r)
						if err != nil {
							log.Fatal("remove old type in mc " + err.Error())
						}

						placeRadiusCoordinate(newType, idMap, true)
					}
				}

				if oldType.ImpactRadius != 0 {
					removeImpact(newType, idMap)
					_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET id_type = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
						idType, idMap, q, r)
					if err != nil {
						log.Fatal("update type in mc " + err.Error())
					}
				}
			}
		}
	} else {

		newType.Level = defaultLevel

		if newType.ImpactRadius == 0 {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed, "+
				"x_offset, y_offset)"+
				" VALUES ($1, $2, $3, $4, $5, '', $6, $7, $8, $9)",
				idMap, idType, q, r, defaultLevel, 0, 60, 0, 0)
			if err != nil {
				log.Fatal("add new type in mc " + err.Error())
			}
		} else {
			if checkRadiusCoordinate(newType, idMap) {
				placeRadiusCoordinate(newType, idMap, true)
			}
		}
	}
}

func removeImpact(removeCoordinate *coordinate.Coordinate, idMap int) {
	keyString := parseImpactToString(removeCoordinate)
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE id_map = $1 AND impact= $2",
		idMap, keyString)
	if err != nil {
		log.Fatal("remove impact type in mc " + err.Error())
	}
}

func checkRadiusCoordinate(placeCoordinate *coordinate.Coordinate, idMap int) bool {

	radiusCoordinates := coordinate.GetCoordinatesRadius(placeCoordinate, placeCoordinate.ImpactRadius)

	passed := true

	for _, coor := range radiusCoordinates {
		mapCoor := getMapCoordinateInMC(idMap, coor.Q, coor.R)
		if mapCoor != nil {
			if mapCoor.Level != placeCoordinate.Level || mapCoor.Impact != nil {
				passed = false
			}
		}
	}

	return passed
}

func placeRadiusCoordinate(placeCoordinate *coordinate.Coordinate, idMap int, add bool) {

	radiusCoordinates := coordinate.GetCoordinatesRadius(placeCoordinate, placeCoordinate.ImpactRadius)

	qSize, rSize := getSizeMap(idMap)

	for _, coor := range radiusCoordinates {

		if coor.Q < 0 || coor.R < 0 {
			// координата провалилась за пределы карты, удаляем ее
			continue
		}

		if coor.Q > qSize-1 || coor.R > rSize-1 {
			// координата провалилась за пределы карты, удаляем ее
			//continue
		}

		if !(coor.Q == placeCoordinate.Q && coor.R == placeCoordinate.R) {
			mapCoor := getMapCoordinateInMC(idMap, coor.Q, coor.R)

			if mapCoor != nil {
				// добавляем тип координате такой же как у то что влияет, типо она под влияющей.
				_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET id_type = $1, impact = $5 "+
					" WHERE id_map = $2 AND q=$3 AND r = $4",
					placeCoordinate.ID, idMap, coor.Q, coor.R, parseImpactToString(placeCoordinate))
				if err != nil {
					log.Fatal("update radius impact type in mc " + err.Error())
				}

			} else {

				// добавляем тип координате такой же как у то что влияет, типо она под влияющей.
				_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, "+
					"animate_speed, x_offset, y_offset) "+
					" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
					idMap, placeCoordinate.ID, coor.Q, coor.R, placeCoordinate.Level, parseImpactToString(placeCoordinate), 0, 60, 0, 0)
				if err != nil {
					log.Fatal("add new radius impact type in mc " + err.Error())
				}

			}
		}
	}

	if add {
		// у самой влияющей координаты нет значения impact и это говорит клиенту что рисовать обьект надо именно тут а не в подчиненных
		_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact, rotate, animate_speed,"+
			" x_offset, y_offset) "+
			"VALUES ($1, $2, $3, $4, $5, '', $6, $7)",
			idMap, placeCoordinate.ID, placeCoordinate.Q, placeCoordinate.R, placeCoordinate.Level, 0, 60, 0, 0)
		if err != nil {
			log.Fatal("add new impact type in mc " + err.Error())
		}
	}
}

func parseImpactToString(targetCoordinate *coordinate.Coordinate) string {
	var target string

	if targetCoordinate != nil {
		target = strconv.Itoa(targetCoordinate.Q) + ":" + strconv.Itoa(targetCoordinate.R)
	}

	return target
}
