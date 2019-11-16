package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

const reservoirRadius = 15 // TODO

func searchStaticMapCollisionByRect(x, y int, mp *_map.Map, min bool, rect *Polygon, bodyID, minDist int) (bool, bool) {
	xZone, yZone := x/game_math.DiscreteSize, y/game_math.DiscreteSize

	if mp.GeoZones[xZone] == nil || mp.GeoZones[xZone][yZone] == nil {
		return false, true
	}

	obstacles := mp.GeoZones[xZone][yZone].Obstacle

	fastFindObstacle := func() (bool, bool) {

		possibleMove := false
		Front := true
		stopFind := false

		go func() {

			defer func() { stopFind = true }()

			for i := 0; i < len(obstacles); i++ {

				if stopFind {
					return
				}

				obstacle := obstacles[i]

				distToObstacle := game_math.GetBetweenDist(x, y, obstacle.X, obstacle.Y)
				if int(distToObstacle) < obstacle.Radius+minDist {
					if rect.detectCollisionRectToCircle(&point{x: float64(obstacle.X), y: float64(obstacle.Y)}, obstacle.Radius) {
						possibleMove = false
						Front = true

						if min {
							addCacheCoordinate(mp.Id, bodyID, x, y, false)
						}

						return
					}
				}
			}

			possibleMove = true
			Front = true
			return
		}()

		go func() {

			defer func() { stopFind = true }()

			for i := len(obstacles) - 1; i >= 0; i-- {

				if stopFind {
					return
				}

				obstacle := obstacles[i]

				distToObstacle := game_math.GetBetweenDist(x, y, obstacle.X, obstacle.Y)
				if int(distToObstacle) < obstacle.Radius+minDist {
					if rect.detectCollisionRectToCircle(&point{x: float64(obstacle.X), y: float64(obstacle.Y)}, obstacle.Radius) {
						possibleMove = false
						Front = true

						if min {
							addCacheCoordinate(mp.Id, bodyID, x, y, false)
						}

						return
					}
				}
			}

			possibleMove = true
			Front = true

			if min {
				addCacheCoordinate(mp.Id, bodyID, x, y, true)
			}

			return
		}()

		for !stopFind {
			time.Sleep(time.Millisecond)
		}

		return possibleMove, Front
	}

	return fastFindObstacle()
}

func checkMapReservoir(mp *_map.Map, rect *Polygon) (bool, bool) {

	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {

			if reservoir.Move() {
				continue
			}

			if rect.detectCollisionRectToCircle(&point{x: float64(reservoir.X), y: float64(reservoir.Y)}, reservoirRadius) {
				return false, true
			}
		}
	}

	return true, true
}

func checkObjectsMap(mp *_map.Map, rect *Polygon) (bool, bool) {
	// TODO ужасная оптимизация, точнее ее отсутсвие. Сильно сказывается на поиске пути
	// 	но сейчас у меня нет сил делать оптимизацию С:
	for _, q := range mp.StaticObjects {
		for _, object := range q {
			for _, geoPoint := range object.GeoData {
				if rect.detectCollisionRectToCircle(&point{x: float64(geoPoint.X), y: float64(geoPoint.Y)}, geoPoint.Radius) {
					return true, true
				}
			}
		}
	}

	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, object := range x {
			for _, geoPoint := range object.GeoData {
				if rect.detectCollisionRectToCircle(&point{x: float64(geoPoint.X), y: float64(geoPoint.Y)}, geoPoint.Radius) {
					return true, true
				}
			}
		}
	}

	return false, false
}

func checkCollisionsBoxes(mapID int, rect *Polygon, undergroundBox bool) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	for _, mapBox := range boxs {

		if mapBox.Height == 0 || mapBox.Width == 0 {
			// todo если ящик состоит из 1 точки то колизия происходит всегда
			mapBox.Height = 5
			mapBox.Width = 5
		}

		rectBox := GetCenterRect(float64(mapBox.X), float64(mapBox.Y), float64(mapBox.Height), float64(mapBox.Width))
		rectBox.Rotate(mapBox.Rotate)

		// поздемные ящики не имеют колизий
		if mapBox.Underground && !undergroundBox {
			continue
		}

		if rect.detectCollisionRectToRect(rectBox) {
			return mapBox
		}
	}
	return nil
}

func checkCollisionsUnits(rect *Polygon, units map[int]*unit.ShortUnitInfo, mapID int, excludeUnitID int) bool {
	for _, otherUnit := range units {
		if otherUnit == nil {
			continue
		}

		if mapID != otherUnit.MapID {
			continue
		}

		if excludeUnitID == otherUnit.ID {
			continue
		}

		if otherUnit != nil {
			userRect := GetBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate, false, false)
			if rect.detectCollisionRectToRect(userRect) {
				return false
			}
		}
	}

	return true
}
