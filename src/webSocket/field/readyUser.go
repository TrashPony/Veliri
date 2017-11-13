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

			for yLine := range client.Units {
				for _, unit := range client.Units[yLine] {
					unit.Action = true

					if phase == "move" {
						unit.Target = ""
					}

					var unitsParameter InitUnit
					unitsParameter.initUnit(unit, client.Login)
				}
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
							mechanics.DelUnit(sortUnits[i].Id)
							// TODO оповещалка об атаки в сокет
							DelUnit(sortUnits[i], activeUser)
						} else {
							mechanics.UpdateUnit(sortUnits[i].Id, sortUnits[i].Hp)
							// TODO оповещалка об атаки в сокет
							UpdateUnit(sortUnits[i], activeUser)
						}
					}
				}
			}
		}
		mechanics.UpdateTarget(unit.Id) // TODO иногда не все цели сбрасываются
	}
}

func UpdateUnit(unit objects.Unit, activeUser []*Clients) {
	for _, client := range activeUser {
		if unit.NameUser == client.Login {

			client.addUnit(&unit)
			var unitsParameter InitUnit
			unitsParameter.initUnit(&unit, client.Login)

		} else {
			_, ok := client.HostileUnits[unit.X][unit.Y]
			if ok {
				client.addHostileUnit(&unit)
				var unitsParameter InitUnit
				unitsParameter.initUnit(&unit, client.Login)
			}
		}
	}
}

func DelUnit(unit objects.Unit, activeUser []*Clients) {
	for _, client := range activeUser {
		if unit.NameUser == client.Login {
			_, ok := client.Units[unit.X][unit.Y]
			if ok {
				delete(client.Units[unit.X], unit.Y)
				resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: unit.X, Y: unit.Y}
				coordiante <- resp
				// TODO обновление зоны видимости
			} else {
				_, ok := client.HostileUnits[unit.X][unit.Y]
				if ok {
					delete(client.HostileUnits[unit.X], unit.Y)
					resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: unit.X, Y: unit.Y}
					coordiante <- resp
				}
			}
		}
	}
}