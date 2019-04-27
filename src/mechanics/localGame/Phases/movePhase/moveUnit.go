package movePhase

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/db/localGame/update"
	squadUpdate "github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
)

type TruePatchNode struct {
	WatchNode  *watchZone.UpdaterWatchZone `json:"watch_node"`
	PathNode   *coordinate.Coordinate      `json:"path_node"`
	UnitRotate int                         `json:"unit_rotate"`
	ToMC       bool                        `json:"to_mc"`
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

		errorMove, unitRotate, watchNode := Move(gameUnit, pathNode, client, game, deleteUnit, event)

		if errorMove != nil {

			if errorMove.Error() == "ToMC" {
				truePatchNode := TruePatchNode{}

				truePatchNode.WatchNode = watchNode // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
				truePatchNode.PathNode = pathNode   // добавляем ячейку в путь
				truePatchNode.UnitRotate = unitRotate
				truePatchNode.ToMC = true

				gameUnit.Rotate = unitRotate

				path = append(path, &truePatchNode)
				break
			}

			if errorMove.Error() == "cell is busy" {
				break
			}

			if errorMove.Error() == "find hostile" { // если нашли юнита то выходим из цикла но добавляем последние клетки
				truePatchNode := TruePatchNode{}

				truePatchNode.WatchNode = watchNode // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
				truePatchNode.PathNode = pathNode   // добавляем ячейку в путь
				truePatchNode.UnitRotate = unitRotate
				gameUnit.Rotate = unitRotate

				path = append(path, &truePatchNode)
				break
			}
		} else {
			truePatchNode := TruePatchNode{}

			truePatchNode.WatchNode = watchNode // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
			truePatchNode.PathNode = pathNode   // добавляем ячейку в путь
			truePatchNode.UnitRotate = unitRotate
			gameUnit.Rotate = unitRotate

			path = append(path, &truePatchNode)
		}
	}

	if !gameUnit.FindHostile || gameUnit.ActionPoints == 0 {
		gameUnit.Move = false
		gameUnit.ActionPoints = 0
		QueueMove(game) // определяет какой игрок будет ходить следующим
	}

	gameUnit.FindHostile = false

	// если мы идем в трюм то юнит не может быть на карте
	if event != "ToMC" {
		gameUnit.OnMap = true
	}

	squadUpdate.Squad(client.GetSquad(), true)
	update.Player(client)

	return
}

func Move(gameUnit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, game *localGame.Game, deleteUnit bool, event string) (error, int, *watchZone.UpdaterWatchZone) {

	obstacleUnit, ok := game.GetUnit(pathNode.Q, pathNode.R)

	if event == "ToMC" && ok && obstacleUnit.Body.MotherShip && obstacleUnit.Owner == client.GetLogin() {
		// убераем юнита из игры и кладем в трюм
		gameUnit.OnMap = false

		game.DelUnit(gameUnit)
		client.DelUnit(gameUnit, false)

		game.AddUnitToStorage(gameUnit)
		client.AddUnitStorage(gameUnit)

		return errors.New("ToMC"), 0, watchZone.UpdateWatchZone(game, client)
	}

	if ok || !checkMSPlace(client, pathNode, gameUnit, event == "ToMC") {
		return errors.New("cell is busy"), 0, nil // если клетка занято то выходит из этого пути и генерить новый
	}

	if deleteUnit {
		game.DelUnit(gameUnit) // Удаляем юнита со старых позиций
		client.DelUnit(gameUnit, false)
	}

	rotate := findDirection(pathNode, gameUnit)

	gameUnit.Q = pathNode.Q // даем новые координаты юниту
	gameUnit.R = pathNode.R
	gameUnit.ActionPoints -= 1

	game.SetUnit(gameUnit)
	client.AddUnit(gameUnit) // добавляем новую позицию юнита

	watchNode := watchZone.UpdateWatchZone(game, client) // смотри открыл он новых вражеских юнитов
	if len(watchNode.OpenUnit) > 0 {

		for _, openHostileUnit := range watchNode.OpenUnit { // добавляем всех открытых юнитов в увидиные пользователем
			client.AddNewMemoryHostileUnit(*openHostileUnit)
		}

		gameUnit.FindHostile = true
		return errors.New("find hostile"), rotate, watchNode
	}

	return nil, rotate, watchNode
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
