package movePhase

import (
	"../../../db/localGame/update"
	"../../../localGame"
	"../../../player"
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

func QueueMove(client *player.Player, game *localGame.Game) {

	client.Move = false
	client.SubMove = true
	update.Player(client)

	newMoveCycle := true

	if canMoveUser(game) {

		for _, gameUser := range game.GetPlayers() {
			if !gameUser.SubMove {
				newMoveCycle = false
			}
		}

		if newMoveCycle {
			for i := 1; i <= len(game.GetPlayers()); i++ {
				check := false
				for _, gameUser := range game.GetPlayers() {
					if gameUser.QueueMovePos == i && canMoveUnit(gameUser) {
						gameUser.Move = true
						gameUser.SubMove = false
						update.Player(gameUser)
						check = true
					}
				}
				if check {
					break
				}
			}

			for _, gameUser := range game.GetPlayers() {
				if canMoveUnit(gameUser) {
					gameUser.SubMove = false
					update.Player(gameUser)
				}
			}
		} else {
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
		}
	} else {
		println("все походили")
		//Все походили смена фазы т.к. больше чего не кто не сможет сделать
		for _, gameUser := range game.GetPlayers() {
			gameUser.SetReady(true)
		}
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

	if client.Ready {
		return false
	}

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

	for _, unit := range client.GetUnitsStorage() {
		if unit.ActionPoints > 0 {
			return true
		}
	}

	return false
}
