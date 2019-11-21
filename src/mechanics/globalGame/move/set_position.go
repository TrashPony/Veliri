package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"math"
	"time"
)

func SetPosition(moveUnit *unit.Unit, pathUnit *unit.PathUnit, deltaTime int64) {

	var setAngle, setX, setY bool

	// todo дельта не помогает ногда прогрешности аж до 15 мс доходят
	pathUnit.Millisecond -= int(deltaTime) // учитываем время расчетов

	go func() {
		defer func() {
			setAngle = true
		}()
		// горутина создана что бы оружием и тело имели актуальное положение независимо от тика сервера
		// иначе угол и положение орудия берется не правильно, из за чего в движение огонь выглядит странно

		diffRotate := int(math.Abs(float64(pathUnit.Rotate - moveUnit.Rotate)))

		if diffRotate > 0 {
			msToOneAngle := pathUnit.Millisecond / diffRotate

			for {
				oldRotate := moveUnit.Rotate
				if RotateUnit(&moveUnit.Rotate, &pathUnit.Rotate, 1) == 0 {
					return
				}
				moveUnit.GunRotate -= oldRotate - moveUnit.Rotate
				time.Sleep(time.Millisecond * time.Duration(msToOneAngle))
			}
		}
	}()

	setPos := func(value *int, needValue int, flag *bool) {

		defer func() {
			*flag = true
		}()

		diff := int(math.Abs(float64(pathUnit.Y - *value)))

		if diff > 0 {

			msToOneAngle := pathUnit.Millisecond / diff

			for {

				if *value < needValue {
					*value++
				}

				if *value > needValue {
					*value--
				}

				time.Sleep(time.Millisecond * time.Duration(msToOneAngle))

				if *value == needValue {
					return
				}
			}
		} else {
			time.Sleep(time.Millisecond * time.Duration(pathUnit.Millisecond))
		}
	}

	go setPos(&moveUnit.X, pathUnit.X, &setX)
	go setPos(&moveUnit.Y, pathUnit.Y, &setY)

	for !setAngle || !setX || !setY {
		time.Sleep(time.Millisecond)
	}
}
