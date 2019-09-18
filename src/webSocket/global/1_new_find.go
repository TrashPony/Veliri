package global

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"github.com/satori/go.uuid"
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

func LeftHandAlgorithm(moveUnit *unit.Unit, startX, startY, ToX, ToY, maxSpeed float64, user *player.Player, uuid string) ([]*unit.PathUnit, error) {

	mp, _ := maps.Maps.GetByID(moveUnit.MapID)
	rotate := 180
	rectSize := 25

	CreateRect("green", int(startX), int(startY), rectSize, moveUnit.MapID, user)
	CreateLine("white", int(startX), int(startY), int(ToX), int(ToY), rectSize, moveUnit.MapID, user)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	collision := SearchCollisionInLine(startX, startY, ToX, ToY, mp, moveUnit.Body)
	if !collision {
		return noCollision(moveUnit, startX, startY, maxSpeed, user, rectSize, rotate)
	} else {
		// 0.3 на прямой были найдены препятвия
		entryPoint, outPoint, collisionPoints, collision, endIsObstacle := BetweenLine(startX, startY, ToX, ToY, mp, moveUnit.Body, true, rectSize)
		DrawPoints(entryPoint, outPoint, collisionPoints, rectSize, mp.Id, user)

		// 0.1 если конечная точка находится в препятсвие то смотрим куда ближе идти ко входу или к выходу
		if endIsObstacle {
			// последние точки это колизия вокруг точки назначения
			lastEntryX, lastEntryY := &entryPoint[len(entryPoint)-1].X, &entryPoint[len(entryPoint)-1].Y
			lastOutX, lastOutY := &outPoint[len(outPoint)-1].X, &outPoint[len(outPoint)-1].Y

			// ищем ближайшую точку которая не в колизии
			EndIsObstacle(&ToX, &ToY, lastEntryX, lastEntryY, lastOutX, lastOutY, &collision, moveUnit, len(entryPoint))
		}

		if !collision {
			return noCollision(moveUnit, startX, startY, maxSpeed, user, rectSize, rotate)
		} else {
			return startFind(moveUnit, int(startX), int(startY), ToX, ToY, maxSpeed, user, uuid, rectSize, rotate, mp)
		}
	}
}

func noCollision(moveUnit *unit.Unit, startX, startY, maxSpeed float64, user *player.Player, size, rotate int) ([]*unit.PathUnit, error) {
	// 0.2 на прямой между стартом и концом нет прептявий
	CreateRect("green", int(moveUnit.ToX), int(moveUnit.ToY), size, moveUnit.MapID, user)
	CreateLine("green", int(startX), int(startY), int(moveUnit.ToX), int(moveUnit.ToY), size, moveUnit.MapID, user)

	err, path := To2(startX, startY, maxSpeed, moveUnit.ToX, moveUnit.ToY, moveUnit.Rotate, rotate, 200)
	return path, err
}

func startFind(moveUnit *unit.Unit, x, y int, ToX, ToY, maxSpeed float64, user *player.Player, uuid string, size, rotate int, mp *_map.Map) ([]*unit.PathUnit, error) {

	// TODO кеширование пройденных клеток, т.к. это может быть 1 очень кривое препятвие и придется обходить 1 и теже точки
	path := make([]*unit.PathUnit, 0)
	startX, startY := x, y

	appendPath := func(appendPath []*unit.PathUnit) {
		for i := 0; i < len(appendPath); i++ {
			path = append(path, appendPath[i])
		}
	}

	createPath := func(toX, toY int) {
		if len(path) != 0 {
			_, aPath := To2(float64(path[len(path)-1].X), float64(path[len(path)-1].Y), maxSpeed, float64(toX), float64(toY), path[len(path)-1].Rotate, rotate, 200)
			appendPath(aPath)
			CreateLine("green", int(path[len(path)-1].X), int(path[len(path)-1].Y), int(toX), int(toY), size, moveUnit.MapID, user)
		} else {
			_, path = To2(float64(startX), float64(startY), maxSpeed, float64(toX), float64(toY), moveUnit.Rotate, rotate, 200)
		}
	}

	for moveUnit.MoveUUID == uuid {

		entryPoint, outPoint, collisionPoints, collision, _ := BetweenLine(float64(x), float64(y), ToX, ToY, mp, moveUnit.Body, false, size)
		DrawPoints(entryPoint, outPoint, collisionPoints, size, mp.Id, user)

		if collision {
			//  2.1.2 если между координатой истиного пути и целью есть препятсивия запоминаем координату, переходим к пункту 1 и формируем новую

			// находим все препятвия на пути
			obstacles := DetectObstacles(entryPoint, outPoint, collisionPoints, moveUnit, size, user, uuid, mp)
			// берем первое препятвие и обходим его
			points, err := ObstacleAvoidance(mp, obstacles[0], moveUnit, size, user, uuid)
			// находим максимальную отдаленную точку куда может попать юнит
			x, y = SearchPoint(&points, x, y, size, mp, user, moveUnit)

			// TODO в некоторых случаях нули возвращаются в проходимые проходы, так не должно быть
			if err != nil || (x == 0 && y == 0) {
				return nil, err
			}

			CreateRect("green", int(x), int(y), size, moveUnit.MapID, user)
			createPath(x, y)
	} else {
		CreateRect("green", int(moveUnit.ToX), int(moveUnit.ToY), size, moveUnit.MapID, user)
			//  2.1.1 если между координатой истиного пути и целью нет препятсвий формируем путь. Выходим из функции.
			createPath(int(moveUnit.ToX), int(moveUnit.ToY))
			break
		}
	}

	return path, nil
}

func EndIsObstacle(ToX, ToY *float64, xCollision, yCollision, xEnd, yEnd *int, collision *bool, moveUnit *unit.Unit, countCollision int) {

	collisionStartDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *xCollision, *yCollision)
	collisionEndDist := game_math.GetBetweenDist(int(*ToX), int(*ToY), *xEnd, *yEnd)

	// если то старта колизии ближе чем до конца то считаем что маршрут без колизий
	if collisionStartDist < collisionEndDist {
		moveUnit.ToX, moveUnit.ToY = float64(*xCollision), float64(*yCollision)

		// говорим что нет колизий если она всего одна
		if countCollision == 1 {
			*collision = false
		}

	} else {
		// иначе переназначаем конечный пункт что бы не искать путь вечно
		*ToX, *ToY = float64(*xCollision), float64(*yCollision)
	}

	*xCollision, *yCollision = int(*ToX), int(*ToY)
}

func DrawPoints(entryPoints, collisionPoints, outPoints []*coordinate.Coordinate, size int, mapID int, user *player.Player) {
	for i, coor := range collisionPoints {
		CreateRect("red", coor.X, coor.Y, size, mapID, user)

		if len(outPoints)-1 >= i {
			CreateRect("red", outPoints[i].X, outPoints[i].Y, size, mapID, user)
			CreateLine("red", coor.X, coor.Y, outPoints[i].X, outPoints[i].Y, size, mapID, user)
		}
	}
}
