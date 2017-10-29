package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
	"time"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := findUnit(msg, ws)
	idGame := usersFieldWs[ws].GameStat.Id
	if find {
		respawn := usersFieldWs[ws].Respawn
		coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
		toX := msg.ToX
		toY := msg.ToY
		var passed bool
		for _, coordinate := range coordinates{
			if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
					obstacles := objects.GetUnitsCoordinate(unit.WatchUnit) // TODO: добавить еще не проходимые учатки когда добавлю непроходимые участки
					start := objects.Coordinate{ X:unit.X, Y:unit.Y }
					end := objects.Coordinate{ X: toX, Y: toY}
					path := mechanics.FindPath(usersFieldWs[ws].Map, start, end, obstacles)
					go Move(unit, path, idGame, msg, ws)
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

func Move(unit *objects.Unit, path []objects.Coordinate, idGame int, msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse
	for _, pathNode := range path {
		x := unit.X
		y := unit.Y
		toX,toY, err := mechanics.MoveUnit(idGame, unit, pathNode.X, pathNode.Y)
		if err != nil {
			println(err.Error())
			break
		} else {
			// TODO изменить юнита в бд присвоить ему координаты и выдать их в пул

			unit.X = toX
			unit.Y = toY
			unit.Watch, unit.WatchUnit, err = sendPermissionCoordinates(idGame, ws, unit)
			if err != nil {
				break
			}

			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: x, Y: y, ToX: toX, ToY: toY}
			fieldPipe <- resp
			time.Sleep(1000 * time.Millisecond)
		}
	}
}