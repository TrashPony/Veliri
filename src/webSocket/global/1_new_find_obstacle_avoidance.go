package global

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
)

//gameMap *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, scaleMap int, allUnits map[int]*unit.ShortUnitInfo
func ObstacleAvoidance(mp *_map.Map, obstacle *Obstacle, gameUnit *unit.Unit, size int, user *player.Player, uuid string) ([]*coordinate.Coordinate, error) {

	// start - это валидная координата перед препятсивем (точка входа)
	// end - это валидная координата после препяствия (точка выхода)
	// unitX, unitY - текущее положение юнита
	// gameUnit - сам юнит, от сюда мы берем тушу для колизии
	// size - размер дискретного квадрата поля

	// 1. обходим препятвие с 2х сторон, "держась за него правой и левой рукой"
	// 	1.1 если мы фиксимуем что препятвие на клетку "провалилось" ближе к цели и дальше от юнита кладем координату в "возможный путь"
	// 	1.2 замеряем дальность каждой координаты до цели (лучший кандидат если конец в препятсиве)
	//  1.3 когда 1 рука дошла до точки выхода из препятвия - это который путь. уничтожаем вторую руку. Выходим из цикла

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

	// TODO иногда поиск идет не на то препятсвие
	// уобьекта obstacle уже есть конту ненадо его опять искать жуком, просто обхгодим то что уже есть
	ClearVisiblePath(mp.Id, user)
	CreateRect("white", int(obstacle.Entry.X), int(obstacle.Entry.Y), size, mp.Id, user)
	CreateRect("red", int(obstacle.EntryCollision.X), int(obstacle.EntryCollision.Y), size, mp.Id, user)

	// если обьект не полный значит у него заполнилась только 1 сторона
	if obstacle.NoFull {
		return obstacle.Contour, nil
	}

	pointsOne := make([]*coordinate.Coordinate, 0)
	pointsTwo := make([]*coordinate.Coordinate, 0)

	// смотрим в какую сторону надо пройти путь меньше
	for i := len(obstacle.Contour) - 1; i >= 0; i-- {
		distToEnd := game_math.GetBetweenDist(obstacle.Contour[i].X, obstacle.Contour[i].Y, obstacle.Out.X, obstacle.Out.Y)
		if int(distToEnd) < size+10 {
			break
		}
		pointsOne = append(pointsOne, obstacle.Contour[i])
	}

	for i := 0; i < len(obstacle.Contour); i++ {
		distToEnd := game_math.GetBetweenDist(obstacle.Contour[i].X, obstacle.Contour[i].Y, obstacle.Out.X, obstacle.Out.Y)
		if int(distToEnd) < size+10 {
			break
		}
		pointsTwo = append(pointsTwo, obstacle.Contour[i])
	}

	if len(pointsOne) < len(pointsTwo) {
		return pointsOne, nil
	} else {
		return pointsTwo, nil
	}
}

