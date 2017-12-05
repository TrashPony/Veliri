package game

import (
	"errors"
)

func InitMove(unit *Unit, toX int, toY int , client *Player, game *Game) (truePath map[Coordinate]*UpdaterWatchZone, pathNodes []Coordinate) {
	truePath = make(map[Coordinate]*UpdaterWatchZone)
	pathNodes = make([]Coordinate,0)
	idGame := client.GetGameID()
	moveTrigger := true

	for {
		obstacles := GetObstacles(client, game)

		start := Coordinate{X: unit.X, Y: unit.Y}
		end := Coordinate{X: toX, Y: toY}

		mp := game.GetMap()

		path := FindPath(mp, start, end, obstacles)

		for _, pathNode := range path {

			errorMove := Move(unit, pathNode, client, end, game)

			if errorMove != nil && errorMove.Error() == "cell is busy" {
				moveTrigger = false
				break
			} else {
				truePath[pathNode] = client.UpdateWatchZone(game) // обновляем у клиента открытые ячейки, удаляем закрытые кидаем в карту
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

func Move(unit *Unit, pathNode Coordinate, client *Player, end Coordinate, game *Game) (error) {

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

		unit.X = pathNode.X // даем новые координаты юниту
		unit.Y = pathNode.Y

		game.SetUnit(unit)
		client.AddUnit(unit) // добавляем новую позицию юнита

		return nil
}



func GetMoveCoordinate(radius []*Coordinate, unit *Unit, obstaclesMatrix map[int]map[int]*Coordinate) (res []*Coordinate) { // берет все соседние клетки от текущей
	start := Coordinate{X: unit.X, Y: unit.Y}
	openCoordinate := make(map[int]map[int]*Coordinate)
	closeCoordinate := make(map[int]map[int]*Coordinate)

	startMatrix := generateNeighboursCoord(&start, obstaclesMatrix)

	for _, xline := range startMatrix {
		for _, coordiante := range xline {
			addCoordIfValid(openCoordinate, obstaclesMatrix, coordiante.X, coordiante.Y)
		}
	}

	for i := 0; i < unit.MoveSpeed-1; i++ {
		for _, xline := range openCoordinate {
			for _, coordinate := range xline {
				matrix := generateNeighboursCoord(coordinate, obstaclesMatrix)
				for _, xline := range matrix {
					for _, coordinate := range xline {
						_, ok := openCoordinate[coordinate.X][coordinate.Y]
						if !ok {
							addCoordIfValid(closeCoordinate, obstaclesMatrix, coordinate.X, coordinate.Y)
						}
					}
				}
			}
		}

		for _, xline := range closeCoordinate {
			for _, coordinate := range xline {
				addCoordIfValid(openCoordinate, obstaclesMatrix, coordinate.X, coordinate.Y)
			}
		}
	}


	for _, coordinate := range radius {
		_, ok := openCoordinate[coordinate.X][coordinate.Y]
		if ok {
			res = append(res, coordinate)
		}
	}

	return
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
