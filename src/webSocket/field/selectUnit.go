package field

import (
	"websocket-master"
	"strconv"
	"../../game/objects"
	"../../game/mechanics"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].Units[strconv.Itoa(msg.X) + ":" + strconv.Itoa(msg.Y)]
	if find {
		respawn := usersFieldWs[ws].Respawn
		if usersFieldWs[ws].GameStat.Phase == "move" {
			coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			unitsCoordinate := objects.GetUnitsCoordinate(usersFieldWs[ws].Units)
			hostileCoordinate := objects.GetUnitsCoordinate(usersFieldWs[ws].HostileUnits)

			for _, hostile := range hostileCoordinate {
				unitsCoordinate = append(unitsCoordinate, hostile)
			}

			responseCoordinate := subtraction(coordinates, unitsCoordinate)

			for i := 0; i < len(responseCoordinate); i++ {
				if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
					var createCoordinates= FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase,
						X: responseCoordinate[i].X, Y: responseCoordinate[i].Y}
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
