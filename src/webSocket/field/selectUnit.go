package field

import (
	"websocket-master"
	"strconv"
	"../../game/objects"
	"../../game/mechanics"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse

	unit, find := findUnit(msg, ws)

	if find {
		respawn := usersFieldWs[ws].Respawn

		if usersFieldWs[ws].GameStat.Phase == "move" {
			coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			unitsCoordinate := objects.GetUnitsCoordinate(usersFieldWs[ws].Units)
			responseCoordinate := subtraction(coordinates, unitsCoordinate)

			for i := 0; i < len(responseCoordinate); i++ {
				if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
					var createCoordinates= FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase,
						X: strconv.Itoa(responseCoordinate[i].X), Y: strconv.Itoa(responseCoordinate[i].Y)}
					fieldPipe <- createCoordinates
				}
			}
		}
		if usersFieldWs[ws].GameStat.Phase == "targeting" {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase, RangeAttack: strconv.Itoa(unit.RangeAttack)}
			fieldPipe <- resp
		}
	}
}
