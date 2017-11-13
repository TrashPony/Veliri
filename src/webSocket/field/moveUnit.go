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
	if find {
		if unit.Action {
			respawn := usersFieldWs[ws].Respawn
			coordinates := mechanics.GetCoordinates(unit.X, unit.Y, unit.MoveSpeed)
			var passed bool
			for _, coordinate := range coordinates {
				if !(coordinate.X == respawn.X && coordinate.Y == respawn.Y) {
					if coordinate.X == msg.ToX && coordinate.Y == msg.ToY {
						resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login}
						fieldPipe <- resp
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
			resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "unit already move"}
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
		obstacles := getObstacles(unit.WatchUnit) // TODO: добавить еще не проходимые учатки когда добавлю непроходимые участки
		start := objects.Coordinate{X: unit.X, Y: unit.Y}
		end := objects.Coordinate{X: toX, Y: toY}
		path := mechanics.FindPath(usersFieldWs[ws].Map, start, end, obstacles)
		x, y, errorMove := Move(unit, path, idGame, msg, ws, end)
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

func Move(unit *objects.Unit, path []objects.Coordinate, idGame int, msg FieldMessage, ws *websocket.Conn, end objects.Coordinate) (int, int, error) {
	client, ok := usersFieldWs[ws]
	if !ok {
		return 0,0, errors.New("connect is lost")
	}

	units := objects.GetAllUnits(client.GameStat.Id)
	activeUser := ActionGameUser(usersFieldWs[ws].Players)

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := unit.WatchUnit[end.X][end.Y]
			if ok {
				unit.Action = false
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
					HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
				initUnit <- unitsParametr
				return unit.X, unit.Y, errors.New("end cell is busy")
			}
		} else {
			_, ok := unit.WatchUnit[pathNode.X][pathNode.Y]
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

		oldWatchZone := unit.Watch

		oldWatchUnit := unit.WatchUnit

		delete(units, strconv.Itoa(x) + ":" + strconv.Itoa(y))
		units[strconv.Itoa(unit.X) + ":" + strconv.Itoa(unit.Y)] = unit

		delete(client.Units[x], y)          // удаляем в карте общий юнитов старое место расположение
		client.addUnit(unit)                // добавляем новое

		UpdateWatchZone(*client, *unit, units, oldWatchZone) // отправляем открытые ячейки, удаляем закрытые
		UpdateHostile(*client, oldWatchUnit, *unit) 		 // добавляем и удаляем вражских юнитов по мере их открытия/закрытия
		go UpdateWatchHostileUser(*client, *unit, x, y, activeUser)		 // добавляем и удаляем нашего юнита у врагов на карте

		time.Sleep(100 * time.Millisecond)
	}

	return unit.X, unit.Y, nil
}

func UpdateWatchZone(client Clients, unitMove objects.Unit, units map[string]*objects.Unit, oldWatchZone map[string]*objects.Coordinate)  {
	for yLine := range client.Units {
		for _, unitWatch := range client.Units[yLine] {
			var err error
			unitWatch.Watch, unitWatch.WatchUnit, err = PermissionCoordinates(client, unitWatch, units) //обновляем зону видимости всех мобов
			if err != nil {
				continue
			}
		}
	}

	unit := client.Units[unitMove.X][unitMove.Y]

	for _, newCoordinate := range unit.Watch { // отправляем все новые поля
		_, ok := unit.WatchUnit[newCoordinate.X][newCoordinate.Y]
		if !ok {
			resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: newCoordinate.X, Y: newCoordinate.Y}
			coordiante <- resp
		}
	}

	for _, oldCoordinate := range oldWatchZone {
		find := findCoordinate(client, *oldCoordinate)
		if !find {
			resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: oldCoordinate.X, Y: oldCoordinate.Y} // удаляем старое поле доступа
			coordiante <- resp
		}
	}

	var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
		HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
	initUnit <- unitsParametr
}

func UpdateWatchHostileUser(client Clients, unit objects.Unit, x,y int, activeUser []*Clients) {
	for _, user := range activeUser {
		if user.Login != client.Login {
			for _, xLine := range user.Units {
				for _, userUnits := range xLine {

					_, okGetUnit := userUnits.WatchUnit[x][y]

					if okGetUnit {
						delete(userUnits.WatchUnit[x], y)                                                      // если удалось взять вражеского юнита по старым координатам то удаляем его
						userUnits.Watch[strconv.Itoa(x)+":"+strconv.Itoa(y)] = &objects.Coordinate{X: x, Y: y} // и добавлдяем на его место пустую зону
						delete(user.HostileUnits[x], y)                                                        // и удаляем в общей карте вражеских юнитов
						resp := Coordinate{Event: "OpenCoordinate", UserName: user.Login, X: x, Y: y}          // и остылаем событие удаление юнита
						coordiante <- resp
					}

					_, okGetXY := userUnits.Watch[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]

					if okGetXY { 																								 // если следующая клетка юнита в зоне видимости
						delete(userUnits.Watch, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))                                   // удаляем пустую клетку
						userUnits.AddWatchUnit(&unit)                                                                            // и добавляем юнита в видимость юнита
						user.addHostileUnit(&unit)                                                                               // и в общую карту вражескию юнитов
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
	}
}

func UpdateHostile(client Clients, oldWatchUnit map[int]map[int]*objects.Unit, unit objects.Unit) {
	for _, xLine := range unit.WatchUnit { // добавляев новых открытых вражеских юнитов
		for _, hostile := range xLine {
			if hostile.NameUser != client.Login {
				_, ok := oldWatchUnit[hostile.X][hostile.Y]
				if !ok {
					client.addHostileUnit(hostile)                                                                                          // если появился новый враг
					var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: hostile.NameType, UserOwned: hostile.NameUser,
						HP: hostile.Hp, UnitAction: strconv.FormatBool(hostile.Action), Target: hostile.Target, X: hostile.X, Y: hostile.Y} // остылаем событие добавления юнита
					initUnit <- unitsParametr
					continue
				}
			} else {
				continue
			}
		}
	}

	for _, xLine := range unit.WatchUnit { // добавляев новых открытых вражеских юнитов
		for _, hostile := range xLine {
			deleteUnit := true
			for _, xLine := range client.Units {
				for _, userUnit := range xLine {
					if hostile.NameUser != client.Login {
						_, ok := userUnit.WatchUnit[hostile.X][hostile.Y]
						if ok {
							deleteUnit = false
							break
						}
					} else {
						deleteUnit = false
						continue
					}
				}
			}
			if deleteUnit {
				delete(client.HostileUnits[hostile.X], hostile.Y)                   // если раньше видили врага а сейчас нет
				resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: hostile.X, Y: hostile.Y} // то остылаем событие удаление юнита
				coordiante <- resp
			}
		}
	}
}

func getObstacles(units map[int]map[int]*objects.Unit)([]*objects.Coordinate)  {
	coordinates := make([]*objects.Coordinate,0)
	for yLine := range units {
		for _, unit := range units[yLine] {
			var coordinate objects.Coordinate
			coordinate.X = unit.X
			coordinate.Y = unit.Y
			coordinates = append(coordinates, &coordinate)
		}
	}
	return coordinates
}

func findCoordinate(client Clients, coordinate objects.Coordinate) (find bool) {
	for _, xLine := range client.Units {
		for _, userUnit := range xLine {
			if userUnit.NameUser == client.Login {
				_, ok := userUnit.Watch[strconv.Itoa(coordinate.X)+":"+strconv.Itoa(coordinate.Y)]
				if ok {
					find = true
					return
				}
			} else {
				continue
			}
		}
	}
	find = false
	return
}