func GetStartBugOptions(xStart, yStart, xCollision, yCollision int, mp *_map.Map, gameUnit *unit.Unit, size int, user *player.Player) (int, int, int, error) {
	// TODO весь метод не правильный
	// к сожалению по неведымим причинам иногда координата xStart, yStart бывает недоступной, из за чего алгоритм не работает
	// поэтому оступаем назад пока не найдем нужную координату по вектору и обязательно проверяем что бы спереди было препятвие
	// после нахождения точки надо надо обязательно найти направление к точке входа в препятвиея, иначе алгоритм опять же не запустится
	// порой точка входа находится по диагонали поэтому якорь найти не удается и приходится изменять дискретность
	angleStart := game_math.GetBetweenAngle(float64(xStart), float64(yStart), float64(xCollision), float64(yCollision))
	radian := float64(angleStart) * math.Pi / 180

	count := 0
	for {
		count++
		if count > 4 {
			return 0, 0, 0, errors.New("no hook")
		}

		ok, _ := checkRect(xStart, yStart, gameUnit.Body, mp)
		if !ok {
			CreateRect("white", int(xStart), int(yStart), size, mp.Id, user)
			x3, y3 := float64(size)*math.Cos(radian), float64(size)*math.Sin(radian)

			xStart += int(math.Round(x3))
			yStart += int(math.Round(y3))

		} else {

			findHook := false
			// находим угол входа!
			for angleStart = 0; angleStart < 360; angleStart += 90 {

				radian := float64(angleStart) * math.Pi / 180
				x3, y3 := float64(size)*math.Cos(radian), float64(size)*math.Sin(radian)

				ok, _ := checkRect(xStart+int(math.Round(x3)), yStart+int(math.Round(y3)), gameUnit.Body, mp)
				if !ok {
					findHook = true
					break
				}
			}

			if !findHook {
				// алгоритм не смог найти точку входа
				// увеличить дискретность и попробовать снова
				return 0, 0, 0, errors.New("no hook")
				// todo надо исправить но я не ибу как
			}

			break
		}
	}

	return xStart, yStart, angleStart, nil
}

func SearchPoint(points *[]*coordinate.Coordinate, unitX, unitY, size int, mp *_map.Map, user *player.Player, gameUnit *unit.Unit) (int, int) {
	// 2. смотрим каждую координату в масиве "возможный путь" (от дальней)
	// 	2.1 если между юнитом и координатой нет препятсвий то это истиный путь, запоминаем координату
	// идем по масиву с конца что бы найти самую дальную валидну точку
	// TODO самый дорогой метод в алгоритме
	for i := len(*points) - 1; i >= 0; i-- {
		collision := SearchCollisionInLine(float64((*points)[i].X), float64((*points)[i].Y), float64(unitX), float64(unitY), mp, gameUnit.Body)
		if !collision {
			CreateRect("green", (*points)[i].X, (*points)[i].Y, size, mp.Id, user)
			return (*points)[i].X, (*points)[i].Y
		}
	}
	return 0, 0
}

func Hand(side float64, x, y int, stopFlag *bool, exitFlag, noPath, noMap *bool, points *[]*coordinate.Coordinate, angleStart,
	size int, mp *_map.Map, user *player.Player, gameUnit *unit.Unit, uuid string, xEnd, yEnd int) {

	// угол в радианах
	angle := float64(angleStart) * math.Pi / 180
	// сохраняет стартовые точки, если рука вернулась к ней и не нашла точку выхода то пути не существует
	startX, startY := x, y

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

		distToStart := game_math.GetBetweenDist(x, y, startX, startY)
		if int(distToStart) < size+10 && step > 10 {
			// мы вернулись на старт и не встрели выхода
			*noPath = true
			*stopFlag = true
			return
		}

		ok, coorNoMap := checkRect(x, y, gameUnit.Body, mp)
		if coorNoMap {
			*stopFlag = true
			*noMap = true
			return
		}

		if ok {
			// поворачиваем направо
			angle += 1.5708 * side // +90 градусов
			CreateRect("white", int(x), int(y), size, mp.Id, user)
			x3, y3 := float64(size)*math.Cos(angle), float64(size)*math.Sin(angle)
			x += int(math.Round(x3))
			y += int(math.Round(y3))

			// TODO помечать координаты которые изменяют свое положение относительно оси,
			//  а то сканировать все поинты дорого
			*points = append(*points, &coordinate.Coordinate{X: x, Y: y})

			distToEnd := game_math.GetBetweenDist(x, y, xEnd, yEnd)
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

func checkRect(x, y int, body *detail.Body, mp *_map.Map) (bool, bool) {

	q, r := game_math.GetQRfromXY(x, y)
	_, find := mp.OneLayerMap[q][r]
	if !find {
		return false, true
	}

	possibleMove, _, _, _ := collisions.CheckCollisionsOnStaticMap(x, y, 0, mp, body, false, true)
	return possibleMove, false
}
