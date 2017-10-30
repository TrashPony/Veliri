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

	unit, find := usersFieldWs[ws].Units[strconv.Itoa(msg.X) + ":" + strconv.Itoa(msg.Y)]
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
		if errorMove == nil || errorMove.Error() == "end cell is busy" || errorMove.Error() == "connect is lost"{
			break
		}
	}
}

func Move(unit *objects.Unit, path []objects.Coordinate, idGame int, msg FieldMessage, ws *websocket.Conn, end objects.Coordinate) (error) {
	var resp FieldResponse

	client, ok := usersFieldWs[ws]
	if !ok {
		return errors.New("connect is lost")
	}

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := unit.WatchUnit[strconv.Itoa(end.X)+":"+strconv.Itoa(end.Y)]
			if ok {
				return errors.New("end cell is busy")
			}
		} else {
			_, ok := unit.WatchUnit[strconv.Itoa(pathNode.X)+":"+strconv.Itoa(pathNode.Y)]
			if ok {
				return errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		x := unit.X
		y := unit.Y

		unit.X = pathNode.X
		unit.Y = pathNode.Y

		oldWatchZone := unit.WatchUnit

		mechanics.MoveUnit(idGame, unit, pathNode.X, pathNode.Y)
		units := objects.GetAllUnits(client.GameStat.Id)

		delete(client.Units,strconv.Itoa(x) + ":" + strconv.Itoa(y)) // удаляем в карте общий юнитов старое место расположение
		client.Units[strconv.Itoa(unit.X) + ":" + strconv.Itoa(unit.Y)] = unit // добавляем новое

		for _, unitWatch := range client.Units {
			var err error
			unitWatch.Watch, unitWatch.WatchUnit, err = PermissionCoordinates(*client, unitWatch, units) //обновляем зону видимости идущего моба TODO :/ не обновляются враги
			if err != nil{
				continue
			}
		}

		for _, hostile := range unit.WatchUnit { // добавляев новых открытых вражеских юнитов
			if hostile.NameUser != client.Login {
				_, ok := oldWatchZone[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)]
				if !ok {
					client.HostileUnits[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)] = hostile
					continue
				}
			} else {
				continue
			}
		}

		for _, oldHostile := range oldWatchZone { // удаляем старых закрытых вражеских юнитов
			if oldHostile.NameUser != client.Login {
				_, ok := unit.WatchUnit[strconv.Itoa(oldHostile.X)+":"+strconv.Itoa(oldHostile.Y)]
				if !ok {
					delete(client.HostileUnits, strconv.Itoa(oldHostile.X)+":"+strconv.Itoa(oldHostile.Y))
					continue
				} else {
				}
			} else {
				continue
			}
		}

		resp = FieldResponse{Event: msg.Event, UserName: client.Login, X: x, Y: y, ToX: unit.X, ToY: unit.Y}
		fieldPipe <- resp
		time.Sleep(1 * time.Millisecond)
	}
	return nil
}
