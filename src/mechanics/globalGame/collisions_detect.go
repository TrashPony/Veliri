package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/boxInMap"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"math"
)

func CheckCollisionsOnStaticMap(x, y, rotate int, mp *_map.Map, body *detail.Body) (bool, int, int, bool) {
	startCoordinate := GetQRfromXY(x, y, mp)

	if body == nil {
		return true, startCoordinate.Q, startCoordinate.R, true
	}

	rect := getBodyRect(body, float64(x), float64(y), rotate)
	for _, obstacle := range mp.GeoData {
		if rect.detectCollisionRectToCircle(&point{x: obstacle.X, y: obstacle.Y}, obstacle.Radius) {
			return false, startCoordinate.Q, startCoordinate.R, true
		}
	}

	const reservoirRadius = 15
	for _, qLine := range mp.Reservoir {
		for _, reservoir := range qLine {

			if reservoir.Move() {
				continue
			}

			reservoirX, reservoirY := GetXYCenterHex(reservoir.Q, reservoir.R)
			if rect.detectCollisionRectToCircle(&point{x: reservoirX, y: reservoirY}, reservoirRadius) {
				return false, startCoordinate.Q, startCoordinate.R, true
			}
		}
	}

	return true, startCoordinate.Q, startCoordinate.R, true
}

func CheckCollisionsBoxes(x, y, rotate, mapID int, body *detail.Body) *boxInMap.Box {
	boxs := boxes.Boxes.GetAllBoxByMapID(mapID)

	const boxRadius = 5

	rect := getBodyRect(body, float64(x), float64(y), rotate)
	for _, mapBox := range boxs {

		// поздемные ящики не имеют колизий
		if mapBox.Underground {
			continue
		}

		xBox, yBox := GetXYCenterHex(mapBox.Q, mapBox.R)
		if rect.detectCollisionRectToCircle(&point{x: xBox, y: yBox}, boxRadius) {
			return mapBox
		}
	}
	return nil
}

func CheckCollisionsPlayers(moveUnit *unit.Unit, x, y, rotate int, units map[int]*unit.ShortUnitInfo) (bool, *unit.ShortUnitInfo) {

	for _, otherUnit := range units {

		if otherUnit == nil {
			continue
		}

		if moveUnit.MapID != otherUnit.MapID {
			continue
		}

		if otherUnit != nil && (moveUnit.ID != otherUnit.ID) { // todo && !user.GetSquad().Evacuation

			mUserRect := getBodyRect(moveUnit.Body, float64(x), float64(y), rotate)
			userRect := getBodyRect(otherUnit.Body, float64(otherUnit.X), float64(otherUnit.Y), otherUnit.Rotate)

			if mUserRect.centerX == userRect.centerX && mUserRect.centerY == userRect.centerY {
				// при одинаковом прямоугольнике и одинаковым центром, не будет пересечений и колизия будет не найдена
				// поэтому это тут
				return false, otherUnit
			}

			if mUserRect.detectCollisionRectToRect(&userRect, float64(rotate), float64(otherUnit.Rotate)) {
				return false, otherUnit
			}
		}
	}

	return true, nil
}

func getBodyRect(body *detail.Body, x, y float64, rotate int) rect {

	/*
		squad.rectDebag.moveTo(-50, -25);
		squad.rectDebag.lineTo(-50, +25);

		squad.rectDebag.lineTo(-50, +25);
		squad.rectDebag.lineTo(+50, +25);

		squad.rectDebag.lineTo(+50, +25);
		squad.rectDebag.lineTo(+50, -25);

		squad.rectDebag.lineTo(+50, -25);
		squad.rectDebag.lineTo(-50, -25);

		// A - [0] B - [1] C = [2] D = [3]
	*/

	heightBody, widthBody := float64(body.Height), float64(body.Width)

	bodyRec := rect{
		sides: []*sideRec{
			// A 									// B
			{x1: x - widthBody, y1: y - heightBody, x2: x - widthBody, y2: y + heightBody},
			// B									// C
			{x1: x - widthBody, y1: y + heightBody, x2: x + widthBody, y2: y + heightBody},
			// C									// D
			{x1: x + widthBody, y1: y + heightBody, x2: x + widthBody, y2: y - heightBody},
			// D									// A
			{x1: x + widthBody, y1: y - heightBody, x2: x - widthBody, y2: y - heightBody},
		},
		centerX: float64(x),
		centerY: float64(y),
	}

	bodyRec.rotate(rotate)
	return bodyRec
}

type rect struct {
	sides            []*sideRec
	centerX, centerY float64
}

type sideRec struct {
	x1, y1 float64
	x2, y2 float64
}

type point struct {
	x int
	y int
}

func (r *rect) rotate(rotate int) {

	// поворачиваем квадрат по формуле (x0:y0 - центр)
	//X = (x — x0) * cos(alpha) — (y — y0) * sin(alpha) + x0;
	//Y = (x — x0) * sin(alpha) + (y — y0) * cos(alpha) + y0;

	rotatePoint := func(x, y, x0, y0 float64, rotate int) (newX, newY float64) {
		alpha := float64(rotate) * math.Pi / 180
		newX = (x-x0)*math.Cos(float64(alpha)) - (y-y0)*math.Sin(float64(alpha)) + x0
		newY = (x-x0)*math.Sin(float64(alpha)) + (y-y0)*math.Cos(float64(alpha)) + y0
		return
	}

	rotateSide := func(side *sideRec, x0, y0 float64, rotate int) {
		side.x1, side.y1 = rotatePoint(side.x1, side.y1, x0, y0, rotate)
		side.x2, side.y2 = rotatePoint(side.x2, side.y2, x0, y0, rotate)
	}

	for _, side := range r.sides {
		rotateSide(side, r.centerX, r.centerY, rotate)
	}
}

