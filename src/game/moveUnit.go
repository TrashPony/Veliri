package game

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
