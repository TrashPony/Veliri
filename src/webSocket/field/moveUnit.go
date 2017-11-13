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

	unit, find := usersFieldWs[ws].Units[msg.X][msg.Y]
	client, ok := usersFieldWs[ws]
	if find && ok {
		if unit.Action {
			respawn := client.Respawn
			coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			var passed bool
			for _, coordinate := range coordinates {
				if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
					if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
						resp = FieldResponse{Event: msg.Event, UserName: client.Login}
						fieldPipe <- resp
						go InitMove(unit, msg, client)
						passed = true
					}
				}
			}

			if !passed {
				resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
				fieldPipe <- resp
			}
		} else {
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "unit already move"}
			fieldPipe <- resp
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not found unit"}
		fieldPipe <- resp
	}
}

func InitMove(unit *objects.Unit, msg FieldMessage, client *Clients )  {

	idGame := client.GameStat.Id
	toX := msg.ToX
	toY := msg.ToY

	for {
		obstacles := getObstacles(client)

		start := objects.Coordinate{X: unit.X, Y: unit.Y}
		end := objects.Coordinate{X: toX, Y: toY}
		path := mechanics.FindPath(client.Map, start, end, obstacles)
		x, y, errorMove := Move(unit, path, idGame, msg, client, end)
		if errorMove != nil {
			if errorMove.Error() != "cell is busy" {
				mechanics.MoveUnit(idGame, unit, x, y)
				break
			}
		} else {
			mechanics.MoveUnit(idGame, unit, x, y)
			break
		}
	}
}

func Move(unit *objects.Unit, path []objects.Coordinate, idGame int, msg FieldMessage, client *Clients, end objects.Coordinate) (int, int, error) {

	units := objects.GetAllUnits(client.GameStat.Id)
	activeUser := ActionGameUser(client.Players)

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := client.HostileUnits[end.X][end.Y]
			if ok {
				unit.Action = false
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
					HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
				initUnit <- unitsParametr
				return unit.X, unit.Y, errors.New("end cell is busy")
			}
		} else {
			_, ok := client.HostileUnits[pathNode.X][pathNode.Y]
			if ok {
				return 0,0, errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		x := unit.X
		y := unit.Y

		unit.X = pathNode.X
		unit.Y = pathNode.Y

		if (end.X == pathNode.X) && (end.Y == pathNode.Y){
			unit.Action = false
		}

		oldWatchZone := client.Watch
		oldWatchUnit := client.HostileUnits

		delete(units, strconv.Itoa(x) + ":" + strconv.Itoa(y))
		units[strconv.Itoa(unit.X) + ":" + strconv.Itoa(unit.Y)] = unit

		delete(client.Units[x], y)          // удаляем в карте старое место расположение юнита
		client.addUnit(unit)                // добавляем новое

		UpdateWatchZone(client, *unit, units, oldWatchZone, oldWatchUnit) // отправляем открытые ячейки, удаляем закрытые
		go UpdateWatchHostileUser(*client, *unit, x, y, activeUser)		 // добавляем и удаляем нашего юнита у врагов на карте

		time.Sleep(100 * time.Millisecond)
	}

	return unit.X, unit.Y, nil
}

func UpdateWatchZone(client *Clients, unitMove objects.Unit, units map[string]*objects.Unit, oldWatchZone map[int]map[int]*objects.Coordinate, oldWatchUnit map[int]map[int]*objects.Unit) {

	client.Watch = nil
	client.HostileUnits = nil

	client.getAllWatchObject(units)
	updateOpenCoordinate(client, oldWatchZone)
	updateHostileUnit(client, oldWatchUnit)

	unit := client.Units[unitMove.X][unitMove.Y]                                                             // отсылаем новое место юнита
	var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
		HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
	initUnit <- unitsParametr
}

