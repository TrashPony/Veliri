package game

import (
	"sort"
)

func AttackPhase(Units map[int]map[int]*Unit) (sortUnits []*Unit) {
	sortUnits = createQueueAttack(Units)
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

func DelUnit(id int) {
	_, err := db.Exec("DELETE FROM action_game_unit WHERE id=$1", id)
	if err != nil {
		println("нет такого юнита") // TODO сбрасывать инициативу до дефолта
	}
}

func UpdateUnit(id int, hp int) {
	_, err := db.Exec("UPDATE action_game_unit SET hp=$2 WHERE id=$1", id, hp)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}

func UpdateTarget(id int) {
	_, err := db.Exec("UPDATE action_game_unit SET target=$2, queue_attack=$3 WHERE id=$1", id, "", 0)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}
