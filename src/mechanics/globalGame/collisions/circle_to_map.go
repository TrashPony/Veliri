package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func CircleAllCollisionCheck(xCenter, yCenter, radius int, mp *_map.Map, units map[int]*unit.ShortUnitInfo, boxs []*boxInMap.Box) bool {

	if units == nil {
		units = globalGame.Clients.GetAllShortUnits(mp.Id)
	}

	if boxs == nil {
		boxs = boxes.Boxes.GetAllBoxByMapID(mp.Id)
	}

	// статичные обьекты на карте
	if CircleStaticMap(xCenter, yCenter, radius, mp) {
		return true
	}

	//динамические обьекты
	collision := CircleDynamicMap(xCenter, yCenter, radius, mp)
	if collision {
		return true
	}

	// глобальная гео дата
	if CircleGlobalGeoDataMap(xCenter, yCenter, radius, mp) {
		return true
	}

	// руды
	if CircleReservoirMap(xCenter, yCenter, radius, mp) {
		return true
	}

	// все юниты
	if CircleUnits(xCenter, yCenter, radius, units) {
		return true
	}

	// все ящики
	if CircleBoxes(xCenter, yCenter, radius, boxs) {
		return true
	}

	return false
}

func CircleStaticMap(xCenter, yCenter, radius int, mp *_map.Map) bool {
	// статичные обьекты на карте
	for _, x := range mp.StaticObjects {
		for _, sObj := range x {
			for _, sGeoPoint := range sObj.GeoData {
				distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, sGeoPoint.X, sGeoPoint.Y)
				if int(distToObstacle) < sGeoPoint.Radius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
					return true
				}
			}
		}
	}

	return false
}

func CircleDynamicMap(xCenter, yCenter, radius int, mp *_map.Map) bool {
	// динамические обьекты
	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, sObj := range x {
			for _, sGeoPoint := range sObj.GeoData {
				distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, sGeoPoint.X, sGeoPoint.Y)
				if int(distToObstacle) < sGeoPoint.Radius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
					return true
				}
			}
		}
	}

	return false
}

func CircleGlobalGeoDataMap(xCenter, yCenter, radius int, mp *_map.Map) bool {
	for _, sGeoPoint := range mp.GeoData {
		distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, sGeoPoint.X, sGeoPoint.Y)
		if int(distToObstacle) < sGeoPoint.Radius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
			return true
		}
	}
	return false
}

func CircleReservoirMap(xCenter, yCenter, radius int, mp *_map.Map) bool {

	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {
			distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, reservoir.X, reservoir.Y)
			if int(distToObstacle) < reservoirRadius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
				return true
			}
		}
	}
	return false
}

func CircleUnits(xCenter, yCenter, radius int, units map[int]*unit.ShortUnitInfo) bool {

	for _, gameUnit := range units {
		rect := GetBodyRect(gameUnit.Body, float64(gameUnit.X), float64(gameUnit.Y), gameUnit.Rotate, false, false)
		if rect.detectCollisionRectToCircle(&point{x: float64(xCenter), y: float64(yCenter)}, radius) {
			return true
		}
	}

	return false
}

func CircleBoxes(xCenter, yCenter, radius int, boxs []*boxInMap.Box) bool {
	for _, mapBox := range boxs {
		rect := GetCenterRect(float64(mapBox.X), float64(mapBox.Y), float64(mapBox.Height), float64(mapBox.Width))
		rect.Rotate(mapBox.Rotate)
		if rect.detectCollisionRectToCircle(&point{x: float64(xCenter), y: float64(yCenter)}, radius) {
			return true
		}
	}

	return false
}
