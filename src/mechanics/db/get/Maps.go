package get

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
	"strconv"
	"strings"
)

func Maps() map[int]*_map.Map {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"Select " +
		"id, " +
		"name, " +
		"q_size, " +
		"r_size, " +
		"id_type, " +
		"level, " +
		"specification, " +
		"global, " +
		"in_game " +
		"" +
		"FROM maps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	allMap := make(map[int]*_map.Map)

	for rows.Next() {

		mp := &_map.Map{}

		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification,
			&mp.Global, &mp.InGame)
		if err != nil {
			log.Fatal(err)
		}

		CoordinatesMap(mp)
		GeoData(mp)
		Beams(mp)
		Emitters(mp)

		allMap[mp.Id] = mp
	}

	return allMap
}

func MapByID(id int) *_map.Map {
	rows, err := dbConnect.GetDBConnect().Query(""+
		"Select "+
		"id, "+
		"name, "+
		"q_size, "+
		"r_size, "+
		"id_type, "+
		"level, "+
		"specification, "+
		"global, "+
		"in_game "+
		""+
		"FROM maps WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var mp _map.Map

		err := rows.Scan(&mp.Id, &mp.Name, &mp.QSize, &mp.RSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification,
			&mp.Global, &mp.InGame)
		if err != nil {
			log.Fatal(err)
		}

		CoordinatesMap(&mp)
		GeoData(&mp)
		Beams(&mp)
		Emitters(&mp)

		return &mp
	}

	return nil
}

func Emitters(mp *_map.Map) {
	mp.Emitters = make([]*_map.Emitter, 0)
	rows, err := dbConnect.GetDBConnect().Query(""+
		"Select "+
		"id, "+
		"x, "+
		"y, "+
		"min_scale, "+
		"max_scale, "+
		"min_speed, "+
		"max_speed, "+
		"ttl, "+
		"width, "+
		"height, "+
		"color, "+
		"frequency, "+
		"min_alpha, "+
		"max_alpha, "+
		"animate, "+
		"animate_speed, "+
		"name_particle, "+
		"alpha_loop_time, "+
		"yoyo "+
		""+
		"FROM map_emitters WHERE id_map = $1", mp.Id)
	if err != nil {
		log.Fatal(err.Error() + "db get emitters")
	}

	for rows.Next() {
		var emitter _map.Emitter

		err := rows.Scan(&emitter.ID, &emitter.X, &emitter.Y, &emitter.MinScale, &emitter.MaxScale, &emitter.MinSpeed,
			&emitter.MaxSpeed, &emitter.TTL, &emitter.Width, &emitter.Height, &emitter.Color, &emitter.Frequency,
			&emitter.MinAlpha, &emitter.MaxAlpha, &emitter.Animate, &emitter.AnimateSpeed, &emitter.NameParticle,
			&emitter.AlphaLoopTime, &emitter.Yoyo)
		mp.Emitters = append(mp.Emitters, &emitter)
		if err != nil {
			log.Fatal(err.Error() + "scan emitters")
		}
	}
	defer rows.Close()
}

func Beams(mp *_map.Map) {
	mp.Beams = make([]*_map.Beam, 0)
	rows, err := dbConnect.GetDBConnect().Query(""+
		"Select "+
		"id, "+
		"x_start, "+
		"y_start, "+
		"x_end, "+
		"y_end, "+
		"color "+
		""+
		"FROM map_beams WHERE id_map = $1", mp.Id)
	if err != nil {
		log.Fatal(err.Error() + "db get beam")
	}

	for rows.Next() {
		var beam _map.Beam

		err := rows.Scan(&beam.ID, &beam.XStart, &beam.YStart, &beam.XEnd, &beam.YEnd, &beam.Color)
		mp.Beams = append(mp.Beams, &beam)
		if err != nil {
			log.Fatal(err.Error() + "scan beam")
		}
	}
	defer rows.Close()
}

