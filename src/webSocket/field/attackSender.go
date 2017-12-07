package field

import (
	"../../game"
	"time"
)

func attack(activeGame *game.Game, activeUser []*game.Player, msg FieldMessage, phase string)  {
	var resp FieldResponse
	for _, player := range activeUser {
		resp = FieldResponse{Event: msg.Event, UserName: player.GetLogin(), Phase: phase}
		fieldPipe <- resp
	}

	resultBattle := game.AttackPhase(activeGame, activeUser)

	for _, attack := range resultBattle {
		attackSender(&attack.AttackUnit, activeUser)
		if attack.Delete {
			DelUnit(&attack.TargetUnit, activeUser)
		} else {
			UpdateUnit(&attack.TargetUnit, activeUser)
		}
	}
}



func attackSender(unit *game.Unit, activeUser []*game.Player) {

	for _, client := range activeUser {
		attackInfo := FieldResponse{Event: "Attack", UserName: client.GetLogin(), X: unit.X, Y: unit.Y, ToX: unit.Target.X, ToY: unit.Target.Y}
		fieldPipe <- attackInfo
	}

	time.Sleep(1000 * time.Millisecond)

	for _, client := range activeUser {
		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.GetLogin())
	}
}

func UpdateUnit(unit *game.Unit, activeUser []*game.Player) {
	for _, client := range activeUser {
		var unitsParameter InitUnit
		unitsParameter.initUnit(unit, client.GetLogin())
	}
}

func DelUnit(unit *game.Unit, activeUser []*game.Player) {
	for _, client := range activeUser {
		openCoordinate(client.GetLogin(), unit.X, unit.Y)
		UpdateWatchZone(client, Games[client.GetGameID()], nil)
	}
}
