package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
)

func GetUnitPos(units []int, mapID, toX, toY int) []*coordinate.Coordinate {

	toPos := make([]*coordinate.Coordinate, 0)

	mp, _ := maps.Maps.GetByID(mapID)
	center := GetQRfromXY(toX, toY, mp)

	getNearCoordinate := func() *coordinate.Coordinate {

		// проверяем что бы ее уже небыло в добавленных
		check := func(q, r int) bool {
			for _, coor := range toPos {
				if q == coor.Q && r == coor.R {
					return true
				}
			}
			return false
		}

		// максимальное количество юнитов у игрока 7 поэтому радиус 2 должно хватить, но это не точно
		radius := coordinate.GetCoordinatesRadius(center, 3)

		for _, coor1 := range radius {

			// ищем самую ближнюю координту к toX, toY
			min := true

			x1, y1 := GetXYCenterHex(coor1.Q, coor1.R)
			dist1 := GetBetweenDist(x1, y1, toX, toY)

			for _, coor3 := range radius {

				x3, y3 := GetXYCenterHex(coor3.Q, coor3.R)
				dist3 := GetBetweenDist(x3, y3, toX, toY)

				if dist1 > dist3 && !check(coor3.Q, coor3.R) {
					min = false
					break
				}
			}

			// если координата не добавлени и имеет мин растояние то ретуним
			if !check(coor1.Q, coor1.R) && min {
				return coor1
			}
		}

		return nil
	}

	for i := 0; i < len(units); i++ {
		if i == 0 {
			x, y := GetXYCenterHex(center.Q, center.R)
			toPos = append(toPos, &coordinate.Coordinate{Q: center.Q, R: center.R, X: x, Y: y})
		} else {
			coordinateNear := getNearCoordinate()
			x, y := GetXYCenterHex(coordinateNear.Q, coordinateNear.R)
			toPos = append(toPos, &coordinate.Coordinate{Q: coordinateNear.Q, R: coordinateNear.R, X: x, Y: y})
		}
	}

	return toPos
}
