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
		err := rows.Scan(&gameCoordinate.Q, &gameCoordinate.R, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.PassableEdges, &gameCoordinate.Level)
		if err != nil {
			log.Fatal(err)
		}

		gameCoordinate.GameID = game.Id
		CoordinateEffects(&gameCoordinate)

		// в бд карта храниться в хексовых координатах
		gameCoordinate.X = gameCoordinate.Q - (gameCoordinate.R - (gameCoordinate.R & 1)) / 2
		gameCoordinate.Z = gameCoordinate.R
		gameCoordinate.Y = -gameCoordinate.X - gameCoordinate.Z

		if oneLayerMap[gameCoordinate.Q] != nil {
			oneLayerMap[gameCoordinate.Q][gameCoordinate.R] = &gameCoordinate
		} else {
			oneLayerMap[gameCoordinate.Q] = make(map[int]*coordinate.Coordinate)
			oneLayerMap[gameCoordinate.Q][gameCoordinate.R] = &gameCoordinate
		}
	}

	defaultCoordinate := DefaultCoordinateType(mp)

	for q := 0; q < mp.QSize; q++ { // заполняем карту пустыми клетками тоесть дефолтными по карте
		for r := 0; r < mp.RSize; r++ {
			_, find := oneLayerMap[q][r]
			if !find {

				var gameCoordinate coordinate.Coordinate

				gameCoordinate = defaultCoordinate

				gameCoordinate.Q = q
				gameCoordinate.R = r
				// в бд карта храниться в хексовых координатах
				gameCoordinate.X = q - (r - (r & 1)) / 2
				gameCoordinate.Z = r
				gameCoordinate.Y = -gameCoordinate.X - gameCoordinate.Z

				gameCoordinate.GameID = game.Id
				CoordinateEffects(&gameCoordinate)

				if oneLayerMap[gameCoordinate.Q] != nil {
					oneLayerMap[gameCoordinate.Q][gameCoordinate.R] = &gameCoordinate
				} else {
					oneLayerMap[gameCoordinate.Q] = make(map[int]*coordinate.Coordinate)
					oneLayerMap[gameCoordinate.Q][gameCoordinate.R] = &gameCoordinate
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
