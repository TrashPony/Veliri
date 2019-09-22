package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"math"
)

type Polygon struct {
	sides            []*sideRec
	centerX, centerY float64
}

type sideRec struct {
	x1, y1 float64
	x2, y2 float64
}

func (r *Polygon) rotate(rotate int) {

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

func (r *Polygon) detectCollisionRectToCircle(centerCircle *point, radius int) bool {
	// A - [0]1 B - [1]1 C = [2]1 D = [3]1

	if r.detectPointInRectangle(centerCircle.x, centerCircle.y) {
		// цент находится внутри прямоуголника, пересекается
		return true
	}

	/*
			intersectCircle(S, (A, B)) or
		    intersectCircle(S, (B, C)) or
		    intersectCircle(S, (C, D)) or
		    intersectCircle(S, (D, A))
	*/

	a := &point{x: r.sides[0].x1, y: r.sides[0].y1}
	b := &point{x: r.sides[1].x1, y: r.sides[1].y1}
	c := &point{x: r.sides[2].x1, y: r.sides[2].y1}
	d := &point{x: r.sides[3].x1, y: r.sides[3].y1}

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

func (r *Polygon) detectCollisionRectToRect(r2 *Polygon, alpha11, alpha22 float64) bool {

	intersection := func(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
		v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
		v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
		v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
		v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)

		return (v1*v2 < 0) && (v3*v4 < 0)
	}

	for _, side1 := range r.sides {
		for _, side2 := range r2.sides {
			if intersection(side1.x1, side1.y1, side1.x2, side1.y2, side2.x1, side2.y1, side2.x2, side2.y2) {
				return true
			}
		}
	}

	return false
}

func (r *Polygon) detectPointInRectangle(x, y float64) bool {
	dot := func(u, v *point) float64 {
		return u.x*v.x + u.y*v.y
	}

	// A - [0]1 B - [1]1 C = [2]1 D = [3]1
	//0 ≤ AP·AB ≤ AB·AB and 0 ≤ AP·AD ≤ AD·AD
	AB := vector(&point{x: r.sides[0].x1, y: r.sides[0].y1}, &point{x: r.sides[1].x1, y: r.sides[1].y1})
	AM := vector(&point{x: r.sides[0].x1, y: r.sides[0].y1}, &point{x: x, y: y})
	BC := vector(&point{x: r.sides[1].x1, y: r.sides[1].y1}, &point{x: r.sides[2].x1, y: r.sides[2].y1})
	BM := vector(&point{x: r.sides[1].x1, y: r.sides[1].y1}, &point{x: x, y: y})

	return 0 <= dot(AB, AM) && dot(AB, AM) <= dot(AB, AB) && 0 <= dot(BC, BM) && dot(BC, BM) <= dot(BC, BC)
}

func getBodyRect(body *detail.Body, x, y float64, rotate int, full, min bool) *Polygon {

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

	if full {
		if heightBody > widthBody {
			widthBody = heightBody
		} else {
			heightBody = widthBody
		}
	}

	if min {
		if heightBody < widthBody {
			widthBody = heightBody
		} else {
			heightBody = widthBody
		}
	}

	bodyRec := getCenterRect(x, y, heightBody*2, widthBody*2)
	bodyRec.rotate(rotate)
	return bodyRec
}

func getCenterRect(x, y, height, width float64) *Polygon {

	// делем на 2 что бы центр квадрата был в х у
	height = height / 2
	width = width / 2

	return &Polygon{
		sides: []*sideRec{
			// A 									// B
			{x1: x - width, y1: y - height, x2: x - width, y2: y + height},
			// B									// C
			{x1: x - width, y1: y + height, x2: x + width, y2: y + height},
			// C									// D
			{x1: x + width, y1: y + height, x2: x + width, y2: y - height},
			// D									// A
			{x1: x + width, y1: y - height, x2: x - width, y2: y - height},
		},
		centerX: float64(x),
		centerY: float64(y),
	}
}

func getRect(x, y, height, width float64) *Polygon {
	return &Polygon{
		sides: []*sideRec{
			// A 									// B
			{x1: x, y1: y, x2: x, y2: y + height},
			// B									// C
			{x1: x, y1: y + height, x2: x + width, y2: y + height},
			// C									// D
			{x1: x + width, y1: y + height, x2: x + width, y2: y},
			// D									// A
			{x1: x + width, y1: y, x2: x, y2: y},
		},
		centerX: float64(x + width/2),
		centerY: float64(y + height/2),
	}
}
