package global

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"github.com/satori/go.uuid"
	"math"
	"time"
)

func Unit(moveUnit *unit.Unit, ToX, ToY float64, user *player.Player) ([]*unit.PathUnit, error) {

	moveUUID := uuid.NewV1().String()
	moveUnit.MoveUUID = moveUUID
	moveUnit.ToX = ToX
	moveUnit.ToY = ToY

	startX := float64(moveUnit.X)
	startY := float64(moveUnit.Y)

	maxSpeed := float64(moveUnit.Speed)
	if moveUnit.Body.MotherShip {
		efficiency := move.WorkOutThorium(moveUnit.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
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
	if moveUUID == moveUnit.MoveUUID {
		return LeftHandAlgorithm(moveUnit, startX, startY, ToX, ToY, maxSpeed, user, moveUUID)
	} else {
		return nil, errors.New("wrong uuid")
	}
}

func To2(forecastX, forecastY, speed, ToX, ToY float64, rotate, rotateAngle, ms int) (error, []*unit.PathUnit) {

	// TODO искать предварительно часть пути до тех пока не будет разница в углу 0
	// TODO в разных вариантах, разворот на скорости или развород на месте, выбирать то где мешье частей пути (выше скорость)
	//  или если 1 из способов не пройти
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
		move.RotateUnit(&rotate, &needRotate, rotateAngle)

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

func CreateRect(color string, startX, startY int, rectSize, mapID int, user *player.Player) {
	go SendMessage(Message{Event: "CreateRect", Color: color, RectSize: rectSize,
		X: int(startX), Y: int(startY), IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})

	time.Sleep(20 * time.Millisecond)
}

func ClearVisiblePath(mapID int, user *player.Player) {
	go SendMessage(Message{Event: "ClearPath", IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})
}

func CreateLine(color string, X, Y, ToX, ToY int, rectSize, mapID int, user *player.Player) {
	go SendMessage(Message{Event: "CreateLine", Color: color, RectSize: rectSize,
		X: int(X), Y: int(Y), ToX: float64(ToX), ToY: float64(ToY), IDUserSend: user.GetID(), IDMap: mapID, Bot: user.Bot})

	time.Sleep(20 * time.Millisecond)
}

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY, maxSpeed float64, user *player.Player, uuid string) ([]*unit.PathUnit, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)
	rotate := 180

	rectSize := 25

	path := make([]*unit.PathUnit, 0)
	appendPath := func(appendPath []*unit.PathUnit) {
		for i := 0; i < len(appendPath); i++ {
			path = append(path, appendPath[i])
		}
	}
	//
	CreateRect("green", int(startX), int(startY), rectSize, moveUnit.MapID, user)
	CreateLine("white", int(startX), int(startY), int(ToX), int(ToY), rectSize, moveUnit.MapID, user)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	xCollision, yCollision, xEnd, yEnd, xObstacle, yObstacle, collision, endIsObstacle := BetweenLine(startX, startY, ToX, ToY, mp, moveUnit.Body, true)

	// 0.1 если конечная точка находится в первом препятсвие то смотрим куда ближе идти ко входу или к выходу
	if endIsObstacle {
		EndIsObstacle(&ToX, &ToY, &xCollision, &yCollision, &xEnd, &yEnd, &collision, moveUnit)
	}

	if !collision {
		// 0.2 на прямой между стартом и концом нет прептявий
		CreateRect("green", int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)
		CreateLine("green", int(startX), int(startY), int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)

		err, path := To2(startX, startY, maxSpeed, moveUnit.ToX, moveUnit.ToY, moveUnit.Rotate, rotate, 200)
		return path, err
	} else {
		// 0.3 на прямой были найдены препятвия

		//  todo кеширование пройденных клеток, т.к. это может быть 1 очень кривое препятвие и придется обходить 1 и теже точки

		CreateRect("red", xCollision, yCollision, rectSize, moveUnit.MapID, user)
		CreateRect("red", xEnd, yEnd, rectSize, moveUnit.MapID, user)
		CreateLine("red", int(xCollision), int(yCollision), int(xEnd), int(yEnd), rectSize, moveUnit.MapID, user)

		x, y, err := ObstacleAvoidance(mp, &coordinate.Coordinate{X: xCollision, Y: yCollision}, &coordinate.Coordinate{X: xEnd, Y: yEnd},
			xObstacle, yObstacle, moveUnit.X, moveUnit.Y, moveUnit, rectSize, user, uuid)

		if err != nil {
			return nil, err
		}

		_, path = To2(startX, startY, maxSpeed, float64(x), float64(y), moveUnit.Rotate, rotate, 200)

		CreateRect("green", x, y, rectSize, moveUnit.MapID, user)
		CreateLine("green", int(startX), int(startY), int(x), int(y), rectSize, moveUnit.MapID, user)

		for moveUnit.MoveUUID == uuid {

			xCollision, yCollision, xEnd, yEnd, xObstacle, yObstacle, collision, endIsObstacle = BetweenLine(float64(x), float64(y), ToX, ToY, mp, moveUnit.Body, false)
			if endIsObstacle {
				EndIsObstacle(&ToX, &ToY, &xCollision, &yCollision, &xEnd, &yEnd, &collision, moveUnit)
			}

			if !collision {
				//  2.1.1 если между координатой истиного пути и целью нет препятсвий формируем путь. Выходим из функции.
				if len(path) != 0 {
					_, aPath := To2(float64(path[len(path)-1].X), float64(path[len(path)-1].Y), maxSpeed, moveUnit.ToX, moveUnit.ToY, path[len(path)-1].Rotate, rotate, 200)
					appendPath(aPath)
					CreateLine("green", int(path[len(path)-1].X), int(path[len(path)-1].Y), int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)
				} else {
					_, path = To2(float64(startX), float64(startY), maxSpeed, moveUnit.ToX, moveUnit.ToY, moveUnit.Rotate, rotate, 200)
				}

				CreateRect("green", int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)

				break
			} else {
				//  2.1.2 если между координатой истиного пути и целью есть препятсивия запоминаем координату, переходим к пункту 1 и формируем новую

				CreateRect("red", xCollision, yCollision, rectSize, moveUnit.MapID, user)
				CreateRect("red", xEnd, yEnd, rectSize, moveUnit.MapID, user)
				CreateLine("red", int(xCollision), int(yCollision), int(xEnd), int(yEnd), rectSize, moveUnit.MapID, user)

				x, y, err = ObstacleAvoidance(mp, &coordinate.Coordinate{X: xCollision, Y: yCollision}, &coordinate.Coordinate{X: xEnd, Y: yEnd},
					xObstacle, yObstacle, x, y, moveUnit, rectSize, user, uuid)
				if err != nil {
					return nil, err
				}

				if len(path) != 0 {
					_, aPath := To2(float64(path[len(path)-1].X), float64(path[len(path)-1].Y), maxSpeed, float64(x), float64(y), path[len(path)-1].Rotate, rotate, 200)
					appendPath(aPath)
					CreateLine("green", int(path[len(path)-1].X), int(path[len(path)-1].Y), int(x), int(y), rectSize, moveUnit.MapID, user)
				} else {
					_, path = To2(float64(startX), float64(startY), maxSpeed, float64(x), float64(y), moveUnit.Rotate, rotate, 200)
				}

				CreateRect("green", int(x), int(y), rectSize, moveUnit.MapID, user)

			}
		}

		return path, nil
	}
}

func EndIsObstacle(ToX, ToY *float64, xCollision, yCollision, xEnd, yEnd *int, collision *bool, moveUnit *unit.Unit) {
	collisionStartDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *xCollision, *yCollision)
	collisionEndDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *xEnd, *yEnd)

	// если то старта колизии ближе чем до конца то считаем что маршрут без колизий
	if collisionStartDist < collisionEndDist {
		moveUnit.ToX, moveUnit.ToY = float64(*xCollision), float64(*yCollision)
		*collision = false
	} else {
		// иначе переназначаем конечный пункт что бы не искать путь вечно
		*ToX, *ToY = float64(*xCollision), float64(*yCollision)
	}
}

