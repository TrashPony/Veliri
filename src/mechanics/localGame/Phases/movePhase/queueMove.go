package movePhase

import (
	"../../../localGame"
	"../../../player"
	"../../../gameObjects/unit"
	"math/rand"
	"time"
)

func QueueMove(game *localGame.Game) {

	if canMoveUser(game) {

		maxInitiative := 0
		var maxUnit *unit.Unit

		for _, q := range game.GetUnits() { //находим юнита с макс инициативой
			for _, gameUnit := range q {
				if gameUnit.ActionPoints > 0 && gameUnit.Initiative > maxInitiative {
					maxInitiative = gameUnit.Initiative
					maxUnit = gameUnit
				}
			}
		}

		for _, gameUnit := range game.GetUnitsStorage() { //находим юнита с макс инициативой
			if gameUnit.ActionPoints > 0 && gameUnit.Initiative > maxInitiative {
				maxInitiative = gameUnit.Initiative
				maxUnit = gameUnit
			}
		}

		moveUnits := make([]*unit.Unit, 0)

		for _, q := range game.GetUnits() { //ищем юнитов с такойже инициативой
			for _, gameUnit := range q {
				if gameUnit.ActionPoints > 0 && gameUnit.Initiative == maxUnit.Initiative {
					moveUnits = append(moveUnits, gameUnit)
				}
			}
		}

		for _, gameUnit := range game.GetUnitsStorage() { //ищем юнитов с такойже инициативой
			if gameUnit.ActionPoints > 0 && gameUnit.Initiative == maxUnit.Initiative {
				moveUnits = append(moveUnits, gameUnit)
			}
		}

		if len(moveUnits) > 1 {
			randomUnitMove(moveUnits).Move = true
		} else {
			moveUnits[0].Move = true
		}
	} else {
		println("все походили")
		//Все походили смена фазы т.к. больше чего не кто не сможет сделать
		for _, gameUser := range game.GetPlayers() {
			gameUser.SetReady(true)
		}
	}
}

func randomUnitMove(moveUnits []*unit.Unit) *unit.Unit {
	//Генератор случайных чисел обычно нужно рандомизировать перед использованием, иначе, он, действительно,
	// будет выдавать одну и ту же последовательность.
	rand.Seed(time.Now().UnixNano())
	numberUnit := rand.Intn(len(moveUnits))

	return moveUnits[numberUnit]
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
		for _, gameUnit := range q {
			if gameUnit.ActionPoints > 0 {
				return true
			}
		}
	}

	for _, gameUnit := range client.GetUnitsStorage() {
		if gameUnit.ActionPoints > 0 {
			return true
		}
	}

	return false
}
