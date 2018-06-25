package movePhase

import (
	"../../map/coordinate"
	"../../../gameObjects/unit"
	"../../../player"
	"../../map/watchZone"
	"../../../db"
	"errors"
	"math"
	"../../../localGame"
)

type TruePatchNode struct {
	WatchNode  *watchZone.UpdaterWatchZone `json:"watch_node"`
	PathNode   *coordinate.Coordinate      `json:"path_node"`
	UnitRotate int                         `json:"unit_rotate"`
}

func InitMove(gameUnit *unit.Unit, toX int, toY int, client *player.Player, game *localGame.Game) (path []*TruePatchNode) {
	moveTrigger := true

	mp := game.GetMap()
	start, _ := mp.GetCoordinate(gameUnit.X, gameUnit.Y)
	end, _ := mp.GetCoordinate(toX, toY)

	path = make([]*TruePatchNode, 0)

	startPoint := TruePatchNode{WatchNode: nil, PathNode: start, UnitRotate: gameUnit.Rotate}
	path = append(path, &startPoint)

	for {

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
			gameUnit.QueueAttack = queue

			db.UpdateUnit(gameUnit)

			return
		}
	}
}

func Move(gameUnit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, end *coordinate.Coordinate, game *localGame.Game) (error, int) {

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		_, ok := client.GetHostileUnit(end.X, end.Y)
		if ok {
			gameUnit.Action = false // todo должно быть true но для тестов пока будет false
			return errors.New("end cell is busy"), 0
		}
	} else {
		_, ok := client.GetHostileUnit(pathNode.X, pathNode.Y)
		if ok {
			return errors.New("cell is busy"), 0 // если клетка занято то выходит из этого пути и генерить новый
		}
	}

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		gameUnit.Action = false  // todo должно быть true но для тестов пока будет false
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

func findDirection(pathNode *coordinate.Coordinate, unit *unit.Unit) int {

	rotate := math.Atan2(float64(pathNode.Y - unit.Y), float64(pathNode.X - unit.X))

	rotate = rotate * 180/math.Pi

	if rotate < 0 {
		rotate = 360 + rotate
	}

	return int(rotate)
}
