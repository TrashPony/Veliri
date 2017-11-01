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

	for _, pathNode := range path {
		if (end.X == pathNode.X) && (end.Y == pathNode.Y) {
			_, ok := unit.WatchUnit[strconv.Itoa(end.X)+":"+strconv.Itoa(end.Y)]
			if ok {
				return unit.X, unit.Y, errors.New("end cell is busy")
			}
		} else {
			_, ok := unit.WatchUnit[strconv.Itoa(pathNode.X)+":"+strconv.Itoa(pathNode.Y)]
			if ok {
				return 0,0, errors.New("cell is busy") // если клетка занято то выходит из этого пути и генерить новый
			}
		}

		x := unit.X
		y := unit.Y

		unit.X = pathNode.X
		unit.Y = pathNode.Y

		oldWatchZone := unit.Watch

		oldWatchUnit := unit.WatchUnit // TODO А нахуя я вообще вношу изменния в бд каждую итерацию?! ретурном возвращаем значение и его уже вносим блять!

		delete(units, strconv.Itoa(x) + ":" + strconv.Itoa(y))
		units[strconv.Itoa(unit.X) + ":" + strconv.Itoa(unit.Y)] = unit

		delete(client.Units, strconv.Itoa(x)+":"+strconv.Itoa(y))          // удаляем в карте общий юнитов старое место расположение
		client.Units[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = unit // добавляем новое

		UpdateWatchZone(*client, *unit, units, oldWatchZone) // отправляем открытые ячейки, удаляем закрытые
		UpdateHostile(*client, oldWatchUnit, *unit) 		 // добавляем и удаляем вражских юнитов по мере их открытия/закрытия
		go UpdateWatchHostileUser(*client, *unit, x, y)		 // добавляем и удаляем нашего юнита у врагов на карте

		time.Sleep(100 * time.Millisecond)
	}

	return unit.X, unit.Y, nil
}

func UpdateWatchZone(client Clients, unitMove objects.Unit, units map[string]*objects.Unit, oldWatchZone map[string]*objects.Coordinate)  {
	// TODO проверка клетки на ее изначальное состояние
	for _, unitWatch := range client.Units {
		var err error
		unitWatch.Watch, unitWatch.WatchUnit, err = PermissionCoordinates(client, unitWatch, units) //обновляем зону видимости всех мобов
		if err != nil {
			continue
		}
	}

	unit := client.Units[strconv.Itoa(unitMove.X) + ":" + strconv.Itoa(unitMove.Y)]

	for _, newCoordinate := range unit.Watch { // отправляем все новые поля
		_, ok := unit.WatchUnit[strconv.Itoa(newCoordinate.X)+":"+strconv.Itoa(newCoordinate.Y)]
		if !ok {
			resp := FieldResponse{Event: "OpenCoordinate", UserName: client.Login, X: newCoordinate.X, Y: newCoordinate.Y}
			fieldPipe <- resp
		}
	}

	for _, oldCoordinate := range oldWatchZone {
		deleteCell := true
		for _, userUnit := range client.Units {
			if userUnit.NameUser == client.Login {
				_, ok := userUnit.Watch[strconv.Itoa(oldCoordinate.X)+":"+strconv.Itoa(oldCoordinate.Y)]
				if ok {
					deleteCell = false
					break
				}
			} else {
				deleteCell = false
				continue
			}
		}
		if deleteCell {
			resp := FieldResponse{Event: "DellCoordinate", UserName: client.Login, X: oldCoordinate.X, Y: oldCoordinate.Y} // удаляем старое поле доступа
			fieldPipe <- resp
		}
	}

	var unitsParametr = FieldResponse{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
		HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
	fieldPipe <- unitsParametr
}

func UpdateWatchHostileUser(client Clients, unit objects.Unit, x,y int)  {
	for _, gameUser := range client.Players {
		for _, user := range usersFieldWs {
			if user.Login != client.Login {
				if gameUser.Name == client.Login {
					for _, userUnits := range user.Units {
						_, okGetUnit := userUnits.WatchUnit[strconv.Itoa(x)+":"+strconv.Itoa(y)]
						if okGetUnit {
							delete(userUnits.WatchUnit, strconv.Itoa(x)+":"+strconv.Itoa(y))                       // если удалось взять вражеского юнита по старым координатам то удаляем его
							userUnits.Watch[strconv.Itoa(x)+":"+strconv.Itoa(y)] = &objects.Coordinate{X: x, Y: y} // и добавлдяем на его место пустую зону
							delete(user.HostileUnits, strconv.Itoa(x)+":"+strconv.Itoa(y))                         // и удаляем в общей карте вражеских юнитов
							resp := FieldResponse{Event: "OpenCoordinate", UserName: user.Login, X: x, Y: y}       // и остылаем событие удаление юнита
							fieldPipe <- resp
						}
						_, okGetXY := userUnits.Watch[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
						if okGetXY { // если следующая клетка юнита в зоне видимости
							delete(userUnits.Watch, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))     // удаляем пустую клетку
							userUnits.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit // и добавляем юнита в видимость юнита
							user.HostileUnits[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit   // и в общую карту вражескию юнитов
							var unitsParametr = FieldResponse{Event: "InitUnit", UserName: user.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
								HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: strconv.Itoa(unit.Target), X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
							fieldPipe <- unitsParametr
						}
						if okGetUnit && !okGetXY { // если удалось взять юнита по старым параметрам и не удалось взять координату открытую
							resp := FieldResponse{Event: "OpenCoordinate", UserName: user.Login, X: x, Y: y} // то остылаем событие удаление юнита
							fieldPipe <- resp
						}
					}
				}
			}
		}
	}
}

func UpdateHostile(client Clients, oldWatchUnit map[string]*objects.Unit, unit objects.Unit) {
	for _, hostile := range unit.WatchUnit { // добавляев новых открытых вражеских юнитов
		if hostile.NameUser != client.Login {
			_, ok := oldWatchUnit[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)]
			if !ok {
				client.HostileUnits[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)] = hostile                                                    // если появился новый враг
				var unitsParametr = FieldResponse{Event: "InitUnit", UserName: client.Login, TypeUnit: hostile.NameType, UserOwned: hostile.NameUser,
					HP: hostile.Hp, UnitAction: strconv.FormatBool(hostile.Action), Target: strconv.Itoa(hostile.Target), X: hostile.X, Y: hostile.Y} // остылаем событие добавления юнита
				fieldPipe <- unitsParametr
				continue
			}
		} else {
			continue
		}
	}

	for _, hostile := range oldWatchUnit {
		deleteUnit := true
		for _, userUnit := range client.Units {
			if hostile.NameUser != client.Login {
				_, ok := userUnit.WatchUnit[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)]
				if ok {
					deleteUnit = false
					break
				}
			} else {
				deleteUnit = false
				continue
			}
		}
		if deleteUnit {
			delete(client.HostileUnits, strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y))                   // если раньше видили врага а сейчас нет
			resp := FieldResponse{Event: "DellCoordinate", UserName: client.Login, X: hostile.X, Y: hostile.Y} // то остылаем событие удаление юнита
			fieldPipe <- resp
		}
	}
}
