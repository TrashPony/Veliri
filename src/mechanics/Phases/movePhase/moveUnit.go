package movePhase

/*func InitMove(unit *Unit, toX int, toY int , client *Player, game *Game) (watchNode map[string]*UpdaterWatchZone, pathNodes []Coordinate) {
	watchNode = make(map[string]*UpdaterWatchZone)
	idGame := client.GetGameID()
	moveTrigger := true

	pathNodes = make([]Coordinate,0) // создаем пустую болванку для пути
	pathNodes = append(pathNodes, Coordinate{X: unit.X, Y: unit.Y}) // кладем в него стартовую ячейку

	for {
		obstacles := GetObstacles(client, game)

		start := Coordinate{X: unit.X, Y: unit.Y}
		end := Coordinate{X: toX, Y: toY}

		mp := game.GetMap()

		path := FindPath(mp, start, end, obstacles)

		for _, pathNode := range path {

			errorMove := Move(unit, &pathNode, client, end, game)

			if errorMove != nil && errorMove.Error() == "cell is busy" {
				moveTrigger = false
				break
			} else {
				watchNode[strconv.Itoa(pathNode.X) + ":" + strconv.Itoa(pathNode.Y)] = client.UpdateWatchZone(game) // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
				pathNodes = append(pathNodes, pathNode)           // создать пройденный путь
			}
		}

		if moveTrigger {
			queue := MoveUnit(idGame, unit, end.X, end.Y)
			unit.Queue = queue
			return
		}
	}
}

func Move(unit *Unit, pathNode *Coordinate, client *Player, end Coordinate, game *Game) (error) {

		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := client.GetHostileUnit(end.X,end.Y)
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



func MoveUnit(idGame int, unit *Unit, toX int, toY int) int {

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
}

func findDirection(pathNode *Coordinate, unit *Unit)  {
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
}*/