// возвращает xКолизии, yКолизии, хВыхода из колизии, yВыхода из колизии прошли без столкновений
func BetweenLine(startX, startY, ToX, ToY float64, mp *_map.Map, body *detail.Body, startMove bool) (
	xCollision, yCollision, xEnd, yEnd, xObstacle, yObstacle int, collision, endIsObstacle bool) {

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
				return 0, 0, 0, 0, 0, 0, false, false
			} else {
				endIsObstacle = true
			}
		}

		stopX, stopY := float64(speed)*math.Cos(radian), float64(speed)*math.Sin(radian)

		possibleMove, _, _, _ := collisions.CheckCollisionsOnStaticMap(int(currentX), int(currentY), angle, mp, body, true)
		if !possibleMove {
			// если юнит по каким то причинам стартует из колизии то дать ему выйти и потом уже искать колизию
			if !(distToStart < speed+2 && startMove) {
				// фиксируем 1 точку колизии
				if !collision {
					xCollision = int(currentX)
					yCollision = int(currentY)
					xObstacle = int(currentX + stopX)
					yObstacle = int(currentY + stopY)
					collision = true
				}
			}
		} else {
			// фиксируем точку выхода из колизии
			if collision {
				// TODO однако если препятсвий много или препятвие С обратное то смотрим путь дальше и отдаем самую последнюю точку препятсвия
				//xEnd = int(currentX)
				//yEnd = int(currentY)
				//collision = false
				return xCollision, yCollision, int(currentX), int(currentY), xObstacle, yObstacle, collision, endIsObstacle
			}
		}

		currentX += stopX
		currentY += stopY
	}
}

//gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int, allUnits map[int]*unit.ShortUnitInfo
func ObstacleAvoidance(mp *_map.Map, start, end *coordinate.Coordinate, xObstacle, yObstacle, unitX, unitY int, gameUnit *unit.Unit, size int, user *player.Player, uuid string) (int, int, error) {

	// start - это валидная координата перед препятсивем (точка входа)
	// end - это валидная координата после препяствия (точка выхода)
	// unitX, unitY - текущее положение юнита
	// gameUnit - сам юнит, от сюда мы берем тушу для колизии
	// size - размер дискретного квадрата поля

	// 1. обходим препятвие с 2х сторон, "держась за него правой и левой рукой"
	// 	1.1 если мы фиксимуем что препятвие на клетку "провалилось" ближе к цели и дальше от юнита кладем координату в "возможный путь"
	// 	1.2 замеряем дальность каждой координаты до цели (лучший кандидат если конец в препятсиве)
	//  1.3 когда 1 рука дошла до точки выхода из препятвия - это который путь. уничтожаем вторую руку. Выходим из цикла
	// 2. смотрим каждую координату в масиве "возможный путь" (от дальней)
	// 	2.1 если между юнитом и координатой нет препятсвий то это истиный путь, запоминаем координату

	//CreateRect("white", int(start.X), int(start.Y), size, mp.Id, user)
	//CreateRect("red", int(xObstacle), int(yObstacle), size, mp.Id, user)

	/*
			находим координату припятсвий, их может быть много но мы берем ток первую
			_________________________________
			||		   ||		||         ||
			|| x-1 y+1 || x y+1 || x+1 y+1 ||
			||		   ||		||         ||
			---------------------------------
			||	       ||		||         ||
			|| x-1 y   || x y   || x+1 y   ||
			||		   ||		||         ||
			---------------------------------
		    ||		   ||		||         ||
			|| x-1 y-1 || x y-1 || x+1 y-1 ||
			||		   ||		||         ||
			---------------------------------

		Жук начинает движение с белой области по направлению к черной, Как только он попадает на черный элемент,
		он поворачивает налево и переходит к следующему элементу. Если этот элемент белый, то жук поворачивается направо,
		иначе - налево. Процедура повторяется до тех пор, пока жук не вернется в исходную точку.
	*/
	ClearVisiblePath(mp.Id, user)

	x1, y1 := start.X, start.Y
	xO, yO := xObstacle, yObstacle
	xEnd, yEnd := end.X, end.Y

	// к сожалению по неведымим причинам иногда координата x1, y1 бывает недоступной, из за чего алгоритм не работает
	// поэтому оступаем назад пока не найдем нужную координату по вектору и обязательно проверяем что бы спереди было препятвие
	// после нахождения точки надо надо обязательно найти направление к точке входа в препятвиея, иначе алгоритм опять же не запустится
	// порой точка входа находится по диагонали поэтому якорь найти не удается и приходится изменять дискретность
	angleStart := game_math.GetBetweenAngle(float64(x1), float64(y1), float64(xO), float64(yO))
	radian := float64(angleStart) * math.Pi / 180

	for {
		if !checkRect(x1, y1, gameUnit.Body, mp) {
			CreateRect("white", int(x1), int(y1), size, mp.Id, user)
			x3, y3 := float64(size)*math.Cos(radian), float64(size)*math.Sin(radian)

			x1 += int(math.Round(x3))
			y1 += int(math.Round(y3))
		} else {

			findHook := false
			// находим угол входа!
			for angleStart = 0; angleStart < 360; angleStart += 90 {

				radian := float64(angleStart) * math.Pi / 180
				x3, y3 := float64(size)*math.Cos(radian), float64(size)*math.Sin(radian)
				if !checkRect(x1+int(math.Round(x3)), y1+int(math.Round(y3)), gameUnit.Body, mp) {
					findHook = true
					break
				}
			}

			if !findHook {
				// алгоритм не смог найти точку входа
				// увеличить дискретность и попробовать снова
				println("no hook")
				extX, extY := 0, 0
				for extX == 0 && extY == 0 && uuid == gameUnit.MoveUUID {
					extX, extY, _ = ObstacleAvoidance(mp, start, end, xObstacle, yObstacle, unitX, unitY, gameUnit, size+5, user, uuid)
				}
				return extX, extY, nil
			}

			break
		}
	}

	CreateRect("red", int(xO), int(yO), size, mp.Id, user)

	hand := func(side float64, x, y int, stopFlag *bool, exitFlag *bool, points *[]*coordinate.Coordinate) {

		// угол в радианах
		angle := float64(angleStart) * math.Pi / 180

		// первый вход в препятвие нельзя убирать в общий фор т.к. тут работает найденное для якоря направление
		// а в форе оно сразу будет изменено
		CreateRect("white", int(x), int(y), size, mp.Id, user)
		x3, y3 := float64(size)*math.Cos(angle), float64(size)*math.Sin(angle)
		x += int(math.Round(x3))
		y += int(math.Round(y3))

		for step := 0; true; step++ {

			if uuid != gameUnit.MoveUUID {
				return
			}

			if *exitFlag {
				// если какаято рука звершила поиск то вторая уничтожается
				return
			}

			if checkRect(x, y, gameUnit.Body, mp) {
				// поворачиваем направо
				angle += 1.5708 * side // +90 градусов
				CreateRect("white", int(x), int(y), size, mp.Id, user)
				x3, y3 := float64(size)*math.Cos(angle), float64(size)*math.Sin(angle)
				x += int(math.Round(x3))
				y += int(math.Round(y3))

				distToEnd := game_math.GetBetweenDist(x, y, xEnd, yEnd)

				// TODO класть сюда только те координаты которые изменяют свое положение относительно оси,
				//  а то сканировать все поинты дорого
				*points = append(*points, &coordinate.Coordinate{X: x, Y: y})

				if int(distToEnd) < size+10 {
					// мы дошли до точки выхода
					*stopFlag = true
					return
				}
			} else {
				// поворачиваем направо
				angle -= 1.5708 * side // -90 градусов
				CreateRect("white", int(x), int(y), size, mp.Id, user)
				x3, y3 := float64(size)*math.Cos(angle), float64(size)*math.Sin(angle)
				x += int(math.Round(x3))
				y += int(math.Round(y3))
			}
		}
	}

	oneHandPoints := make([]*coordinate.Coordinate, 0)
	twoHandPoints := make([]*coordinate.Coordinate, 0)

	oneHandStop := false
	twoHandStop := false
	exit := false

	go hand(-1, x1, y1, &oneHandStop, &exit, &oneHandPoints)
	go hand(1, x1, y1, &twoHandStop, &exit, &twoHandPoints)

	for !oneHandStop && !twoHandStop {
		time.Sleep(time.Millisecond)
	}
	exit = true

	searchPoint := func(points *[]*coordinate.Coordinate) (int, int) {
		// идем по масиву с конца что бы найти самую дальную валидну точку
		for i := len(*points) - 1; i >= 0; i-- {
			_, _, _, _, _, _, collision, _ := BetweenLine(float64((*points)[i].X), float64((*points)[i].Y), float64(unitX), float64(unitY), mp, gameUnit.Body, false)
			if !collision {
				CreateRect("green", (*points)[i].X, (*points)[i].Y, size, mp.Id, user)
				return (*points)[i].X, (*points)[i].Y
			}
		}
		return 0, 0
	}

	extX, extY := 0, 0
	if oneHandStop {
		extX, extY = searchPoint(&oneHandPoints)
	} else {
		extX, extY = searchPoint(&twoHandPoints)
	}

	for extX == 0 && extY == 0 && uuid == gameUnit.MoveUUID {
		extX, extY, _ = ObstacleAvoidance(mp, start, end, xObstacle, yObstacle, unitX, unitY, gameUnit, size-5, user, uuid)
	}

	return extX, extY, nil
}

func checkRect(x, y int, body *detail.Body, mp *_map.Map) bool {
	possibleMove, _, _, _ := collisions.CheckCollisionsOnStaticMap(x, y, 0, mp, body, true)
	return possibleMove
}
