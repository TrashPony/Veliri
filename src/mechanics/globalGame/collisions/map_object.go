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

	units := globalGame.Clients.GetAllShortUnits(mp.Id)
	boxs := boxes.Boxes.GetAllBoxByMapID(mp.Id)

	for _, GeoPoint := range obj.GeoData {

		// статичные обьекты на карте
		if CircleStaticMap(GeoPoint.X, GeoPoint.Y, GeoPoint.Radius, mp) {
			return true
		}

		// динамические обьекты
		for _, x := range mp.GetCopyMapDynamicObjects() {
			for _, sObj := range x {
				for _, sGeoPoint := range sObj.GeoData {

					if sObj.ID == obj.ID {
						continue
					}

					distToObstacle := game_math.GetBetweenDist(GeoPoint.X, GeoPoint.Y, sGeoPoint.X, sGeoPoint.Y)
					if int(distToObstacle) < sGeoPoint.Radius+GeoPoint.Radius {
						return true
					}
				}
			}
		}

		// глобальная гео дата
		if CircleGlobalGeoDataMap(GeoPoint.X, GeoPoint.Y, GeoPoint.Radius, mp) {
			return true
		}

		// руды
		if CircleReservoirMap(GeoPoint.X, GeoPoint.Y, 100-reservoirRadius, mp) {
			return true
		}

		// все юниты
		if CircleUnits(GeoPoint.X, GeoPoint.Y, GeoPoint.Radius, units) {
			return true
		}

		// все ящики
		if CircleBoxes(GeoPoint.X, GeoPoint.Y, GeoPoint.Radius, boxs) {
			return true
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
