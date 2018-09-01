package movePhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../player"
	"../../map/watchZone"
	"errors"
	"../../../localGame"
	"../../../db/updateSquad"
)

type TruePatchNode struct {
	WatchNode  *watchZone.UpdaterWatchZone `json:"watch_node"`
	PathNode   *coordinate.Coordinate      `json:"path_node"`
	UnitRotate int                         `json:"unit_rotate"`
}

func InitMove(gameUnit *unit.Unit, toQ int, toR int, client *player.Player, game *localGame.Game) (path []*TruePatchNode) {
	moveTrigger := true

	mp := game.GetMap()
	start, _ := mp.GetCoordinate(gameUnit.Q, gameUnit.R)
	end, _ := mp.GetCoordinate(toQ, toR)

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

			updateSquad.Squad(client.GetSquad())

			return
		}
	}
}

func Move(gameUnit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, end *coordinate.Coordinate, game *localGame.Game) (error, int) {

	if (end.Q == pathNode.Q) && (end.R == pathNode.R) {
		_, ok := client.GetHostileUnit(end.Q, end.R)
		if ok {
			gameUnit.Action = false // todo должно быть true но для тестов пока будет false
			return errors.New("end cell is busy"), 0
		}
	} else {
		_, ok := client.GetHostileUnit(pathNode.Q, pathNode.R)
		if ok {
			return errors.New("cell is busy"), 0 // если клетка занято то выходит из этого пути и генерить новый
		}
	}

	if (end.Q == pathNode.Q) && (end.R == pathNode.R) {
		gameUnit.Action = false // todo должно быть true но для тестов пока будет false
	}

	game.DelUnit(gameUnit) // Удаляем юнита со старых позиций
	client.DelUnit(gameUnit.Q, gameUnit.R)

	rotate := findDirection(pathNode, gameUnit)

	gameUnit.Q = pathNode.Q // даем новые координаты юниту
	gameUnit.R = pathNode.R

	game.SetUnit(gameUnit)
	client.AddUnit(gameUnit) // добавляем новую позицию юнита

	return nil, rotate
}

func findDirection(pathNode *coordinate.Coordinate, unit *unit.Unit) int {

	if unit.Q < pathNode.Q && unit.R == pathNode.R { return 0 }
	if unit.Q > pathNode.Q && unit.R == pathNode.R { return 180 }

	if unit.R%2 == 0 {
		if unit.Q == pathNode.Q && unit.R < pathNode.R { return 60 }
		if unit.Q > pathNode.Q && unit.R < pathNode.R {	return 120 }
		if unit.Q > pathNode.Q && unit.R > pathNode.R { return 240 }
		if unit.Q == pathNode.Q && unit.R > pathNode.R { return 300 }
	} else {
		if unit.Q < pathNode.Q && unit.R < pathNode.R { return 60 }
		if unit.Q == pathNode.Q && unit.R < pathNode.R { return 120 }
		if unit.Q == pathNode.Q && unit.R > pathNode.R { return 240 }
		if unit.Q < pathNode.Q && unit.R > pathNode.R { return 300 }
	}

	return 0
}
