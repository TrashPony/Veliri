package globalGame

import "../player"

type PathUnit struct {
	X           int `json:"x"`
	Y           int `json:"y"`
	Rotate      int `json:"rotate"`
	Millisecond int `json:"millisecond"`
}

func MoveTo(user *player.Player) {

}

func RotateUnit(unitRotate, needRotate *int) int {

	if *unitRotate < 0 {
		*unitRotate += 360
	}

	if *unitRotate > 360 {
		*unitRotate -= 360
	}

	if *needRotate < 0 {
		*needRotate += 360
	}

	if *needRotate > 360 {
		*needRotate -= 360
	}

	if unitRotate != needRotate {
		if directionRotate(*unitRotate, *needRotate) {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func directionRotate(unitAngle, needAngle int) bool {
	// true ++
	// false --
	count := 0
	direction := false

	if unitAngle < needAngle {
		for unitAngle < needAngle {
			count++
			direction = true
			unitAngle++
		}
	} else {
		for unitAngle > needAngle {
			count++
			direction = false
			needAngle++
		}
	}

	if direction {
		return count <= 180
	} else {
		return !(count <= 180)
	}
}
