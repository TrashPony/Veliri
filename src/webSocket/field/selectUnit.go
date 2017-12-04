package field

import (
	"../../game"
	"github.com/gorilla/websocket"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].GetUnit(msg.X, msg.Y)
	client, ok := usersFieldWs[ws]
	activeGame, ok := Games[client.GetGameID()]
	respawn := client.GetRespawn()

	if find && ok {
		if activeGame.GetStat().Phase == "move" {
			if unit.Action {

				coordinates := game.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
				obstacles := game.GetObstacles(client, activeGame)
				moveCoordinate := game.GetMoveCoordinate(coordinates, unit, obstacles)

				for i := 0; i < len(moveCoordinate); i++ {
					if !(moveCoordinate[i].X == respawn.X && moveCoordinate[i].Y == respawn.Y) && moveCoordinate[i].X >= 0 && moveCoordinate[i].Y >= 0 {
						var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Phase: activeGame.GetStat().Phase,
							X: moveCoordinate[i].X, Y: moveCoordinate[i].Y}
						fieldPipe <- createCoordinates
					}
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Error: "unit already move"}
				fieldPipe <- resp
			}
		}

		if activeGame.GetStat().Phase == "targeting" {
			coordinates := game.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
			for _, coordinate := range coordinates {
				targetUnit, ok := client.GetHostileStructure(coordinate.X,coordinate.Y)
				if ok && targetUnit.NameUser != client.GetLogin() {
					var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Phase: activeGame.GetStat().Phase,
						X: targetUnit.X, Y: targetUnit.Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	} else {
		if activeGame.GetStat().Phase == "Init"{
			var coordinates []*game.Coordinate

			for _, coordinate := range usersFieldWs[ws].GetCreateZone() {
				_, ok := client.GetUnit(coordinate.X, coordinate.Y)
				if !ok {
					coordinates = append(coordinates, coordinate)
				}
			}

			for i := 0; i < len(coordinates); i++ {
				if !(coordinates[i].X == respawn.X && coordinates[i].Y == respawn.Y) {
					var createCoordinates = FieldResponse{Event: "SelectCoordinateCreate", UserName: usersFieldWs[ws].GetLogin(), X: coordinates[i].X, Y: coordinates[i].Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	}
}


