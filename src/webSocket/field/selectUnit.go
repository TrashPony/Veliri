package field

import (
	"../../game/mechanics"
	"../../game/objects"
	"github.com/gorilla/websocket"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client, ok := usersFieldWs[ws]
	game, ok := Games[client.GameID]

	if find && ok {
		respawn := client.Respawn
		if game.Stat.Phase == "move" {
			if unit.Action {
				coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
				unitsCoordinate := objects.GetUnitsCoordinate(client.Units)
				hostileCoordinate := objects.GetUnitsCoordinate(client.HostileUnits)

				for _, hostile := range hostileCoordinate {
					unitsCoordinate = append(unitsCoordinate, hostile)
				}

				responseCoordinate := subtraction(coordinates, unitsCoordinate)

				for i := 0; i < len(responseCoordinate); i++ {
					if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
						var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: game.Stat.Phase,
							X: responseCoordinate[i].X, Y: responseCoordinate[i].Y}
						fieldPipe <- createCoordinates
					}
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: client.Login, Error: "unit already move"}
				fieldPipe <- resp
			}
		}

		if game.Stat.Phase == "targeting" {
			coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.RangeAttack)
			for _, coordinate := range coordinates {
				targetUnit, ok := client.HostileUnits[coordinate.X][coordinate.Y]
				if ok && targetUnit.NameUser != client.Login {
					var createCoordinates = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: game.Stat.Phase,
						X: targetUnit.X, Y: targetUnit.Y}
					fieldPipe <- createCoordinates
				}
			}
		}
	}
}
