package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body, full bool) (bool, int, int, bool) {

	q, r := game_math.GetQRfromXY(x, y)
	startCoordinate, find := mp.OneLayerMap[q][r]

	if !find {
		return false, 0, 0, true
	}

	if body == nil {
		return true, 0, 0, true
	}

	rect := getBodyRect(body, float64(x), float64(y), rotate, full)
	for _, obstacle := range mp.GeoData {
		if rect.detectCollisionRectToCircle(&point{x: float64(obstacle.X), y: float64(obstacle.Y)}, obstacle.Radius) {
			return false, startCoordinate.Q, startCoordinate.R, true
		}
	}

	const reservoirRadius = 15
	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {

			if reservoir.Move() {
				continue
			}

			reservoirX, reservoirY := game_math.GetXYCenterHex(reservoir.Q, reservoir.R)
			if rect.detectCollisionRectToCircle(&point{x: float64(reservoirX), y: float64(reservoirY)}, reservoirRadius) {
				return false, startCoordinate.Q, startCoordinate.R, true
			}
		}
	}

	return true, startCoordinate.Q, startCoordinate.R, true
}

func CheckCollisionsBoxes(x, y, rotate, mapID int, body *detail.Body) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	const boxRadius = 5

	rect := getBodyRect(body, float64(x), float64(y), rotate, false)
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

			mUserRect := getBodyRect(moveUnit.Body, float64(x), float64(y), rotate, false)
			userRect := getBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate, false)

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
