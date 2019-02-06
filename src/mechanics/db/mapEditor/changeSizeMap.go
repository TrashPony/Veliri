package mapEditor

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"log"
)

// ------ ROWS ------- //
func AddStartRow(mapID int) {
	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	rOffset(1, mapID)
	qSize, rSize := getSizeMap(mapID)
	rSize++
	setNewSizeMap(qSize, rSize, mapID)

	// проверяем занятое пространство обьемников после перестановки
	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func AddEndRow(mapID int) {
	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	qSize, rSize := getSizeMap(mapID)
	rSize++
	setNewSizeMap(qSize, rSize, mapID)

	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func RemoveStartRow(mapID int) {

	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	removeAllRCoordinate(0, mapID)

	rOffset(-1, mapID)

	qSize, rSize := getSizeMap(mapID)
	rSize--
	setNewSizeMap(qSize, rSize, mapID)

	// проверяем занятое пространство обьемников после перестановки
	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func RemoveEndRow(mapID int) {
	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	qSize, rSize := getSizeMap(mapID)
	removeAllRCoordinate(rSize-1, mapID) // -1 потому что отсчет на карте идет с нуля
	rSize--
	setNewSizeMap(qSize, rSize, mapID)

	// проверяем занятое пространство обьемников после перестановки
	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

// ------ Columns ------- //
func AddStartColumn(mapID int) {

	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	qOffset(1, mapID)
	qSize, rSize := getSizeMap(mapID)
	qSize++
	setNewSizeMap(qSize, rSize, mapID)

	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func AddEndColumn(mapID int) {

	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	qSize, rSize := getSizeMap(mapID)
	qSize++
	setNewSizeMap(qSize, rSize, mapID)

	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func RemoveStartColumn(mapID int) {

	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	removeAllQCoordinate(0, mapID)

	qOffset(-1, mapID)

	qSize, rSize := getSizeMap(mapID)
	qSize--
	setNewSizeMap(qSize, rSize, mapID)

	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
}

func RemoveEndColumn(mapID int) {

	// удаляем все большим обьектам импакт что бы они нормально передвинулись
	removeBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)

	qSize, rSize := getSizeMap(mapID)
	removeAllQCoordinate(qSize-1, mapID) // -1 потому что отсчет на карте идет с нуля
	qSize--
	setNewSizeMap(qSize, rSize, mapID)

	clipBigObjectsImpact(getMapALLCoordinateInMC(mapID), mapID)
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

func removeBigObjectsImpact(coordinates []*coordinate.Coordinate, idMap int) {
	for _, mcCoordinates := range coordinates {
		if mcCoordinates.ImpactRadius > 0 && mcCoordinates.Impact == nil {
			removeImpact(mcCoordinates, idMap)
		}
	}
}

func clipBigObjectsImpact(coordinates []*coordinate.Coordinate, idMap int) {
	for _, mcCoordinates := range coordinates {
		if mcCoordinates.ImpactRadius > 0 && mcCoordinates.Impact == nil {
			placeRadiusCoordinate(mcCoordinates, idMap, false)
		}
	}
}
