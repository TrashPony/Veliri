package game

import (
	"sort"
)

func AttackPhase(Units map[int]map[int]*Unit) (resultBattle []ResultBattle) {
	sortUnits := createQueueAttack(Units)
	resultBattle = attack(sortUnits)
	return
}

type ResultBattle struct {
	AttackUnit Unit
	TargetUnit Unit
	Delete bool
}

func attack(sortUnits []*Unit) (resultBattle []ResultBattle) {
	resultBattle = make([]ResultBattle, 0)

// TODO добавить обьект оповещения боя

	for _, unit := range sortUnits {
		if unit.Hp > 0 {
			if unit.Target != nil {
				for i, target := range sortUnits {
					if target.X == unit.Target.X && target.Y == unit.Target.Y {

						sortUnits[i].Hp = target.Hp - unit.Damage

						deleteUnit := false
						if sortUnits[i].Hp <= 0 {
							dbDelUnit(sortUnits[i].Id)
							deleteUnit = true
						} else {
							dbUpdateHpUnit(sortUnits[i].Id, sortUnits[i].Hp)
						}

						result := ResultBattle{AttackUnit: *unit, TargetUnit: *sortUnits[i], Delete: deleteUnit}
						resultBattle = append(resultBattle, result)
					}
				}
			}
		}
		dbUpdateTargetUnit(unit.Id)
		unit.Target = nil
		unit.Queue = 0
	}

	return
}

func createQueueAttack(Units map[int]map[int]*Unit) (sortUnits []*Unit) {

	for _, xLine := range Units {
		for _, unit := range xLine {
			unit.Initiative = unit.Initiative + unit.Queue
			sortUnits = append(sortUnits, unit)
		}
	}

	sort.Slice(sortUnits, func(i, j int) bool {
		return sortUnits[i].Initiative > sortUnits[j].Initiative
	})

	return
}

func dbDelUnit(id int) {
	_, err := db.Exec("DELETE FROM action_game_unit WHERE id=$1", id)
	if err != nil {
		println("нет такого юнита") // TODO сбрасывать инициативу до дефолта
	}
}

func dbUpdateHpUnit(id int, hp int) {
	_, err := db.Exec("UPDATE action_game_unit SET hp=$2 WHERE id=$1", id, hp)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}

func dbUpdateTargetUnit(id int) {
	_, err := db.Exec("UPDATE action_game_unit SET target=$2, queue_attack=$3 WHERE id=$1", id, "", 0)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}
