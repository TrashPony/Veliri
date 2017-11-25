package objects

import (
	"strconv"
	"log"
)

func GetMap(idMap int) (oneLayerMap  map[int]map[int]*Coordinate)  {
	oneLayerMap = make(map[int]map[int]*Coordinate)
	rows, err := db.Query("Select x, y, type, texture FROM map_constructor WHERE id_map =" + strconv.Itoa(idMap))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var coordinate Coordinate
		err := rows.Scan(&coordinate.X, &coordinate.Y, &coordinate.Type, &coordinate.Texture)
		if err != nil {
			log.Fatal(err)
		}

		if oneLayerMap[coordinate.X] != nil {
			oneLayerMap[coordinate.X][coordinate.Y] = &coordinate
		} else {
			oneLayerMap[coordinate.X] = make(map[int]*Coordinate)
			oneLayerMap[coordinate.X][coordinate.Y] = &coordinate
		}
	}

	return
}
