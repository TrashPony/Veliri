package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func CircleAllCollisionCheck(xCenter, yCenter, radius int, mp *_map.Map, units map[int]*unit.ShortUnitInfo, boxs []*boxInMap.Box) (bool, string, int) {

	if units == nil {
		units = globalGame.Clients.GetAllShortUnits(mp.Id)
	}

	if boxs == nil {
		boxs = boxes.Boxes.GetAllBoxByMapID(mp.Id)
	}

	// статичные обьекты на карте
	if CircleStaticMap(xCenter, yCenter, radius, mp) {
		return true, "static", 0
	}

	//динамические обьекты
	collision, obj := CircleDynamicMap(xCenter, yCenter, radius, mp)
	if collision {
		return true, "object", obj.ID
	}

	// глобальная гео дата
	if CircleGlobalGeoDataMap(xCenter, yCenter, radius, mp) {
		return true, "static", 0
	}

	// руды
	collision, reservoir := CircleReservoirMap(xCenter, yCenter, radius, mp)
	if collision {
		return true, "reservoir", reservoir.ID
	}

	// все юниты
	collision, gameUnit := CircleUnits(xCenter, yCenter, radius, units)
	if collision {
		return true, "unit", gameUnit.ID
	}

	// все ящики
	collision, mapBox := CircleBoxes(xCenter, yCenter, radius, boxs)
	if collision {
		return true, "box", mapBox.ID
	}

	return false, "", 0
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

func CircleDynamicMap(xCenter, yCenter, radius int, mp *_map.Map) (bool, *dynamic_map_object.Object) {
	// динамические обьекты
	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, sObj := range x {
			for _, sGeoPoint := range sObj.GeoData {
				distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, sGeoPoint.X, sGeoPoint.Y)
				if int(distToObstacle) < sGeoPoint.Radius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
					return true, sObj
				}
			}
		}
	}

	return false, nil
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

func CircleReservoirMap(xCenter, yCenter, radius int, mp *_map.Map) (bool, *resource.Map) {

	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {
			distToObstacle := game_math.GetBetweenDist(xCenter, yCenter, reservoir.X, reservoir.Y)
			if int(distToObstacle) < reservoirRadius+radius { // если растония меньше чем обра радиуса значит окружности пересекается
				return true, reservoir
			}
		}
	}
	return false, nil
}

func CircleUnits(xCenter, yCenter, radius int, units map[int]*unit.ShortUnitInfo) (bool, *unit.ShortUnitInfo) {

	for _, gameUnit := range units {
		rect := GetBodyRect(gameUnit.Body, float64(gameUnit.X), float64(gameUnit.Y), gameUnit.Rotate, false, false)
		if rect.detectCollisionRectToCircle(&point{x: float64(xCenter), y: float64(yCenter)}, radius) {
			return true, gameUnit
		}
	}

	return false, nil
}

func CircleBoxes(xCenter, yCenter, radius int, boxs []*boxInMap.Box) (bool, *boxInMap.Box) {
	for _, mapBox := range boxs {
		rect := GetCenterRect(float64(mapBox.X), float64(mapBox.Y), float64(mapBox.Height), float64(mapBox.Width))
		rect.Rotate(mapBox.Rotate)
		if rect.detectCollisionRectToCircle(&point{x: float64(xCenter), y: float64(yCenter)}, radius) {
			return true, mapBox
		}
	}

	return false, nil
}
