package collisions

import "math"

func vector(p1, p2 *point) *point {
	return &point{
		x: p2.x - p1.x,
		y: p2.y - p1.y,
	}
}

func IntersectVectorToCircle(a, b, centerCircle *point, radius int) (intersect bool, point1, point2 *point) {
	// https://stackoverflow.com/questions/1073336/circle-line-segment-collision-detection-algorithm
	// вычисляем расстояние между A и B
	var LAB = int(math.Sqrt(math.Pow(float64(b.x-a.x), 2) + math.Pow(float64(b.y-a.y), 2)))

	// вычислить вектор направления D от A до B
	var Dx = (b.x - a.x) / LAB
	var Dy = (b.y - a.y) / LAB

	// compute the value t of the closest point to the circle center (Cx, Cy)
	var t = Dx*(centerCircle.x-a.x) + Dy*(centerCircle.y-a.y)

	// This is the projection of C on the line from A to B.

	// вычислить координаты точки E на прямой и ближайшей к C
	E := &point{x: (t * Dx) + a.x, y: (t * Dy) + a.y}

	// высчитывает растояние от E до центра круга
	var LEC = int(math.Sqrt(math.Pow(float64(E.x-centerCircle.x), 2) + math.Pow(float64(E.y-centerCircle.y), 2)))

	// проверяем что бы проекционная точка была ближе радиуса
	if int(LEC) < radius {
		// compute distance from t to circle intersection point
		var dt = int(math.Sqrt(math.Pow(float64(radius), 2) - math.Pow(float64(LEC), 2)))
		// ищем первую точку пересечения
		point1 := &point{x: (t-dt)*Dx + a.x, y: (t-dt)*Dy + a.y}
		// и вторую
		point2 := &point{x: (t+dt)*Dx + a.x, y: (t+dt)*Dy + a.y}

		intersect := point1.pointInVector(a, b) || point2.pointInVector(a, b)
		return intersect, point1, point2
	}

	if int(LEC) == radius { // else test if the line is tangent to circle
		// прямая прилегает к окружности 1 точка пересечени
		return E.pointInVector(a, b), E, nil
	}

	return false, nil, nil
}
