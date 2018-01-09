package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	activeGame, ok := Games[client.GetGameID()]
	//activeUser := ActionGameUser(Games[client.GetGameID()].GetPlayers())

	if find && ok {
		if unit.Action && !activeGame.GetUserReady(client.GetLogin()) {

			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			obstacles := game.GetObstacles(client, Games[client.GetGameID()])
			moveCoordinate := game.GetMoveCoordinate(coordinates, unit, obstacles)

			var passed bool

			for _, coordinate := range moveCoordinate {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {

					_, pathNodes := game.InitMove(unit, msg.ToX, msg.ToY, client, activeGame)
					moves := Move{Event: msg.Event, UnitX:msg.X, UnitY:msg.Y, UserName:client.GetLogin(), PathNodes: pathNodes}
					move <- moves

					passed = true
				}
			}

			if !passed {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "not allow"}
				ws.WriteJSON(resp)
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "unit already move"}
			ws.WriteJSON(resp)
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, ErrorType: "not found unit"}
		ws.WriteJSON(resp)
	}
}

func skipMoveUnit(msg FieldMessage, ws *websocket.Conn) {
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

/*func Move(unit *game.Unit, client *game.Player, updaterWatchZone *game.UpdaterWatchZone, x, y int)  {

	resp := FieldResponse{Event: "MoveUnit", UserName: client.GetLogin(), X: x, Y: y, ToX: unit.X, ToY:unit.Y}
	fieldPipe <- resp


	sendNewHostileUnit(updaterWatchZone.OpenUnit, client.GetLogin())
	sendNewHostileStructure(updaterWatchZone.OpenStructure, client.GetLogin())
	UpdateOpenCoordinate(updaterWatchZone.OpenCoordinate, updaterWatchZone.CloseCoordinate, client.GetLogin())
}*/

func updateWatchHostileUser(client game.Player, unit *game.Unit, x, y int, activeUser []*game.Player) {
	var unitsParameter InitUnit

	for _, user := range activeUser {
		if user.GetLogin() != client.GetLogin() {

			_, okGetUnit := user.GetHostileUnit(x, y)

			if okGetUnit {
				//coordinate, _ := activeGame.GetMap().GetCoordinate(x,y)
				//if find {  TODO полностью инициализировать карту
				coordinate := game.Coordinate{X: x, Y: y}
				user.AddCoordinate(&coordinate) // добавляем на место старого места юнита пустую зону
				//}
				user.DelHostileUnit(x, y)             // и удаляем в общей карте вражеских юнитов
				openCoordinate(user.GetLogin(), x, y) // и остылаем событие удаление юнита

			}

			_, okGetXY := user.GetWatchCoordinate(unit.X, unit.Y)

			if okGetXY { // если следующая клетка юнита в зоне видимости
				user.DelWatchCoordinate(unit.X, unit.Y) // удаляем пустую клетку
				user.AddHostileUnit(unit)               // и добавляем в общую карту вражеских юнитов
				unitsParameter.initUnit(unit, user.GetLogin())
			}
		}
	}
}
