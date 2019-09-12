package move

func RotateUnit(unitRotate, needRotate *int, step int) {

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

	if !(*unitRotate > *needRotate && *unitRotate-*needRotate > step || *needRotate > *unitRotate && *needRotate-*unitRotate > step) {
		if *unitRotate >= *needRotate {
			step = *unitRotate - *needRotate
		} else {
			step = *needRotate - *unitRotate
		}
	}

	if unitRotate != needRotate {
		if directionRotate(*unitRotate, *needRotate) {
			*unitRotate += step
		} else {
			*unitRotate -= step
		}
	}
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
