package move

func RotateUnit(unitRotate, needRotate *int, step int) {

	if *unitRotate < 0 {
		*unitRotate += 360
	}

	if *unitRotate >= 360 {
		*unitRotate -= 360
	}

	if *needRotate < 0 {
		*needRotate += 360
	}

	if *needRotate >= 360 {
		*needRotate -= 360
	}

	for i := 0; i < step; i++ {

		if *unitRotate != *needRotate {

			if directionRotate(*unitRotate, *needRotate) {
				*unitRotate++
				if *unitRotate >= 360 {
					*unitRotate -= 360
				}
			} else {
				*unitRotate--
				if *unitRotate < 0 {
					*unitRotate += 360
				}
			}

		} else {
			return
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
