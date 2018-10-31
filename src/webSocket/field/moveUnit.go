package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/movePhase"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
	"strconv"
)

type Move struct {
	Event    string                     `json:"event"`
	UserName string                     `json:"user_name"`
	GameID   int                        `json:"game_id"`
	Unit     *unit.Unit                 `json:"unit"`
	Path     []*movePhase.TruePatchNode `json:"path"`
	Error    string                     `json:"error"`
	Move     bool                       `json:"move"`
}

/*
TODO улучшить метод движения за счет общения бекенда и фронтенда
TODO юниты при передвежение будет говорить бекенду свои координаты
TODO и бекенд будет решать за счет этого кому из игроков говорить где и как двигается юнит
TODO тогда беда с туманом войны и баг с проебом координаты решается на все 100%
*/

func MoveUnit(msg Message, ws *websocket.Conn) {
	var event string

	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	gameUnit, findUnit := client.GetUnitStorage(msg.UnitID)
	if !findUnit {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
	} else {
		event = "SelectStorageUnit"
	}

	if findUnit && findClient && findGame {
		if !client.GetReady() && gameUnit.ActionPoints > 0 {

			moveCoordinate := movePhase.GetMoveCoordinate(gameUnit, client, activeGame, event)
			_, find := moveCoordinate[strconv.Itoa(msg.ToQ)][strconv.Itoa(msg.ToR)]

			if find {
				path := movePhase.InitMove(gameUnit, msg.ToQ, msg.ToR, client, activeGame, event)
				client.DelUnitStorage(gameUnit.ID)

				ws.WriteJSON(Move{Event: msg.Event, Unit: gameUnit, UserName: client.GetLogin(), Path: path})
				updateWatchHostileUser(client, activeGame, gameUnit, path, event)
				QueueSender(activeGame, ws)
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

func QueueSender(game *localGame.Game, ws *websocket.Conn) {

	allReady := true

	for _, user := range game.GetPlayers() {
		for _, q := range user.GetUnits() {
			for _, gameUnit := range q {
				if gameUnit.Move {
					moves := Move{Event: "QueueMove", UserName: user.GetLogin(), GameID: game.Id, Unit: gameUnit}
					move <- moves
				}
			}
			if !user.Ready {
				allReady = false
			}
		}

		for _, gameUnit := range user.GetUnitsStorage() {
			if gameUnit.Move {
				moves := Move{Event: "QueueMove", UserName: user.GetLogin(), GameID: game.Id, Unit: gameUnit}
				move <- moves
			}
		}
	}

	if allReady {
		Ready(ws)
	}
}

func SkipMoveUnit(msg Message, ws *websocket.Conn) {

	gameUnit, findUnit := usersFieldWs[ws].GetUnit(msg.Q, msg.R)
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	if findUnit && findClient && findGame {
		movePhase.SkipMove(gameUnit, activeGame, client)
		ws.WriteJSON(Move{Event: "MoveUnit", Unit: gameUnit, UserName: client.GetLogin()})

		QueueSender(activeGame, ws)
	} else {
		resp := ErrorMessage{Event: "MoveUnit", Error: "not found unit or game or player"}
		ws.WriteJSON(resp)
	}
}

func updateWatchHostileUser(client *player.Player, activeGame *localGame.Game, gameUnit *unit.Unit, pathNodes []*movePhase.TruePatchNode, event string) {

	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {

			// пытаемся взять юнита по начальной координате
			_, okGetUnit := user.GetHostileUnitByID(gameUnit.ID)

			// если юзер видит юнита то удаляем его со строго места
			if okGetUnit {
				user.DelHostileUnit(gameUnit.ID)
			}

			// пытаемся взять юнита по конечной координате
			_, okGetEndQR := user.GetWatchCoordinate(gameUnit.Q, gameUnit.R)

			// если конечная точка пути видима то добавляем юнита
			if okGetEndQR {
				user.AddHostileUnit(gameUnit)
			}

			send := false
			okSecondNode := false
			okEarlyNode := false
			// тут происходит формирование пути для пользователя который может видеть не весь путь юнита
			for i, pathNode := range pathNodes {

				pathNode.WatchNode = nil

				firstNode, okFirstNode := user.GetWatchCoordinate(pathNode.PathNode.Q, pathNode.PathNode.R)

				if len(pathNodes) > i+1 {
					_, okSecondNode = user.GetWatchCoordinate(pathNodes[i+1].PathNode.Q, pathNodes[i+1].PathNode.R)
				}
				if 0 < i {
					_, okEarlyNode = user.GetWatchCoordinate(pathNodes[i-1].PathNode.Q, pathNodes[i-1].PathNode.R)
				}

				if event == "SelectStorageUnit" && i == 0 && okFirstNode {
					// если машина выходит из мазер шипа говорим что как бе он вышел из тумана что бы она появилась на карте)
					firstNode.Type = "outFog"
				}

				// если юзер не видит координату то скрваем ее
				if !okFirstNode {
					var fakeNode coordinate.Coordinate

					if okSecondNode {
						fakeNode.Type = "outFog"
						fakeNode.Q = pathNode.PathNode.Q
						fakeNode.R = pathNode.PathNode.R
					} else {
						if (okGetUnit && i == 0) || okEarlyNode {
							fakeNode.Type = "inToFog"
							fakeNode.Q = pathNode.PathNode.Q
							fakeNode.R = pathNode.PathNode.R
						} else {
							fakeNode.Type = "hide"
							fakeNode.Q = 0
							fakeNode.R = 0
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
