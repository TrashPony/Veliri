package field

import (
	"../../mechanics/gameObjects/coordinate"
	"../../mechanics/gameObjects/unit"
	"../../mechanics/localGame"
	"../../mechanics/localGame/Phases/movePhase"
	"../../mechanics/player"
	"github.com/gorilla/websocket"
)

func SelectUnit(msg Message, ws *websocket.Conn) {
	var findUnit bool
	var gameUnit *unit.Unit

	client, findClient := usersFieldWs[ws]
	activeGame, findGame := Games.Get(client.GetGameID())

	if msg.Event == "SelectStorageUnit" {
		// т.к. юнит в корпусе берем координаты мс и присваиваем их юниту.
		gameUnit, findUnit = client.GetUnitStorage(msg.UnitID)
		if findUnit {
			gameUnit.Q = client.GetSquad().MatherShip.Q
			gameUnit.R = client.GetSquad().MatherShip.R
			SelectMove(client, gameUnit, activeGame, ws, msg.Event)
		}
	} else {
		gameUnit, findUnit = client.GetUnit(msg.Q, msg.R)
	}

	if findClient && findUnit && findGame {
		if activeGame.Phase == "move" {
			SelectMove(client, gameUnit, activeGame, ws, msg.Event)
		}
	}
}

func SelectMove(client *player.Player, gameUnit *unit.Unit, actionGame *localGame.Game, ws *websocket.Conn, event string) {
	if !client.GetReady() {
		if gameUnit.ActionPoints > 0 && gameUnit.Move {
			ws.WriteJSON(MoveCoordinate{Event: "SelectMoveUnit", Unit: gameUnit, Move: movePhase.GetMoveCoordinate(gameUnit, client, actionGame, event)})
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
