package field

import (
	"websocket-master"
	"strconv"
	"../../game/objects"
	"../../game/mechanics"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := findUnit(msg, ws)

	if find {
		respawn := usersFieldWs[ws].Respawn
		coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
		toX, _  := strconv.Atoi(msg.ToX)
		toY, _  := strconv.Atoi(msg.ToY)
		var passed bool
		for _, coordinate := range coordinates{
			if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
				if strconv.Itoa(coordinate.X) == msg.ToX && strconv.Itoa(coordinate.Y) == msg.ToY {
					obstacles := objects.GetUnitsCoordinate(unit.WatchUnit) // TODO: добавить еще не проходимые учатки когда добавлю непроходимые участки
					start := objects.Coordinate{X:unit.X, Y:unit.Y}
					end := objects.Coordinate{ X: toX, Y: toY}
					path := mechanics.FindPath(usersFieldWs[ws].Map, start, end, obstacles)
					for _, pathNode :=range path {
						println(pathNode.X, pathNode.Y)
					}
					passed = true
				}
			}
		}

		if !passed {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
			fieldPipe <- resp
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not found unit"}
		fieldPipe <- resp
	}
}