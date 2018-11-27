package mapEditor

import (
	"../../../dbConnect"
	"log"
)

// TODO генерить новые типы сразу в движок

func CreateNewTerrain(terrainName string) bool {
	findType := getTypeByTerrainAndObject(terrainName, "", "")
	if findType != nil {
		return false
	} else {

		AddNewTypeCoordinate("", terrainName, "",
			"", false, true, true, true, 0)

		return true
	}
}

func CreateNewObject(objectName, animateName string, move, watch, attack bool, radius int) bool {
	if objectName != "" {
		rows, err := dbConnect.GetDBConnect().Query("SELECT id FROM coordinate_type WHERE texture_object=$1", objectName)
		if err != nil {
			println("get by Object coordinates in map editor")
			log.Fatal(err)
		}

		var id int
		if rows.Next() {
			rows.Scan(&id)
		}

		if id != 0 {
			return false
		} else {
			// desert тип по умолчанию
			AddNewTypeCoordinate("", "desert", objectName,
				"", false, move, watch, attack, radius)

			return true
		}
	} else {
		// todo
		return false
	}
}