func GeoData(mp *_map.Map) {
	mp.GeoData = make([]*_map.ObstaclePoint, 0)
	rows, err := dbConnect.GetDBConnect().Query(""+
		"Select "+
		"id, "+
		"x, "+
		"y, "+
		"radius "+
		""+
		"FROM global_geo_data WHERE id_map = $1", mp.Id)
	if err != nil {
		log.Fatal(err.Error() + "db get geo data")
	}

	for rows.Next() { // заполняем карту значащами клетками
		var obstaclePoint _map.ObstaclePoint
		err := rows.Scan(&obstaclePoint.ID, &obstaclePoint.X, &obstaclePoint.Y, &obstaclePoint.Radius)
		mp.GeoData = append(mp.GeoData, &obstaclePoint)
		if err != nil {
			log.Fatal(err.Error() + "scan geo data")
		}
	}
	defer rows.Close()
}

func CoordinatesMap(mp *_map.Map) {
	oneLayerMap := make(map[int]map[int]*coordinate.Coordinate)

	rows, err := dbConnect.GetDBConnect().Query("SELECT ct.id, mc.q, mc.r, ct.type, ct.texture_flore, "+
		"ct.texture_object, ct.move, ct.view, ct.attack, mc.level, ct.animate_sprite_sheets, ct.animate_loop, "+
		"ct.impact_radius, mc.impact, ct.scale, ct.shadow, mc.rotate, mc.animate_speed, mc.x_offset, mc.y_offset, "+
		"ct.unit_overlap, mc.texture_over_flore, mc.transport, mc.handler, mc.to_q, mc.to_r, mc.to_base_id, mc.to_map_id "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id;", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err.Error() + "map constructor")
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var impact string
		var gameCoordinate coordinate.Coordinate
		err := rows.Scan(&gameCoordinate.ID, &gameCoordinate.Q, &gameCoordinate.R, &gameCoordinate.Type,
			&gameCoordinate.TextureFlore, &gameCoordinate.TextureObject, &gameCoordinate.Move, &gameCoordinate.View,
			&gameCoordinate.Attack, &gameCoordinate.Level, &gameCoordinate.AnimateSpriteSheets,
			&gameCoordinate.AnimateLoop, &gameCoordinate.ImpactRadius, &impact, &gameCoordinate.Scale,
			&gameCoordinate.Shadow, &gameCoordinate.ObjRotate, &gameCoordinate.AnimationSpeed, &gameCoordinate.XOffset,
			&gameCoordinate.YOffset, &gameCoordinate.UnitOverlap, &gameCoordinate.TextureOverFlore,
			&gameCoordinate.Transport, &gameCoordinate.Handler, &gameCoordinate.ToQ, &gameCoordinate.ToR,
			&gameCoordinate.ToBaseID, &gameCoordinate.ToMapID)
		if err != nil {
			log.Fatal(err.Error() + "scan map constructor")
		}

		gameCoordinate.Impact = ParseTarget(impact)

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
	rows, err := dbConnect.GetDBConnect().Query("SELECT type, texture_flore, texture_object, move, view, "+
		"attack, animate_sprite_sheets, animate_loop, impact_radius "+
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
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.AnimateSpriteSheets,
			&gameCoordinate.AnimateLoop, &gameCoordinate.ImpactRadius)
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
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_flore, texture_object, move, view, " +
		"attack, animate_sprite_sheets, animate_loop, impact_radius, scale, shadow, unit_overlap FROM coordinate_type")
	if err != nil {
		println("get all type coordinates")
		log.Fatal(err)
	}

	coordinates := make([]*coordinate.Coordinate, 0)

	for rows.Next() {
		var gameCoordinate coordinate.Coordinate

		rows.Scan(&gameCoordinate.ID, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.AnimateSpriteSheets,
			&gameCoordinate.AnimateLoop, &gameCoordinate.ImpactRadius, &gameCoordinate.Scale, &gameCoordinate.Shadow,
			&gameCoordinate.UnitOverlap)

		CoordinateEffects(&gameCoordinate)
		coordinates = append(coordinates, &gameCoordinate)
	}

	return coordinates
}

func ParseTarget(targetKey string) *coordinate.Coordinate {
	targetCell := strings.Split(targetKey, ":")

	if len(targetCell) > 1 { // устанавливаем таргет если он есть
		q, ok := strconv.Atoi(targetCell[0])
		r, ok := strconv.Atoi(targetCell[1])
		if ok == nil {
			target := coordinate.Coordinate{Q: q, R: r}
			return &target
		} else {
			return nil
		}
	} else {
		return nil
	}
}
