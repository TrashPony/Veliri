package field

import (
	"../../game/mechanics"
	"../../game/objects"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client, ok := usersFieldWs[ws]
	if find && ok {
		if unit.Action {

			coordinates := objects.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			obstacles := getObstacles(client)
			moveCoordinate := mechanics.GetMoveCoordinate(coordinates, unit, obstacles)

			var passed bool
			for _, coordinate := range moveCoordinate {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
					resp = FieldResponse{Event: msg.Event, UserName: client.Login}
					fieldPipe <- resp
					go InitMove(unit, msg, client)
					passed = true
				}
			}

			if !passed {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
				fieldPipe <- resp
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "unit already move"}
			fieldPipe <- resp
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not found unit"}
		fieldPipe <- resp
	}
}

func InitMove(unit *objects.Unit, msg FieldMessage, client *Clients) {

	idGame := client.GameID
	toX := msg.ToX
	toY := msg.ToY

	for {
		obstacles := getObstacles(client)

		start := objects.Coordinate{X: unit.X, Y: unit.Y}
		end := objects.Coordinate{X: toX, Y: toY}

		mp := Games[client.GameID].getMap()

		path := mechanics.FindPath(mp, start, end, obstacles)

		x, y, errorMove := Move(unit, path, client, end)
		if errorMove != nil {
			if errorMove.Error() != "cell is busy" {
				queue := mechanics.MoveUnit(idGame, unit, x, y)
				unit.Queue = queue
				break
			}
		} else {
			queue := mechanics.MoveUnit(idGame, unit, x, y)
			unit.Queue = queue
			break
		}
	}
}

func Move(unit *objects.Unit, path []objects.Coordinate, client *Clients, end objects.Coordinate) (int, int, error) {

	game := Games[client.GameID]
	players := Games[client.GameID].getPlayers()
	activeUser := ActionGameUser(players)

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := client.HostileUnits[end.X][end.Y]
			if ok {
				unit.Action = false
				var unitsParameter InitUnit
				unitsParameter.initUnit(unit, client.Login)
				return unit.X, unit.Y, errors.New("end cell is busy")
			}
		} else {
			_, ok := client.HostileUnits[pathNode.X][pathNode.Y]
			if ok {
				return 0, 0, errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		game.delUnit(unit) // TODO сделать интерфейс для ходьбы

		x := unit.X
		y := unit.Y

		unit.X = pathNode.X
		unit.Y = pathNode.Y

		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			unit.Action = false
		}

		game.setUnit(unit)

		delete(client.Units[x], y)
		client.addUnit(unit) // добавляем новое

		client.updateWatchZone(game) // отправляем открытые ячейки, удаляем закрытые
		go updateWatchHostileUser(*client, *unit, x, y, activeUser)  // добавляем и удаляем нашего юнита у врагов на карте

		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.Login) // отсылаем новое место юнита

		time.Sleep(200 * time.Millisecond)
	}

	return unit.X, unit.Y, nil
}

func updateWatchHostileUser(client Clients, unit objects.Unit, x, y int, activeUser []*Clients) {
	var unitsParameter InitUnit

	for _, user := range activeUser {
		if user.Login != client.Login {
			_, okGetUnit := user.HostileUnits[x][y]

			if okGetUnit {
				user.Watch[x][y] = &objects.Coordinate{X: x, Y: y} // добавлдяем на место старого места юнита пустую зону
				delete(user.HostileUnits[x], y)                    // и удаляем в общей карте вражеских юнитов
				openCoordinate(user.Login, x, y)                   // и остылаем событие удаление юнита
			}

			_, okGetXY := user.Watch[unit.X][unit.Y]

			if okGetXY { // если следующая клетка юнита в зоне видимости
				delete(user.Watch[unit.X], unit.Y) // удаляем пустую клетку
				user.addHostileUnit(&unit)         // и добавляем в общую карту вражеских юнитов
				unitsParameter.initUnit(&unit, user.Login)
			}
		}
	}
}

func getObstacles(client *Clients) (obstaclesMatrix map[int]map[int]*objects.Coordinate) { // TODO: это все очень странно
	coordinates := make([]*objects.Coordinate, 0)
	obstaclesMatrix = make(map[int]map[int]*objects.Coordinate)

	// TODO переделать создание сразу в карту
	for _, xLine := range client.Units {
		for _, unit := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range client.HostileUnits {
		for _, unit := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range client.Structure {
		for _, structure := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = structure.X
			coordinate.Y = structure.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range client.HostileStructure {
		for _, structure := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = structure.X
			coordinate.Y = structure.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range Games[client.GameID].coordinate {
		for _, obstacles := range xLine {
			if obstacles.Type == "obstacle" {
				var coordinate objects.Coordinate
				coordinate.X = obstacles.X
				coordinate.Y = obstacles.Y
				coordinates = append(coordinates, &coordinate)
			}
		}
	}


	for _, obstacle := range coordinates{
		if obstaclesMatrix[obstacle.X] != nil {
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		} else {
			obstaclesMatrix[obstacle.X] = make(map[int]*objects.Coordinate)
			obstaclesMatrix[obstacle.X][obstacle.Y] = obstacle
		}
	}

	return
}
