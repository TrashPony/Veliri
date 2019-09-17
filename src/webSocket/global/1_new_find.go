package global

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
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

	path := make([]*unit.PathUnit, 0)
	appendPath := func(appendPath []*unit.PathUnit) {
		for i := 0; i < len(appendPath); i++ {
			path = append(path, appendPath[i])
		}
	}

	CreateRect("green", int(startX), int(startY), rectSize, moveUnit.MapID, user)
	CreateLine("white", int(startX), int(startY), int(ToX), int(ToY), rectSize, moveUnit.MapID, user)

	// 0 пытаемя проложить путь от начала пути до конечной точки по прямой
	entryPoint, outPoint, collisionPoints, collision, endIsObstacle := BetweenLine(startX, startY, ToX, ToY, mp, moveUnit.Body, true)
	DrawPoints(entryPoint, outPoint, collisionPoints, rectSize, mp.Id, user)

	// 0.1 если конечная точка находится в препятсвие то смотрим куда ближе идти ко входу или к выходу
	if endIsObstacle {
		// последние точки это колизия вокруг точки назначения
		lastEntryX, lastEntryY := &entryPoint[len(entryPoint)-1].X, &entryPoint[len(entryPoint)-1].Y
		lastOutX, lastOutY := &outPoint[len(outPoint)-1].X, &outPoint[len(outPoint)-1].Y
		EndIsObstacle(&ToX, &ToY, lastEntryX, lastEntryY, lastOutX, lastOutY, &collision, moveUnit)
	}

	if !collision {
		// 0.2 на прямой между стартом и концом нет прептявий
		CreateRect("green", int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)
		CreateLine("green", int(startX), int(startY), int(moveUnit.ToX), int(moveUnit.ToY), rectSize, moveUnit.MapID, user)

		err, path := To2(startX, startY, maxSpeed, moveUnit.ToX, moveUnit.ToY, moveUnit.Rotate, rotate, 200)
		return path, err
	} else {
		// 0.3 на прямой были найдены препятвия

		// TODO кеширование пройденных клеток, т.к. это может быть 1 очень кривое препятвие и придется обходить 1 и теже точки
		// TODO надо определить это 2+ препятвия или 1 но С/E образного типа! Это можно сделать только предварительным сканирование обьекта

		// находим все препятвия на пути
		obstacles := DetectObstacles(entryPoint, outPoint, collisionPoints, moveUnit, rectSize, user, uuid, mp)
		// берем первое препятвие и обходим его
		points, err := ObstacleAvoidance(mp, obstacles[0].Entry, obstacles[0].Out, obstacles[0].EntryCollision, moveUnit, rectSize, user, uuid)
		// находим максимальную отдаленную точку куда может попать юнит
		x, y := SearchPoint(&points, moveUnit.X, moveUnit.Y, rectSize, mp, user, moveUnit)

		if err != nil {
			return nil, err
		}

		_, path = To2(startX, startY, maxSpeed, float64(x), float64(y), moveUnit.Rotate, rotate, 200)

		CreateRect("green", x, y, rectSize, moveUnit.MapID, user)
		CreateLine("green", int(startX), int(startY), int(x), int(y), rectSize, moveUnit.MapID, user)

		for moveUnit.MoveUUID == uuid {

			entryPoint, outPoint, collisionPoints, collision, endIsObstacle = BetweenLine(float64(x), float64(y), ToX, ToY, mp, moveUnit.Body, false)
			DrawPoints(entryPoint, outPoint, collisionPoints, rectSize, mp.Id, user)

			if endIsObstacle {
				// последние точки это колизия вокруг точки назначения
				//lastEntryX, lastEntryY := &entryPoint[len(entryPoint)-1].X, &entryPoint[len(entryPoint)-1].Y
				//lastOutX, lastOutY := &outPoint[len(outPoint)-1].X, &outPoint[len(outPoint)-1].Y
				//EndIsObstacle(&ToX, &ToY, lastEntryX, lastEntryY, lastOutX, lastOutY, &collision, moveUnit)
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

				// находим все препятвия на пути
				obstacles = DetectObstacles(entryPoint, outPoint, collisionPoints, moveUnit, rectSize, user, uuid, mp)
				// берем первое препятвие и обходим его
				points, err = ObstacleAvoidance(mp, obstacles[0].Entry, obstacles[0].Out, obstacles[0].EntryCollision, moveUnit, rectSize, user, uuid)
				// находим максимальную отдаленную точку куда может попать юнит
				x, y = SearchPoint(&points, x, y, rectSize, mp, user, moveUnit)

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
		// TODO если колизий больше чем одна то мы не может отключать колизию
		*collision = false
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