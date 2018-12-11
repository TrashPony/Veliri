package userReady

import (
	"../../db/localGame/update"
	squadUpdate "../../db/squad/update"
	"../../localGame"
	"../../localGame/Phases/movePhase"
	"../../player"
)

func UserReady(client *player.Player) {
	client.SetReady(true)
	update.Player(client)
	squadUpdate.Squad(client.GetSquad())
}

func AllReady(actionGame *localGame.Game) bool {
	allReady := true

	for _, user := range actionGame.GetPlayers() {
		if user.GetReady() == false {
			allReady = false
			break
		}
	}

	if allReady {
		changeGamePhase(actionGame)
	} else {
		if actionGame.Phase == "move" {
			/* -- это тут для фазы движени а то будет жопка */
			movePhase.QueueMove(actionGame)
		}
	}

	return allReady
}

func changeGamePhase(actionGame *localGame.Game) {
	if actionGame.Phase == "attack" {
		actionGame.Step = actionGame.Step + 1
		actionGame.Phase = "move"
	} else {
		if actionGame.Phase == "move" {
			actionGame.Phase = "targeting"
		} else {
			if actionGame.Phase == "targeting" {
				actionGame.Phase = "attack"
			}
		}
	}

	update.Game(actionGame)

	for _, user := range actionGame.GetPlayers() {
		user.SetReady(false)
		update.Player(user)
	}
}
