package move

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"strconv"
)

func Unit(moveUnit *unit.Unit, ToX, ToY float64) ([]*unit.PathUnit, error) {
	moveUnit.ToX = ToX
	moveUnit.ToY = ToY

	startX := float64(moveUnit.X)
	startY := float64(moveUnit.Y)

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

	return LeftHandAlgorithm(moveUnit, startX, startY, ToX, ToY, maxSpeed)
}

func To2(forecastX, forecastY, speed, ToX, ToY float64, rotate, rotateAngle, ms int) (error, []*unit.PathUnit) {

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

	if len(path) > 1 {
		path[len(path)-1].Speed = 0 // на последней точке машина останавливается
	}

	return nil, path
}

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY, maxSpeed float64) ([]*unit.PathUnit, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)
	rotate := 180

	xCollision, yCollision, xEnd, yEnd, collision, endIsObstacle := BetweenLine(startX, startY, ToX, ToY, mp, moveUnit.Body, true)

	path := make([]*unit.PathUnit, 0)
	appendPath := func(appendPath []*unit.PathUnit) {
		for i := 0; i < len(appendPath); i++ {
			path = append(path, appendPath[i])
		}
	}

	// если конечная точка нгаходится в первом препятсвие то смотрим куда ближе идти ко входу или к выходу
	if endIsObstacle {
		collisionStartDist := game_math.GetBetweenDist(int(ToX), int(ToY), xCollision, yCollision)
		collisionEndDist := game_math.GetBetweenDist(int(ToX), int(ToY), xEnd, yEnd)

		// если то старта колизии ближе чем до конца то считаем что маршрут без колизий
		if collisionStartDist < collisionEndDist {
			moveUnit.ToX, moveUnit.ToY = float64(xCollision), float64(yCollision)
			collision = false
		} else {
			// иначе переназначаем конечный пункт что бы не искать путь вечно
			ToX, ToY = float64(xCollision), float64(yCollision)
		}
	}

	if !collision {
		//на прямой между стартом и концом нет прептявий

		err, path := To2(startX, startY, maxSpeed, moveUnit.ToX, moveUnit.ToY, moveUnit.Rotate, rotate, 200)
		return path, err
	} else {
		// на прямой были найдены препятвия

		x, y, err := FindExtremum(mp, &coordinate.Coordinate{X: xCollision, Y: yCollision}, &coordinate.Coordinate{X: xEnd, Y: yEnd}, moveUnit.X, moveUnit.Y, moveUnit)
		if err != nil {
			return nil, err
		}

		_, path = To2(startX, startY, maxSpeed, float64(x), float64(y), moveUnit.Rotate, rotate, 200)

		for {

			xCollision, yCollision, xEnd, yEnd, collision, endIsObstacle = BetweenLine(float64(x), float64(y), ToX, ToY, mp, moveUnit.Body, false)
			if !collision {
				_, aPath := To2(float64(path[len(path)-1].X), float64(path[len(path)-1].Y), maxSpeed, moveUnit.ToX, moveUnit.ToY, path[len(path)-1].Rotate, rotate, 200)
				appendPath(aPath)
				break
			} else {

				x, y, err = FindExtremum(mp, &coordinate.Coordinate{X: xCollision, Y: yCollision}, &coordinate.Coordinate{X: xEnd, Y: yEnd}, x, y, moveUnit)
				if err != nil {
					return nil, err
				}

				_, aPath := To2(float64(path[len(path)-1].X), float64(path[len(path)-1].Y), maxSpeed, float64(x), float64(y), path[len(path)-1].Rotate, rotate, 200)
				appendPath(aPath)
			}
		}

		return path, nil
	}
}

// возвращает xКолизии, yКолизии, хВыхода из колизии, yВыхода из колизии прошли без столкновений
func BetweenLine(startX, startY, ToX, ToY float64, mp *_map.Map, body *detail.Body, startMove bool) (xCollision, yCollision, xEnd, yEnd int, collision, endIsObstacle bool) {

	// TODO учитывать угол поворота юнита
	// идем по линии со скорость 10 рх
	speed := 10.0

	// угол от старта до конца
	angle := game_math.GetBetweenAngle(ToX, ToY, startX, startY)
	radian := float64(angle) * math.Pi / 180

	// текущее положение курсора
	currentX, currentY := startX, startY

	for {
		// находим длинную вектора до цели
		distToEnd := game_math.GetBetweenDist(int(currentX), int(currentY), int(ToX), int(ToY))
		distToStart := game_math.GetBetweenDist(int(currentX), int(currentY), int(startX), int(startY))

		if distToEnd < speed+2 {
			if !collision {
				return 0, 0, 0, 0, false, false
			} else {
				endIsObstacle = true
			}
		}

		stopX, stopY := float64(speed)*math.Cos(radian), float64(speed)*math.Sin(radian)

		possibleMove, _ := collisions.CheckCollisionsOnStaticMap(int(currentX), int(currentY), angle, mp, body, true, false)
		if !possibleMove {
			// если юнит по каким то причинам стартует из колизии то дать ему выйти и потом уже искать колизию
			if !(distToStart < speed+2 && startMove) {
				// фиксируем 1 точку колизии
				if !collision {
					xCollision = int(currentX)
					yCollision = int(currentY)
					collision = true
				}
			}
		} else {
			// фиксируем точку выхода из колизии
			if collision {
				return xCollision, yCollision, int(currentX), int(currentY), collision, endIsObstacle
			}
		}

		currentX += stopX
		currentY += stopY
	}
}

//gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int, allUnits map[int]*unit.ShortUnitInfo
func FindExtremum(mp *_map.Map, start, end *coordinate.Coordinate, unitX, unitY int, gameUnit *unit.Unit) (int, int, error) {

	println(
		" xStart: " + strconv.Itoa(start.X) + " yStart: " + strconv.Itoa(start.Y) +
			" xEnd: " + strconv.Itoa(end.X) + " yEnd: " + strconv.Itoa(end.Y))

	// todo крайне дорогой метод из за поиска А* который тут особо и не нужен
	err, path := find_path.FindPath(mp, start, end, gameUnit, 30, nil)
	if err != nil {
		return 0, 0, errors.New(err.Error() +
			" xStart: " + strconv.Itoa(start.X) + " yStart: " + strconv.Itoa(start.Y) +
			" xEnd: " + strconv.Itoa(end.X) + " yEnd: " + strconv.Itoa(end.Y))
	}

	x, y := 0, 0

	for _, point := range path {
		_, _, _, _, collision, _ := BetweenLine(float64(unitX), float64(unitY), float64(point.X), float64(point.Y), mp, gameUnit.Body, false)
		if !collision {
			x = point.X
			y = point.Y
		}
	}

	return x, y, nil
}
