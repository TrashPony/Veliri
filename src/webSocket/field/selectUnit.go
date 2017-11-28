package field

import (
	"../../game/objects"
	"github.com/gorilla/websocket"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client, ok := usersFieldWs[ws]
	game, ok := Games[client.GameID]

	if find && ok {
		respawn := client.Respawn
		if game.stat.Phase == "move" {
			if unit.Action {
				coordinates := objects.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
				obstacles := getObstacles(client)

				moveCoordinate := getMoveCoordinate(coordinates, unit, obstacles)

				for i := 0; i < len(moveCoordinate); i++ {
					if !(moveCoordinate[i].X == respawn.X && moveCoordinate[i].Y == respawn.Y) && moveCoordinate[i].X >= 0 && moveCoordinate[i].Y >= 0 {
						var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: game.stat.Phase,
							X: moveCoordinate[i].X, Y: moveCoordinate[i].Y}
						fieldPipe <- createCoordinates
					}
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: client.Login, Error: "unit already move"}
				fieldPipe <- resp
			}
		}

		if game.stat.Phase == "targeting" {
			coordinates := objects.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
			for _, coordinate := range coordinates {
				targetUnit, ok := client.HostileUnits[coordinate.X][coordinate.Y]
				if ok && targetUnit.NameUser != client.Login {
					var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: game.stat.Phase,
						X: targetUnit.X, Y: targetUnit.Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	} else {
		if game.stat.Phase == "Init"{
			var coordinates []*objects.Coordinate
			respawn := usersFieldWs[ws].Respawn

			for _, coordinate := range usersFieldWs[ws].CreateZone {
				_, ok := usersFieldWs[ws].Units[coordinate.X][coordinate.Y]
				if !ok {
					coordinates = append(coordinates, coordinate)
				}
			}

			for i := 0; i < len(coordinates); i++ {
				if !(coordinates[i].X == respawn.X && coordinates[i].Y == respawn.Y) {
					var createCoordinates = FieldResponse{Event: "SelectCoordinateCreate", UserName: usersFieldWs[ws].Login, X: coordinates[i].X, Y: coordinates[i].Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	}
}

func getMoveCoordinate(radius []*objects.Coordinate, unit *objects.Unit, obstacles []*objects.Coordinate) (res []*objects.Coordinate) { // берет все соседние клетки от текущей
	start := objects.Coordinate{X: unit.X, Y: unit.Y}
	openCoordinate := make(map[int]map[int]*objects.Coordinate)
	closeCoordinate := make(map[int]map[int]*objects.Coordinate)
	obstaclesMatrix := make(map[int]map[int]*objects.Coordinate)

	startMatrix := generateNeighboursCoord(&start, obstaclesMatrix)


	for _, obstacle := range obstacles{
		if obstaclesMatrix[obstacle.X] != nil {
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		} else {
			obstaclesMatrix[obstacle.X] = make(map[int]*objects.Coordinate)
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		}
	}

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
						addCoordIfValid(closeCoordinate, obstaclesMatrix, coordinate.X, coordinate.Y)
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

func generateNeighboursCoord(curr *objects.Coordinate, obstacles map[int]map[int]*objects.Coordinate) (res map[int]map[int]*objects.Coordinate) { // берет все соседние клетки от текущей
	res = make(map[int]map[int]*objects.Coordinate)
	//верх лево
	addCoordIfValid(res, obstacles, curr.X-1, curr.Y+1)
	//верх центр
	addCoordIfValid(res, obstacles, curr.X, curr.Y+1)
	//верх право
	addCoordIfValid(res, obstacles, curr.X+1, curr.Y+1)

	//строго лево
	addCoordIfValid(res, obstacles, curr.X-1, curr.Y)
	//строго право
	addCoordIfValid(res, obstacles, curr.X+1, curr.Y)

	//низ лево
	addCoordIfValid(res, obstacles, curr.X-1, curr.Y-1)
	//низ центр
	addCoordIfValid(res, obstacles, curr.X, curr.Y-1)
	//низ право
	addCoordIfValid(res, obstacles, curr.X+1, curr.Y-1)

	return
}

func addCoordIfValid(res map[int]map[int]*objects.Coordinate, obstacles map[int]map[int]*objects.Coordinate, x int, y int) {
	coor := objects.Coordinate{X:x , Y:y}
	_, ok := obstacles[x][y]
	if !ok {
		if res[x] != nil {
			res[x][y] = &coor
		} else {
			res[x] = make(map[int]*objects.Coordinate)
			res[x][y] = &coor
		}
	}
}
