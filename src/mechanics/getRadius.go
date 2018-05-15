package mechanics

import "./coordinate"

var coordinates = make([]*coordinate.Coordinate, 0)

func GetCoordinates(xCenter int, yCenter int, WatchZone int) []*coordinate.Coordinate {
	circle(xCenter, yCenter, WatchZone, false) // метод отрисовывает только растовый полукруг что бы получить полную фигуруз надо у и х поменять местами и прогнать еще раз
	circle(yCenter, xCenter, WatchZone, true)

	zx := xCenter - WatchZone
	zy := yCenter - WatchZone

	for y := zy; y <= (WatchZone*2+WatchZone)+yCenter; y++ {
		xMax, xMin := xMaxMin(y)
		for x := zx; x <= (WatchZone*2+(WatchZone-1))+xCenter; x++ {
			if xMin < x && xMax > x {
				coordinates = append(coordinates, &coordinate.Coordinate{X: x, Y: y})
			}
		}
	}

	sendCoordinates := removeDuplicates(coordinates)

	defer delCoorinates() // удаляем собранные координаты после ретурна
	return sendCoordinates
}

func xMaxMin(y int) (int, int) {
	var xMax, xMin int

	for i := 0; i < len(coordinates); i++ {
		if i == 0 {
			xMax = coordinates[i].X
			xMin = coordinates[i].X
		} else {
			if coordinates[i].Y == y {
				if coordinates[i].X > xMax {
					xMax = coordinates[i].X
				}
				if coordinates[i].X < xMin {
					xMin = coordinates[i].X
				}
			}
		}
	}
	return xMax, xMin
}

func removeDuplicates(elements []*coordinate.Coordinate) []*coordinate.Coordinate {
	encountered := map[coordinate.Coordinate]bool{}
	result := make([]*coordinate.Coordinate, 0)

	for v := range elements {
		if encountered[*elements[v]] == true {
		} else {
			encountered[*elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func circle(xCenter, yCenter, radius int, invert bool) {
	var x, y, delta int
	x = 0
	y = radius
	delta = 3 - 2*radius

	for x < y { // инопланетные технологии взятые из С++ для формирования растовых окружностей алгоритмом Брезенхэма
		putCoordinates(x, y, xCenter, yCenter, invert)
		putCoordinates(x, y, xCenter, yCenter, invert)
		if delta < 0 {
			delta += 4*x + 6
		} else {
			delta += 4*(x-y) + 10
			y--
		}
		x++
	}
	if x == y {
		putCoordinates(x, y, xCenter, yCenter, invert)
	}
}

func putCoordinates(x int, y int, xCenter int, yCenter int, invert bool) {
	if !invert { // метод отрисовывает только растовый полукруг что бы получить полную фигуруз надо у и х поменять местами и прогнать еще раз
		coordinates = append(coordinates, &coordinate.Coordinate{X: xCenter + x, Y: yCenter + y})
		coordinates = append(coordinates, &coordinate.Coordinate{X: xCenter + x, Y: yCenter - y})
		coordinates = append(coordinates, &coordinate.Coordinate{X: xCenter - x, Y: yCenter + y})
		coordinates = append(coordinates, &coordinate.Coordinate{X: xCenter - x, Y: yCenter - y})
	} else {
		coordinates = append(coordinates, &coordinate.Coordinate{X: yCenter + y, Y: xCenter + x})
		coordinates = append(coordinates, &coordinate.Coordinate{X: yCenter - y, Y: xCenter + x})
		coordinates = append(coordinates, &coordinate.Coordinate{X: yCenter + y, Y: xCenter - x})
		coordinates = append(coordinates, &coordinate.Coordinate{X: yCenter - y, Y: xCenter - x})
	}
}

func delCoorinates() {
	coordinates = nil
}
