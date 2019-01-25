package globalGame

import (
	"../factories/boxes"
	"../gameObjects/boxInMap"
	"../gameObjects/coordinate"
	"../gameObjects/map"
	"../player"
	"github.com/gorilla/websocket"
	"math"
)

// TODO тут все крайне не правильно но работает

const bodyRadius = 31 // размеры подобраны методом тыка)
const coordinateRadius = HexagonHeight / 2

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map) (bool, int, int, bool) {

	startCoordinate := GetQRfromXY(x, y, mp)
	checkCoordinate := coordinate.GetCoordinatesRadius(startCoordinate, 2)

	for _, faceCoordinate := range checkCoordinate {

		mapCoordinate, find := mp.GetCoordinate(faceCoordinate.Q, faceCoordinate.R)
		if find {
			xc, yc := GetXYCenterHex(mapCoordinate.Q, mapCoordinate.R)

			//находим растояние координаты от места остановки
			dist := int(GetBetweenDist(x, y, xc, yc))

			// если координата находиться в теоритическом радиусе радиусе то проверяем на колизии
			if dist < coordinateRadius*3 {

				for i := rotate - 20; i < rotate+20; i++ { // смотрим колизии на самой морде
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius+15)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius+15)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius {
						if !mapCoordinate.Move || checkLevelViewCoordinate(startCoordinate, mapCoordinate) {
							return false, startCoordinate.Q, startCoordinate.R, true
						}
					}
				}

				for i := rotate - 40; i < rotate+40; i++ { // смотри колизии ближе к бокам
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius {
						if !mapCoordinate.Move || checkLevelViewCoordinate(startCoordinate, mapCoordinate) {
							return false, startCoordinate.Q, startCoordinate.R, true
						}
					}
				}

				//for i := 0; i < 360; i++ { // смотри колизии везде по радиусу боков
				//	rad := float64(i) * math.Pi / 180
				//	bX := int(float64(40)*math.Cos(rad)) + x
				//	bY := int(float64(40)*math.Sin(rad)) + y
				//
				//	dist := int(GetBetweenDist(bX, bY, xc, yc))
				//	if dist < coordinateRadius {
				//		if !mapCoordinate.Move {
				//			return false, q, r, false
				//		}
				//	}
				//}

				for i := rotate - 220; i < rotate-140; i++ { // смотри колизии ближе к бокам от зада
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius/2)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius/2)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius || checkLevelViewCoordinate(startCoordinate, mapCoordinate) {
						if !mapCoordinate.Move {
							return false, startCoordinate.Q, startCoordinate.R, false
						}
					}
				}

				for i := rotate + 200; i < rotate+160; i++ { // смотрим колизии на жопке
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius+20)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius+20)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius {
						if !mapCoordinate.Move || checkLevelViewCoordinate(startCoordinate, mapCoordinate) {
							return false, startCoordinate.Q, startCoordinate.R, false
						}
					}
				}

			}
		}
	}
	return true, startCoordinate.Q, startCoordinate.R, true
}

func CheckCollisionsPlayers(moveUser *player.Player, x, y, rotate, mapID int, users map[*websocket.Conn]*player.Player) bool {
	for _, user := range users {
		if user != nil && user.GetSquad().MapID == mapID && moveUser.GetID() != user.GetID() && !user.GetSquad().Evacuation {

			dist := int(GetBetweenDist(x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

			if dist < bodyRadius*5 {

				for i := rotate - 20; i < rotate+20; i++ { // смотрим колизии на самой морде
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius+15)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius+15)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))
					if dist < 5 {
						return false
					}
				}

				for i := rotate - 40; i < rotate+40; i++ { // смотри колизии ближе к бокам
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))
					if dist < 50 {
						return false
					}
				}

				for i := 0; i < 360; i++ { // смотри колизии везде по радиусу боков
					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius-10)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius-10)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))
					if dist < 50 {
						return false
					}
				}
			}
		}
	}
	return true
}

func CheckCollisionsBoxes(x, y, rotate, mapID int) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	for _, mapBox := range boxs {
		xBox, yBox := GetXYCenterHex(mapBox.Q, mapBox.R)
		dist := int(GetBetweenDist(x, y, xBox, yBox))

		if dist < bodyRadius*5 && !mapBox.Underground {

			for i := rotate - 20; i < rotate+20; i++ { // смотрим колизии на самой морде
				rad := float64(i) * math.Pi / 180
				bX := int(float64(bodyRadius+15)*math.Cos(rad)) + x
				bY := int(float64(bodyRadius+15)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 5 {
					return mapBox
				}
			}

			for i := rotate - 40; i < rotate+40; i++ { // смотри колизии ближе к бокам
				rad := float64(i) * math.Pi / 180
				bX := int(float64(bodyRadius)*math.Cos(rad)) + x
				bY := int(float64(bodyRadius)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 5 {
					return mapBox
				}
			}

			for i := 0; i < 360; i++ { // смотри колизии везде по радиусу боков
				rad := float64(i) * math.Pi / 180
				bX := int(float64(bodyRadius-10)*math.Cos(rad)) + x
				bY := int(float64(bodyRadius-10)*math.Sin(rad)) + y

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
