package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"time"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	phase, err, phaseChange := game.UserReady(usersFieldWs[ws].Id, msg.IdGame)
	activeGame := Games[usersFieldWs[ws].GameID]
	players := activeGame.getPlayers()
	activeUser := ActionGameUser(players)
	if phase != "" { // TODO
		activeGame.stat.Phase = phase
	}
	if phase == "attack" {
		sortUnits := game.AttackPhase(activeGame.getUnits())
		attack(sortUnits, activeUser)

		for _, client := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase}
			fieldPipe <- resp
		}

		phaseChange = true
		phase, _ = game.PhaseСhange(msg.IdGame)
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
			activeGame.stat.Phase = phase

			if phase == "move" {
				resp = FieldResponse{Event: msg.Event, UserName: client.Login, Phase: phase, GameStep: activeGame.stat.Step + 1}
				activeGame.stat.Step += 1
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

func attack(sortUnits []*game.Unit, activeUser []*Clients) {
	for _, unit := range sortUnits {
		if unit.Hp > 0 {
			if unit.Target != nil {
				for i, target := range sortUnits {
					if target.X == unit.Target.X && target.Y == unit.Target.Y {
						sortUnits[i].Hp = target.Hp - unit.Damage
						if sortUnits[i].Hp <= 0 {
							game.DelUnit(sortUnits[i].Id)
							attackSender(unit, activeUser)
							DelUnit(sortUnits[i], activeUser)
						} else {
							game.UpdateUnit(sortUnits[i].Id, sortUnits[i].Hp)
							attackSender(unit, activeUser)
							UpdateUnit(sortUnits[i], activeUser)
						}
					}
				}
			}
		}
		game.UpdateTarget(unit.Id)
		unit.Target = nil
		unit.Queue = 0
	}
}

func attackSender(unit *game.Unit, activeUser []*Clients) {

	for _, client := range activeUser {
		attackInfo := FieldResponse{Event: "Attack", UserName: client.Login, X: unit.X, Y: unit.Y, ToX: unit.Target.X, ToY: unit.Target.Y}
		fieldPipe <- attackInfo
	}

	time.Sleep(2000 * time.Millisecond)

	for _, client := range activeUser {
		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.Login)
	}
}

func UpdateUnit(unit *game.Unit, activeUser []*Clients) {
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

func DelUnit(unit *game.Unit, activeUser []*Clients) {
	for _, client := range activeUser {
		if unit.NameUser == client.Login {
			_, ok := client.Units[unit.X][unit.Y]
			if ok {
				delete(client.Units[unit.X], unit.Y)
				Games[client.GameID].delUnit(unit)

				openCoordinate(client.Login, unit.X, unit.Y)
				client.updateWatchZone(Games[client.GameID])
			}
		} else {
			_, ok := client.HostileUnits[unit.X][unit.Y]
			if ok {
				delete(client.HostileUnits[unit.X], unit.Y)
				openCoordinate(client.Login, unit.X, unit.Y)
			}
		}
	}
}
