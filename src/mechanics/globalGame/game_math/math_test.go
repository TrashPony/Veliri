package game_math

import "testing"

func TestCoordinateTranslation(t *testing.T) {
	x, y := GetXYCenterHex(10, 10)
	q, r := GetQRfromXY(x, y)
	if q != 10 || r != 10 {
		t.Error("coordinate translation test 1 failed ")
	}

	x, y = GetXYCenterHex(17, 12)
	q, r = GetQRfromXY(x, y)
	if q != 17 || r != 12 {
		t.Error("coordinate translation test 2 failed ")
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {

			q, r := GetQRfromXY(x, y)
			xTest, yTest := GetXYCenterHex(q, r)
			qTest, rTest := GetQRfromXY(xTest, yTest)

			if qTest != q || rTest != r {
				t.Error("coordinate test 3 failed ")
				t.Error("x:", x, "y:", y)
				t.Error("newX:", xTest, "newY:", yTest)
				t.Error("q:", q, "r:", r)
				t.Error("newQ:", qTest, "newR:", rTest)
			}
		}
	}

	for q := 0; q < 100; q++ {
		for r := 0; r < 100; r++ {

			xTest, yTest := GetXYCenterHex(q, r)
			qTest, rTest := GetQRfromXY(xTest, yTest)

			if qTest != q || rTest != r {
				t.Error("coordinate test 4 failed ")
				t.Error("newX:", xTest, "newY:", yTest)
				t.Error("q:", q, "r:", r)
				t.Error("newQ:", qTest, "newR:", rTest)
			}
		}
	}
}
