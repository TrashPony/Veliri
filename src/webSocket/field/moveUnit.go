package field

import (
	"websocket-master"
	"../../game/objects"
	"../../game/mechanics"
	"strconv"
	"errors"
	"time"
)

func MoveUnit(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse

	unit, find := findUnit(msg, ws)
	if find {
		respawn := usersFieldWs[ws].Respawn
		coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
		var passed bool
		for _, coordinate := range coordinates{
			if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
				if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
					go InitMove(unit, msg, ws)
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

func InitMove(unit *objects.Unit, msg FieldMessage, ws *websocket.Conn )  {

	idGame := usersFieldWs[ws].GameStat.Id
	toX := msg.ToX
	toY := msg.ToY
	for {
		obstacles := objects.GetUnitsCoordinate(unit.WatchUnit) // TODO: добавить еще не проходимые учатки когда добавлю непроходимые участки
		start := objects.Coordinate{X: unit.X, Y: unit.Y}
		end := objects.Coordinate{X: toX, Y: toY}
		path := mechanics.FindPath(usersFieldWs[ws].Map, start, end, obstacles)
		errorMove := Move(unit, path, idGame, msg, ws, end)
		if errorMove == nil || errorMove.Error() == "end cell is busy" {
			break
		}
	}
}

func Move(unit *objects.Unit, path []objects.Coordinate, idGame int, msg FieldMessage, ws *websocket.Conn, end objects.Coordinate) (error) {
	var resp FieldResponse
	for _, pathNode := range path {
		x := unit.X
		y := unit.Y

		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := unit.WatchUnit[strconv.Itoa(end.X) + ":" + strconv.Itoa(end.Y)]
			if ok {
				return errors.New("end cell is busy") // если конечная клетка занята то оставновиться перед ней
			}
		} else {
			_, ok := unit.WatchUnit[strconv.Itoa(pathNode.X) + ":" + strconv.Itoa(pathNode.Y)]
			if ok {
				return errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		toX,toY, err := mechanics.MoveUnit(idGame, unit, pathNode.X, pathNode.Y)
		if err != nil {
			println(err.Error())
			break
		}

		unit.X = toX
		unit.Y = toY

		for _, unitWatch := range usersFieldWs[ws].Units {
			// удаляем юнита со старого места и добавляем в новое у всех союзных юнитов
			_, ok := unitWatch.WatchUnit[strconv.Itoa(x) + ":" + strconv.Itoa(y)]
			if ok {
				delete(unitWatch.WatchUnit, strconv.Itoa(x) + ":" + strconv.Itoa(y))
			}
			unitWatch.WatchUnit[strconv.Itoa(toX) + ":" + strconv.Itoa(toY)] = unit
		}

		unit.Watch, unit.WatchUnit, err = sendPermissionCoordinates(idGame, ws, unit) //TODO очень медленная функция

		if err != nil {
			break
		}

		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: x, Y: y, ToX: toX, ToY: toY}
		fieldPipe <- resp
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}