package game

import (
	"strconv"
	"log"
)

type Map struct {
	Id	    	 int
	Name  		 string
	Xsize 		 int
	Ysize 		 int
	Type		 string
	OneLayerMap  map[int]map[int]*Coordinate
}

func GetMap(idMap int) Map {

	rows, err := db.Query("Select * FROM maps WHERE id =" + strconv.Itoa(idMap))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.Xsize, &mp.Ysize, &mp.Type)
		if err != nil {
			log.Fatal(err)
		}
	}

	oneLayerMap := GetCoordinateMap(idMap)
	mp.OneLayerMap = oneLayerMap

	return mp
}

func (mp *Map) GetCoordinate(x, y int) (coordinate *Coordinate, find bool)  {
	coordinate, find = mp.OneLayerMap[x][y]
	return
}


func GetCoordinateMap(idMap int) (oneLayerMap  map[int]map[int]*Coordinate)  {
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
