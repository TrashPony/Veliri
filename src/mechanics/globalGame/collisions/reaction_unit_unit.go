package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

func InitCheckCollision(moveUnit *unit.Unit, pathUnit *unit.PathUnit) (bool, *unit.ShortUnitInfo, int, int, int) {
	// вынесено в отдельную функцию что бы можно было беспробленнмно сделать defer rLock.Unlock()
	units := globalGame.Clients.GetAllShortUnits(moveUnit.MapID, true)

	noCollision, collisionUnit := CheckCollisionsPlayers(moveUnit, pathUnit.X, pathUnit.Y, pathUnit.Rotate, units,
		false, false, false, false, true, nil)

	x, y := moveUnit.X, moveUnit.Y
	percent := 0
	//if !noCollision {
	//	x, y, noCollision, percent = detailCheckCollision(moveUnit, pathUnit, units)
	//}
	return noCollision, collisionUnit, x, y, percent
}

func detailCheckCollision(moveUnit *unit.Unit, pathUnit *unit.PathUnit, units map[int]*unit.ShortUnitInfo) (int, int, bool, int) {

	// юнит имеет высокую скорость последние точки делить его путь что бы адекватно обработать колизии
	x, y := float64(moveUnit.X), float64(moveUnit.Y)

	// перменная для контроля зависаний, если дальность начала возрастать значит алгоритм проебал точку выхода
	startDist := game_math.GetBetweenDist(int(x), int(y), pathUnit.X, pathUnit.Y)
	minDist := startDist

	for {

		dist := game_math.GetBetweenDist(int(x), int(y), pathUnit.X, pathUnit.Y)
		if dist <= 1 || minDist < dist {
			return int(x), int(y), true, 100
		}

		radRotate := float64(pathUnit.Rotate) * math.Pi / 180
		stopX := float64(1) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(1) * math.Sin(radRotate)

		noCollision, _ := CheckCollisionsPlayers(moveUnit, int(x), int(y), pathUnit.Rotate, units,
			false, false, false, false, false, nil)
		if !noCollision {
			return int(x), int(y), false, int((dist * 100) / startDist)
		}

		x += stopX
		y += stopY
		minDist = dist
	}
}

func UnitToUnitCollisionReaction(takeUnit, toUnit *unit.Unit) (*unit.PathUnit, *unit.PathUnit) {
	// задаем переменные массы шаров
	mass1 := takeUnit.Body.CapacitySize
	mass2 := toUnit.Body.CapacitySize

	if takeUnit.CurrentSpeed < 2 {
		takeUnit.CurrentSpeed = 2
	}

	// задаем переменные скорости
	// расчет для первой машины
	radRotate1 := float64(takeUnit.Rotate) * math.Pi / 180
	xVel1 := float64(takeUnit.CurrentSpeed) * math.Cos(radRotate1) // идем по вектору движения корпуса
	yVel1 := float64(takeUnit.CurrentSpeed) * math.Sin(radRotate1)

	// расчет для второй машины
	radRotate2 := float64(toUnit.Rotate) * math.Pi / 180
	xVel2 := float64(toUnit.CurrentSpeed) * math.Cos(radRotate2) // идем по вектору движения корпуса
	yVel2 := float64(toUnit.CurrentSpeed) * math.Sin(radRotate2)

	//Угол между осью х и линией действия
	needRad := math.Atan2(float64(toUnit.Y-takeUnit.Y), float64(toUnit.X-takeUnit.X))
	cosAlfa := math.Cos(needRad)
	sinAlfa := math.Sin(needRad)

	// находим скорости вдоль линии действия
	xVel1prime := xVel1*cosAlfa + yVel1*sinAlfa
	xVel2prime := xVel2*cosAlfa + yVel2*sinAlfa

	// находим скорости перпендикулярные линии действия
	yVel1prime := yVel1*cosAlfa - xVel1*sinAlfa
	yVel2prime := yVel2*cosAlfa - xVel2*sinAlfa

	//// применяем законы сохранения
	P := float64(mass1)*xVel1prime + float64(mass2)*xVel2prime
	V := xVel1prime - xVel2prime
	v2f := (P + float64(mass1)*V) / (float64(mass1) + float64(mass2))
	v1f := v2f - xVel1prime + xVel2prime
	xVel1prime = v1f
	xVel2prime = v2f

	// Проецируем обратно на оси Х и У.
	xVel1 = xVel1prime*cosAlfa - yVel1prime*sinAlfa
	yVel1 = yVel1prime*cosAlfa + xVel1prime*sinAlfa

	xVel2 = xVel2prime*cosAlfa - yVel2prime*sinAlfa
	yVel2 = yVel2prime*cosAlfa + xVel2prime*sinAlfa

	speed1 := (math.Sqrt((xVel1 * xVel1) + (yVel1 * yVel1))) / 5
	speed2 := (math.Sqrt((xVel2 * xVel2) + (yVel2 * yVel2))) / 5

	takeUnit.CurrentSpeed = speed1
	takeUnit.X += int(float64(-speed1) * math.Cos(needRad))
	takeUnit.Y += int(float64(-speed1) * math.Sin(needRad))

	// проверка нового места толкаемого юзера на колизию в статичной карте
	mp, _ := maps.Maps.GetByID(takeUnit.MapID)

	possibleMove, _ := BodyCheckCollisionsOnStaticMap(
		int(toUnit.X+int(float64(speed2)*math.Cos(needRad))),
		int(toUnit.Y+int(float64(speed2)*math.Sin(needRad))),
		toUnit.Rotate,
		mp,
		toUnit.Body,
		false,
		false,
	)

	// проверка нового места толкаемого юзера на колизию с другими юзерами // TODO не отдебажено
	noCollision, _, _, _, _ := InitCheckCollision(toUnit, &unit.PathUnit{
		X:      int(toUnit.X + int(float64(speed2)*math.Cos(needRad))),
		Y:      int(toUnit.Y + int(float64(speed2)*math.Sin(needRad))),
		Rotate: toUnit.Rotate,
	})

	if possibleMove && noCollision {
		toUnit.X += int(float64(speed2) * math.Cos(needRad))
		toUnit.Y += int(float64(speed2) * math.Sin(needRad))
	} else {
		// оталкиваем игрока инициализирующего столкновение иначе они застрянут
		takeUnit.X += int(float64(-speed2) * math.Cos(needRad))
		takeUnit.Y += int(float64(-speed2) * math.Sin(needRad))
	}

	userPath := unit.PathUnit{
		X:           takeUnit.X,
		Y:           takeUnit.Y,
		Rotate:      takeUnit.Rotate,
		Millisecond: 200,
		Speed:       takeUnit.CurrentSpeed,
	}

	toUserPath := unit.PathUnit{
		X:           toUnit.X,
		Y:           toUnit.Y,
		Rotate:      toUnit.Rotate,
		Millisecond: 200,
		Speed:       speed2,
	}

	return &userPath, &toUserPath
}
