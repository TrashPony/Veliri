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

	q, r := game_math.GetQRfromXY(x, y)
	_, find := mp.OneLayerMap[q][r]

	if !find {
		return false, true
	}

	if body == nil {
		return true, true
	}

	noCach := func() (bool, bool) {
		possibleMove, front := searchStaticMapCollision(x, y, rotate, mp, body, full, min)
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

func searchStaticMapCollision(x, y, rotate int, mp *_map.Map, body *detail.Body, full, min bool) (bool, bool) {
	xZone, yZone := x/100, y/100

	if mp.GeoZone[xZone] == nil || mp.GeoZone[xZone][yZone] == nil {
		return false, true
	}

	zone := mp.GeoZone[xZone][yZone]
	rect := getBodyRect(body, float64(x), float64(y), rotate, full, min)

	fastFindObstacle := func() (bool, bool) {

		possibleMove := false
		Front := true
		stopFind := false

		go func() {

			defer func() { stopFind = true }()

			for i := 0; i < len(zone); i++ {

				if stopFind {
					return
				}

				obstacle := zone[i]

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

			for i := len(zone) - 1; i >= 0; i-- {

				if stopFind {
					return
				}

				obstacle := zone[i]

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

			reservoirX, reservoirY := game_math.GetXYCenterHex(reservoir.Q, reservoir.R)
			if rect.detectCollisionRectToCircle(&point{x: float64(reservoirX), y: float64(reservoirY)}, reservoirRadius) {
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

		xBox, yBox := game_math.GetXYCenterHex(mapBox.Q, mapBox.R)
		if rect.detectCollisionRectToCircle(&point{x: float64(xBox), y: float64(yBox)}, boxRadius) {
			return mapBox
		}
	}
	return nil
}

func CheckCollisionsPlayers(moveUnit *unit.Unit, x, y, rotate int, units map[int]*unit.ShortUnitInfo) (bool, *unit.ShortUnitInfo) {

	for _, otherUnit := range units {

		if otherUnit == nil {
			continue
		}

		if moveUnit.MapID != otherUnit.MapID {
			continue
		}

		if otherUnit != nil && (moveUnit.ID != otherUnit.ID) { // todo && !user.GetSquad().Evacuation

			mUserRect := getBodyRect(moveUnit.Body, float64(x), float64(y), rotate, false, false)
			userRect := getBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate, false, false)

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
