package userReady

import (
	"../../player"
	"../../db/localGame/update"
	"../../localGame"
	"../../db/updateSquad"
)

func UserReady(client *player.Player, actionGame *localGame.Game) (bool) {
	client.SetReady(true)
	update.Player(client)

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

	updateSquad.Squad(client.GetSquad())

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

	update.Game(actionGame)

	for _, user := range actionGame.GetPlayers() {
		user.SetReady(false)
		update.Player(user)
	}

	for _, xLine := range actionGame.GetUnits() {
		for _, unit := range xLine {
			unit.Action = false
			unit.Target = nil
		}
	}
}
