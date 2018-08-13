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

	for slot, slotUnit := range units {

		if units[slot].Unit == nil {
			_, err := tx.Exec("DELETE FROM squad_units WHERE id_squad=$1 AND slot = $2",
				squad.ID, slot)
			if err != nil {
				log.Fatal("delete unit " + err.Error())
			}

			/*_, err = tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1 AND id_squad_unit = $2",
				squad.ID, slotUnit.Unit.ID)
			if err != nil {
				log.Fatal("delete unit equip " + err.Error())
			} todo */
		}

		if units[slot].Unit != nil && slotUnit.Unit.ID == 0 { // если ид 0 значит этого юнита создали в програме и его еще нет в бд
			id := 0
			err := tx.QueryRow("INSERT INTO squad_units (id_squad, id_body, slot, x, y, rotate, on_map, action, target, queue_attack, hp) "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squad.ID, slotUnit.Unit.Body.ID, slot, slotUnit.Unit.X, slotUnit.Unit.Y, slotUnit.Unit.Rotate, slotUnit.Unit.OnMap, slotUnit.Unit.Action,
				parseTarget(slotUnit.Unit), slotUnit.Unit.QueueAttack, slotUnit.Unit.HP).Scan(&id)
			if err != nil {
				log.Fatal("add new unit to squad " + err.Error())
			}

			slotUnit.Unit.ID = id
			UpdateBody(units[slot].Unit, squad.ID, "squad_units_equipping", tx)
			if squad.InGame {
				//todo обновление эфектов
				//UnitEffects(squadUnit)
			}
		}

		if units[slot].Unit != nil && slotUnit.Unit.ID != 0 {
			_, err := tx.Exec(
				"UPDATE squad_units "+
					"SET id_body = $1, x = $2, y = $3, rotate = $4, action = $5, target = $6, queue_attack = $7, hp = $8 "+
					"WHERE id_squad = $9 AND slot = $10",
				slotUnit.Unit.Body.ID, slotUnit.Unit.X, slotUnit.Unit.Y, slotUnit.Unit.Rotate, slotUnit.Unit.Action, parseTarget(slotUnit.Unit), slotUnit.Unit.QueueAttack, slotUnit.Unit.HP, squad.ID, slot)

			if err != nil {
				log.Fatal("update unit squad" + err.Error())
			}

			UpdateBody(units[slot].Unit, squad.ID, "squad_units_equipping", tx)
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
