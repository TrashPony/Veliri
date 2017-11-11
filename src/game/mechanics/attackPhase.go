package mechanics

import (
	"../objects"
	"sort"
)

func AttackPhase(idGame int) (sortUnits []objects.Unit){
	Units := objects.GetAllUnits(idGame)
	sortUnits = createQueueAttack(Units)
	return
}

func createQueueAttack(Units map[string]*objects.Unit)(sortUnits []objects.Unit)  {

	for _, unit := range Units {
		unit.Init = unit.Init + unit.Queue
		sortUnits = append(sortUnits, *unit)
	}

	sort.Slice(sortUnits, func(i, j int) bool {
		return sortUnits[i].Init > sortUnits[j].Init
	})

	return
}

func DelUnit(id int)  {
	_, err := db.Exec("DELETE FROM action_game_unit WHERE id=$1", id)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}

func UpdateUnit(id int, hp int) {
	_, err := db.Exec("UPDATE action_game_unit SET hp=$2 WHERE id=$1", id, hp)
	if err != nil {
		println("нет такого юнита") // TODO
	}
}

func UpdateTarget(id int)  {
	_, err := db.Exec("UPDATE action_game_unit SET target=$2 WHERE id=$1", id, "")
	if err != nil {
		println("нет такого юнита") // TODO
	}
}
