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
					truePath, pathNodes := game.InitMove(unit, msg.X, msg.Y, client, activeGame)
					for i, pathNode := range pathNodes {

						updateWatch, ok := truePath[pathNode]
						if ok {
							UpdateWatchZone(client, activeGame, updateWatch)
							time.Sleep(200 * time.Millisecond)
						}

						tmpUnit := game.Unit{X: pathNode.X, Y: pathNode.Y}

						if i == 0 {
							go updateWatchHostileUser(*client, &tmpUnit, msg.X, msg.Y, activeUser)
						}

						if i > 0 && i < len(pathNodes) - 2 {
							go updateWatchHostileUser(*client, &tmpUnit, pathNodes[i-1].X, pathNodes[i-1].Y, activeUser)
						}

						if i == len(pathNodes){
							unit, find := activeGame.GetUnit(tmpUnit.X, tmpUnit.Y)
							if find {
								go updateWatchHostileUser(*client, unit, pathNodes[i-1].X, pathNodes[i-1].Y, activeUser)
							}
						}
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

func updateWatchHostileUser(client game.Player, unit *game.Unit, x, y int, activeUser []*game.Player) {
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
				user.AddHostileUnit(unit)               // и добавляем в общую карту вражеских юнитов
				unitsParameter.initUnit(unit, user.GetLogin())
			}
		}
	}
}