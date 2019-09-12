package collisions

import "testing"

func TestCollisionRectToCircle(t *testing.T) {
	bodyRec := Polygon{
		sides: []*sideRec{
			// A 									// B
			{x1: 1425, y1: 1057, x2: 1425, y2: 1087},
			// B									// C
			{x1: 1425, y1: 1087, x2: 1459, y2: 1087},
			// C									// D
			{x1: 1459, y1: 1087, x2: 1459, y2: 1057},
			// D									// A
			{x1: 1459, y1: 1057, x2: 1425, y2: 1057},
		},
		centerX: float64(1442),
		centerY: float64(1072),
	}

	if bodyRec.detectCollisionRectToCircle(&point{2907, 1095}, 8) {
		t.Error("intersect rect to circle test 1 failed ")
	}
}

func TestIntersectVectorToCircle(t *testing.T) {
	//
	// прямая не пересекает окружность
	intersect, point1, point2 := IntersectVectorToCircle(&point{x: 20, y: 10}, &point{x: 10, y: 50}, &point{x: 30, y: 30}, 10)
	if intersect || point1 != nil || point2 != nil {
		t.Error("intersect vector to circle test 1 failed ")
	}
	//
	//// прямая пересекает окружность
	intersect, point1, point2 = IntersectVectorToCircle(&point{x: 50, y: 10}, &point{x: 10, y: 50}, &point{x: 30, y: 30}, 10)
	if !intersect || int(point1.x) != 37 || int(point1.y) != 22 || int(point2.x) != 22 || int(point2.y) != 37 {
		t.Error("intersect vector to circle test 2 failed ")
	}
	//
	//// прямая пересекает окружность
	intersect, point1, point2 = IntersectVectorToCircle(&point{x: 50, y: 10}, &point{x: 15, y: 30}, &point{x: 30, y: 30}, 10)
	println(intersect, int(point1.x), int(point1.y), int(point2.x), int(point2.y))
	if !intersect || int(point1.x) != 32 || int(point1.y) != 20 || int(point2.x) != 20 || int(point2.y) != 26 {
		t.Error("intersect vector to circle test 3 failed ")
	}

	// прямая прилегает к окружности
	intersect, point1, point2 = IntersectVectorToCircle(&point{x: 50, y: 10}, &point{x: 10, y: 28}, &point{x: 30, y: 30}, 10)
	if !intersect || int(point1.x) != 25 || int(point1.y) != 20 || point2 != nil {
		t.Error("intersect vector to circle test 4 failed ")
	}

	// прямая не в окружности но вектор пересекает ее
	intersect, point1, point2 = IntersectVectorToCircle(&point{x: 50, y: 10}, &point{x: 40, y: 20}, &point{x: 30, y: 30}, 10)
	if intersect || int(point1.x) != 37 || int(point1.y) != 22 || int(point2.x) != 22 || int(point2.y) != 37 {
		t.Error("intersect vector to circle test 5 failed ")
	}

	// прямая не в окружности но вектор пересекает ее в 1 точке
	intersect, point1, point2 = IntersectVectorToCircle(&point{x: 50, y: 20}, &point{x: 40, y: 20}, &point{x: 30, y: 30}, 10)
	if intersect && int(point1.x) != 30 || int(point1.y) != 20 || point2 != nil {
		t.Error("intersect vector to circle test 6 failed ")
	}
}

func TestPointInVector(t *testing.T) {

	// принадлежит отрезку
	testPoint := &point{x: 45, y: 20}
	if !testPoint.pointInVector(&point{x: 50, y: 20}, &point{x: 40, y: 20}) {
		t.Error("intersect point in vector test 1 failed ")
	}

	// не принадлежит отрезку
	if testPoint.pointInVector(&point{x: 50, y: 25}, &point{x: 40, y: 20}) {
		t.Error("intersect point in vector test 2 failed ")
	}

	// не принадлежит отрезку
	if testPoint.pointInVector(&point{x: 50, y: 5}, &point{x: 40, y: 20}) {
		t.Error("intersect point in vector test 3 failed ")
	}

	// не принадлежит отрезку
	testPoint = &point{x: 2841, y: 662}
	if testPoint.pointInVector(&point{x: 1438, y: 986}, &point{x: 1467, y: 980}) {
		t.Error("intersect point in vector test 4 failed ")
	}
}
