package field

import (
	"strconv"
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/movePhase"
	"../../mechanics/unit"
	"../../mechanics/game"
	"../../mechanics/player"
	"../../mechanics/coordinate"
)

type Move struct {
	Event    string                     `json:"event"`
	UserName string                     `json:"user_name"`
	GameID   int                        `json:"game_id"`
	Unit     *unit.Unit                 `json:"unit"`
	Path     []*movePhase.TruePatchNode `json:"path"`
	Error    string                     `json:"error"`
}

func MoveUnit(msg Message, ws *websocket.Conn) {

	gameUnit, findUnit := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games[client.GetGameID()]

	if findUnit && findClient && findGame {
		if !gameUnit.Action && !client.GetReady() {

			moveCoordinate := movePhase.GetMoveCoordinate(gameUnit, client, activeGame)
			_, find := moveCoordinate[strconv.Itoa(msg.ToX)][strconv.Itoa(msg.ToY)]

			if find {
				path := movePhase.InitMove(gameUnit, msg.ToX, msg.ToY, client, activeGame)

				ws.WriteJSON(Move{Event: msg.Event, Unit: gameUnit, UserName: client.GetLogin(), Path: path})

				updateWatchHostileUser(client, activeGame, gameUnit, path)
			} else {
				resp := ErrorMessage{Event: msg.Event, Error: "not allow"}
				ws.WriteJSON(resp)
			}
		} else {
			resp := ErrorMessage{Event: msg.Event, Error: "unit already move"}
			ws.WriteJSON(resp)
		}
	} else {
		resp := ErrorMessage{Event: msg.Event, Error: "not found unit"}
		ws.WriteJSON(resp)
	}
}

/*func skipMoveUnit(msg Message, ws *websocket.Conn) { // todo
	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	activeGame, ok := Games[client.GetGameID()]

	if find && ok {
		if unit.Action {
			unit.Action = false

			queue := Mechanics.MoveUnit(activeGame.GetStat().Id, unit, unit.X, unit.Y)
			unit.Queue = queue

			var unitsParameter InitUnit
			unitsParameter.initUnit("InitUnit", unit, client.GetLogin())
		}
	}
}*/

func updateWatchHostileUser(client *player.Player, activeGame *game.Game, gameUnit *unit.Unit, pathNodes []*movePhase.TruePatchNode) {

	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {

			// пытаемся взять юнита по начальной координате
			_, okGetUnit := user.GetHostileUnitByID(gameUnit.Id)

			// если юзер видит юнита то удаляем его со строго места
			if okGetUnit {
				user.DelHostileUnit(gameUnit.Id)
			}

			// пытаемся взять юнита по конечной координате
			_, okGetEndXY := user.GetWatchCoordinate(gameUnit.X, gameUnit.Y)

			// если конечная точка пути видима то добавляем юнита
			if okGetEndXY {
				user.AddHostileUnit(gameUnit)
			}

			send := false
			okSecondNode := false
			okEarlyNode := false
			// тут происходит формирование пути для пользователя который может видеть не весь путь юнита
			for i, pathNode := range pathNodes {
				pathNode.WatchNode = nil

				_, okFirstNode := user.GetWatchCoordinate(pathNode.PathNode.X, pathNode.PathNode.Y)

				if len(pathNodes) > i+1 {
					_, okSecondNode = user.GetWatchCoordinate(pathNodes[i+1].PathNode.X, pathNodes[i+1].PathNode.Y)
				}
				if 0 < i {
					_, okEarlyNode = user.GetWatchCoordinate(pathNodes[i-1].PathNode.X, pathNodes[i-1].PathNode.Y)
				}

				// если юзер не видит координату то скрваем ее
				if !okFirstNode {
					var fakeNode coordinate.Coordinate

					if okSecondNode {
						fakeNode.Type = "outFog"
						fakeNode.X = pathNode.PathNode.X
						fakeNode.Y = pathNode.PathNode.Y
					} else {
						if (okGetUnit && i == 0) || okEarlyNode {
							fakeNode.Type = "inToFog"
							fakeNode.X = pathNode.PathNode.X
							fakeNode.Y = pathNode.PathNode.Y
						} else {
							fakeNode.Type = "hide"
							fakeNode.X = 0
							fakeNode.Y = 0
						}
					}

					pathNode.PathNode = &fakeNode
				} else {
					send = true
				}
			}

			// отправляем только тем кто видит хотя бы 1 клетку пути
			if send || okGetUnit {
				moves := Move{Event: "HostileUnitMove", Unit: gameUnit, UserName: user.GetLogin(), GameID: activeGame.Id, Path: pathNodes}
				move <- moves
			}
		}
	}
}
