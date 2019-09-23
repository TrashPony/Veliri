package move

import (
	"errors"
	"fmt"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/satori/go.uuid"
	"math"
	"strconv"
	"time"
)

func Unit(moveUnit *unit.Unit, ToX, ToY, StartX, StartY float64) ([]*unit.PathUnit, error) {

	start := time.Now()

	defer func() {
		// TODO идиальное время 200 мс :С
		if debug.Store.Move {
			elapsed := time.Since(start)
			fmt.Println("time all path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
			fmt.Println("--------------------------------------------------")
		}
	}()

	moveUUID := uuid.NewV1().String()
	moveUnit.MoveUUID = moveUUID
	moveUnit.ToX = ToX
	moveUnit.ToY = ToY

	startX := StartX
	startY := StartY
	rotate := 90

	maxSpeed := float64(moveUnit.Speed)
	if moveUnit.Body.MotherShip {
		efficiency := WorkOutThorium(moveUnit.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
		maxSpeed = (maxSpeed * efficiency) / 100
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

	// что бы игрок не смог сгенерить одновременно много путей
	pathPoints, err := find_path.LeftHandAlgorithm(moveUnit, startX, startY, ToX, ToY, moveUUID)
	if err != nil {
		return nil, err
	}

	if moveUUID == moveUnit.MoveUUID {
		return CreatePath(pathPoints, startX, startY, maxSpeed, moveUnit, rotate), nil
	} else {
		return nil, errors.New("wrong uuid")
	}
}

func CreatePath(pathPoints []*coordinate.Coordinate, startX, startY, maxSpeed float64, moveUnit *unit.Unit, rotate int) []*unit.PathUnit {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time create path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	path := make([]*unit.PathUnit, 0)

	timeUnit := 250

	appendPath := func(appendPath []*unit.PathUnit) {
		for i := 0; i < len(appendPath); i++ {
			path = append(path, appendPath[i])
		}
	}

	for i, pathPoint := range pathPoints {
		if i == 0 || len(path) == 0 {
			_, path = UnitTo(float64(startX), float64(startY), maxSpeed, float64(pathPoint.X), float64(pathPoint.Y), moveUnit.Rotate, rotate, timeUnit)
		} else {

			lastX, lastY, lastAngle := float64(path[len(path)-1].X), float64(path[len(path)-1].Y), path[len(path)-1].Rotate

			_, aPath := UnitTo(lastX, lastY, maxSpeed, float64(pathPoint.X), float64(pathPoint.Y), lastAngle, rotate, timeUnit)
			appendPath(aPath)
		}
	}

	return path
}

func UnitTo(forecastX, forecastY, speed, ToX, ToY float64, rotate, rotateAngle, ms int) (error, []*unit.PathUnit) {

	// TODO искать предварительно часть пути до тех пока не будет разница в углу 0
	// TODO в разных вариантах, разворот на скорости или развород на месте, выбирать то где мешье частей пути (выше скорость)
	//  или если 1 из способов не пройти

	// TODO если юнит имеет высокую скорость последние точки делить его путь что бы адекватно обработать колизии

	speed = speed / float64(1000/ms)
	path := make([]*unit.PathUnit, 0)

	for {
		// находим длинную вектора до цели
		dist := game_math.GetBetweenDist(int(forecastX), int(forecastY), int(ToX), int(ToY))
		if dist < speed+5 {
			break
		}

		radRotate := float64(rotate) * math.Pi / 180
		stopX := float64(speed) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(speed) * math.Sin(radRotate)

		//находим какой угол необходимо принять телу
		needRotate := game_math.GetBetweenAngle(ToX, ToY, forecastX, forecastY)
		RotateUnit(&rotate, &needRotate, rotateAngle)

		forecastX = forecastX + stopX
		forecastY = forecastY + stopY

		q, r := game_math.GetQRfromXY(int(forecastX), int(forecastY))
		path = append(path, &unit.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: ms,
			Q: q, R: r, Speed: speed, Animate: true})
	}

	return nil, path
}
