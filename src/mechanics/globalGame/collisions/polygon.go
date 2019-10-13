package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

type Polygon struct {
	Sides            []*SideRec `json:"sides"`
	centerX, centerY float64
	Height, Width    float64
	Angle            int
}

type SideRec struct {
	X1 float64 `json:"x_1"`
	Y1 float64 `json:"y_1"`
	X2 float64 `json:"x_2"`
	Y2 float64 `json:"y_2"`
}

func (r *Polygon) Rotate(rotate int) {

	r.Angle = rotate

	rotateSide := func(side *SideRec, x0, y0 float64, rotate int) {
		side.X1, side.Y1 = game_math.RotatePoint(side.X1, side.Y1, x0, y0, rotate)
		side.X2, side.Y2 = game_math.RotatePoint(side.X2, side.Y2, x0, y0, rotate)
	}

	for _, side := range r.Sides {
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

	a := &point{x: r.Sides[0].X1, y: r.Sides[0].Y1}
	b := &point{x: r.Sides[1].X1, y: r.Sides[1].Y1}
	c := &point{x: r.Sides[2].X1, y: r.Sides[2].Y1}
	d := &point{x: r.Sides[3].X1, y: r.Sides[3].Y1}

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

func (r *Polygon) detectCollisionRectToRect(r2 *Polygon) bool {

	if r.detectPointInRectangle(float64(r2.centerX), float64(r2.centerY)) {
		// цент находится внутри прямоуголника, пересекается
		return true
	}

	if r2.detectPointInRectangle(float64(r.centerX), float64(r.centerY)) {
		// цент находится внутри прямоуголника, пересекается
		return true
	}

	if r.centerX == r2.centerX && r.centerY == r2.centerY {
		// при одинаковом прямоугольнике и одинаковым центром, не будет пересечений и колизия будет не найдена
		// поэтому это тут
		return true
	}

	intersection := func(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
		v1 := (bx2-bx1)*(ay1-by1) - (by2-by1)*(ax1-bx1)
		v2 := (bx2-bx1)*(ay2-by1) - (by2-by1)*(ax2-bx1)
		v3 := (ax2-ax1)*(by1-ay1) - (ay2-ay1)*(bx1-ax1)
		v4 := (ax2-ax1)*(by2-ay1) - (ay2-ay1)*(bx2-ax1)

		return (v1*v2 < 0) && (v3*v4 < 0)
	}

	for _, side1 := range r.Sides {
		for _, side2 := range r2.Sides {
			if intersection(side1.X1, side1.Y1, side1.X2, side1.Y2, side2.X1, side2.Y1, side2.X2, side2.Y2) {
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
	AB := vector(&point{x: r.Sides[0].X1, y: r.Sides[0].Y1}, &point{x: r.Sides[1].X1, y: r.Sides[1].Y1})
	AM := vector(&point{x: r.Sides[0].X1, y: r.Sides[0].Y1}, &point{x: x, y: y})
	BC := vector(&point{x: r.Sides[1].X1, y: r.Sides[1].Y1}, &point{x: r.Sides[2].X1, y: r.Sides[2].Y1})
	BM := vector(&point{x: r.Sides[1].X1, y: r.Sides[1].Y1}, &point{x: x, y: y})

	return 0 <= dot(AB, AM) && dot(AB, AM) <= dot(AB, AB) && 0 <= dot(BC, BM) && dot(BC, BM) <= dot(BC, BC)
}

func GetBodyRect(body *detail.Body, x, y float64, rotate int, full, min bool) *Polygon {

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

	bodyRec := GetCenterRect(x, y, heightBody*2, widthBody*2)
	bodyRec.Rotate(rotate)
	return bodyRec
}

func GetCenterRect(x, y, height, width float64) *Polygon {

	// делем на 2 что бы центр квадрата был в х у
	height = height / 2
	width = width / 2

	return GetRect(x, y, height, width)
}

func GetRect(x, y, height, width float64) *Polygon {
	return &Polygon{
		Sides: []*SideRec{
			// A 									// B
			{X1: x - width, Y1: y - height, X2: x - width, Y2: y + height},
			// B									// C
			{X1: x - width, Y1: y + height, X2: x + width, Y2: y + height},
			// C									// D
			{X1: x + width, Y1: y + height, X2: x + width, Y2: y - height},
			// D									// A
			{X1: x + width, Y1: y - height, X2: x - width, Y2: y - height},
		},
		centerX: float64(x),
		centerY: float64(y),
		Height:  height,
		Width:   width,
	}
}
