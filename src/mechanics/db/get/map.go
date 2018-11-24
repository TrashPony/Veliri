package get

import (
	"../../../dbConnect"
	"../../gameObjects/coordinate"
	"../../gameObjects/effect"
	"../../gameObjects/map"
	"log"
	"strconv"
)

func Map(mapID int) _map.Map {

	rows, err := dbConnect.GetDBConnect().Query("Select id, name, q_size, r_size, id_type, level, specification FROM maps WHERE id = $1", mapID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mp _map.Map
	for rows.Next() {
		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification)
		if err != nil {
			log.Fatal(err)
		}
	}

	CoordinatesMap(&mp)

	return mp
}

func CoordinatesMap(mp *_map.Map) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := dbConnect.GetDBConnect().Query("SELECT ct.id, mc.q, mc.r, ct.type, ct.texture_flore, ct.texture_object, ct.move, ct.view, ct.attack, ct.passable_edges, mc.level "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id;", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		err := rows.Scan(&gameCoordinate.ID, &gameCoordinate.Q, &gameCoordinate.R, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.PassableEdges, &gameCoordinate.Level)
		if err != nil {
			log.Fatal(err)
		}

		CoordinateEffects(&gameCoordinate)

		// в бд карта храниться в хексовых координатах
		gameCoordinate.X = gameCoordinate.Q - (gameCoordinate.R-(gameCoordinate.R&1))/2
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
				gameCoordinate.X = q - (r-(r&1))/2
				gameCoordinate.Z = r
				gameCoordinate.Y = -gameCoordinate.X - gameCoordinate.Z

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

func CoordinateEffects(mapCoordinate *coordinate.Coordinate) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT et.id, et.name, et.level, et.type, et.parameter, et.quantity, et.percentages, et.forever "+
		"FROM effects_type et, coordinate_type_effect cte, coordinate_type ct "+
		"WHERE et.id = cte.id_effect AND ct.id=cte.id_type AND ct.id = $1;", mapCoordinate.ID)
	if err != nil {
		println("get coordinate effects")
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var coordinateEffect effect.Effect

		err := rows.Scan(&coordinateEffect.TypeID, &coordinateEffect.Name, &coordinateEffect.Level, &coordinateEffect.Type,
			&coordinateEffect.Parameter, &coordinateEffect.Quantity, &coordinateEffect.Percentages, &coordinateEffect.Forever)
		if err != nil {
			println("get coordinate effects")
			log.Fatal(err)
		}

		mapCoordinate.Effects = append(mapCoordinate.Effects, &coordinateEffect)
	}
}

func AllTypeCoordinate() []*coordinate.Coordinate {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_flore, texture_object, move, view, attack FROM coordinate_type")
	if err != nil {
		println("get all type coordinates")
		log.Fatal(err)
	}

	coordinates := make([]*coordinate.Coordinate, 0)

	for rows.Next() {
		var gameCoordinate coordinate.Coordinate

		rows.Scan(&gameCoordinate.ID, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject, &gameCoordinate.Move,
			&gameCoordinate.View, &gameCoordinate.Attack)

		CoordinateEffects(&gameCoordinate)
		coordinates = append(coordinates, &gameCoordinate)
	}

	return coordinates
}
