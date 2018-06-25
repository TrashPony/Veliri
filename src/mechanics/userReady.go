package mechanics

import (
	"./localGame"
	"./player"
	"./db"
)

func UserReady(client *player.Player, actionGame *localGame.Game) (bool) {
	client.SetReady(true)
	db.UpdatePlayer(client)

	allReady := true

	for _, user := range actionGame.GetPlayers() {
		if user.GetReady() == false {
			allReady = false
			break
		}
	}

	if allReady {
		changeGamePhase(actionGame)
	}

	return allReady
}

func changeGamePhase(actionGame *localGame.Game) {
	if actionGame.Phase == "Init" || actionGame.Phase == "attack" {
		actionGame.Phase = "move"
	} else {
		if actionGame.Phase == "move" {
			actionGame.Phase = "targeting"
		} else {
			if actionGame.Phase == "targeting" {
				actionGame.Phase = "attack"
				// todo запуск фазы атаки
				actionGame.Step = actionGame.Step + 1
			}
		}
	}

	db.UpdateGame(actionGame)

	for _, user := range actionGame.GetPlayers() {
		user.SetReady(false)
		db.UpdatePlayer(user)
	}

	for _, xLine := range actionGame.GetUnits() {
		for _, unit := range xLine {
			unit.Action = false
			unit.Target = nil
			db.UpdateUnit(unit)
		}
	}
}
