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
func DetectObstacles(entryPoints, outPoints, collisions []*coordinate.Coordinate, gameUnit *unit.Unit, size int, user *player.Player, uuid string, mp *_map.Map) []*Obstacle {
	// пройти по контуру препятсвия start[0] пометить все точки которые встретялся на пути и находятся в масивах start, end
	// пройти по контору препятсвия со следующей не помеченой точки из масива start

	ClearVisiblePath(mp.Id, user)
	obstacles := make([]*Obstacle, 0)
	for i := 0; i < len(entryPoints); i++ {
		if !entryPoints[0].Find {
			points := GetObstaclePoints(entryPoints[0], collisions[0], gameUnit, size, user, uuid, mp)
			obstacle := DetectObstacle(entryPoints, outPoints, collisions, points, size)
			obstacles = append(obstacles, obstacle)
		}
	}

	return obstacles
}

func GetObstaclePoints(start, collisions *coordinate.Coordinate, gameUnit *unit.Unit, size int, user *player.Player, uuid string, mp *_map.Map) []*coordinate.Coordinate {
	points := make([]*coordinate.Coordinate, 0)
	handStop := false
	exit := false
	noPath := false

	x1, y1, angleStart := GetStartBugOptions(start.X, start.Y, collisions.X, collisions.Y, mp, gameUnit, size, user)
	//пускаем 1 руку по контуру всего обьекта без конечной точки что бы собрать все координаты контура
	go Hand(1, x1, y1, &handStop, &exit, &noPath, &points, angleStart, size, mp, user, gameUnit, uuid, 0, 0)
	for !handStop {
		// ожидаем пока рука пройдет весь путь
		time.Sleep(time.Millisecond)
	}

	return points
}

func DetectObstacle(entryPoints, outPoints, collisions, obstaclePoints []*coordinate.Coordinate, size int) *Obstacle {

	obstacle := &Obstacle{}
	firstEntryID := len(entryPoints) // мы ищем самую первую точку входа
	lastOutID := 0                   // мы ищем самую последнюю точку выхода

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
	// TODO по идеи тут не должно возникнуть ошибки но хз)
	obstacle.Entry = entryPoints[firstEntryID]
	obstacle.EntryCollision = collisions[firstEntryID]
	obstacle.Out = outPoints[lastOutID]

	return obstacle
}

type Obstacle struct {
	Entry          *coordinate.Coordinate
	EntryCollision *coordinate.Coordinate
	Out            *coordinate.Coordinate
}
