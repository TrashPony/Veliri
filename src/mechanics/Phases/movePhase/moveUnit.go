package movePhase

import (
	"../../coordinate"
	"../../unit"
	"../../player"
	"../../game"
	"../../watchZone"
	"strconv"
	"errors"
)

func InitMove(unit *unit.Unit, toX int, toY int, client *player.Player, game *game.Game) (watchNode map[string]*watchZone.UpdaterWatchZone, pathNodes []*coordinate.Coordinate) {
	watchNode = make(map[string]*watchZone.UpdaterWatchZone)
	moveTrigger := true

	pathNodes = make([]*coordinate.Coordinate, 0)

	for {
		mp := game.GetMap()

		start, _ := mp.GetCoordinate(unit.X, unit.Y)
		end, _ := mp.GetCoordinate(toX, toY)

		path := FindPath(client, mp, start, end)

		for _, pathNode := range path {

			errorMove := Move(unit, pathNode, client, end, game)

			if errorMove != nil && errorMove.Error() == "cell is busy" {
				moveTrigger = false
				break
			} else {
				watchNode[strconv.Itoa(pathNode.X)+":"+strconv.Itoa(pathNode.Y)] = watchZone.UpdateWatchZone(game, client) // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
				pathNodes = append(pathNodes, pathNode)                                                                    // создать пройденный путь
			}
		}

		if moveTrigger {
			//queue := MoveUnit(idGame, unit, end.X, end.Y)
			//unit.Queue = queue
			return
		}
	}
}

func Move(unit *unit.Unit, pathNode *coordinate.Coordinate, client *player.Player, end *coordinate.Coordinate, game *game.Game) (error) {

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		_, ok := client.GetHostileUnit(end.X, end.Y)
		if ok {
			unit.Action = false
			return errors.New("end cell is busy")
		}
	} else {
		_, ok := client.GetHostileUnit(pathNode.X, pathNode.Y)
		if ok {
			return errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
		}
	}

	if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
		unit.Action = false
	}

	game.DelUnit(unit) // Удаляем юнита со старых позиций
	client.DelUnit(unit.X, unit.Y)

	findDirection(pathNode, unit)

	unit.X = pathNode.X // даем новые координаты юниту
	unit.Y = pathNode.Y

	game.SetUnit(unit)
	client.AddUnit(unit) // добавляем новую позицию юнита

	return nil
}

/*func MoveUnit(idGame int, unit *unit.Unit, toX int, toY int) int {

	rows, err := db.Query("Select  MAX(queue_attack) FROM action_game_unit WHERE id_game=$1", idGame)
	if err != nil {
		return 0
	}
	defer rows.Close()

	var queue int

	for rows.Next() {
		err := rows.Scan(&queue)
		if err != nil {
			return 0
		}
	}
	queue += 1

	// устанавливает фраг готовности пользователя и ставить очередь
	_, err = db.Query("UPDATE action_game_unit  SET x = $1, y = $2, action = $5, queue_attack = $6  WHERE id=$3 AND id_game=$4", toX, toY, unit.Id, idGame, false, queue)
	if err != nil {
		return queue
	} else {
		return queue
	}
}*/

func findDirection(pathNode *coordinate.Coordinate, unit *unit.Unit) {
	//TODO//////////// проверка направления юнита ///////////////

	if pathNode.X < unit.X && pathNode.Y == unit.Y {
		println("Идет ровно влево")
	}

	if pathNode.X > unit.X && pathNode.Y == unit.Y {
		println("Идет ровно вправо")
	}

	if pathNode.X == unit.X && pathNode.Y > unit.Y {
		println("Идет ровно вниз")
	}

	if pathNode.X == unit.X && pathNode.Y < unit.Y {
		println("Идет ровно вверх")
	}

	//TODO///////////////////////////////////////////////////////

	if pathNode.X < unit.X && pathNode.Y < unit.Y {
		println("Идет верх влево")
	}

	if pathNode.X > unit.X && pathNode.Y < unit.Y {
		println("Идет верх вправо")
	}

	if pathNode.X < unit.X && pathNode.Y > unit.Y {
		println("Идет вниз влево")
	}

	if pathNode.X > unit.X && pathNode.Y > unit.Y {
		println("Идет вниз вправо")
	}
}
