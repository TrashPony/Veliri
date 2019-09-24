package move

func RotateUnit(unitRotate, needRotate *int, step int) int {

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

	countRotateAngle := 0

	for i := 0; i < step; i++ {

		if *unitRotate != *needRotate {

			if directionRotate(*unitRotate, *needRotate) {
				*unitRotate++
				if *unitRotate >= 360 {
					*unitRotate -= 360
				}

				countRotateAngle++

			} else {
				*unitRotate--
				if *unitRotate < 0 {
					*unitRotate += 360
				}

				countRotateAngle++

			}

		} else {
			return countRotateAngle
		}
	}

	return countRotateAngle
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
