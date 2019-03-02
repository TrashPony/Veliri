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

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body) (bool, int, int, bool) {

	startCoordinate := GetQRfromXY(x, y, mp)

	if body == nil {
		return true, startCoordinate.Q, startCoordinate.R, true
	}

	for _, obstacle := range mp.GeoData {
		dist := int(GetBetweenDist(x, y, obstacle.X, obstacle.Y))

		if dist < body.FrontRadius+obstacle.Radius {
			// проверяем фронт машины
			for i := rotate - body.LeftFrontAngle; i < rotate+body.RightFrontAngle; i++ {

				rad := float64(i) * math.Pi / 180
				bX := int(float64(body.FrontRadius)*math.Cos(rad)) + x
				bY := int(float64(body.FrontRadius)*math.Sin(rad)) + y

				distToObstacle := int(GetBetweenDist(bX, bY, obstacle.X, obstacle.Y))
				if distToObstacle < obstacle.Radius {
					return false, startCoordinate.Q, startCoordinate.R, true
				}
			}
		}

		if dist < body.BackRadius+obstacle.Radius {
			// проверяем жопку
			for i := (rotate + 180) - body.LeftBackAngle; i < (rotate+180)+body.RightBackAngle; i++ {

				rad := float64(i) * math.Pi / 180
				bX := int(float64(body.BackRadius)*math.Cos(rad)) + x
				bY := int(float64(body.BackRadius)*math.Sin(rad)) + y

				distToObstacle := int(GetBetweenDist(bX, bY, obstacle.X, obstacle.Y))
				if distToObstacle < obstacle.Radius {
					println("back")
					return false, startCoordinate.Q, startCoordinate.R, false
				}
			}
		}

		if dist < body.SideRadius+obstacle.Radius {
			// проверяем бока машины
			for i := 0; i < 360; i++ {

				rad := float64(i) * math.Pi / 180
				bX := int(float64(body.SideRadius)*math.Cos(rad)) + x
				bY := int(float64(body.SideRadius)*math.Sin(rad)) + y

				distToObstacle := int(GetBetweenDist(bX, bY, obstacle.X, obstacle.Y))
				if distToObstacle < obstacle.Radius {
					println("side")
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
		if user != nil && user.GetSquad().MapID == mapID && moveUser.GetID() != user.GetID() && !user.GetSquad().Evacuation {

			bodyUser := user.GetSquad().MatherShip.Body

			dist := int(GetBetweenDist(x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

			if dist < bodyMove.SideRadius+bodyUser.SideRadius {
				// проверяем боковые радиусы
				for i := 0; i < 360; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.FrontRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.FrontRadius)*math.Sin(rad)) + y

					distToObstacle := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))
					if distToObstacle < bodyUser.SideRadius {
						return false
					}
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

			if dist < bodyMove.FrontRadius+bodyUser.FrontRadius {
				// проверяем морду идущего и бока стоящего
				for i := rotate - bodyMove.LeftFrontAngle; i < rotate+bodyMove.RightFrontAngle; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyMove.FrontRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyMove.FrontRadius)*math.Sin(rad)) + y
					distToObstacle := GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY)
					if distToObstacle < float64(bodyUser.SideRadius) {
						return false
					}
				}
			}
			// TODO обработать жопу машины, обьеденить код в метод
		}
	}
	return true
}

func CheckCollisionsBoxes(x, y, rotate, mapID int) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	for _, mapBox := range boxs {
		xBox, yBox := GetXYCenterHex(mapBox.Q, mapBox.R)
		dist := int(GetBetweenDist(x, y, xBox, yBox))

		if dist < 43*5 && !mapBox.Underground {

			for i := rotate - 20; i < rotate+20; i++ { // смотрим колизии на самой морде
				rad := float64(i) * math.Pi / 180
				bX := int(float64(43+15)*math.Cos(rad)) + x
				bY := int(float64(43+15)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 5 {
					return mapBox
				}
			}

			for i := rotate - 40; i < rotate+40; i++ { // смотри колизии ближе к бокам
				rad := float64(i) * math.Pi / 180
				bX := int(float64(43)*math.Cos(rad)) + x
				bY := int(float64(43)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 5 {
					return mapBox
				}
			}

			for i := 0; i < 360; i++ { // смотри колизии везде по радиусу боков
				rad := float64(i) * math.Pi / 180
				bX := int(float64(43-10)*math.Cos(rad)) + x
				bY := int(float64(43-10)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 5 {
					return mapBox
				}
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
