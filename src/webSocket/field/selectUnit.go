package field

import (
	"github.com/gorilla/websocket"
	"../../mechanics"
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
	}

	/*if find && ok && !activeGame.GetUserReady(client.GetLogin()) {



				for i := 0; i < len(moveCoordinate); i++ {
					if !(moveCoordinate[i].X == respawn.X && moveCoordinate[i].Y == respawn.Y) && moveCoordinate[i].X >= 0 && moveCoordinate[i].Y >= 0 && moveCoordinate[i].X < 10 && moveCoordinate[i].Y < 10 {
						var createCoordinates = Response{Event: msg.Event, UserName: client.GetLogin(), Phase: activeGame.GetStat().Phase, // TODO не до 10ти а до края карты
							X: moveCoordinate[i].X, Y: moveCoordinate[i].Y}
						fieldPipe <- createCoordinates
					}
				}



		if activeGame.GetStat().Phase == "targeting" {
			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
			for _, coordinate := range coordinates {
				targetUnit, ok := client.GetHostileUnit(coordinate.X, coordinate.Y)
				if ok && targetUnit.Owner != client.GetLogin() {
					var createCoordinates = Response{Event: msg.Event, UserName: client.GetLogin(), Phase: activeGame.GetStat().Phase,
						X: targetUnit.X, Y: targetUnit.Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	} else {
		if activeGame.GetStat().Phase == "Init" {
			var coordinates []*game.Coordinate

			for _, coordinate := range usersFieldWs[ws].GetCreateZone() {
				_, ok := client.GetUnit(coordinate.X, coordinate.Y)
				if !ok {
					coordinates = append(coordinates, coordinate)
				}
			}

			for i := 0; i < len(coordinates); i++ {
				if !(coordinates[i].X == respawn.X && coordinates[i].Y == respawn.Y) {
					var createCoordinates = Response{Event: "SelectCoordinateCreate", UserName: usersFieldWs[ws].GetLogin(), X: coordinates[i].X, Y: coordinates[i].Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	}*/
}

func SelectMove(client *player.Player, gameUnit *unit.Unit, actionGame *game.Game, ws *websocket.Conn) {
	if !gameUnit.Action {
		ws.WriteJSON(MoveCoordinate{Event: "SelectMoveUnit", Move: mechanics.GetMoveCoordinate(gameUnit, client, actionGame)})
	} else {
		ws.WriteJSON(ErrorMessage{Event: "Error", Error: "unit already move"})
	}
}

type MoveCoordinate struct {
	Event string                                       `json:"event"`
	Move  map[string]map[string]*coordinate.Coordinate `json:"move"`
}
