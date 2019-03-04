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

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body) (bool, int, int, bool) {

	startCoordinate := GetQRfromXY(x, y, mp)

	if body == nil {
		return true, startCoordinate.Q, startCoordinate.R, true
	}

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
			if checkCollision(0, 0, 360, body.SideRadius, x, y, obstacle.X, obstacle.Y, obstacle.Radius) {
				return false, startCoordinate.Q, startCoordinate.R, false
			}
		}
	}

	return CheckMapResource(x, y, rotate, mp, body, startCoordinate)
}

func CheckMapResource(x, y, rotate int, mp *_map.Map, body *detail.Body, startCoordinate *coordinate.Coordinate) (bool, int, int, bool) {
	const reservoirRadius = 15
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
				if checkCollision(0, 0, 360, body.SideRadius, x, y, reservoirX, reservoirY, reservoirRadius) {
					return false, startCoordinate.Q, startCoordinate.R, false
				}
			}
		}
	}

	return true, startCoordinate.Q, startCoordinate.R, true
}

func CheckCollisionsPlayers(moveUser *player.Player, x, y, rotate, mapID int, users map[*websocket.Conn]*player.Player) bool {
	bodyMove := moveUser.GetSquad().MatherShip.Body

	for _, user := range users {
		if user != nil && user.GetSquad().MapID == mapID &&
			(moveUser.GetID() > 0 && moveUser.GetID() != user.GetID() || moveUser.UUID != "" && moveUser.UUID != user.UUID) &&
			!user.GetSquad().Evacuation {

			bodyUser := user.GetSquad().MatherShip.Body

			dist := int(GetBetweenDist(x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

			if dist < bodyMove.SideRadius+bodyUser.SideRadius {
				// проверяем боковые радиусы
				if checkCollision(0, 0, 360, bodyMove.SideRadius, x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY, bodyUser.SideRadius) {
					return false
				}
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
							return false
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
							return false
						}
					}
				}
			}

			if dist < bodyMove.FrontRadius+bodyUser.SideRadius {
				// проверяем морду идущего и бока стоящего
				if checkCollision(rotate, bodyMove.LeftFrontAngle, bodyMove.RightFrontAngle, bodyMove.FrontRadius, x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY, bodyUser.SideRadius) {
					return false
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
							return false
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
							return false
						}
					}
				}
			}
		}
	}
	return true
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
