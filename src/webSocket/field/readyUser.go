package field

import (
	"../../game"
	"github.com/gorilla/websocket"
	"time"
)

func Ready(msg FieldMessage, ws *websocket.Conn) {
	var resp FieldResponse
	phase, err, phaseChange := game.UserReady(usersFieldWs[ws].GetID(), msg.IdGame)
	client := usersFieldWs[ws]
	activeGame := Games[client.GetGameID()]
	players := activeGame.GetPlayers()
	activeUser := ActionGameUser(players)
	if phase != "" { // TODO
		activeGame.GetStat().Phase = phase
	}

	if phase == "attack" {
		resultBattle := game.AttackPhase(activeGame.GetUnits())

		for _, player := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase}
			fieldPipe <- resp
		}

		for _, attack := range resultBattle {
			attackSender(&attack.AttackUnit, activeUser)
			if attack.Delete {
				DelUnit(&attack.TargetUnit, activeUser)
			} else {
				UpdateUnit(&attack.TargetUnit, activeUser)
			}
		}

		phaseChange = true
		phase, _ = game.PhaseСhange(msg.IdGame)
	}

	if err != nil {
		resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Error: err.Error()}
		fieldPipe <- resp
		return
	}

	if 0 == len(usersFieldWs[ws].GetUnits()) {
		resp = FieldResponse{Event: msg.Event, UserName: client.GetLogin(), Error: "not units"}
		fieldPipe <- resp
		// TODO добавить окончание игры
		return
	}

	if phaseChange {
		for _, player := range activeUser {
			resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase}
			fieldPipe <- resp
			activeGame.GetStat().Phase = phase

			if phase == "move" {
				resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase, GameStep: activeGame.GetStat().Step + 1}
				activeGame.GetStat().Step += 1
			}

			for yLine := range player.GetUnits() { // TODO Нахера?
				for _, unit := range player.GetUnits()[yLine] {
					unit.Action = true

					if phase == "move" {
						unit.Target = nil
					}

					var unitsParameter InitUnit
					unitsParameter.initUnit(unit, player.GetLogin())
				}
			}
		}
	} else {
		resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].GetLogin(), Phase: phase}
		fieldPipe <- resp
	}
}



func attackSender(unit *game.Unit, activeUser []*game.Player) {

	for _, client := range activeUser {
		attackInfo := FieldResponse{Event: "Attack", UserName: client.GetLogin(), X: unit.X, Y: unit.Y, ToX: unit.Target.X, ToY: unit.Target.Y}
		fieldPipe <- attackInfo
	}

	time.Sleep(2000 * time.Millisecond)

	for _, client := range activeUser {
		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.GetLogin())
	}
}

func UpdateUnit(unit *game.Unit, activeUser []*game.Player) {
	for _, client := range activeUser {
		if unit.NameUser == client.GetLogin() {

			client.AddUnit(unit)

			var unitsParameter InitUnit
			unitsParameter.initUnit(unit, client.GetLogin())

		} else {
			_, ok := client.GetHostileUnit(unit.X, unit.Y)
			if ok {
				client.AddHostileUnit(unit)
				var unitsParameter InitUnit
				unitsParameter.initUnit(unit, client.GetLogin())
			}
		}
	}
}

func DelUnit(unit *game.Unit, activeUser []*game.Player) {
	for _, client := range activeUser {
		if unit.NameUser == client.GetLogin() {
			_, ok := client.GetUnit(unit.X, unit.Y)
			if ok {
				client.DelUnit(unit.X, unit.Y)
				Games[client.GetGameID()].DelUnit(unit)

				openCoordinate(client.GetLogin(), unit.X, unit.Y)
				UpdateWatchZone(client, Games[client.GetGameID()], nil)
			}
		} else {
			_, ok := client.GetHostileUnit(unit.X, unit.Y)
			if ok {
				client.DelUnit(unit.X, unit.Y)
				openCoordinate(client.GetLogin(), unit.X, unit.Y)
			}
		}
	}
}
