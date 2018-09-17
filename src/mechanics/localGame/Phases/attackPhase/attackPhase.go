package attackPhase

import (
	"../../../../mechanics/db/updateSquad"
	"../../../../mechanics/gameObjects/unit"
	"../../../localGame"
	"sort"
)

func AttackPhase(game *localGame.Game) (resultBattle []*ResultBattle) {

	sortUnits := createQueueAttack(game.GetUnits())
	resultBattle = attack(sortUnits, game)
	// todo отыгрыш эквипа
	// todo влияние бафов
	// todo востановление power
	// todo ломание эквипа при попадание в юнита

	for _, player := range game.GetPlayers() {
		updateSquad.Squad(player.GetSquad())
	}

	return
}

type ResultBattle struct {
	Map         bool        `json:"map"`
	AttackUnit  unit.Unit   `json:"attack_unit"`
	TargetUnit  unit.Unit   `json:"target_unit"`
	TargetsUnit []unit.Unit `json:"targets_unit"`
	Error       string      `json:"error"`
}

func attack(sortUnits []*unit.Unit, game *localGame.Game) (resultBattle []*ResultBattle) {
	resultBattle = make([]*ResultBattle, 0)

	for _, gameUnit := range sortUnits {
		if gameUnit.HP > 0 {
			if gameUnit.Target != nil {

				target, findCoordinate := game.Map.GetCoordinate(gameUnit.Target.Q, gameUnit.Target.R)

				if findCoordinate {
					result := InitAttack(gameUnit, target, game)
					resultBattle = append(resultBattle, result)
				}
			}
		}

		gameUnit.Target = nil
		gameUnit.QueueAttack = 0
	}

	return
}

func createQueueAttack(Units map[int]map[int]*unit.Unit) (sortUnits []*unit.Unit) {

	for _, xLine := range Units {
		for _, gameUnit := range xLine {
			gameUnit.QueueAttack += gameUnit.Body.Initiative
			sortUnits = append(sortUnits, gameUnit)
		}
	}

	sort.Slice(sortUnits, func(i, j int) bool {
		return sortUnits[i].QueueAttack > sortUnits[j].QueueAttack
	})

	return
}
