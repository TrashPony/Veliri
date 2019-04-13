package hexLineDraw

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"math"
)

func Draw(a, b *coordinate.Coordinate, game *localGame.Game) []*coordinate.Coordinate {
	/* SOURCE Algorithm: https://www.redblobgames.com/grids/hexagons/#line-drawing
		1) First we calculate N=1 to be the hex distance between the endpoints.
	    2) Then evenly sample N+1 points between point A and point B. Using linear interpolation, each point
			will be A + (B - A) * 1.0/N * i, for values of i from 0 to N, inclusive. In the diagram these sample points
			are the dark blue dots. This results in floating point coordinates.
	    3) Convert each sample point (float) back into a hex (int). The algorithm is called cube_round.
	*/
	result := make([]*coordinate.Coordinate, 0)

	b.CalculateXYZ()
	a.CalculateXYZ()

	distance := int(cubeDistance(a, b))

	for i := 0; i <= distance; i++ {
		x, y, z := cubeLerp(a, b, 1.0/float64(distance)*float64(i))
		q, r := cubeRound(x, y, z)

		pathCell, find := game.GetMap().GetCoordinate(q, r)
		if find {
			result = append(result, pathCell)
		}
	}

	return result
}

func cubeRound(x, y, z float64) (int, int) {
	var rx = int(x)
	var ry = int(y)
	var rz = int(z)

	var xDiff = math.Abs(float64(rx) - x)
	var yDiff = math.Abs(float64(ry) - y)
	var zDiff = math.Abs(float64(rz) - z)

	if xDiff > yDiff && xDiff > zDiff {
		rx = -ry - rz
	} else {
		if yDiff > zDiff {
			ry = -rx - rz
		} else {
			rz = -rx - ry
		}
	}

	col := rx + (rz-(rz&1))/2
	row := rz
	return col, row
}
func cubeLerp(a, b *coordinate.Coordinate, t float64) (float64, float64, float64) {
	x := lerp(a.X, b.X, t)
	y := lerp(a.Y, b.Y, t)
	z := lerp(a.Z, b.Z, t)

	return x, y, z
}

func lerp(a, b int, t float64) float64 {

	return float64(a) + (float64(b)-float64(a))*t
}
func cubeDistance(a, b *coordinate.Coordinate) float64 {
	return max(math.Abs(float64(a.X)-float64(b.X)), math.Abs(float64(a.Y)-float64(b.Y)), math.Abs(float64(a.Z)-float64(b.Z)))
}

func max(a, b, c float64) float64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	if c >= a && c >= b {
		return c
	}
	return 0
}
