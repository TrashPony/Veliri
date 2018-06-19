package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics/Phases/movePhase"
	"../../mechanics/Phases/targetPhase"
	"../../mechanics/player"
	"../../mechanics/coordinate"
	"../../mechanics/game"
	"../../mechanics/unit"
)

func SelectUnit(msg Message, ws *websocket.Conn) {

	client, findClient := usersFieldWs[ws]
	gameUnit, findUnit := client.GetUnit(msg.X, msg.Y)
	activeGame, findGame := Games.Get(client.GetGameID())

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
			ws.WriteJSON(MoveCoordinate{Event: "SelectMoveUnit", Unit: gameUnit, Move: movePhase.GetMoveCoordinate(gameUnit, client, actionGame)})
		} else {
			ws.WriteJSON(ErrorMessage{Event: "Error", Error: "unit already move"})
		}
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type MoveCoordinate struct {
	Event string                                       `json:"event"`
	Unit  *unit.Unit                                   `json:"unit"`
	Move  map[string]map[string]*coordinate.Coordinate `json:"move"`
}

func SelectTarget(client *player.Player, gameUnit *unit.Unit, actionGame *game.Game, ws *websocket.Conn) {
	if !client.GetReady() {
		ws.WriteJSON(TargetCoordinate{Event: "GetTargets", Unit: gameUnit, Targets: targetPhase.GetTargetCoordinate(gameUnit, actionGame)})
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "you ready"})
	}
}

type TargetCoordinate struct {
	Event   string                                       `json:"event"`
	Unit    *unit.Unit                                   `json:"unit"`
	Targets map[string]map[string]*coordinate.Coordinate `json:"targets"`
}
