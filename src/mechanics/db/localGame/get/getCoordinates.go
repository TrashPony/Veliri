package get

import (
	"strconv"
	"log"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/map"
	"../../../localGame"
	"../../../../dbConnect"
)

func CoordinatesMap(mp *_map.Map, game *localGame.Game) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := dbConnect.GetDBConnect().Query("SELECT mc.q, mc.r, ct.type, ct.texture_flore, ct.texture_object, ct.move, ct.view, ct.attack, ct.passable_edges, mc.level "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id;", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		err := rows.Scan(&gameCoordinate.X, &gameCoordinate.Z, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.PassableEdges, &gameCoordinate.Level)
		if err != nil {
			log.Fatal(err)
		}

		gameCoordinate.GameID = game.Id
		CoordinateEffects(&gameCoordinate)

		gameCoordinate.Y = -gameCoordinate.X - gameCoordinate.Z

		if oneLayerMap[gameCoordinate.X] != nil {
			oneLayerMap[gameCoordinate.X][gameCoordinate.Z] = &gameCoordinate
		} else {
			oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
			oneLayerMap[gameCoordinate.X][gameCoordinate.Z] = &gameCoordinate
		}
	}

	defaultCoordinate := DefaultCoordinateType(mp)

	for q := 0; q < mp.QSize; q++ { // заполняем карту пустыми клетками тоесть дефолтными по карте
		for r := 0; r < mp.RSize; r++ {
			_, find := oneLayerMap[q][r]
			if !find {

				var gameCoordinate coordinate.Coordinate

				gameCoordinate = defaultCoordinate

				gameCoordinate.X = q
				gameCoordinate.Z = r
				gameCoordinate.Y = -q-r

				gameCoordinate.GameID = game.Id
				CoordinateEffects(&gameCoordinate)

				if oneLayerMap[gameCoordinate.X] != nil {
					oneLayerMap[gameCoordinate.X][gameCoordinate.Z] = &gameCoordinate
				} else {
					oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
					oneLayerMap[gameCoordinate.X][gameCoordinate.Z] = &gameCoordinate
				}
			}
		}
	}

	mp.OneLayerMap = oneLayerMap
}

func DefaultCoordinateType(mp *_map.Map) coordinate.Coordinate {
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
