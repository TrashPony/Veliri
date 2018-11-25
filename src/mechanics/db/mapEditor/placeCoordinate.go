package mapEditor

import (
	"../../../dbConnect"
	"log"
)

func PlaceCoordinate(idMap, idType, q, r int) {

	defaultLevel, _ := getDefaultMap(idMap)

	oldType := getMapCoordinateInMC(idMap, q, r)
	newType := getTypeByID(idType)

	if oldType != nil {
		if oldType.ImpactRadius == 0 && newType.ImpactRadius == 0 {
			_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET id_type = $1 WHERE id_map = $2 AND q=$3 AND r = $4",
				idType, idMap, q, r)
			if err != nil {
				log.Fatal("update type in mc " + err.Error())
			}
		} else {
			if oldType.ImpactRadius != 0 {
				// todo убераем старое воздействие
			}
			if newType.ImpactRadius != 0 {
				// todo берем радиус
				// todo смотрим что бы у всех была одинакеовая высота и они были свободны
				// todo затем обновляем, добавляем в конструктор координаты в радиусе
			}
		}
	} else {
		if newType.ImpactRadius == 0 {
			_, err := dbConnect.GetDBConnect().Exec("INSERT INTO map_constructor (id_map, id_type, q, r, level, impact) VALUES ($1, $2, $3, $4, $5, '')",
				idMap, idType, q, r, defaultLevel)
			if err != nil {
				log.Fatal("add new type in mc " + err.Error())
			}
		} else {
			// todo берем радиус
			// todo смотрим что бы у всех была одинакеовая высота и они были свободны
			// todo затем обновляем, добавляем в конструктор координаты в радиусе
		}
	}
}
