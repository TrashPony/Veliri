package field

import (
	"../../mechanics/localGame/Phases/movePhase"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/player"
	"../../mechanics/localGame"
	"../../mechanics/gameObjects/coordinate"
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
}

/*
TODO улучшить метод движения за счет общения бекенда и фронтенда
TODO юниты при передвежение будет говорить бекенду свои координаты
TODO и бекенд будет решать за счет этого кому из игроков говорить где и как двигается юнит
TODO тогда беда с туманом войны и баг с проебом координаты решается на все 100%
 */

func MoveUnit(msg Message, ws *websocket.Conn) {

	gameUnit, findUnit := usersFieldWs[ws].GetUnit(msg.Q, msg.R)
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

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

func SkipMoveUnit(msg Message, ws *websocket.Conn) {

	gameUnit, findUnit := usersFieldWs[ws].GetUnit(msg.Q, msg.R)
	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	if findUnit && findClient && findGame {
		if !gameUnit.Action {
			movePhase.SkipMove(gameUnit, activeGame, client)
			ws.WriteJSON(Move{Event: "MoveUnit", Unit: gameUnit, UserName: client.GetLogin()})
		} else {
			resp := ErrorMessage{Event: "MoveUnit", Error: "unit already"}
			ws.WriteJSON(resp)
		}
	} else {
		resp := ErrorMessage{Event: "MoveUnit", Error: "not found unit or game or player"}
		ws.WriteJSON(resp)
	}
}

func updateWatchHostileUser(client *player.Player, activeGame *localGame.Game, gameUnit *unit.Unit, pathNodes []*movePhase.TruePatchNode) {

	for _, user := range activeGame.GetPlayers() {
		if user.GetLogin() != client.GetLogin() {

			// пытаемся взять юнита по начальной координате
			_, okGetUnit := user.GetHostileUnitByID(gameUnit.ID)

			// если юзер видит юнита то удаляем его со строго места
			if okGetUnit {
				user.DelHostileUnit(gameUnit.ID)
			}

			// пытаемся взять юнита по конечной координате
			_, okGetEndXY := user.GetWatchCoordinate(gameUnit.Q, gameUnit.R)

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
