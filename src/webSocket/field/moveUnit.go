package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"time"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	activeGame, ok := Games[client.GetGameID()]
	players := Games[client.GetGameID()].GetPlayers()
	activeUser := ActionGameUser(players)

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

					truePath, pathNodes := game.InitMove(unit, msg.ToX, msg.ToY, client, activeGame)

					for i, pathNode := range pathNodes {
						updateWatch, ok := truePath[pathNode]

						unit.X = pathNode.X
						unit.Y = pathNode.Y

						if ok && i == 0 {
							UpdateWatchZone(client, activeGame, updateWatch)
							go updateWatchHostileUser(*client, unit, msg.X, msg.Y, activeUser)
						}

						if ok && i > 0 {
							UpdateWatchZone(client, activeGame, updateWatch)
							go updateWatchHostileUser(*client, unit, pathNodes[i-1].X, pathNodes[i-1].Y, activeUser)
						}

						time.Sleep(200 * time.Millisecond)
					}
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

func skipMoveUnit(msg FieldMessage, ws *websocket.Conn)  {
	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	activeGame, ok := Games[client.GetGameID()]

	if find && ok {
		if unit.Action {
			unit.Action = false

			queue := game.MoveUnit(activeGame.GetStat().Id, unit, unit.X, unit.Y)
			unit.Queue = queue

			var unitsParameter InitUnit
			unitsParameter.initUnit(unit, client.GetLogin())
		}
	}
}

func updateWatchHostileUser(client game.Player, unit *game.Unit, x, y int, activeUser []*game.Player) {
	var unitsParameter InitUnit

	for _, user := range activeUser {
		if user.GetLogin() != client.GetLogin() {

			_, okGetUnit := user.GetHostileUnit(x,y)

			if okGetUnit {
				//coordinate, _ := activeGame.GetMap().GetCoordinate(x,y)
				//if find {  TODO полностью инициализировать карту
				coordinate := game.Coordinate{X: x, Y: y}
				user.AddCoordinate(&coordinate) // добавляем на место старого места юнита пустую зону
				//}
				user.DelHostileUnit(x,y) 		                   // и удаляем в общей карте вражеских юнитов
				openCoordinate(user.GetLogin(), x, y)              // и остылаем событие удаление юнита

			}

			_, okGetXY := user.GetWatchCoordinate(unit.X, unit.Y)

			if okGetXY {                                 // если следующая клетка юнита в зоне видимости
				user.DelWatchCoordinate(unit.X, unit.Y)  // удаляем пустую клетку
				user.AddHostileUnit(unit)               // и добавляем в общую карту вражеских юнитов
				unitsParameter.initUnit(unit, user.GetLogin())
			}
		}
	}
}