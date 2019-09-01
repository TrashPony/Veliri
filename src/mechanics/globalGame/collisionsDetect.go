package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"math"
)

// TODO тут все крайне не правильно но работает
// TODO требует рефакторинга, оч много повторяющегося кода

func checkCollision(rotate, fromAngle, toAngle, radius, x, y, obstacleX, obstacleY, obstacleRadius int) bool {
	for i := rotate - fromAngle; i < rotate+toAngle; i++ {

		rad := float64(i) * math.Pi / 180
		bX := int(float64(radius)*math.Cos(rad)) + x
		bY := int(float64(radius)*math.Sin(rad)) + y

		distToObstacle := int(GetBetweenDist(bX, bY, obstacleX, obstacleY))
		if distToObstacle < obstacleRadius {
			return true
		}
	}
	return false
}

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body, fast bool) (bool, int, int, bool) {
	// TODO так как карта статична, нет необходимости считать это каждый раз заного, можно запомнить и использовать пресеты
	startCoordinate := GetQRfromXY(x, y, mp)

	if body == nil {
		return true, startCoordinate.Q, startCoordinate.R, true
	}

	if fast {
		for _, obstacle := range mp.GeoData {
			dist := int(GetBetweenDist(x, y, obstacle.X, obstacle.Y))
			if dist < body.FrontRadius+obstacle.Radius {
				return false, startCoordinate.Q, startCoordinate.R, true
			}
		}
	} else {
		for _, obstacle := range mp.GeoData {
			dist := int(GetBetweenDist(x, y, obstacle.X, obstacle.Y))

			if dist < body.FrontRadius+obstacle.Radius {
				// проверяем фронт машины
				if checkCollision(rotate, body.LeftFrontAngle, body.RightFrontAngle, body.FrontRadius, x, y, obstacle.X, obstacle.Y, obstacle.Radius) {
					return false, startCoordinate.Q, startCoordinate.R, true
				}
			}

			if dist < body.BackRadius+obstacle.Radius {
				// проверяем жопку
				if checkCollision(rotate+180, body.LeftBackAngle, body.RightBackAngle, body.BackRadius, x, y, obstacle.X, obstacle.Y, obstacle.Radius) {
					return false, startCoordinate.Q, startCoordinate.R, false
				}
			}

			if dist < body.SideRadius+obstacle.Radius {
				// проверяем бока машины
				return false, startCoordinate.Q, startCoordinate.R, true
			}
		}
	}

	return CheckMapResource(x, y, rotate, mp, body, startCoordinate, fast)
}

func CheckMapResource(x, y, rotate int, mp *_map.Map, body *detail.Body, startCoordinate *coordinate.Coordinate, fast bool) (bool, int, int, bool) {
	const reservoirRadius = 15

	if fast {
		for _, qLine := range mp.Reservoir {
			for _, reservoir := range qLine {

				if reservoir.Move() {
					continue
				}

				reservoirX, reservoirY := GetXYCenterHex(reservoir.Q, reservoir.R)
				dist := int(GetBetweenDist(x, y, reservoirX, reservoirY))
				if dist < body.FrontRadius+reservoirRadius {
					// проверяем бока машины
					return false, startCoordinate.Q, startCoordinate.R, true
				}
			}
		}
	} else {
		for _, qLine := range mp.Reservoir {
			for _, reservoir := range qLine {

				if reservoir.Move() {
					continue
				}

				reservoirX, reservoirY := GetXYCenterHex(reservoir.Q, reservoir.R)

				dist := int(GetBetweenDist(x, y, reservoirX, reservoirY))

				if dist < body.FrontRadius+reservoirRadius {
					// проверяем фронт машины
					if checkCollision(rotate, body.LeftFrontAngle, body.RightFrontAngle, body.FrontRadius, x, y, reservoirX, reservoirY, reservoirRadius) {
						return false, startCoordinate.Q, startCoordinate.R, true
					}
				}

				if dist < body.BackRadius+reservoirRadius {
					// проверяем жопку
					if checkCollision(rotate+180, body.LeftBackAngle, body.RightBackAngle, body.BackRadius, x, y, reservoirX, reservoirY, reservoirRadius) {
						return false, startCoordinate.Q, startCoordinate.R, false
					}
				}

				if dist < body.SideRadius+reservoirRadius {
					// проверяем бока машины
					return false, startCoordinate.Q, startCoordinate.R, true
				}
			}
		}
	}

	return true, startCoordinate.Q, startCoordinate.R, true
}

