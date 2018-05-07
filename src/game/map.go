package game

import (
	"strconv"
	"log"
)

type Map struct {
	Id            int
	Name          string
	Xsize         int
	Ysize         int
	Type          string
	Specification string
	OneLayerMap   map[int]map[int]*Coordinate
}

func GetMap(idMap int) Map {

	rows, err := db.Query("Select * FROM maps WHERE id =" + strconv.Itoa(idMap))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Xsize, &mp.Ysize, &mp.Type, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
	}

	mp.GetCoordinatesMap()

	return mp
}

func (mp *Map) GetCoordinate(x, y int) (coordinate *Coordinate, find bool) {
	coordinate, find = mp.OneLayerMap[x][y]
	return
}

func (mp *Map) GetCoordinatesMap() {
	oneLayerMap := make(map[int]map[int]*Coordinate)

	rows, err := db.Query("Select x, y, type, texture FROM map_constructor WHERE id_map =" + strconv.Itoa(mp.Id))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
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

	for x := 0; x < mp.Xsize; x++ { // заполняем карту пустыми клетками
		for y := 0; y < mp.Xsize; y++ {
			_, find := oneLayerMap[x][y]
			if !find {
				var coordinate Coordinate
				coordinate = Coordinate{X:x, Y:y}
				oneLayerMap[x][y] = &coordinate
			}
		}
	}

	mp.OneLayerMap = oneLayerMap
}
