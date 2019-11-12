package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"log"
)

func getTypeByID(idType int) *coordinate.Coordinate {
	//rows, err := dbConnect.GetDBConnect().Query("SELECT "+
	//	"id, "+
	//	"type, "+
	//	"texture_flore, "+
	//	"texture_object, "+
	//	"animate_sprite_sheets, "+
	//	"animate_loop, "+
	//	"unit_overlap "+
	//	""+
	//	"FROM coordinate_type WHERE id=$1", idType)
	//if err != nil {
	//	println("get by id type coordinates in map editor")
	//	log.Fatal(err)
	//}
	//
	//var coordinateType coordinate.Coordinate
	//
	//for rows.Next() {
	//	err := rows.Scan(&coordinateType.ID, &coordinateType.Type, &coordinateType.TextureFlore, &coordinateType.TextureObject,
	//		&coordinateType.AnimateSpriteSheets, &coordinateType.AnimateLoop, &coordinateType.UnitOverlap)
	//	if err != nil {
	//		log.Fatal("getTypeByID() " + err.Error())
	//	}
	//}
	//
	//return &coordinateType
	return nil // TODO
}

func getTypeByTerrainAndObject(textureFlore, textureObject, animate string) *coordinate.Coordinate {
	//
	//rows, err := dbConnect.GetDBConnect().Query("SELECT id, type, texture_flore, texture_object, "+
	//	" animate_sprite_sheets, animate_loopFROM coordinate_type "+
	//	"WHERE texture_flore=$1 AND texture_object=$2 AND animate_sprite_sheets=$3",
	//	textureFlore, textureObject, animate)
	//if err != nil {
	//	println("get by Flore and Object coordinates in map editor")
	//	log.Fatal(err)
	//}
	//
	//var coordinateType coordinate.Coordinate
	//
	//for rows.Next() {
	//	err := rows.Scan(&coordinateType.ID, &coordinateType.Type, &coordinateType.TextureFlore, &coordinateType.TextureObject,
	//		&coordinateType.AnimateSpriteSheets, &coordinateType.AnimateLoop)
	//	if err != nil {
	//		log.Fatal("getTypeByTerrainAndObject() " + err.Error())
	//	}
	//}
	//
	//if coordinateType.ID != 0 {
	//	return &coordinateType
	//} else {
	//	return nil
	//}

	// TODO
	return nil
}

// берет координату из таблицы map_constructor, если ее там нет то вернут nil
func getMapCoordinateInMC(idMap, x, y int) *coordinate.Coordinate {

	var idType int
	var id int
	var rotate int
	var animateSpeed int

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, id_type, rotate, animate_speed "+
		"FROM map_constructor "+
		"WHERE id_map = $1 AND x=$2 AND y = $3",
		idMap, x, y)
	if err != nil {
		log.Fatal("get mc coordinate in editor map " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &idType, &rotate, &animateSpeed)
		if err != nil {
			log.Fatal("getMapCoordinateInMC() " + err.Error())
		}
	}

	if id == 0 {
		return nil
	} else {
		mcCoordinate := getTypeByID(idType)
		mcCoordinate.X = x
		mcCoordinate.Y = y
		//mcCoordinate.ObjRotate = rotate
		//mcCoordinate.AnimationSpeed = animateSpeed

		return mcCoordinate
	}
}

func getSizeMap(idMap int) (xSize, ySize int) {

	rows, err := dbConnect.GetDBConnect().Query("SELECT x_size, y_size FROM maps WHERE id = $1", idMap)
	if err != nil {
		log.Fatal("get default level and type map in editor map " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&xSize, &ySize)
	}

	return xSize, ySize
}
