package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

// функция принимает набор точек входа и выхода, и должна определить
// 1 принадлежатли эти точки одному или множеству препятсвий
// 2 если множеству то пометить какая точка какому препятсвию
// 3 вернуть обьекты препятсвий
func DetectObstacles(entryPoints, outPoints, collisions []*coordinate.Coordinate, gameUnit *unit.Unit, size int, user *player.Player, uuid string, mp *_map.Map) ([]*Obstacle, error) {
	// пройти по контуру препятсвия start[0] пометить все точки которые встретялся на пути и находятся в масивах start, end
	// todo пройти по контору препятсвия со следующей не помеченой точки из масива start

	ClearVisiblePath(mp.Id, user)

	obstacles := make([]*Obstacle, 0)
	for i := 0; i < len(entryPoints); i++ {
		if !entryPoints[0].Find {

			points, noPath, err := GetObstaclePoints(entryPoints[0], collisions[0], outPoints[0], gameUnit, size, user, uuid, mp)
			if err != nil {
				return nil, err
			}

			obstacle := DetectObstacle(entryPoints, outPoints, collisions, points, size, noPath)
			obstacles = append(obstacles, obstacle)
		}
	}

	return obstacles, nil
}

func GetObstaclePoints(start, collisions, out *coordinate.Coordinate, gameUnit *unit.Unit, size int, user *player.Player, uuid string, mp *_map.Map) ([]*coordinate.Coordinate, bool, error) {

	x1, y1, angleStart, err := GetStartBugOptions(start.X, start.Y, collisions.X, collisions.Y, mp, gameUnit, size, user)
	if err != nil {
		return nil, false, err
	}

	oneHandPoints := make([]*coordinate.Coordinate, 0)
	twoHandPoints := make([]*coordinate.Coordinate, 0)
	oneHandStop := false
	twoHandStop := false
	exit := false
	noPath := false
	noMap := false

	// TODO неправильно работают пути к координатам которые за картой
	go Hand(-1, x1, y1, &oneHandStop, &exit, &noPath, &noMap, &oneHandPoints, angleStart, size, mp, user, gameUnit, uuid, 0, 0)
	go Hand(1, x1, y1, &twoHandStop, &exit, &noPath, &noMap, &twoHandPoints, angleStart, size, mp, user, gameUnit, uuid, 0, 0)

	for !oneHandStop || !twoHandStop {
		time.Sleep(time.Millisecond)
	}
	exit = true

	// находим наиболее полный контур, неполный контур возможен если например обьект уходит за пределы карты
	if len(oneHandPoints) > len(twoHandPoints) {
		return oneHandPoints, noMap, nil
	} else {
		return twoHandPoints, noMap, nil
	}
}

func DetectObstacle(entryPoints, outPoints, collisions, points []*coordinate.Coordinate, size int, noPath bool) *Obstacle {
	obstacle := &Obstacle{}

	// мы ищем самую последнюю точку выхода
	searchPoints := func(obstaclePoints []*coordinate.Coordinate) {

		firstEntryID := len(entryPoints) // мы ищем самую первую точку входа
		lastOutID := 0

		for _, point := range obstaclePoints {
			for i, startPoint := range entryPoints {
				dist := game_math.GetBetweenDist(point.X, point.Y, startPoint.X, startPoint.Y)
				println(int(dist))
				if int(dist) < 100 { // todo нихуя не 100
					startPoint.Find = true
					if i < firstEntryID {
						firstEntryID = i
					}
				}
			}

			for i, endPoint := range outPoints {
				dist := game_math.GetBetweenDist(point.X, point.Y, endPoint.X, endPoint.Y)
				if int(dist) < 100 { // todo нихуя не 100
					endPoint.Find = true
					if i > lastOutID {
						lastOutID = i
					}
				}
			}
		}

		// на каждую точку входа всегда есть точка колизии
		// TODO должно возникнуть ошибки надо обработать с умом
		obstacle.Entry = entryPoints[firstEntryID]
		obstacle.EntryCollision = collisions[firstEntryID]
		obstacle.Out = outPoints[lastOutID]
	}

	if len(points) > 0 {
		searchPoints(points)
	}

	obstacle.Contour = points
	obstacle.NoFull = noPath

	return obstacle
}

type Obstacle struct {
	Entry          *coordinate.Coordinate
	EntryCollision *coordinate.Coordinate
	Out            *coordinate.Coordinate
	Contour        []*coordinate.Coordinate
	NoFull         bool
}
