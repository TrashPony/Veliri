package mechanics

import (
	"./game"
	"./player"
	"./db"
)

func UserReady(client *player.Player, actionGame *game.Game) () {
	client.SetReady(true)
	db.UpdatePlayer(client)

	var allReady bool

	for _, user := range actionGame.GetPlayers() {
		if user.GetReady() == false {
			allReady = false
			break
		}
	}

	if allReady {
		changeGamePhase(actionGame)
	}
}

func changeGamePhase(actionGame *game.Game) {
	if actionGame.Phase == "Init" || actionGame.Phase == "attack" {
		actionGame.Phase = "move"
	} else {
		if actionGame.Phase == "move" {
			actionGame.Phase = "targeting"
		} else {
			if actionGame.Phase == "targeting" {
				actionGame.Phase = "attack" // todo запуск фазы атаки
			}
		}
	}

	actionGame.Step = actionGame.Step + 1
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
