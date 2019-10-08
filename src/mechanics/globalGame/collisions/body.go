package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
)

func BodyCheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body, full, min bool) (bool, bool) {

	if x < 0 || y < 0 || x > mp.XSize || y > mp.YSize {
		return false, true
	}

	if body == nil {
		return true, true
	}

	rect := getBodyRect(body, float64(x), float64(y), rotate, full, min)

	noCach := func() (bool, bool) {
		possibleMove, front := searchStaticMapCollisionByRect(x, y, mp, min, rect, body.ID, body.Height*2)
		if !possibleMove {
			return possibleMove, front
		} else {
			return checkMapReservoir(mp, rect)
		}
	}

	if min {
		find, value := getCacheCoordinate(mp.Id, body.ID, x, y)
		if find {
			if value {
				return checkMapReservoir(mp, rect)
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

func BodyCheckCollisionBoxes(moveUnit *unit.Unit, body *detail.Body, path *unit.PathUnit) (*boxInMap.Box, int, int, int) {
	rectBody := getBodyRect(body, float64(path.X), float64(path.Y), path.Rotate, false, false)
	mapBox := checkCollisionsBoxes(moveUnit.MapID, rectBody, false)

	if mapBox != nil {

		rectBody = getBodyRect(body, float64(moveUnit.X), float64(moveUnit.Y), moveUnit.Rotate, false, false)
		rectBox := getCenterRect(float64(mapBox.X), float64(mapBox.Y), float64(mapBox.Height), float64(mapBox.Width))
		rectBox.rotate(mapBox.Rotate)

		x, y, free, percent := detailCheckCollisionPolygons(rectBody, rectBox, path)
		if !free {
			return mapBox, x, y, percent
		}
	}

	return nil, 0, 0, 0
}

func CheckCollisionsPlayers(moveUnit *unit.Unit, x, y, rotate int, units map[int]*unit.ShortUnitInfo,
	min, max, hostileMax, onlyStanding, nextPoint bool, excludeIds []int) (bool, *unit.ShortUnitInfo) {

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

		if excludeIds != nil && findExcludeUnit(otherUnit, excludeIds) {
			continue
		}

		if otherUnit != nil && (moveUnit.ID != otherUnit.ID) { // todo && !user.GetSquad().Evacuation

			mUserRect := getBodyRect(moveUnit.Body, float64(x), float64(y), rotate, max, min)
			userRect := getBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate, hostileMax, false)

			if mUserRect.detectCollisionRectToRect(userRect) {
				return false, otherUnit
			}

			if nextPoint && otherUnit.ActualPathCell != nil {
				mUserRect = getBodyRect(moveUnit.Body, float64(x), float64(y), rotate, max, min)
				userRect = getBodyRect(otherUnit.Body, float64(otherUnit.ActualPathCell.X),
					float64(otherUnit.ActualPathCell.Y), otherUnit.ActualPathCell.Rotate, hostileMax, false)

				if mUserRect.detectCollisionRectToRect(userRect) {
					return false, otherUnit
				}
			}
		}
	}

	return true, nil
}

func findExcludeUnit(unit *unit.ShortUnitInfo, ids []int) bool {
	for _, id := range ids {
		if id == unit.ID {
			return true
		}
	}
	return false
}
