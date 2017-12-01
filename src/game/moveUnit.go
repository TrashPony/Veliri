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

func generateNeighboursCoord(curr *Coordinate, obstacles map[int]map[int]*Coordinate) (res map[int]map[int]*Coordinate) { // берет все соседние клетки от текущей
	res = make(map[int]map[int]*Coordinate)

	//строго лево
	_, left := obstacles[curr.X-1][curr.Y]
	addCoordIfValid(res, obstacles, curr.X-1, curr.Y)
	//строго право
	_, right := obstacles[curr.X+1][curr.Y]
	addCoordIfValid(res, obstacles, curr.X+1, curr.Y)
	//верх центр
	_, top := obstacles[curr.X][curr.Y-1]
	addCoordIfValid(res, obstacles, curr.X, curr.Y-1)
	//низ центр
	_, bottom := obstacles[curr.X][curr.Y+1]
	addCoordIfValid(res, obstacles, curr.X, curr.Y+1)


	//верх лево/    ЛЕВО И верх
	if !(left || top) {
		addCoordIfValid(res, obstacles, curr.X-1, curr.Y-1)
	}
	//верх право/   ПРАВО И верх
	if !(right || top) {
		addCoordIfValid(res, obstacles, curr.X+1, curr.Y-1)
	}
	//низ лево/  если ЛЕВО И низ
	if !(left || bottom) {
		addCoordIfValid(res, obstacles, curr.X-1, curr.Y+1)
	}
	//низ право/  низ И ВЕРХ
	if !(right || bottom) {
		addCoordIfValid(res, obstacles, curr.X+1, curr.Y+1)
	}

	return
}

func addCoordIfValid(res map[int]map[int]*Coordinate, obstacles map[int]map[int]*Coordinate, x int, y int) {
	coor := Coordinate{X:x , Y:y}

	_, ok := obstacles[x][y]
	if !ok && x >= 0 && y >= 0 {
		if res[x] != nil {
			res[x][y] = &coor
		} else {
			res[x] = make(map[int]*Coordinate)
			res[x][y] = &coor
		}
	}
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
