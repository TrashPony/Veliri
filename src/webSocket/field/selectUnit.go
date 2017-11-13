package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
)

func SelectUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]

	if find {
		respawn := usersFieldWs[ws].Respawn
		if usersFieldWs[ws].GameStat.Phase == "move" {
			if unit.Action {
				coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
				unitsCoordinate := objects.GetUnitsCoordinate(usersFieldWs[ws].Units)
				hostileCoordinate := objects.GetUnitsCoordinate(usersFieldWs[ws].HostileUnits)

				for _, hostile := range hostileCoordinate {
					unitsCoordinate = append(unitsCoordinate, hostile)
				}

				responseCoordinate := subtraction(coordinates, unitsCoordinate)

				for i := 0; i < len(responseCoordinate); i++ {
					if !(responseCoordinate[i].X == respawn.X && responseCoordinate[i].Y == respawn.Y) {
						var createCoordinates = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase,
							X: responseCoordinate[i].X, Y: responseCoordinate[i].Y}
						fieldPipe <- createCoordinates
					}
				}
			} else {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Error: "unit already move"}
				fieldPipe <- resp
			}
		}

		if usersFieldWs[ws].GameStat.Phase == "targeting" {
			// TODO атака может быть дальше чем видимость.

			for _, xLine := range unit.WatchUnit {
				for _, targetUnit := range xLine {
					if targetUnit.NameUser != usersFieldWs[ws].Login {
						var createCoordinates = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase,
							X: targetUnit.X, Y: targetUnit.Y}
						fieldPipe <- createCoordinates
					}
				}
			}

			/*unit, find := usersFieldWs[ws].Units[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
			if find {
				unitTarget, ok := unit.WatchUnit[strconv.Itoa(msg.X)+":"+strconv.Itoa(msg.Y)]
				if ok {
					var createCoordinates = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: usersFieldWs[ws].GameStat.Phase,
						X: unitTarget.X, Y: unitTarget.Y}
					fieldPipe <- createCoordinates
				}
			}*/
		}
	}
}