func detectPointInRectangle(r *rect, x, y int) bool {
	// A - [0]1 B - [1]1 C = [2]1 D = [3]1
	//0 ≤ AP·AB ≤ AB·AB and 0 ≤ AP·AD ≤ AD·AD
	AB := vector(&point{x: int(r.sides[0].x1), y: int(r.sides[0].y1)}, &point{x: int(r.sides[1].x1), y: int(r.sides[1].y1)})
	AM := vector(&point{x: int(r.sides[0].x1), y: int(r.sides[0].y1)}, &point{x: int(x), y: int(y)})
	BC := vector(&point{x: int(r.sides[1].x1), y: int(r.sides[1].y1)}, &point{x: int(r.sides[2].x1), y: int(r.sides[2].y1)})
	BM := vector(&point{x: int(r.sides[1].x1), y: int(r.sides[1].y1)}, &point{x: int(x), y: int(y)})

	return 0 <= dot(AB, AM) && dot(AB, AM) <= dot(AB, AB) && 0 <= dot(BC, BM) && dot(BC, BM) <= dot(BC, BC)
}

func dot(u, v *point) int {
	return u.x*v.x + u.y*v.y
}

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
	var Ex = (t * Dx) + a.x
	var Ey = (t * Dy) + a.y

	// высчитывает растояние от E до центра круга
	var LEC = int(math.Sqrt(math.Pow(float64(Ex-centerCircle.x), 2) + math.Pow(float64(Ey-centerCircle.y), 2)))

	// проверяем что бы проекционная точка была ближе радиуса
	if int(LEC) < radius {
		// compute distance from t to circle intersection point
		var dt = int(math.Sqrt(math.Pow(float64(radius), 2) - math.Pow(float64(LEC), 2)))
		// ищем первую точку пересечения
		point1 := &point{x: (t-dt)*Dx + a.x, y: (t-dt)*Dy + a.y}
		// и вторую
		point2 := &point{x: (t+dt)*Dx + a.x, y: (t+dt)*Dy + a.y}

		intersect := pointInVector(a, b, point1) || pointInVector(a, b, point2)
		return intersect, point1, point2
	}

	if int(LEC) == radius { // else test if the line is tangent to circle
		// прямая прилегает к окружности 1 точка пересечени
		return pointInVector(a, b, &point{x: Ex, y: Ey}), &point{x: Ex, y: Ey}, nil
	}

	return false, nil, nil
}

func pointInVector(a, b, textPoint *point) bool {

	dx1 := int(b.x) - int(a.x)
	dy1 := int(b.y) - int(a.y)

	dx := int(textPoint.x) - int(a.x)
	dy := int(textPoint.y) - int(a.y)

	// небольшая погрешность S <= 1 && S >= -1
	s := dx1*dy - dx*dy1

	return s == 0
}

func (r *rect) detectCollisionRectToCircle(centerCircle *point, radius int) bool {
	// A - [0]1 B - [1]1 C = [2]1 D = [3]1

	if detectPointInRectangle(r, centerCircle.x, centerCircle.y) {
		// цент находится внутри прямоуголника, пересекается
		return true
	}

	/*
		intersectCircle(S, (A, B)) or
	    intersectCircle(S, (B, C)) or
	    intersectCircle(S, (C, D)) or
	    intersectCircle(S, (D, A))
	*/

	a := &point{x: int(r.sides[0].x1), y: int(r.sides[0].y1)}
	b := &point{x: int(r.sides[1].x1), y: int(r.sides[1].y1)}
	c := &point{x: int(r.sides[2].x1), y: int(r.sides[2].y1)}
	d := &point{x: int(r.sides[3].x1), y: int(r.sides[3].y1)}

	intersect1, _, _ := IntersectVectorToCircle(a, b, centerCircle, radius)
	intersect2, _, _ := IntersectVectorToCircle(b, c, centerCircle, radius)
	intersect3, _, _ := IntersectVectorToCircle(c, d, centerCircle, radius)
	intersect4, _, _ := IntersectVectorToCircle(d, a, centerCircle, radius)

	// пересекается 1 из сторон
	if intersect1 || intersect2 || intersect3 || intersect4 {
		return true
	}

	return false
}

func (r *rect) detectCollisionRectToRect(r2 *rect, alpha11, alpha22 float64) bool {
	for _, side1 := range r.sides {
		for _, side2 := range r2.sides {
			if intersection(side1.x1, side1.y1, side1.x2, side1.y2, side2.x1, side2.y1, side2.x2, side2.y2) {
				return true
			}
		}
	}

	return false
}

func intersection(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
	v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
	v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
	v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
	v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)

	return (v1*v2 < 0) && (v3*v4 < 0)
}
