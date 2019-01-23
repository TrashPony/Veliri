package globalGame

import (
	"../../mechanics/factories/maps"
	"time"
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/globalGame"
	"../../mechanics/gameObjects/map"
)

func HandlersLife() {
	allMaps := maps.Maps.GetAllMap()

	for _, mp := range allMaps {
		for _, coor := range mp.HandlersCoordinates {
			if coor.Handler == "sector" {
				go entranceMonitor(coor, mp)
			}
		}
	}
}

func entranceMonitor(coor *coordinate.Coordinate, mp *_map.Map) {
	for {
		time.Sleep(100 * time.Millisecond)

		xEntry, yEntry := globalGame.GetXYCenterHex(coor.Q, coor.R)
		checkHandlerCoordinate(xEntry, yEntry, mp.Id)

		xOut, yOut := globalGame.GetXYCenterHex(coor.ToQ, coor.ToR)
		checkHandlerCoordinate(xOut, yOut, mp.Id)

		// TODO мониторинг входов что бы выталкивать афк игроков под эвакуаторы
		// TODO мониторинг выходов что бы отключать включать телепорт
	}
}

func checkHandlerCoordinate(x, y, mapID int) bool {

	return false
}
