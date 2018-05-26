package movePhase

import (
	"../../coordinate"
	"../../unit"
	"../../player"
	"../../game"
	"../../watchZone"
	"../../db"
	"errors"
)

type TruePatchNode struct {
	WatchNode  *watchZone.UpdaterWatchZone `json:"watch_node"`
	PathNode   *coordinate.Coordinate      `json:"path_node"`
	UnitRotate int                         `json:"unit_rotate"`
}

func InitMove(gameUnit *unit.Unit, toX int, toY int, client *player.Player, game *game.Game) (path []*TruePatchNode) {
	moveTrigger := true

	path = make([]*TruePatchNode, 0)

	for {
		mp := game.GetMap()

		start, _ := mp.GetCoordinate(gameUnit.X, gameUnit.Y)
		end, _ := mp.GetCoordinate(toX, toY)

		pathNodes := FindPath(client, mp, start, end)

		for _, pathNode := range pathNodes {

			errorMove, unitRotate := Move(gameUnit, pathNode, client, end, game)

			if errorMove != nil && errorMove.Error() == "cell is busy" {
				moveTrigger = false
				break
			} else {
				truePatchNode := TruePatchNode{}

				truePatchNode.WatchNode = watchZone.UpdateWatchZone(game, client) // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
				truePatchNode.PathNode = pathNode                                 // добавляем ячейку в путь
				truePatchNode.UnitRotate = unitRotate
				gameUnit.Rotate = unitRotate

				path = append(path, &truePatchNode)
				moveTrigger = true
			}
		}

		if moveTrigger {

			queue := Queue(game)
			gameUnit.Queue = queue

			db.UpdateUnit(gameUnit)

			return
		}
	}
}

func Move(gameUnit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, end *coordinate.Coordinate, game *game.Game) (error, int) {

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		_, ok := client.GetHostileUnit(end.X, end.Y)
		if ok {
			gameUnit.Action = true
			return errors.New("end cell is busy"), 0
		}
	} else {
		_, ok := client.GetHostileUnit(pathNode.X, pathNode.Y)
		if ok {
			return errors.New("cell is busy"), 0 // если клетка занято то выходит из этого пути и генерить новый
		}
	}

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		gameUnit.Action = true
	}

	game.DelUnit(gameUnit) // Удаляем юнита со старых позиций
	client.DelUnit(gameUnit.X, gameUnit.Y)

	rotate := findDirection(pathNode, gameUnit)

	gameUnit.X = pathNode.X // даем новые координаты юниту
	gameUnit.Y = pathNode.Y

	game.SetUnit(gameUnit)
	client.AddUnit(gameUnit) // добавляем новую позицию юнита

	return nil, rotate
}

func Queue(game *game.Game) int {
	queue := 0

	for _, xLine := range game.GetUnits() {
		for _, gameUnit := range xLine {
			if gameUnit.Action {
				if gameUnit.Queue > queue {
					queue = gameUnit.Queue
				}
			}
		}
	}

	return queue + 1
}

func findDirection(pathNode *coordinate.Coordinate, unit *unit.Unit) int {

	if pathNode.X < unit.X && pathNode.Y == unit.Y {
		return 180
	}

	if pathNode.X > unit.X && pathNode.Y == unit.Y {
		return 0
	}

	if pathNode.X == unit.X && pathNode.Y > unit.Y {
		return 90
	}

	if pathNode.X == unit.X && pathNode.Y < unit.Y {
		return 270
	}

	if pathNode.X < unit.X && pathNode.Y < unit.Y {
		return 225
	}

	if pathNode.X > unit.X && pathNode.Y < unit.Y {
		return 315
	}

	if pathNode.X < unit.X && pathNode.Y > unit.Y {
		return 125
	}

	if pathNode.X > unit.X && pathNode.Y > unit.Y {
		return 45
	}

	return 0
}
