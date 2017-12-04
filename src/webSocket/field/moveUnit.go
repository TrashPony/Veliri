package field

import (
	"../../game"
	"errors"
	"github.com/gorilla/websocket"
	"time"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	if find && ok {
		if unit.Action {

			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			obstacles := game.GetObstacles(client, Games[client.GetGameID()])
			moveCoordinate := game.GetMoveCoordinate(coordinates, unit, obstacles)

			var passed bool
			for _, coordinate := range moveCoordinate {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
					resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin()}
					fieldPipe <- resp
					go InitMove(unit, msg, client)
					passed = true
				}
			}

			if !passed {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "not allow"}
				fieldPipe <- resp
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "unit already move"}
			fieldPipe <- resp
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "not found unit"}
		fieldPipe <- resp
	}
}

func InitMove(unit *game.Unit, msg FieldMessage, client *game.Player) {

	idGame := client.GetGameID()
	toX := msg.ToX
	toY := msg.ToY

	for {
		obstacles := game.GetObstacles(client, Games[idGame])

		start := game.Coordinate{X: unit.X, Y: unit.Y}
		end := game.Coordinate{X: toX, Y: toY}

		mp := Games[client.GetGameID()].GetMap()

		path := game.FindPath(mp, start, end, obstacles)

		x, y, errorMove := Move(unit, path, client, end)
		if errorMove != nil {
			if errorMove.Error() != "cell is busy" {
				queue := game.MoveUnit(idGame, unit, x, y)
				unit.Queue = queue
				break
			}
		} else {
			queue := game.MoveUnit(idGame, unit, x, y)
			unit.Queue = queue
			break
		}
	}
}

func Move(unit *game.Unit, path []game.Coordinate, client *game.Player, end game.Coordinate) (int, int, error) {

	activeGame := Games[client.GetGameID()]
	players := Games[client.GetGameID()].GetPlayers()
	activeUser := ActionGameUser(players)

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := client.GetHostileUnit(end.X,end.Y)
			if ok {
				unit.Action = false
				var unitsParameter InitUnit
				unitsParameter.initUnit(unit, client.GetLogin())
				return unit.X, unit.Y, errors.New("end cell is busy")
			}
		} else {
			_, ok := client.GetHostileUnit(pathNode.X, pathNode.Y)
			if ok {
				return 0, 0, errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		activeGame.DelUnit(unit) // Удаляем юнита со старых позиций
		client.DelUnit(unit.X, unit.Y)

		x := unit.X
		y := unit.Y

		unit.X = pathNode.X
		unit.Y = pathNode.Y

		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			unit.Action = false
		}

		activeGame.SetUnit(unit)
		client.AddUnit(unit) // добавляем новую позицию юнита

		UpdateWatchZone(client, activeGame)  // отправляем открытые ячейки, удаляем закрытые
		go updateWatchHostileUser(*client, *unit, x, y, activeUser)  // добавляем и удаляем нашего юнита у врагов на карте

		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.GetLogin()) // отсылаем новое место юнита

		time.Sleep(200 * time.Millisecond)
	}

	return unit.X, unit.Y, nil
}

func updateWatchHostileUser(client game.Player, unit game.Unit, x, y int, activeUser []*game.Player) {
	var unitsParameter InitUnit

	for _, user := range activeUser {
		if user.GetLogin() != client.GetLogin() {

			_, okGetUnit := user.GetHostileUnit(x,y)

			if okGetUnit {
				//coordinate, _ := activeGame.GetMap().GetCoordinate(x,y)
				//if find {  TODO полностью инициализировать карту
				coordinate := game.Coordinate{X: x, Y: y}
				user.AddCoordinate(&coordinate) // добавлдяем на место старого места юнита пустую зону
				//}
				user.DelHostileUnit(x,y) 		                   // и удаляем в общей карте вражеских юнитов
				openCoordinate(user.GetLogin(), x, y)              // и остылаем событие удаление юнита

			}

			_, okGetXY := user.GetWatchCoordinate(unit.X, unit.Y)

			if okGetXY {                                 // если следующая клетка юнита в зоне видимости
				user.DelWatchCoordinate(unit.X, unit.Y)  // удаляем пустую клетку
				user.AddHostileUnit(&unit)               // и добавляем в общую карту вражеских юнитов
				unitsParameter.initUnit(&unit, user.GetLogin())
			}
		}
	}
}

