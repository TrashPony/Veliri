package field

import (
	"websocket-master"
	"../../game/mechanics"
	"../../game/objects"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	phase, err, phaseChange := mechanics.UserReady(usersFieldWs[ws].Id, msg.IdGame)
	game := Games[usersFieldWs[ws].GameID]
	players := game.getPlayers()
	activeUser := ActionGameUser(players)
	if phase != "" { // TODO
		game.Stat.Phase = phase
	}
	if phase == "attack" {
		sortUnits := mechanics.AttackPhase(game.getUnits())
		attack(sortUnits, activeUser)

		for _, client := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase}
			fieldPipe <- resp
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
			game.Stat.Phase = phase

			if phase == "move" {
				resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase, GameStep: game.Stat.Step + 1}
				game.Stat.Step += 1
			}

			for yLine := range client.Units { // TODO Нахера?
				for _, unit := range client.Units[yLine] {
					unit.Action = true

					if phase == "move" {
						unit.Target = nil
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


func attack(sortUnits []*objects.Unit, activeUser []*Clients) {
	for _, unit := range sortUnits {
		if unit.Hp > 0 {
			if unit.Target != nil {
				for i, target := range sortUnits {
					if target.X == unit.Target.X && target.Y == unit.Target.Y {
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
		mechanics.UpdateTarget(unit.Id)
		unit.Target = nil
		unit.Queue  = 0
	}
}

func UpdateUnit(unit *objects.Unit, activeUser []*Clients) {
	for _, client := range activeUser {
		if unit.NameUser == client.Login {

			client.addUnit(unit)

			var unitsParameter InitUnit
			unitsParameter.initUnit(unit, client.Login)

		} else {
			_, ok := client.HostileUnits[unit.X][unit.Y]
			if ok {
				client.addHostileUnit(unit)
				var unitsParameter InitUnit
				unitsParameter.initUnit(unit, client.Login)
			}
		}
	}
}

func DelUnit(unit *objects.Unit, activeUser []*Clients) {
	for _, client := range activeUser {
		if unit.NameUser == client.Login {
			_, ok := client.Units[unit.X][unit.Y]
			if ok {
				delete(client.Units[unit.X], unit.Y)
				Games[client.GameID].delUnit(unit)

				resp := Coordinate{Event: "OpenCoordinate", UserName: client.Login, X: unit.X, Y: unit.Y}
				coordiante <- resp
				units := Games[client.GameID].getUnits()
				client.updateWatchZone(units)
			}
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