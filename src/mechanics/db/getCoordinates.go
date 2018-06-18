package db

import (
	"strconv"
	"log"
	"../coordinate"
	"../gameMap"
	"../game"
)

func GetCoordinatesMap(mp *gameMap.Map, game *game.Game) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := db.Query("SELECT mc.x, mc.y, ct.type, ct.texture_flore, ct.texture_object, ct.move, ct.view, ct.attack, ct.passable_edges, mc.level "+
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

	for x := 0; x < mp.XSize; x++ { // заполняем карту пустыми клетками
		for y := 0; y < mp.YSize; y++ {
			_, find := oneLayerMap[x][y]
			if !find {
				var gameCoordinate coordinate.Coordinate
				gameCoordinate = coordinate.Coordinate{X: x, Y: y}

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
