package mechanics

import (
	"./player"
	"./game"
	"./coordinate"
)

func GetObstacles(client *player.Player, game *game.Game) (obstaclesMatrix map[int]map[int]*coordinate.Coordinate) { // TODO: это все очень странно
	coordinates := make([]*coordinate.Coordinate, 0)
	obstaclesMatrix = make(map[int]map[int]*coordinate.Coordinate)

	for _, xLine := range client.GetUnits() {
		for _, unit := range xLine {
			var gameCoordinate coordinate.Coordinate
			gameCoordinate.X = unit.X
			gameCoordinate.Y = unit.Y
			coordinates = append(coordinates, &gameCoordinate)
		}
	}

	for _, xLine := range client.GetHostileUnits() {
		for _, unit := range xLine {
			var gameCoordinate coordinate.Coordinate
			gameCoordinate.X = unit.X
			gameCoordinate.Y = unit.Y
			coordinates = append(coordinates, &gameCoordinate)
		}
	}

	var gameCoordinate coordinate.Coordinate
	gameCoordinate.X = client.GetMatherShip().X
	gameCoordinate.Y = client.GetMatherShip().Y
	coordinates = append(coordinates, &gameCoordinate)

	for _, xLine := range client.GetHostileMatherShips() {
		for _, structure := range xLine {
			var gameCoordinate coordinate.Coordinate
			gameCoordinate.X = structure.X
			gameCoordinate.Y = structure.Y
			coordinates = append(coordinates, &gameCoordinate)
		}
	}

	for _, xLine := range game.GetMap().OneLayerMap {
		for _, obstacles := range xLine {
			if obstacles.Type == "obstacle" || obstacles.Type == "terrain" {
				var gameCoordinate coordinate.Coordinate
				gameCoordinate.X = obstacles.X
				gameCoordinate.Y = obstacles.Y
				coordinates = append(coordinates, &gameCoordinate)
			}
		}
	}

	for _, obstacle := range coordinates {
		if obstaclesMatrix[obstacle.X] != nil {
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		} else {
			obstaclesMatrix[obstacle.X] = make(map[int]*coordinate.Coordinate)
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		}
	}
	return
}
