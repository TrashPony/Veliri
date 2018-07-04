package updateSquad

import (
	"../../gameObjects/squad"
	"../../gameObjects/coordinate"
	"log"
	"strconv"
	"database/sql"
)

func Units(squad *squad.Squad, tx *sql.Tx) {
	units := squad.MatherShip.Units

	for slot, squadUnit := range units {
		if units[slot] == nil {
			_, err := tx.Exec("DELETE FROM squad_units WHERE id_squad=$1 AND slot = $2",
				squad.ID, slot)
			if err != nil {
				log.Fatal("delete unit " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1 AND id_squad_unit = $2",
				squad.ID, squadUnit.ID)
			if err != nil {
				log.Fatal("delete unit equip " + err.Error())
			}
		}

		if units[slot] != nil && squadUnit.ID == 0 { // если ид 0 значит этого юнита создали в програме и его еще нет в бд
			id := 0
			err := tx.QueryRow("INSERT INTO squad_units (id_squad, id_body, slot, x, y, rotate, on_map, action, target, queue_attack, hp) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squad.ID, squadUnit.Body.ID, slot, squadUnit.X, squadUnit.Y, squadUnit.Rotate, squadUnit.OnMap, squadUnit.Action,
				parseTarget(squadUnit), squadUnit.QueueAttack, squadUnit.HP).Scan(&id)
			if err != nil {
				log.Fatal("add new unit to squad " + err.Error())
			}

			squadUnit.ID = id
			UpdateBody(units[slot], squad.ID, "squad_units_equipping", tx)
			if squad.InGame {
				//todo обновление эфектов
				//UnitEffects(squadUnit)
			}
		}

		if units[slot] != nil && squadUnit.ID != 0 {
			_, err := tx.Exec(
				"UPDATE squad_units "+
					"SET id_body = $1, x = $2, y = $3, rotate = $4, action = $5, target = $6, queue_attack = $7, hp = $8 "+
					"WHERE id_squad = $9 AND slot = $10",
				squadUnit.Body.ID, squadUnit.X, squadUnit.Y, squadUnit.Rotate, squadUnit.Action, parseTarget(squadUnit), squadUnit.QueueAttack, squadUnit.HP, squad.ID, slot)

			if err != nil {
				log.Fatal("update unit squad" + err.Error())
			}

			UpdateBody(units[slot], squad.ID, "squad_units_equipping", tx)
			if squad.InGame {
				//todo обновление эфектов
				//UnitEffects(squadUnit)
			}
		}
	}
}

type aimer interface {
	GetTarget() *coordinate.Coordinate
}

func parseTarget(aimer aimer) string {
	var target string

	if aimer.GetTarget() != nil {
		target = strconv.Itoa(aimer.GetTarget().X) + ":" + strconv.Itoa(aimer.GetTarget().Y)
	}

	return target
}
