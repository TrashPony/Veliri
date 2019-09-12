package move

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/getlantern/deepcopy"
)

func Unit(moveUnit *unit.Unit, ToX, ToY float64, mp *_map.Map) ([]unit.PathUnit, error) {

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

	// копируем что бы не произошло вычетание топлива на расчетах
	var fakeThoriumSlots map[int]*detail.ThoriumSlot
	if moveUnit.Body.MotherShip {
		err := deepcopy.Copy(&fakeThoriumSlots, &moveUnit.Body.ThoriumSlots)
		if err != nil || len(fakeThoriumSlots) == 0 {
			println(err.Error())
			return nil, err
		}
	} else {
		// если это юнит то проецируем его энергию в топливо т.к. у юнита нет реактора и для двжиения он тратить свой акум
		if moveUnit.Power <= 0 {
			return nil, errors.New("not thorium")
		}

		fakeThoriumSlots = make(map[int]*detail.ThoriumSlot)
		fakeThoriumSlots[1] = &detail.ThoriumSlot{Number: 1, WorkedOut: float32(moveUnit.Power), Inversion: true, Count: 1}
	}

	if moveUnit.Afterburner { // если форсаж то х2 скорости (доступно только МС)
		maxSpeed = maxSpeed * 2
	}

	err, path := To(startX, startY, maxSpeed, minSpeed, startSpeed, ToX, ToY, rotate, 5, mp, false,
		fakeThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity, moveUnit.Body)

	return path, err
}
