package maps

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/effect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"log"
	"strconv"
)

func Maps() map[int]*_map.Map {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"Select " +
		"id, " +
		"name, " +
		"x_size, " +
		"y_size, " +
		"id_type, " +
		"level, " +
		"specification, " +
		"global, " +
		"in_game, " +
		"x_global, " +
		"y_global," +
		"fraction," +
		"possible_battle " +
		"" +
		"FROM maps")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	allMap := make(map[int]*_map.Map)

	for rows.Next() {

		mp := &_map.Map{}

		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.DefaultTypeID, &mp.DefaultLevel, &mp.Specification,
			&mp.Global, &mp.InGame, &mp.XGlobal, &mp.YGlobal, &mp.Fraction, &mp.PossibleBattle)
		if err != nil {
			log.Fatal(err)
		}

		CoordinatesMap(mp)
		GeoData(mp)
		Anomalies(mp)
		Beams(mp)
		Emitters(mp)

		allMap[mp.Id] = mp
	}

	return allMap
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

func Anomalies(mp *_map.Map) {
	mp.Anomalies = make([]*_map.Anomalies, 0)
	rows, err := dbConnect.GetDBConnect().Query(""+
		"Select "+
		"id, "+
		"x, "+
		"y, "+
		"radius,"+
		"type,"+
		"power "+
		""+
		"FROM map_danger_anomalies WHERE id_map = $1", mp.Id)
	if err != nil {
		log.Fatal(err.Error() + "db get anomalies")
	}

	for rows.Next() { // заполняем карту значащами клетками
		var anomaly _map.Anomalies
		err := rows.Scan(&anomaly.ID, &anomaly.X, &anomaly.Y, &anomaly.Radius, &anomaly.Type, &anomaly.Power)
		mp.Anomalies = append(mp.Anomalies, &anomaly)
		if err != nil {
			log.Fatal(err.Error() + "scan geo data")
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

	rows, err := dbConnect.GetDBConnect().Query("SELECT ct.id, mc.x, mc.y, ct.type, ct.texture_flore, "+
		"ct.texture_object, ct.move, ct.view, ct.attack, mc.level, ct.animate_sprite_sheets, ct.animate_loop, "+
		"mc.scale, mc.shadow, mc.rotate, mc.animate_speed, mc.x_offset, mc.y_offset, "+
		"ct.unit_overlap, mc.texture_over_flore, mc.transport, mc.handler, mc.to_positions, mc.to_base_id, mc.to_map_id, "+
		"mc.x_shadow_offset, mc.y_shadow_offset, mc.shadow_intensity, mc.texture_priority, mc.object_priority, "+
		"ct.object_name, ct.object_description, ct.object_inventory, ct.object_hp "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id;", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err.Error() + "map constructor")
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		var positions []byte

		err := rows.Scan(&gameCoordinate.ID, &gameCoordinate.X, &gameCoordinate.Y, &gameCoordinate.Type,
			&gameCoordinate.TextureFlore, &gameCoordinate.TextureObject, &gameCoordinate.Move, &gameCoordinate.View,
			&gameCoordinate.Attack, &gameCoordinate.Level, &gameCoordinate.AnimateSpriteSheets,
			&gameCoordinate.AnimateLoop, &gameCoordinate.Scale,
			&gameCoordinate.Shadow, &gameCoordinate.ObjRotate, &gameCoordinate.AnimationSpeed, &gameCoordinate.XOffset,
			&gameCoordinate.YOffset, &gameCoordinate.UnitOverlap, &gameCoordinate.TextureOverFlore,
			&gameCoordinate.Transport, &gameCoordinate.Handler, &positions,
			&gameCoordinate.ToBaseID, &gameCoordinate.ToMapID, &gameCoordinate.XShadowOffset,
			&gameCoordinate.YShadowOffset, &gameCoordinate.ShadowIntensity, &gameCoordinate.TexturePriority,
			&gameCoordinate.ObjectPriority, &gameCoordinate.ObjectName, &gameCoordinate.ObjectDescription,
			&gameCoordinate.ObjectInventory, &gameCoordinate.ObjectHP)
		if err != nil {
			log.Fatal(err.Error() + "scan map constructor")
		}

		CoordinateEffects(&gameCoordinate)

		err = json.Unmarshal(positions, &gameCoordinate.Positions)
		if err != nil {
			gameCoordinate.Positions = make([]*coordinate.Coordinate, 0)
		}

		if oneLayerMap[gameCoordinate.X] != nil {
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		} else {
			oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
			oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
		}
	}

	//defaultCoordinate := DefaultCoordinateType(mp)
	//for x := 0; x < mp.XSize; x++ { // заполняем карту пустыми клетками тоесть дефолтными по карте
	//	for y := 0; y < mp.YSize; y++ {
	//		_, find := oneLayerMap[x][y]
	//		if !find {
	//
	//			var gameCoordinate coordinate.Coordinate
	//
	//			gameCoordinate = defaultCoordinate
	//			gameCoordinate.ID = mp.DefaultTypeID
	//
	//			gameCoordinate.X = x
	//			gameCoordinate.Y = y
	//
	//			if oneLayerMap[gameCoordinate.X] != nil {
	//				oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
	//			} else {
	//				oneLayerMap[gameCoordinate.X] = make(map[int]*coordinate.Coordinate)
	//				oneLayerMap[gameCoordinate.X][gameCoordinate.Y] = &gameCoordinate
	//			}
	//		}
	//	}
	//}

	mp.OneLayerMap = oneLayerMap
}

func DefaultCoordinateType(mp *_map.Map) coordinate.Coordinate {
	rows, err := dbConnect.GetDBConnect().Query("SELECT type, texture_flore, texture_object, move, view, "+
		"attack, animate_sprite_sheets, animate_loop "+
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
			&gameCoordinate.AnimateLoop)
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
		"attack, animate_sprite_sheets, animate_loop, unit_overlap FROM coordinate_type")
	if err != nil {
		println("get all type coordinates")
		log.Fatal(err)
	}

	coordinates := make([]*coordinate.Coordinate, 0)

	for rows.Next() {
		var gameCoordinate coordinate.Coordinate

		err := rows.Scan(&gameCoordinate.ID, &gameCoordinate.Type, &gameCoordinate.TextureFlore, &gameCoordinate.TextureObject,
			&gameCoordinate.Move, &gameCoordinate.View, &gameCoordinate.Attack, &gameCoordinate.AnimateSpriteSheets,
			&gameCoordinate.AnimateLoop, &gameCoordinate.UnitOverlap)

		if err != nil {
			println("AllTypeCoordinate()")
			log.Fatal(err)
		}

		CoordinateEffects(&gameCoordinate)
		coordinates = append(coordinates, &gameCoordinate)
	}

	return coordinates
}
