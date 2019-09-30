package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body, full, min bool) (bool, bool) {

	if x < 0 || y < 0 || x > mp.XSize || y > mp.YSize {
		return false, true
	}

	if body == nil {
		return true, true
	}

	noCach := func() (bool, bool) {
		possibleMove, front := searchStaticMapCollisionByBody(x, y, rotate, mp, body, full, min)
		if !possibleMove {
			return possibleMove, front
		} else {
			return CheckMapReservoir(x, y, rotate, mp, body, full, min)
		}
	}

	if min {
		find, value := getCacheCoordinate(mp.Id, body.ID, x, y)
		if find {
			if value {
				return CheckMapReservoir(x, y, rotate, mp, body, full, min)
			} else {
				return false, false
			}
		} else {
			return noCach()
		}
	} else {
		return noCach()
	}
}

func searchStaticMapCollisionByBody(x, y, rotate int, mp *_map.Map, body *detail.Body, full, min bool) (bool, bool) {
	xZone, yZone := x/game_math.DiscreteSize, y/game_math.DiscreteSize

	if mp.GeoZones[xZone] == nil || mp.GeoZones[xZone][yZone] == nil {
		return false, true
	}

	obstacles := mp.GeoZones[xZone][yZone].Obstacle
	rect := getBodyRect(body, float64(x), float64(y), rotate, full, min)

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
				if int(distToObstacle) < obstacle.Radius+body.Height*2 {
					if rect.detectCollisionRectToCircle(&point{x: float64(obstacle.X), y: float64(obstacle.Y)}, obstacle.Radius) {
						possibleMove = false
						Front = true

						if min {
							addCacheCoordinate(mp.Id, body.ID, x, y, false)
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
				if int(distToObstacle) < obstacle.Radius+body.Height*2 {
					if rect.detectCollisionRectToCircle(&point{x: float64(obstacle.X), y: float64(obstacle.Y)}, obstacle.Radius) {
						possibleMove = false
						Front = true

						if min {
							addCacheCoordinate(mp.Id, body.ID, x, y, false)
						}

						return
					}
				}
			}

			possibleMove = true
			Front = true

			if min {
				addCacheCoordinate(mp.Id, body.ID, x, y, true)
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

func CheckMapReservoir(x, y, rotate int, mp *_map.Map, body *detail.Body, full, min bool) (bool, bool) {

	const reservoirRadius = 15

	rect := getBodyRect(body, float64(x), float64(y), rotate, full, min)

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

func CheckCollisionsBoxes(x, y, rotate, mapID int, body *detail.Body) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	const boxRadius = 5

	rect := getBodyRect(body, float64(x), float64(y), rotate, false, false)
	for _, mapBox := range boxs {

		// поздемные ящики не имеют колизий
		if mapBox.Underground {
			continue
		}

		if rect.detectCollisionRectToCircle(&point{x: float64(mapBox.X), y: float64(mapBox.Y)}, boxRadius) {
			return mapBox
		}
	}
	return nil
}

func CheckCollisionsPlayers(moveUnit *unit.Unit, x, y, rotate int, units map[int]*unit.ShortUnitInfo, min, max, onlyStanding bool) (bool, *unit.ShortUnitInfo) {

	for _, otherUnit := range units {

		if otherUnit == nil {
			continue
		}

		if moveUnit.MapID != otherUnit.MapID {
			continue
		}

		if onlyStanding && otherUnit.MoveChecker {
			continue
		}

		if otherUnit != nil && (moveUnit.ID != otherUnit.ID) { // todo && !user.GetSquad().Evacuation

			mUserRect := getBodyRect(moveUnit.Body, float64(x), float64(y), rotate, max, min)
			userRect := getBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate, false, false)

			if mUserRect.detectPointInRectangle(float64(otherUnit.X), float64(otherUnit.Y)) {
				// цент находится внутри прямоуголника, пересекается
				return false, otherUnit
			}

			if userRect.detectPointInRectangle(float64(x), float64(y)) {
				// цент находится внутри прямоуголника, пересекается
				return false, otherUnit
			}

			if mUserRect.centerX == userRect.centerX && mUserRect.centerY == userRect.centerY {
				// при одинаковом прямоугольнике и одинаковым центром, не будет пересечений и колизия будет не найдена
				// поэтому это тут
				return false, otherUnit
			}

			if mUserRect.detectCollisionRectToRect(userRect, float64(rotate), float64(otherUnit.Rotate)) {
				return false, otherUnit
			}
		}
	}

	return true, nil
}
