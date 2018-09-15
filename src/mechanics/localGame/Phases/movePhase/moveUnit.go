package movePhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../../localGame"
	"../../../player"
	"../../map/watchZone"
	"errors"
)

type TruePatchNode struct {
	WatchNode  *watchZone.UpdaterWatchZone `json:"watch_node"`
	PathNode   *coordinate.Coordinate      `json:"path_node"`
	UnitRotate int                         `json:"unit_rotate"`
}

func InitMove(gameUnit *unit.Unit, toQ int, toR int, client *player.Player, game *localGame.Game, event string) (path []*TruePatchNode) {
	mp := game.GetMap()
	start, _ := mp.GetCoordinate(gameUnit.Q, gameUnit.R)
	end, _ := mp.GetCoordinate(toQ, toR)

	path = make([]*TruePatchNode, 0)

	if event != "SelectStorageUnit" {
		startPoint := TruePatchNode{WatchNode: nil, PathNode: start, UnitRotate: gameUnit.Rotate}
		path = append(path, &startPoint)
	}

	err, pathNodes := FindPath(client, mp, start, end, gameUnit, event)
	if err != nil {
		print(err.Error())
		return
	}

	for i, pathNode := range pathNodes {
		deleteUnit := true

		if i == 0 && event == "SelectStorageUnit" {
			deleteUnit = false
		}

		errorMove, unitRotate := Move(gameUnit, pathNode, client, game, deleteUnit)

		if errorMove != nil {
			if errorMove.Error() == "cell is busy" {
				break
			}
		} else {
			truePatchNode := TruePatchNode{}

			truePatchNode.WatchNode = watchZone.UpdateWatchZone(game, client) // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
			truePatchNode.PathNode = pathNode                                 // добавляем ячейку в путь
			truePatchNode.UnitRotate = unitRotate
			gameUnit.Rotate = unitRotate

			path = append(path, &truePatchNode)
		}
	}

	if gameUnit.ActionPoints == 0 {
		gameUnit.Action = true
		queue := Queue(game)
		gameUnit.QueueAttack = queue
	}

	gameUnit.OnMap = true

	updateSquad.Squad(client.GetSquad())

	return
}

func Move(gameUnit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, game *localGame.Game, deleteUnit bool) (error, int) {

	_, ok := game.GetUnit(pathNode.Q, pathNode.R)
	if ok || !checkMSPlace(client, pathNode, gameUnit, false) {
		return errors.New("cell is busy"), 0 // если клетка занято то выходит из этого пути и генерить новый
	}

	if deleteUnit {
		game.DelUnit(gameUnit) // Удаляем юнита со старых позиций
		client.DelUnit(gameUnit.Q, gameUnit.R)
	}

	rotate := findDirection(pathNode, gameUnit)

	gameUnit.Q = pathNode.Q // даем новые координаты юниту
	gameUnit.R = pathNode.R
	gameUnit.ActionPoints -= 1

	game.SetUnit(gameUnit)
	client.AddUnit(gameUnit) // добавляем новую позицию юнита

	return nil, rotate
}

func findDirection(pathNode *coordinate.Coordinate, unit *unit.Unit) int {

	if unit.Q < pathNode.Q && unit.R == pathNode.R {
		return 0
	}
	if unit.Q > pathNode.Q && unit.R == pathNode.R {
		return 180
	}

	if unit.R%2 == 0 {
		if unit.Q == pathNode.Q && unit.R < pathNode.R {
			return 60
		}
		if unit.Q > pathNode.Q && unit.R < pathNode.R {
			return 120
		}
		if unit.Q > pathNode.Q && unit.R > pathNode.R {
			return 240
		}
		if unit.Q == pathNode.Q && unit.R > pathNode.R {
			return 300
		}
	} else {
		if unit.Q < pathNode.Q && unit.R < pathNode.R {
			return 60
		}
		if unit.Q == pathNode.Q && unit.R < pathNode.R {
			return 120
		}
		if unit.Q == pathNode.Q && unit.R > pathNode.R {
			return 240
		}
		if unit.Q < pathNode.Q && unit.R > pathNode.R {
			return 300
		}
	}

	return 0
}
