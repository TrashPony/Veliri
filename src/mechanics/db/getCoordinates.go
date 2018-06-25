package db

import (
	"strconv"
	"log"
	"../coordinate"
	"../gameMap"
	"../game"
	"../../dbConnect"
)

func GetCoordinatesMap(mp *gameMap.Map, game *game.Game) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := dbConnect.GetDBConnect().Query("SELECT mc.x, mc.y, ct.type, ct.texture_flore, ct.texture_object, ct.move, ct.view, ct.attack, ct.passable_edges, mc.level "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id;", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		err := rows.Scan(&gameCoordinate.X, &gameCoordinate.Y, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.PassableEdges, &gameCoordinate.Level)
		if err != nil {
			log.Fatal(err)
		}

		gameCoordinate.GameID = game.Id
		GetCoordinateEffects(&gameCoordinate)

		if oneLayerMap[gameCoordinate.X] != nil {
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		} else {
			oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		}
	}

	defaultCoordinate := GetDefaultCoordinateType(mp)

	for x := 0; x < mp.XSize; x++ { // заполняем карту пустыми клетками тоесть дефолтными по карте
		for y := 0; y < mp.YSize; y++ {
			_, find := oneLayerMap[x][y]
			if !find {

				var gameCoordinate coordinate.Coordinate

				gameCoordinate = defaultCoordinate
				gameCoordinate.X = x
				gameCoordinate.Y = y
				
				gameCoordinate.GameID = game.Id
				GetCoordinateEffects(&gameCoordinate)

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

func GetDefaultCoordinateType(mp *gameMap.Map) coordinate.Coordinate {
	rows, err := dbConnect.GetDBConnect().Query("SELECT type, texture_flore, texture_object, move, view, attack, passable_edges "+
		"FROM coordinate_type "+
		"WHERE id = $1;", strconv.Itoa(mp.DefaultTypeID))

	if err != nil {
		println("Get Default coordinate type")
		log.Fatal(err)
	}

	defer rows.Close()

	gameCoordinate := coordinate.Coordinate{Level: mp.DefaultLevel}

	for rows.Next() {
		err := rows.Scan(&gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.PassableEdges)
		if err != nil {
			println("Get Default coordinate type")
			log.Fatal(err)
		}
	}

	return gameCoordinate
}