func CheckCollisionsPlayers(moveUnit *unit.Unit, x, y, rotate int, units map[int]*unit.ShortUnitInfo) (bool, *unit.ShortUnitInfo) {

	bodyMove := moveUnit.Body
	mX, mY := float64(moveUnit.X), float64(moveUnit.Y)

	for _, otherUnit := range units {

		if otherUnit == nil {
			continue
		}

		if moveUnit.MapID != otherUnit.MapID {
			// по неведомой причине нельзя этот иф класть в общий, он не работает там
			continue
		}

		if otherUnit != nil && (moveUnit.ID > 0 && moveUnit.ID != otherUnit.ID) { // todo && !user.GetSquad().Evacuation

			bodyUser := otherUnit.Body

			dist := int(GetBetweenDist(x, y, otherUnit.X, otherUnit.Y))

			if dist < bodyMove.SideRadius+bodyUser.SideRadius {
				return false, otherUnit
			}

			if dist < bodyMove.FrontRadius+bodyUser.FrontRadius {
				// промеряем морду идущего с мордой стоящего
				for i := rotate - bodyMove.LeftFrontAngle; i < rotate+bodyMove.RightFrontAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.FrontRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.FrontRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, otherUnit.X, otherUnit.Y)

					if distToObstacle < float64(bodyUser.FrontRadius) {
						userRotate := otherUnit.Rotate

						abRad := math.Atan2(float64(bY-otherUnit.X), float64(bX-otherUnit.Y))
						ab := int(abRad * 180 / math.Pi)

						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftFrontAngle && ab < userRotate+bodyUser.RightFrontAngle {
							return false, otherUnit
						}
					}
				}
			}

			if dist < bodyMove.FrontRadius+bodyUser.FrontRadius {
				// проверяем морду идущего и жопу стоящего
				for i := rotate - bodyMove.LeftFrontAngle; i < rotate+bodyMove.RightFrontAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.FrontRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.FrontRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, otherUnit.X, otherUnit.Y)

					if distToObstacle < float64(bodyUser.BackRadius) {

						userRotate := otherUnit.Rotate
						if userRotate >= 180 {
							userRotate -= 180
						} else {
							if userRotate < 180 {
								userRotate += 180
							}
						}

						abRad := math.Atan2(float64(bY-otherUnit.Y), float64(bX-otherUnit.X))
						ab := int(abRad * 180 / math.Pi)
						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftBackAngle && ab < userRotate+bodyUser.RightBackAngle {
							return false, otherUnit
						}
					}
				}
			}

			if dist < bodyMove.FrontRadius+bodyUser.SideRadius {
				// проверяем морду идущего и бока стоящего
				if checkCollision(rotate, bodyMove.LeftFrontAngle, bodyMove.RightFrontAngle, bodyMove.FrontRadius, x, y, otherUnit.X, otherUnit.Y, bodyUser.SideRadius) {
					return false, otherUnit
				}
			}

			if dist < bodyMove.BackRadius+bodyUser.BackRadius {
				// промеряем жопу идущего с жопой стоящего
				for i := (rotate + 180) - bodyMove.LeftBackAngle; i < (rotate+180)+bodyMove.RightBackAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.BackRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.BackRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, otherUnit.X, otherUnit.Y)

					if distToObstacle < float64(bodyUser.BackRadius) {

						userRotate := otherUnit.Rotate
						if userRotate >= 180 {
							userRotate -= 180
						} else {
							if userRotate < 180 {
								userRotate += 180
							}
						}

						abRad := math.Atan2(float64(bY-otherUnit.Y), float64(bX-otherUnit.X))
						ab := int(abRad * 180 / math.Pi)
						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftBackAngle && ab < userRotate+bodyUser.RightBackAngle {
							return false, otherUnit
						}
					}
				}
			}

			if dist < bodyMove.BackRadius+bodyUser.FrontRadius {
				// промеряем жопу идущего с мордой стоящего
				for i := (rotate + 180) - bodyMove.LeftBackAngle; i < (rotate+180)+bodyMove.RightBackAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.BackRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.BackRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, otherUnit.X, otherUnit.Y)
					if distToObstacle < float64(bodyUser.FrontRadius) {
						userRotate := otherUnit.Rotate

						abRad := math.Atan2(float64(bY-otherUnit.Y), float64(bX-otherUnit.X))
						ab := int(abRad * 180 / math.Pi)

						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftFrontAngle && ab < userRotate+bodyUser.RightFrontAngle {
							return false, otherUnit
						}
					}
				}
			}

			/*
				    squad.rectDebag.moveTo(-50, -25);
					squad.rectDebag.lineTo(-50, +25);
					squad.rectDebag.lineTo(-25, +25);
					squad.rectDebag.lineTo(+50, +25);
					squad.rectDebag.lineTo(+50, +25);
					squad.rectDebag.lineTo(+50, -25);
					squad.rectDebag.lineTo(+25, -25);
					squad.rectDebag.lineTo(-50, -25);
			*/

			heightUserMove, widthUserMove := float64(moveUnit.Body.Height), float64(moveUnit.Body.Width)
			heightUser, widthUser := float64(otherUnit.Body.Height), float64(otherUnit.Body.Width)

			uX, uY := float64(otherUnit.X), float64(otherUnit.Y)
			mUserRect := rect{
				sides: []sideRec{
					{x1: mX - widthUserMove, y1: mY - heightUserMove, x2: mX - widthUserMove, y2: mY + heightUserMove},
					{x1: mX - heightUserMove, y1: mY + heightUserMove, x2: mX + widthUserMove, y2: mY + heightUserMove},
					{x1: mX + widthUserMove, y1: mY + heightUserMove, x2: mX + widthUserMove, y2: mY - heightUserMove},
					{x1: mX + heightUserMove, y1: mY - heightUserMove, x2: mX - widthUserMove, y2: mY - heightUserMove},
				},
				centerX: float64(moveUnit.X),
				centerY: float64(moveUnit.Y),
			}

			userRect := rect{
				sides: []sideRec{
					{x1: uX - widthUser, y1: uY - heightUser, x2: uX - widthUser, y2: uY + heightUser},
					{x1: uX - heightUser, y1: uY + heightUser, x2: uX + widthUser, y2: uY + heightUser},
					{x1: uX + widthUser, y1: uY + heightUser, x2: uX + widthUser, y2: uY - heightUser},
					{x1: uX + heightUser, y1: uY - heightUser, x2: uX - widthUser, y2: uY - heightUser},
				},
				centerX: float64(otherUnit.X),
				centerY: float64(otherUnit.Y),
			}

			if mUserRect.detect(&userRect, float64(moveUnit.Rotate), float64(otherUnit.Rotate)) {
				return false, otherUnit
			}
		}
	}
	return true, nil
}

