package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func CreateNewTerrain(terrainName string) bool {
	findType := getTypeByTerrainAndObject(terrainName, "", "")
	if findType != nil {
		return false
	} else {

		AddNewTypeCoordinate("", terrainName, "",
			"", false, true, true, true, 0, 0, false)

		addJSPreloadFile(terrainName, "terrain", 0, 0, 0)

		return true
	}
}

func CreateNewObject(objectName, animateName string, move, watch, attack bool, radius int, scale int, shadow bool, countSprites, xSize, ySize int) bool {
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
				"", false, move, watch, attack, radius, scale, shadow)

			addJSPreloadFile(objectName, "objects", 0, 0, 0)

			return true
		}
	} else {
		rows, err := dbConnect.GetDBConnect().Query("SELECT id FROM coordinate_type WHERE animate_sprite_sheets=$1", animateName)
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
			AddNewTypeCoordinate("", "desert", "",
				animateName, true, move, watch, attack, radius, scale, shadow)

			addJSPreloadFile(animateName, "animate", countSprites, xSize, ySize)

			return true
		}
	}
}

func addJSPreloadFile(fileName, typeName string, countSprites, xSize, ySize int) {
	autoFile, err := ioutil.ReadFile("src/static/game/preloadAutoGenerated.js")
	if err != nil {
		log.Fatalln(err)
	}

	fileString := string(autoFile)
	lines := strings.Split(fileString, "\n")

	for i, line := range lines {
		if strings.Contains(line, "}") {

			if typeName == "objects" || typeName == "terrain" {
				lines[i] = "game.load.image('" + fileName + "', 'http://' " +
					"+ window.location.host + " +
					"'/assets/map/" + typeName + "/" + fileName + ".png'); \n}"
			}

			if typeName == "animate" {
				lines[i] = "game.load.spritesheet('" + fileName + "', 'http://' " +
					"+ window.location.host + " +
					"'/assets/map/" + typeName + "/" + fileName + ".png', " +
					strconv.Itoa(xSize) + "," + strconv.Itoa(ySize) + "," + strconv.Itoa(countSprites) + "); \n}"
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("src/static/game/preloadAutoGenerated.js", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
