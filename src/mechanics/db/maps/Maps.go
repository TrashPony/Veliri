package maps

import (
	"encoding/json"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/obstacle_point"
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
		"level, " +
		"specification, " +
		"global, " +
		"x_global, " +
		"y_global," +
		"fraction," +
		"possible_battle " +
		"" +
		"FROM maps")
	if err != nil {
		log.Fatal(err.Error() + " get all maps")
	}
	defer rows.Close()

	allMap := make(map[int]*_map.Map)

	for rows.Next() {

		mp := &_map.Map{}

		err := rows.Scan(&mp.Id, &mp.Name, &mp.XSize, &mp.YSize, &mp.DefaultLevel, &mp.Specification,
			&mp.Global, &mp.XGlobal, &mp.YGlobal, &mp.Fraction, &mp.PossibleBattle)
		if err != nil {
			log.Fatal(err.Error() + " scan all maps")
		}

		GetFlore(mp)
		CoordinatesMap(mp)
		GeoData(mp)
		Anomalies(mp)
		Beams(mp)
		Emitters(mp)
		mp.StaticObjects = GetObjects(mp, "ct.object_hp < 0;")
		mp.DynamicObjects = GetObjects(mp, "ct.object_hp > -1;")

		allMap[mp.Id] = mp
	}

	return allMap
}

func GetFlore(mp *_map.Map) {
	mp.Flore = make(map[int]map[int]*dynamic_map_object.Flore)

	rows, err := dbConnect.GetDBConnect().Query("SELECT x, y, texture_over_flore, texture_priority "+
		"FROM map_constructor "+
		"WHERE id_map = $1 AND texture_over_flore != ''", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err.Error() + "get map flor")
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var flore dynamic_map_object.Flore

		err := rows.Scan(&flore.X, &flore.Y, &flore.TextureOverFlore, &flore.TexturePriority)
		if err != nil {
			log.Fatal(err.Error() + "scan map flor")
		}

		if mp.Flore[flore.X] != nil {
			mp.Flore[flore.X][flore.Y] = &flore
		} else {
			mp.Flore[flore.X] = make(map[int]*dynamic_map_object.Flore)
			mp.Flore[flore.X][flore.Y] = &flore
		}
	}
}

func GetObjects(mp *_map.Map, objType string) map[int]map[int]*dynamic_map_object.Object {
	objMap := make(map[int]map[int]*dynamic_map_object.Object)

	rows, err := dbConnect.GetDBConnect().Query("SELECT ct.id, mc.x, mc.y, ct.type, "+
		"ct.texture_object, ct.animate_sprite_sheets, ct.animate_loop, "+
		"mc.scale, ct.shadow, mc.rotate, ct.animate_speed, "+
		"ct.unit_overlap, "+
		"mc.x_shadow_offset, mc.y_shadow_offset, ct.shadow_intensity, mc.object_priority, "+
		"ct.object_name, ct.object_description, ct.object_inventory, ct.object_hp, ct.geo_data "+
		"FROM map_constructor mc, coordinate_type ct "+
		"WHERE mc.id_map = $1 AND mc.id_type = ct.id AND "+objType, strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err.Error() + "get map obj")
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var obj dynamic_map_object.Object
		var geoData []byte

		err := rows.Scan(&obj.Type, &obj.X, &obj.Y, &obj.Type,
			&obj.Texture, &obj.AnimateSpriteSheets,
			&obj.AnimateLoop, &obj.Scale, &obj.Shadow, &obj.Rotate,
			&obj.AnimationSpeed, &obj.UnitOverlap,
			&obj.XShadowOffset, &obj.YShadowOffset, &obj.ShadowIntensity,
			&obj.Priority, &obj.Name, &obj.Description,
			&obj.Inventory, &obj.HP, &geoData)
		if err != nil {
			log.Fatal(err.Error() + "scan map obj")
		}

		err = json.Unmarshal(geoData, &obj.GeoData)
		if err != nil {
			obj.GeoData = make([]*obstacle_point.ObstaclePoint, 0)
		} else {
			obj.SetGeoData()
		}

		idString := strconv.Itoa(obj.X) + strconv.Itoa(obj.Y)
		obj.ID, _ = strconv.Atoi(idString)

		if objMap[obj.X] != nil {
			objMap[obj.X][obj.Y] = &obj
		} else {
			objMap[obj.X] = make(map[int]*dynamic_map_object.Object)
			objMap[obj.X][obj.Y] = &obj
		}
	}

	return objMap
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
	mp.GeoData = make([]*obstacle_point.ObstaclePoint, 0)
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
		var obstaclePoint obstacle_point.ObstaclePoint
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

	rows, err := dbConnect.GetDBConnect().Query("SELECT x, y, "+
		"transport, handler, to_positions, to_base_id, to_map_id "+
		"FROM map_constructor "+
		"WHERE id_map = $1", strconv.Itoa(mp.Id))

	if err != nil {
		log.Fatal(err.Error() + "map constructor")
	}

	defer rows.Close()

	for rows.Next() { // заполняем карту значащами клетками
		var gameCoordinate coordinate.Coordinate
		var positions []byte

		err := rows.Scan(&gameCoordinate.X, &gameCoordinate.Y,
			&gameCoordinate.Transport, &gameCoordinate.Handler, &positions,
			&gameCoordinate.ToBaseID, &gameCoordinate.ToMapID)
		if err != nil {
			log.Fatal(err.Error() + "scan map constructor")
		}

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

	mp.OneLayerMap = oneLayerMap
}

func AllTypeCoordinate() []*dynamic_map_object.Object {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_object, " +
		" animate_sprite_sheets, animate_loop, unit_overlap, object_name, object_description," +
		" object_inventory, object_hp, shadow_intensity, animate_speed, shadow, geo_data FROM coordinate_type")
	if err != nil {
		log.Fatal(err.Error() + "get all type object")
	}

	objS := make([]*dynamic_map_object.Object, 0)

	for rows.Next() {
		var obj dynamic_map_object.Object
		var geoData []byte

		err := rows.Scan(&obj.TypeID, &obj.Type, &obj.Texture, &obj.AnimateSpriteSheets, &obj.AnimateLoop,
			&obj.UnitOverlap, &obj.Name, &obj.Description, &obj.Inventory, &obj.HP, &obj.ShadowIntensity,
			&obj.AnimationSpeed, &obj.Shadow, &geoData)

		if err != nil {
			log.Fatal(err.Error() + " scan all type coorinate")
		}

		err = json.Unmarshal(geoData, &obj.GeoData)
		if err != nil {
			obj.GeoData = make([]*obstacle_point.ObstaclePoint, 0)
		} else {
			obj.SetGeoData()
		}

		objS = append(objS, &obj)
	}

	return objS
}
