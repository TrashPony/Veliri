package move

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func Unit(moveUnit *unit.Unit, ToX, ToY float64) ([]*unit.PathUnit, error) {

	moveUnit.ToX = ToX
	moveUnit.ToY = ToY

	startX := float64(moveUnit.X)
	startY := float64(moveUnit.Y)
	rotate := moveUnit.Rotate

	//todo
	moveUnit.MinSpeed = 10

	// т.к. метод расчитывает в секунду а путь строится по 100 мс то скорость делим на 10
	maxSpeed := float64(moveUnit.Speed) / 10
	minSpeed := float64(moveUnit.MinSpeed) / 10

	// если текущая скорость выше стартовой то берем ее
	startSpeed := minSpeed
	if minSpeed < moveUnit.CurrentSpeed {
		startSpeed = moveUnit.CurrentSpeed
	}

	if moveUnit.FollowUnitID != 0 {
		followUnit := globalGame.Clients.GetUnitByID(moveUnit.FollowUnitID)
		dist := game_math.GetBetweenDist(followUnit.X, followUnit.Y, int(moveUnit.X), int(moveUnit.Y))
		if dist < 90 && followUnit.CurrentSpeed > 0 {
			maxSpeed = followUnit.CurrentSpeed
			if followUnit.CurrentSpeed <= 0 {
				return nil, errors.New("follower dont move")
			}
		}
	}

	if moveUnit.Body.MotherShip {
		efficiency := WorkOutThorium(moveUnit.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
		maxSpeed = (maxSpeed * efficiency) / 100
	}

	err, path := To(startX, startY, maxSpeed, minSpeed, startSpeed, ToX, ToY, rotate, 25)

	return path, err
}
