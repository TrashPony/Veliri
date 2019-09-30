package move

import (
	"errors"
	"fmt"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"strconv"
	"time"
)

func Unit(moveUnit *unit.Unit, ToX, ToY, StartX, StartY float64, unitRotate int, uuid string, units map[int]*unit.ShortUnitInfo) ([]*unit.PathUnit, error) {

	if uuid != moveUnit.MoveUUID {
		return nil, errors.New("wrong uuid")
	}

	start := time.Now()

	defer func() {
		// TODO идиальное время 200 мс :С
		if debug.Store.Move {
			elapsed := time.Since(start)
			fmt.Println("time all path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
			fmt.Println("--------------------------------------------------")
		}
	}()

	moveUnit.ToX = ToX
	moveUnit.ToY = ToY

	startX := StartX
	startY := StartY

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
	pathPoints, err := find_path.LeftHandAlgorithm(moveUnit, startX, startY, ToX, ToY, uuid, units)
	if err != nil {
		return nil, err
	}

	if uuid == moveUnit.MoveUUID {
		return CreatePath(pathPoints, startX, startY, maxSpeed, moveUnit, unitRotate, units), nil
	} else {
		return nil, errors.New("wrong uuid")
	}
}

func CreatePath(pathPoints []*coordinate.Coordinate, startX, startY, maxSpeed float64, moveUnit *unit.Unit, unitRotate int, units map[int]*unit.ShortUnitInfo) []*unit.PathUnit {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time create path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	path := make([]*unit.PathUnit, 0)

	timeUnit := 250
	rotate := 30

	for i, pathPoint := range pathPoints {
		if i == 0 || len(path) == 0 {
			path, _ = UnitTo(float64(startX), float64(startY), maxSpeed, float64(pathPoint.X), float64(pathPoint.Y),
				unitRotate, rotate, timeUnit, false, false, true, mp, moveUnit, units)
		} else {

			lastX, lastY, lastAngle := float64(path[len(path)-1].X), float64(path[len(path)-1].Y), path[len(path)-1].Rotate

			aPath, _ := UnitTo(lastX, lastY, maxSpeed, float64(pathPoint.X), float64(pathPoint.Y), lastAngle, rotate,
				timeUnit, false, false, true, mp, moveUnit, units)

			appendPath(aPath, &path)
		}
	}

	return path
}

func appendPath(appendPath []*unit.PathUnit, path *[]*unit.PathUnit) {
	for i := 0; i < len(appendPath); i++ {
		*path = append(*path, appendPath[i])
	}
}

func UnitTo(forecastX, forecastY, speed, ToX, ToY float64, rotate, rotateAngle, ms int, searchCollision, onlyRotate, start bool,
	mp *_map.Map, moveUnit *unit.Unit, units map[int]*unit.ShortUnitInfo) ([]*unit.PathUnit, bool) {

	path := make([]*unit.PathUnit, 0)

	if start {
		// пытаемя сгенерить путь с поворотом в движение
		rotatePath, collision := UnitTo(forecastX, forecastY, speed, ToX, ToY, rotate, rotateAngle, ms,
			true, false, false, mp, moveUnit, units)
		// и с поворотом на месте
		noRotatePath, _ := UnitTo(forecastX, forecastY, speed, ToX, ToY, rotate, rotateAngle, ms,
			false, true, false, mp, moveUnit, units)

		// если с поворотом в движение будет колизия то берем путь без движения
		if collision {
			path = noRotatePath
		} else {
			path = rotatePath
		}

		// берем данные из последней точки пути
		if len(path) > 0 {
			forecastX, forecastY, rotate = float64(path[len(path)-1].X), float64(path[len(path)-1].Y), path[len(path)-1].Rotate
		}
	}

	// скорость машинок указывается в сек, поэтому корректируем ее под время
	if onlyRotate {
		speed = 0
	} else {
		speed = speed / float64(1000/ms)
	}

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
		countRotateAngle := RotateUnit(&rotate, &needRotate, rotateAngle)

		if searchCollision {
			// TODO если юнит имеет высокую скорость последние точки делить его путь что бы адекватно обработать колизии
			possibleMove, _ := collisions.CheckCollisionsOnStaticMap(int(forecastX), int(forecastY), rotate, mp, moveUnit.Body, false, false)
			if !possibleMove {
				return path, true
			}

			if units != nil {
				free, _ := collisions.CheckCollisionsPlayers(moveUnit, int(forecastX), int(forecastY), 0, units, true, false, false)
				if !free {
					return path, true
				}
			}
		}

		forecastX = forecastX + stopX
		forecastY = forecastY + stopY

		path = append(path, &unit.PathUnit{X: int(forecastX), Y: int(forecastY), Rotate: rotate, Millisecond: ms, Speed: speed, Animate: true})

		if countRotateAngle == 0 && onlyRotate {
			return path, false
		}
	}

	return path, false
}
