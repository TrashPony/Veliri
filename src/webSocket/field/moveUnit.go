package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]

	if find && ok {

		activeGame, okGetGame := Games[client.GetGameID()]
		activeUser := ActionGameUser(Games[client.GetGameID()].GetPlayers())

		if okGetGame && unit.Action && !activeGame.GetUserReady(client.GetLogin()) {

			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			obstacles := game.GetObstacles(client, Games[client.GetGameID()])
			moveCoordinate := game.GetMoveCoordinate(coordinates, unit, obstacles)

			var passed bool

			for _, coordinate := range moveCoordinate {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {

					watchNode, pathNodes := game.InitMove(unit, msg.ToX, msg.ToY, client, activeGame)

					moves := Move{Event: msg.Event, Unit: unit, UserName: client.GetLogin(), PathNodes: pathNodes, WatchNode: watchNode}
					move <- moves

					updateWatchHostileUser(msg ,client, activeUser, unit, pathNodes)

					passed = true
				}
			}

			if !passed {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, Error: "not allow"}
				ws.WriteJSON(resp)
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, Error: "unit already move"}
			ws.WriteJSON(resp)
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), X: msg.X, Y: msg.Y, Error: "not found unit"}
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
			unitsParameter.initUnit("InitUnit", unit, client.GetLogin())
		}
	}
}

func updateWatchHostileUser(msg FieldMessage, client *game.Player, activeUser []*game.Player, unit *game.Unit, pathNodes []game.Coordinate) {

	for _, user := range activeUser {
		if user.GetLogin() != client.GetLogin() {

			var truePath []game.Coordinate

			_, okGetXY := user.GetWatchCoordinate(unit.X, unit.Y)

			if okGetXY { // если конечная точка пути видима то добавляем юнита
				user.AddHostileUnit(unit)
			}

			tmpUnit := *unit

			_, okGetUnit := user.GetHostileUnit(pathNodes[0].X, pathNodes[0].Y) // пытаемся взять юнита по начальной координате

			if okGetUnit {
				user.DelHostileUnit(pathNodes[0].X, pathNodes[0].Y) // если юзер видит юнита то удаляем его со строго места
			} else {
				tmpUnit.X = 999
				tmpUnit.Y = 999
			}


			// тут происходит формирование пути для пользователя который может видеть не весь путь юнита
			for i := range pathNodes {
				node, okGetNode := user.GetWatchCoordinate(pathNodes[i].X, pathNodes[i].Y)
				if okGetNode {
					node.Type = "visible"
					truePath = append(truePath, *node)
				} else {
					var fakeNode game.Coordinate
					fakeNode.X = 999
					fakeNode.Y = 999
					fakeNode.Type = "hide"
					truePath = append(truePath, fakeNode)
				}
			}

			moves := Move{Event: msg.Event, Unit: &tmpUnit, UserName: user.GetLogin(), PathNodes: truePath}
			move <- moves
		}
	}
}
