package dbo

import (
	"strconv"
	"log"
	"../coordinate"
	"../gameMap"
)

func GetCoordinatesMap(mp *gameMap.Map) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := db.Query("Select x, y, type, texture FROM map_constructor WHERE id_map =" + strconv.Itoa(mp.Id))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		err := rows.Scan(&gameCoordinate.X, &gameCoordinate.Y, &gameCoordinate.Type, &gameCoordinate.Texture)
		if err != nil {
			log.Fatal(err)
		}

		if oneLayerMap[gameCoordinate.X] != nil {
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		} else {
			oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		}
	}

	for x := 0; x < mp.XSize; x++ { // заполняем карту пустыми клетками
		for y := 0; y < mp.YSize; y++ {
			_, find := oneLayerMap[x][y]
			if !find {
				var gameCoordinate coordinate.Coordinate
				gameCoordinate = coordinate.Coordinate{X:x, Y:y}

				if oneLayerMap[gameCoordinate.X] != nil {
					oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
				} else {
					oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
					oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
				}
			}
		}
	}

	mp.OneLayerMap = oneLayerMap
}
