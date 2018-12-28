package globalGame

import (
	"../factories/boxes"
	"../gameObjects/box"
	"../gameObjects/map"
	"../player"
	"github.com/gorilla/websocket"
	"math"
)

// TODO тут все крайне не правильно но работает

const bodyRadius = 63 // размеры подобраны методом тыка)
const coordinateRadius = HexagonHeight / 2

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map) (bool, int, int) {

	minDist := 999

	var q, r int

	for _, qLine := range mp.OneLayerMap {
		for _, mapCoordinate := range qLine {
			xc, yc := GetXYCenterHex(mapCoordinate.Q, mapCoordinate.R)

			//находим растояние координаты от места остановки
			dist := int(GetBetweenDist(x, y, xc, yc))

			// если координата находиться в теоритическом радиусе радиусе то проверяем на колизии
			if dist < coordinateRadius*3 {

				if minDist > dist {
					minDist = dist
					q = mapCoordinate.Q
					r = mapCoordinate.R
				}

				// корпус состоиз из 3х оружностей, 1 спериди, 2 жопке, 3 бока

				for i := rotate - 15; i < rotate+15; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(90)*math.Cos(rad)) + x
					bY := int(float64(90)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius {
						if !mapCoordinate.Move {
							return false, q, r
						}
					}
				}

				//for i := rotate - 190; i < rotate-170; i++ {
				//	rad := float64(i) * math.Pi / 180
				//	bX := int(float64(90)*math.Cos(rad)) + x
				//	bY := int(float64(90)*math.Sin(rad)) + y
				//
				//	dist := int(GetBetweenDist(bX, bY, xc, yc))
				//	if dist < coordinateRadius {
				//		if !mapCoordinate.Move {
				//			return false, q, r
				//		}
				//	}
				//}

				for i := rotate - 360; i < rotate+360; i++ {
					rad := float64(i) * math.Pi / 180
					bX := int(float64(40)*math.Cos(rad)) + x
					bY := int(float64(40)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, xc, yc))
					if dist < coordinateRadius {
						if !mapCoordinate.Move {
							return false, q, r
						}
					}
				}
			}
		}
	}
	return true, q, r
}

func CheckCollisionsPlayers(moveUser *player.Player, x, y, rotate, mapID int, users map[*websocket.Conn]*player.Player) bool {
	for _, user := range users {
		if user != nil && user.GetSquad().MapID == mapID && moveUser.GetID() != user.GetID() && !user.GetSquad().Evacuation {
			dist := int(GetBetweenDist(x, y, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

			if dist < bodyRadius*5 {
				for i := rotate - 5; i < rotate+5; i++ { // смотрим только предметы по курсу )

					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius)*math.Cos(rad)) + x // точки окружности корпуса
					bY := int(float64(bodyRadius)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

					if dist < bodyRadius*2 {
						return false
					}
				}

				for i := rotate - 25; i < rotate+25; i++ { // смотрим бока по меньшему радиусу

					rad := float64(i) * math.Pi / 180
					bX := int(float64(bodyRadius/2)*math.Cos(rad)) + x
					bY := int(float64(bodyRadius/2)*math.Sin(rad)) + y

					dist := int(GetBetweenDist(bX, bY, user.GetSquad().GlobalX, user.GetSquad().GlobalY))

					if dist < bodyRadius*2 {
						return false
					}
				}
			}
		}
	}
	return true
}

func CheckCollisionsBoxes(x, y, rotate, mapID int) *box.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	for _, mapBox := range boxs {
		xBox, yBox := GetXYCenterHex(mapBox.Q, mapBox.R)
		dist := int(GetBetweenDist(x, y, xBox, yBox))

		if dist < bodyRadius*5 && !mapBox.Underground {
			for i := rotate - 40; i < rotate+45; i++ {

				rad := float64(i) * math.Pi / 180
				bX := int(float64(100)*math.Cos(rad)) + x // точки окружности корпуса
				bY := int(float64(100)*math.Sin(rad)) + y

				dist := int(GetBetweenDist(bX, bY, xBox, yBox))
				if dist < 10 {
					return mapBox
				}
			}
		}
	}
	return nil
}
