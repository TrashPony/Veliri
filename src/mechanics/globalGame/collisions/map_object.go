package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dynamic_map_object"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"strings"
)

func CheckObjectCollision(obj *dynamic_map_object.Object, mp *_map.Map, structCheck bool) bool {
	// это супер мега дорогой метод, однако он используется только для популяции растений и вызывается редко
	// TODO а еще лень думоть

	// статичные обьекты на карте
	for _, x := range mp.StaticObjects {
		for _, sObj := range x {
			for _, sGeoPoint := range sObj.GeoData {
				for _, GeoPoint := range obj.GeoData {
					distToObstacle := game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, sGeoPoint.X, sGeoPoint.Y)
					if int(distToObstacle) < sGeoPoint.Radius+GeoPoint.Radius { // если растония меньше чем обра радиуса значит окружности пересекается
						return true
					}
				}
			}
		}
	}

	// динамические обьекты
	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, sObj := range x {

			if obj.ID == sObj.ID {
				continue // исключаем себя
			}

			for _, sGeoPoint := range sObj.GeoData {
				for _, GeoPoint := range obj.GeoData {
					distToObstacle := game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, sGeoPoint.X, sGeoPoint.Y)
					if int(distToObstacle) < sGeoPoint.Radius+GeoPoint.Radius { // если растония меньше чем обра радиуса значит окружности пересекается
						return true
					}
				}
			}
		}
	}

	// глобальная гео дата
	for _, sGeoPoint := range mp.GeoData {
		for _, GeoPoint := range obj.GeoData {
			distToObstacle := game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, sGeoPoint.X, sGeoPoint.Y)
			if int(distToObstacle) < sGeoPoint.Radius+GeoPoint.Radius { // если растония меньше чем обра радиуса значит окружности пересекается
				return true
			}
		}
	}

	// руды
	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {
			for _, GeoPoint := range obj.GeoData {
				distToObstacle := game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, reservoir.X, reservoir.Y)
				if int(distToObstacle) < 100 { // если растония меньше чем обра радиуса значит окружности пересекается
					return true
				}
			}
		}
	}

	// все юниты
	units := globalGame.Clients.GetAllShortUnits(mp.Id)
	for _, unit := range units {
		rect := GetBodyRect(unit.Body, float64(unit.X), float64(unit.Y), unit.Rotate, false, false)
		for _, GeoPoint := range obj.GeoData {
			if rect.detectCollisionRectToCircle(&point{x: float64(GeoPoint.X), y: float64(GeoPoint.Y)}, GeoPoint.Radius) {
				return true
			}
		}
	}

	// ящики
	boxs := boxes.Boxes.GetAllBoxByMapID(mp.Id)

	for _, mapBox := range boxs {
		rect := GetCenterRect(float64(mapBox.X), float64(mapBox.Y), float64(mapBox.Height), float64(mapBox.Width))
		rect.Rotate(mapBox.Rotate)
		for _, GeoPoint := range obj.GeoData {
			if rect.detectCollisionRectToCircle(&point{x: float64(GeoPoint.X), y: float64(GeoPoint.Y)}, GeoPoint.Radius) {
				return true
			}
		}
	}

	// и 100px от дорог, баз, точек выхода из телепорта и самих телепортов
	if structCheck {

		minDist := 250.0

		for _, xLine := range mp.StaticObjects {
			for _, coordinateMap := range xLine {
				if strings.Contains(coordinateMap.Texture, "road") {

					for _, GeoPoint := range obj.GeoData {
						if game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, coordinateMap.X, coordinateMap.Y) < minDist {
							return true
						}
					}
				}
			}
		}

		for _, handler := range mp.HandlersCoordinates {
			for _, GeoPoint := range obj.GeoData {
				if game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, handler.X, handler.Y) < minDist {
					return true
				}
			}
		}

		for _, base := range bases.Bases.GetBasesByMap(mp.Id) {
			for _, GeoPoint := range obj.GeoData {
				if game_math.GetBetweenDist(GeoPoint.X, GeoPoint.X, base.X, base.Y) < minDist*2 {
					return true
				}
			}
		}

		entryPoints := maps.Maps.GetEntryPointsByMapID(mp.Id)
		for _, exit := range entryPoints {
			for _, GeoPoint := range obj.GeoData {
				if game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, exit.X, exit.Y) < minDist {
					return true
				}
			}
		}
	}

	return false
}
