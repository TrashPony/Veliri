package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/movePhase"
	"../../mechanics/player"
	"../../mechanics/coordinate"
	"../../mechanics/game"
	"../../mechanics/unit"
)

func SelectUnit(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games[client.GetGameID()]

	if findClient && findUnit && findGame {
		if activeGame.Phase == "move" {
			SelectMove(client, gameUnit, activeGame, ws)
		}

		if activeGame.Phase == "targeting" {
			SelectTarget(client, gameUnit, activeGame, ws)
		}
	}
}

func SelectMove(client *player.Player, gameUnit *unit.Unit, actionGame *game.Game, ws *websocket.Conn) {
	if !client.GetReady() {
		if !gameUnit.Action {
			ws.WriteJSON(MoveCoordinate{Event: "SelectMoveUnit", Move: movePhase.GetMoveCoordinate(gameUnit, client, actionGame)})
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "unit already move"})
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type MoveCoordinate struct {
	Event string                                       `json:"event"`
	Move  map[string]map[string]*coordinate.Coordinate `json:"move"`
}

func SelectTarget(client *player.Player, gameUnit *unit.Unit, actionGame *game.Game, ws *websocket.Conn) {
	/*
			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
			for _, coordinate := range coordinates {
				targetUnit, ok := client.GetHostileUnit(coordinate.X, coordinate.Y)
				if ok && targetUnit.Owner != client.GetLogin() {
					var createCoordinates = Response{Event: msg.Event, UserName: client.GetLogin(), Phase: activeGame.GetStat().Phase,
						X: targetUnit.X, Y: targetUnit.Y}
					fieldPipe <- createCoordinates
				}
			}
	*/
}

type TargetCoordinate struct {
	Event  string                                       `json:"event"`
	Target map[string]map[string]*coordinate.Coordinate `json:"target"`
}
