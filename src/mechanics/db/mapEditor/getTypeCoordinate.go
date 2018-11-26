package mapEditor

import (
	"../../../dbConnect"
	"../../gameObjects/coordinate"
	"log"
)

func getTypeByID(idType int) *coordinate.Coordinate {
	rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_flore, texture_object, move, view, "+
		"attack, animate_sprite_sheets, animate_loop, impact_radius FROM coordinate_type WHERE id=$1", idType)
	if err != nil {
		println("get by id type coordinates in map editor")
		log.Fatal(err)
	}

	var coordinateType coordinate.Coordinate

	for rows.Next() {
		rows.Scan(&coordinateType.ID, &coordinateType.Type, &coordinateType.TextureFlore, &coordinateType.TextureObject,
			&coordinateType.Move, &coordinateType.View, &coordinateType.Attack, &coordinateType.AnimateSpriteSheets,
			&coordinateType.AnimateLoop, &coordinateType.ImpactRadius)
	}

	return &coordinateType
}

func getTypeByTerrainAndObject(textureFlore, textureObject, animate string) *coordinate.Coordinate {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_flore, texture_object, move, view, "+
		"attack, animate_sprite_sheets, animate_loop, impact_radius FROM coordinate_type WHERE texture_flore=$1 AND texture_object=$2 AND animate_sprite_sheets=$3",
		textureFlore, textureObject, animate)
	if err != nil {
		println("get by Flore and Object coordinates in map editor")
		log.Fatal(err)
	}

	var coordinateType coordinate.Coordinate

	for rows.Next() {
		rows.Scan(&coordinateType.ID, &coordinateType.Type, &coordinateType.TextureFlore, &coordinateType.TextureObject,
			&coordinateType.Move, &coordinateType.View, &coordinateType.Attack, &coordinateType.AnimateSpriteSheets,
			&coordinateType.AnimateLoop, &coordinateType.ImpactRadius)
	}

	if coordinateType.ID != 0 {
		return &coordinateType
	} else {
		return nil
	}
}

// берет координату из таблицы map_constructor, если ее там нет то вернут nil
func getMapCoordinateInMC(idMap, q, r int) *coordinate.Coordinate {

	var level int
	var idType int
	var id int

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, level, id_type FROM map_constructor WHERE id_map = $1 AND q=$2 AND r = $3",
		idMap, q, r)
	if err != nil {
		log.Fatal("get mc coordinate in editor map " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &level, &idType)
	}

	if id == 0 {
		return nil
	} else {
		mcCoordinate := getTypeByID(idType)
		mcCoordinate.Level = level

		return mcCoordinate
	}
}

func getDefaultMap(idMap int) (level, coordinateType int) {
	var defaultLevel int
	var defaultType int

	rows, err := dbConnect.GetDBConnect().Query("SELECT level, id_type FROM maps WHERE id = $1", idMap)
	if err != nil {
		log.Fatal("get default level and type map in editor map " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&defaultLevel, &defaultType)
	}

	return defaultLevel, defaultType
}

func getSizeMap(idMap int) (qSize, rSize int) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT q_size, r_size FROM maps WHERE id = $1", idMap)
	if err != nil {
		log.Fatal("get default level and type map in editor map " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&qSize, &rSize)
	}

	return qSize, rSize
}
