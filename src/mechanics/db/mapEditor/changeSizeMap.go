package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"log"
)

// ------ ROWS ------- //
func AddStartRow(mapID int) {
	rOffset(1, mapID)
	qSize, rSize := getSizeMap(mapID)
	rSize++
	setNewSizeMap(qSize, rSize, mapID)
}

func AddEndRow(mapID int) {
	qSize, rSize := getSizeMap(mapID)
	rSize++
	setNewSizeMap(qSize, rSize, mapID)
}

func RemoveStartRow(mapID int) {
	removeAllRCoordinate(0, mapID)

	rOffset(-1, mapID)

	qSize, rSize := getSizeMap(mapID)
	rSize--
	setNewSizeMap(qSize, rSize, mapID)
}

func RemoveEndRow(mapID int) {
	qSize, rSize := getSizeMap(mapID)
	removeAllRCoordinate(rSize-1, mapID) // -1 потому что отсчет на карте идет с нуля
	rSize--
	setNewSizeMap(qSize, rSize, mapID)
}

// ------ Columns ------- //
func AddStartColumn(mapID int) {
	qOffset(1, mapID)
	qSize, rSize := getSizeMap(mapID)
	qSize++
	setNewSizeMap(qSize, rSize, mapID)
}

func AddEndColumn(mapID int) {

	qSize, rSize := getSizeMap(mapID)
	qSize++
	setNewSizeMap(qSize, rSize, mapID)
}

func RemoveStartColumn(mapID int) {
	removeAllQCoordinate(0, mapID)

	qOffset(-1, mapID)

	qSize, rSize := getSizeMap(mapID)
	qSize--
	setNewSizeMap(qSize, rSize, mapID)
}

func RemoveEndColumn(mapID int) {
	qSize, rSize := getSizeMap(mapID)
	removeAllQCoordinate(qSize-1, mapID) // -1 потому что отсчет на карте идет с нуля
	qSize--
	setNewSizeMap(qSize, rSize, mapID)
}

// -------------------//
func rOffset(offset, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET r = r + $1 WHERE id_map=$2",
		offset, mapID)
	if err != nil {
		log.Fatal("offset r coordinates map " + err.Error())
	}
}

func qOffset(offset, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE map_constructor SET q = q + $1 WHERE id_map=$2",
		offset, mapID)
	if err != nil {
		log.Fatal("offset q coordinates map " + err.Error())
	}
}

func removeAllQCoordinate(q, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE q = $1 AND id_map = $2",
		q, mapID)
	if err != nil {
		log.Fatal("delete q coordinate " + err.Error())
	}
}

func removeAllRCoordinate(r, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("DELETE FROM map_constructor WHERE r = $1 AND id_map = $2",
		r, mapID)
	if err != nil {
		log.Fatal("delete r coordinate " + err.Error())
	}
}

func setNewSizeMap(qSize, rSize, mapID int) {
	_, err := dbConnect.GetDBConnect().Exec("UPDATE maps SET q_size = $1, r_size = $2 WHERE id=$3",
		qSize, rSize, mapID)
	if err != nil {
		log.Fatal("update size map " + err.Error())
	}
}