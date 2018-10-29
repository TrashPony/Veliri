package movePhase

import (
	"../../../localGame"
	"../../../player"
	"../../../db/localGame/update"
)

func Queue(game *localGame.Game) int {
	queue := 0

	for _, xLine := range game.GetUnits() {
		for _, gameUnit := range xLine {
			if gameUnit.ActionPoints == 0 {
				if gameUnit.QueueAttack > queue {
					queue = gameUnit.QueueAttack
				}
			}
		}
	}

	return queue + 1
}

func queueMove(client *player.Player, game *localGame.Game) {

	client.Move = false
	client.SubMove = true
	update.Player(client)

	newMoveCycle := true

	if canMoveUser(game) {

		for i := client.QueueMovePos + 1; i <= len(game.GetPlayers()); i++ {
			check := false
			for _, gameUser := range game.GetPlayers() {
				if i == gameUser.QueueMovePos && canMoveUnit(gameUser) {
					gameUser.Move = true
					update.Player(gameUser)
					check = true
					break
				}
			}
			if check {
				break
			}
		}

		for _, gameUser := range game.GetPlayers() {
			if !gameUser.SubMove {
				newMoveCycle = false
			}
		}

		if newMoveCycle {

			for i := 1; i <= len(game.GetPlayers()); i++ {
				for _, gameUser := range game.GetPlayers() {

					gameUser.SubMove = false
					update.Player(gameUser)

					if gameUser.QueueMovePos == i && canMoveUnit(gameUser) {
						gameUser.Move = true
						gameUser.SubMove = false
						update.Player(gameUser)
						return
					}
				}
			}
		}
	} else {
		println("все походили")
		//TODO Все походили смена фазы
	}
}

func canMoveUser(game *localGame.Game) bool {
	for _, gameUser := range game.GetPlayers() {
		if canMoveUnit(gameUser) {
			return true
		}
	}

	return false
}

func canMoveUnit(client *player.Player) bool {

	if client.GetSquad().MatherShip.ActionPoints > 0 {
		return true
	}

	for _, q := range client.GetUnits() {
		for _, unit := range q {
			if unit.ActionPoints > 0 {
				return true
			}
		}
	}

	return false
}
