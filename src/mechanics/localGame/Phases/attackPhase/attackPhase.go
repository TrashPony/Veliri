package attackPhase

import (
	"../../../db/updateSquad"
	"../../../gameObjects/unit"
	"../../../localGame"
	"sort"
)

func AttackPhase(game *localGame.Game) (resultBattle []*ResultBattle, resultEquip []*ResultEquip) {

	sortUnits := createQueueAttack(game.GetUnits())
	// отыгрываем бой
	resultBattle = attack(sortUnits, game)
	// отыгрываем снаряжение и эфекты наложеные на юнитов
	resultEquip = wageringEquip(sortUnits)
	// востаналиываем энерги, даем актив поинты и снимаем флаги использованого снаряжения
	recovery(game)

	for _, player := range game.GetPlayers() {
		updateSquad.Squad(player.GetSquad()) // вносим все изменениея в базу данных
	}

	return
}

type ResultBattle struct {
	AttackUnit  unit.Unit   `json:"attack_unit"`
	TargetUnits []unit.Unit `json:"targets_units"`
	Error       string      `json:"error"`
}

type ResultEquip struct {

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
