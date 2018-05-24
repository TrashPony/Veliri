package field

import (
	"strconv"
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/movePhase"
	"../../mechanics/unit"
	"../../mechanics/coordinate"
	"../../mechanics/watchZone"
)

type Move struct {
	Event     string                                 `json:"event"`
	UserName  string                                 `json:"user_name"`
	Unit      *unit.Unit                             `json:"unit"`
	PathNodes []*coordinate.Coordinate                `json:"path_nodes"`
	WatchNode map[string]*watchZone.UpdaterWatchZone `json:"watch_node"`
	Error     string                                 `json:"error"`
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
				watchNodes, pathNodes := movePhase.InitMove(gameUnit, msg.ToX, msg.ToY, client, activeGame)

				println(watchNodes)
				println(pathNodes)

				moves := Move{Event: msg.Event, Unit: gameUnit, UserName: client.GetLogin(), PathNodes: pathNodes, WatchNode: watchNodes}
				move <- moves

				// todo updateWatchHostileUser(msg, client, activeUser, unit, pathNodes)
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

/*func skipMoveUnit(msg Message, ws *websocket.Conn) {
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
}

func updateWatchHostileUser(msg Message, client *Mechanics.Player, activeUser []*Mechanics.Player, unit *Mechanics.Unit, pathNodes []Mechanics.Coordinate) {

	for _, user := range activeUser {
		if user.GetLogin() != client.GetLogin() {

			var truePath []Mechanics.Coordinate

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
					var fakeNode Mechanics.Coordinate
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
}*/