type rect struct {
	sides            []sideRec
	centerX, centerY float64
}

type sideRec struct {
	x1, y1 float64
	x2, y2 float64
}

func (r *rect) detect(r2 *rect, alpha11, alpha22 float64) bool {
	for _, side1 := range r.sides {
		for _, side2 := range r2.sides {
			// поворачиваем квадрат по формуле
			//X = (x — x0) * cos(alpha) — (y — y0) * sin(alpha) + x0;
			//Y = (x — x0) * sin(alpha) + (y — y0) * cos(alpha) + y0;

			alpha1 := float64(alpha11) * math.Pi / 180
			alpha2 := float64(alpha22) * math.Pi / 180

			x1a := (side1.x1-r.centerX)*math.Cos(alpha1) - (side1.y1-r.centerY)*math.Sin(alpha1) + r.centerX
			x2a := (side1.x2-r.centerX)*math.Cos(alpha1) - (side1.y2-r.centerY)*math.Sin(alpha1) + r.centerX
			y1a := (side1.x1-r.centerX)*math.Sin(alpha1) + (side1.y1-r.centerY)*math.Cos(alpha1) + r.centerY
			y2a := (side1.x2-r.centerX)*math.Sin(alpha1) + (side1.y2-r.centerY)*math.Cos(alpha1) + r.centerY

			x1b := (side2.x1-r2.centerX)*math.Cos(alpha2) - (side2.y1-r2.centerY)*math.Sin(alpha2) + r2.centerX
			x2b := (side2.x2-r2.centerX)*math.Cos(alpha2) - (side2.y2-r2.centerY)*math.Sin(alpha2) + r2.centerX
			y1b := (side2.x1-r2.centerX)*math.Sin(alpha2) + (side2.y1-r2.centerY)*math.Cos(alpha2) + r2.centerY
			y2b := (side2.x2-r2.centerX)*math.Sin(alpha2) + (side2.y2-r2.centerY)*math.Cos(alpha2) + r2.centerY

			if Intersection(x1a, y1a, x2a, y2a, x1b, y1b, x2b, y2b) {
				return true
			}
		}
	}

	return false
}

func Intersection(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
	v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
	v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
	v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
	v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)

	return (v1*v2 < 0) && (v3*v4 < 0)
}

func CheckCollisionsBoxes(x, y, rotate, mapID int, body *detail.Body) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	const boxRadius = 5

	for _, mapBox := range boxs {

		// поздемные ящики не имеют колизий
		if mapBox.Underground {
			continue
		}

		xBox, yBox := GetXYCenterHex(mapBox.Q, mapBox.R)
		dist := int(GetBetweenDist(x, y, xBox, yBox))

		if dist < body.FrontRadius+boxRadius {
			// проверяем фронт машины
			if checkCollision(rotate, body.LeftFrontAngle, body.RightFrontAngle, body.FrontRadius, x, y, xBox, yBox, boxRadius) {
				return mapBox
			}
		}

		if dist < body.BackRadius+boxRadius {
			// проверяем жопку
			if checkCollision(rotate+180, body.LeftBackAngle, body.RightBackAngle, body.BackRadius, x, y, xBox, yBox, boxRadius) {
				return mapBox
			}
		}

		if dist < body.SideRadius+boxRadius {
			// проверяем бока машины
			if checkCollision(0, 0, 360, body.SideRadius, x, y, xBox, yBox, boxRadius) {
				return mapBox
			}
		}
	}
	return nil
}

func checkLevelViewCoordinate(one, past *coordinate.Coordinate) bool {
	if one.Level > past.Level {
		diffLevel := one.Level - past.Level
		if diffLevel < 2 {
			return false
		} else {
			return true
		}
	} else {
		diffLevel := past.Level - one.Level
		if diffLevel < 2 {
			return false
		} else {
			return true
		}
	}
}
