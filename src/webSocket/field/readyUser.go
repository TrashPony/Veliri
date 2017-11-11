package field

import (
	"websocket-master"
	"../../game/mechanics"
	"../../game/objects"
	"strconv"
	"strings"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	phase, err, phaseChange := mechanics.UserReady(usersFieldWs[ws].Id, msg.IdGame)
	activeUser := ActionGameUser(usersFieldWs[ws].Players)

	if phase == "attack" {
		sortUnits := mechanics.AttackPhase(msg.IdGame) // TODO обновить информацию юнитов у всех юзеров в сети
		attack(sortUnits, activeUser)

		for _, client := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase}
			fieldPipe <- resp
			client.GameStat.Phase = phase
		}

		phaseChange = true
		phase, _ = mechanics.PhaseСhange(msg.IdGame)
	}

	if err != nil {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Error: err.Error()}
		fieldPipe <- resp
		return
	}

	if 0 == len(usersFieldWs[ws].Units) {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Error: "not units"}
		fieldPipe <- resp
		// TODO добавить окончание игры
		return
	}

	if phaseChange {
		for _, client := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase}
			fieldPipe <- resp
			client.GameStat.Phase = phase

			if phase == "move" {
				resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase, GameStep: client.GameStat.Step + 1}
				client.GameStat.Step += 1
			}

			for _, unit := range client.Units {
				unit.Action = true

				if phase == "move" {
					unit.Target = ""
				}

				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
					HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: unit.Target, X: unit.X, Y: unit.Y}
				initUnit <- unitsParametr // отправляем параметры каждого юнита отдельно
			}
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, Phase: phase}
		fieldPipe <- resp
	}
}


func attack(sortUnits []objects.Unit, activeUser []*Clients) {
	for _, unit := range sortUnits {
		if unit.Hp > 0 {
			if unit.Target != "" {
				targetCoordinate := strings.Split(unit.Target, ":") // X:Y

				x, _ := strconv.Atoi(targetCoordinate[0])
				y, _ := strconv.Atoi(targetCoordinate[1])

				for i, target := range sortUnits {
					if target.X == x && target.Y == y {
						sortUnits[i].Hp = target.Hp - unit.Damage
						if sortUnits[i].Hp <= 0 {
							go mechanics.DelUnit(sortUnits[i].Id)
							// TODO оповещалка об атаки в сокет
							DelUnit(sortUnits[i], activeUser)
						} else {
							go mechanics.UpdateUnit(sortUnits[i].Id, sortUnits[i].Hp)
							// TODO оповещалка об атаки в сокет
							UpdateUnit(sortUnits[i], activeUser)
						}
					}
				}
			}
		}
		go mechanics.UpdateTarget(unit.Id)
	}
}

func UpdateUnit(unit objects.Unit, activeUser []*Clients)  {

	for _, client := range activeUser {
		if unit.NameUser == client.Login {
			realUnit := client.Units[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
			unit.WatchUnit = realUnit.WatchUnit
			unit.Watch = realUnit.Watch

			client.Units[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit
			for _, units := range client.Units {
				_, ok := units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
				if ok { // TODO вынести одинаковые метод по поиску юнитов в отдельный метод, а то чето пиздец
					units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit
				}
			}

			var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
				HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: "", X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
			initUnit <- unitsParametr

		} else {

			_, ok := client.HostileUnits[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
			if ok {
				client.HostileUnits[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit
				var unitsParametr = InitUnit{Event: "InitUnit", UserName: client.Login, TypeUnit: unit.NameType, UserOwned: unit.NameUser,
					HP: unit.Hp, UnitAction: strconv.FormatBool(unit.Action), Target: "", X: unit.X, Y: unit.Y} // остылаем событие добавления юнита
				initUnit <- unitsParametr
			}
			for _, units := range client.Units {
				_, ok := units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
				if ok {
					units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)] = &unit
				}
			}
		}
	}
}

func DelUnit(unit objects.Unit, activeUser []*Clients) {

	for _, client := range activeUser {
		if unit.NameUser == client.Login {

			realUnit := client.Units[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
			unit.WatchUnit = realUnit.WatchUnit
			unit.Watch = realUnit.Watch
			delete(client.Units, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))
			resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: unit.X, Y: unit.Y}
			coordiante <- resp

			for _, units := range client.Units {
				_, ok := units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
				if ok {
					delete(units.WatchUnit, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))
				}
			}

			for _, coor := range realUnit.Watch {
				del := true
				for _, units := range client.Units {
					_, ok := units.Watch[strconv.Itoa(coor.X)+":"+strconv.Itoa(coor.Y)]
					if ok {
						del = false
						break
					}
				}
				if del {
					resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: coor.X, Y: coor.Y} // то остылаем событие удаление юнита
					coordiante <- resp
				}
			}

			for _, hostile := range realUnit.WatchUnit { // TODO неуверен что тут это надо
				del := true
				for _, units := range client.Units {
					_, ok := units.Watch[strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y)]
					if ok {
						del = false
						break
					}
				}
				if del {
					delete(client.HostileUnits, strconv.Itoa(hostile.X)+":"+strconv.Itoa(hostile.Y))
					resp := Coordinate{Event: "DellCoordinate", UserName: client.Login, X: hostile.X, Y: hostile.Y} // то остылаем событие удаление юнита
					coordiante <- resp
				}
			}
		} else {
			resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: unit.X, Y: unit.Y}
			coordiante <- resp

			_, ok := client.HostileUnits[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
			if ok {
				delete(client.HostileUnits, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))
			}
			for _, units := range client.Units {
				_, ok := units.WatchUnit[strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y)]
				if ok {
					delete(units.WatchUnit, strconv.Itoa(unit.X)+":"+strconv.Itoa(unit.Y))
				}
			}
		}
	}
}

