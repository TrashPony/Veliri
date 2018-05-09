package game

func GetObstacles(client *Player, game *Game) (obstaclesMatrix map[int]map[int]*Coordinate) { // TODO: это все очень странно
	coordinates := make([]*Coordinate, 0)
	obstaclesMatrix = make(map[int]map[int]*Coordinate)

	for _, xLine := range client.GetUnits() {
		for _, unit := range xLine {
			var coordinate Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range client.GetHostileUnits() {
		for _, unit := range xLine {
			var coordinate Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	var coordinate Coordinate
	coordinate.X = client.GetMatherShip().X
	coordinate.Y = client.GetMatherShip().Y
	coordinates = append(coordinates, &coordinate)

	for _, xLine := range client.GetHostileMatherShips() {
		for _, structure := range xLine {
			var coordinate Coordinate
			coordinate.X = structure.X
			coordinate.Y = structure.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range game.GetMap().OneLayerMap {
		for _, obstacles := range xLine {
			if obstacles.Type == "obstacle" {
				var coordinate Coordinate
				coordinate.X = obstacles.X
				coordinate.Y = obstacles.Y
				coordinates = append(coordinates, &coordinate)
			}
		}
	}

	for _, obstacle := range coordinates {
		if obstaclesMatrix[obstacle.X] != nil {
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		} else {
			obstaclesMatrix[obstacle.X] = make(map[int]*Coordinate)
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		}
	}
	return
}