func UpdateWatchHostileUser(client Clients, unit objects.Unit, x,y int, activeUser []*Clients) {
	for _, user := range activeUser {
		if user.Login != client.Login {

			_, okGetUnit := user.HostileUnits[x][y]

			if okGetUnit {
				user.Watch[x][y] = &objects.Coordinate{X: x, Y: y}                            // добавлдяем на место старого места юнита пустую зону
				delete(user.HostileUnits[x], y)                                               // и удаляем в общей карте вражеских юнитов
				resp := Coordinate{Event: "OpenCoordinate", UserName: user.Login, X: x, Y: y} // и остылаем событие удаление юнита
				coordiante <- resp
			}

			_, okGetXY := user.Watch[unit.X][unit.Y]

			if okGetXY { // если следующая клетка юнита в зоне видимости
				delete(user.Watch[unit.X], unit.Y)                                                                       // удаляем пустую клетку
				user.addHostileUnit(&unit)                                                                               // и добавляем в общую карту вражеских юнитов
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: user.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
					HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
				initUnit <- unitsParametr
			}

			if okGetUnit && !okGetXY { // если удалось взять юнита по старым параметрам и не удалось взять координату открытую
				resp := Coordinate{Event: "OpenCoordinate", UserName: user.Login, X: x, Y: y} // то остылаем событие удаление юнита
				coordiante <- resp
			}
		}
	}
}

func getObstacles(client *Clients)([]*objects.Coordinate)  { // TODO: добавить еще не проходимые учатки когда добавлю непроходимые участки
	coordinates := make([]*objects.Coordinate,0)

	for _, xLine := range client.Units {
		for _, unit := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	for _, xLine := range client.HostileUnits {
		for _, unit := range xLine {
			var coordinate objects.Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}

	return coordinates
}

func updateOpenCoordinate(client *Clients, oldWatchZone map[int]map[int]*objects.Coordinate)  {
	for _, xLine := range client.Watch { // отправляем все новые координаты, и т.к. старая клетка юнита теперь тоже является координатой то и ее тоже обновляем
		for _, newCoordinate := range xLine {
			_, ok := oldWatchZone[newCoordinate.X][newCoordinate.Y]
			if !ok {
				client.addCoordinate(newCoordinate)
				resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: newCoordinate.X, Y: newCoordinate.Y}
				coordiante <- resp
			}
		}
	}

	for _, xLine := range oldWatchZone { // удаляем старые координаты из зоны видимости
		for _, oldCoordinate := range xLine {
			find := findCoordinate(client, oldCoordinate)
			if !find {
				delete(client.Watch[oldCoordinate.X], oldCoordinate.Y)
				resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: oldCoordinate.X, Y: oldCoordinate.Y} // удаляем старое поле доступа
				coordiante <- resp
			}
		}
	}
}

func updateHostileUnit(client *Clients, oldWatchUnit map[int]map[int]*objects.Unit)  {
	for _, xLine := range client.HostileUnits { // добавляем новые вражеские юниты которых открыли
		for _, hostile := range xLine {
			_, ok := oldWatchUnit[hostile.X][hostile.Y]
			if !ok {
				client.addHostileUnit(hostile)
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: hostile.NameType, UserOwned: hostile.NameUser,
					HP: hostile.Hp, UnitAction: strconv.FormatBool(hostile.Action), Target: hostile.Target, X: hostile.X, Y: hostile.Y}
				initUnit <- unitsParametr
			}
		}
	}

	for _, xLine := range oldWatchUnit {
		for _, hostile := range xLine {
			find := findUnite(client, hostile)
			if !find {
				resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: hostile.X, Y: hostile.Y} // удаляем старое поле доступа
				coordiante <- resp
			}
		}
	}
}

func findCoordinate(client *Clients, coordinate *objects.Coordinate) (find bool) {
	_, find = client.Watch[coordinate.X][coordinate.Y]
	return
}

func findUnite(client *Clients, unit *objects.Unit) (find bool) {
	_, find = client.HostileUnits[unit.X][unit.Y]
	return
}