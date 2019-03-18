package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
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

func CheckCollisionsPlayers(moveUser *player.Player, x, y, rotate int, users map[*websocket.Conn]*player.Player) (bool, *player.Player) {

	if moveUser.GetSquad() == nil {
		return true, nil
	}

	bodyMove := moveUser.GetSquad().MatherShip.Body
	mX, mY := float64(moveUser.GetSquad().GlobalX), float64(moveUser.GetSquad().GlobalY)

	for _, user := range users {

		if user.GetSquad() == nil {
			continue
		}

		if user.GetSquad().MapID != moveUser.GetSquad().MapID {
			// по неведомой причине нельзя этот иф класть в общий, он не работает там
			// TODO тут я тоже не уверен надо дебажить, проблема что боты сталкиваются на разных картах имея одинаковые координаты
			continue
		}

		if user != nil && (moveUser.GetID() > 0 && moveUser.GetID() != user.GetID()) || (moveUser.UUID != "" && moveUser.UUID != user.UUID) && !user.GetSquad().Evacuation {

			if user.GetSquad().MapID != moveUser.GetSquad().MapID {
				println("колизия: ", user.GetID(), user.GetSquad().MapID, moveUser.GetID(), moveUser.GetSquad().MapID, user.GetSquad().MapID == moveUser.GetSquad().MapID)
			}

			bodyUser := user.GetSquad().MatherShip.Body

			dist := int(GetBetweenDist(x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

			if dist < bodyMove.SideRadius+bodyUser.SideRadius {
				return false, user
			}

			if dist < bodyMove.FrontRadius+bodyUser.FrontRadius {
				// промеряем морду идущего с мордой стоящего
				for i := rotate - bodyMove.LeftFrontAngle; i < rotate+bodyMove.RightFrontAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.FrontRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.FrontRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY)

					if distToObstacle < float64(bodyUser.FrontRadius) {
						userRotate := user.GetSquad().MatherShip.Rotate

						abRad := math.Atan2(float64(bY-user.GetSquad().GlobalY), float64(bX-user.GetSquad().GlobalX))
						ab := int(abRad * 180 / math.Pi)

						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftFrontAngle && ab < userRotate+bodyUser.RightFrontAngle {
							return false, user
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
					distToObstacle := GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY)

					if distToObstacle < float64(bodyUser.BackRadius) {

						userRotate := user.GetSquad().MatherShip.Rotate
						if userRotate >= 180 {
							userRotate -= 180
						} else {
							if userRotate < 180 {
								userRotate += 180
							}
						}

						abRad := math.Atan2(float64(bY-user.GetSquad().GlobalY), float64(bX-user.GetSquad().GlobalX))
						ab := int(abRad * 180 / math.Pi)
						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftBackAngle && ab < userRotate+bodyUser.RightBackAngle {
							return false, user
						}
					}
				}
			}

			if dist < bodyMove.FrontRadius+bodyUser.SideRadius {
				// проверяем морду идущего и бока стоящего
				if checkCollision(rotate, bodyMove.LeftFrontAngle, bodyMove.RightFrontAngle, bodyMove.FrontRadius, x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY, bodyUser.SideRadius) {
					return false, user
				}
			}

			if dist < bodyMove.BackRadius+bodyUser.BackRadius {
				// промеряем жопу идущего с жопой стоящего
				for i := (rotate + 180) - bodyMove.LeftBackAngle; i < (rotate+180)+bodyMove.RightBackAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.BackRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.BackRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY)

					if distToObstacle < float64(bodyUser.BackRadius) {

						userRotate := user.GetSquad().MatherShip.Rotate
						if userRotate >= 180 {
							userRotate -= 180
						} else {
							if userRotate < 180 {
								userRotate += 180
							}
						}

						abRad := math.Atan2(float64(bY-user.GetSquad().GlobalY), float64(bX-user.GetSquad().GlobalX))
						ab := int(abRad * 180 / math.Pi)
						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftBackAngle && ab < userRotate+bodyUser.RightBackAngle {
							return false, user
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
					distToObstacle := GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY)
					if distToObstacle < float64(bodyUser.FrontRadius) {
						userRotate := user.GetSquad().MatherShip.Rotate

						abRad := math.Atan2(float64(bY-user.GetSquad().GlobalY), float64(bX-user.GetSquad().GlobalX))
						ab := int(abRad * 180 / math.Pi)

						if ab < 0 && userRotate > 180 {
							ab += 360
						}
						if ab > userRotate-bodyUser.LeftFrontAngle && ab < userRotate+bodyUser.RightFrontAngle {
							return false, user
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

			height, width := float64(25), float64(40)

			uX, uY := float64(user.GetSquad().GlobalX), float64(user.GetSquad().GlobalY)
			mUserRect := rect{
				sides: []sideRec{
					{x1: mX - width, y1: mY - height, x2: mX - width, y2: mY + height},
					{x1: mX - height, y1: mY + height, x2: mX + width, y2: mY + height},
					{x1: mX + width, y1: mY + height, x2: mX + width, y2: mY - height},
					{x1: mX + height, y1: mY - height, x2: mX - width, y2: mY - height},
				},
				centerX: float64(moveUser.GetSquad().GlobalX),
				centerY: float64(moveUser.GetSquad().GlobalY),
			}

			userRect := rect{
				sides: []sideRec{
					{x1: uX - width, y1: uY - height, x2: uX - width, y2: uY + height},
					{x1: uX - height, y1: uY + height, x2: uX + width, y2: uY + height},
					{x1: uX + width, y1: uY + height, x2: uX + width, y2: uY - height},
					{x1: uX + height, y1: uY - height, x2: uX - width, y2: uY - height},
				},
				centerX: float64(user.GetSquad().GlobalX),
				centerY: float64(user.GetSquad().GlobalY),
			}

			if mUserRect.detect(&userRect, float64(moveUser.GetSquad().MatherShip.Rotate), float64(user.GetSquad().MatherShip.Rotate)) {
				return false, user
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
