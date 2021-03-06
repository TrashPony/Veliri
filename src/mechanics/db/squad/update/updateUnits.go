package update

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"log"
)

func Units(squad *squad.Squad, tx *sql.Tx) {
	units := squad.MatherShip.Units

	for slot, slotUnit := range units {

		if units[slot].Unit == nil {
			id := 0
			rows, err := tx.Query("Select id FROM squad_units WHERE id_squad=$1 AND slot = $2", squad.ID, slot)
			if err != nil {
				log.Fatal("get id deleting unit " + err.Error())
			}

			for rows.Next() {
				err := rows.Scan(&id)
				if err != nil {
					log.Fatal("get id deleting unit " + err.Error())
				}
			}

			_, err = tx.Exec("DELETE FROM user_memory_unit WHERE id_unit = $1", id)
			if err != nil {
				log.Fatal("delete memory unit " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1 AND id_squad_unit = $2", squad.ID, id)
			if err != nil {
				log.Fatal("delete unit equip " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units_inventory WHERE id_unit=$1", id)
			if err != nil {
				log.Fatal("delete inventory unit " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units WHERE id_squad=$1 AND slot = $2", squad.ID, slot)
			if err != nil {
				log.Fatal("delete unit " + err.Error())
			}
		}

		if units[slot].Unit != nil && slotUnit.Unit.ID == 0 { // если ид 0 значит этого юнита создали в програме и его еще нет в бд
			id := 0
			err := tx.QueryRow("INSERT INTO squad_units ("+
				"id_squad, "+
				"id_body, "+
				"slot, "+
				"x, "+
				"y, "+
				"rotate, "+
				"on_map, "+
				"hp, "+
				"power, "+
				"mother_ship, "+
				"action_point, "+
				"defend, "+
				"move,"+
				"id_map "+
				""+
				") "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
				squad.ID,
				slotUnit.Unit.Body.ID,
				slot,
				slotUnit.Unit.X,
				slotUnit.Unit.Y,
				slotUnit.Unit.Rotate,
				slotUnit.Unit.OnMap,
				slotUnit.Unit.HP,
				slotUnit.Unit.Power,
				false, //mother_ship
				slotUnit.Unit.ActionPoints,
				slotUnit.Unit.Defend,
				slotUnit.Unit.Move,
				slotUnit.Unit.MapID,
			).Scan(&id)
			if err != nil {
				log.Fatal("add new unit to squad " + err.Error())
			}

			slotUnit.Unit.ID = id
			UpdateBody(units[slot].Unit, squad.ID, tx)
		}

		if units[slot].Unit != nil && slotUnit.Unit.ID != 0 {
			_, err := tx.Exec(
				"UPDATE squad_units SET "+
					"id_body = $1, "+
					"x = $2, "+
					"y = $3, "+
					"rotate = $4, "+
					"hp = $5, "+
					"on_map = $6, "+
					"power = $7, "+
					"action_point = $8, "+
					"defend = $9, "+
					"move = $12, "+
					"id_game = $13, "+
					"body_color_1 = $14,"+
					"body_color_2 = $15,"+
					"weapon_color_1 = $16,"+
					"weapon_color_2 = $17,"+
					"id_map = $18 "+
					"WHERE id_squad = $10 AND slot = $11",
				slotUnit.Unit.Body.ID,
				slotUnit.Unit.X,
				slotUnit.Unit.Y,
				slotUnit.Unit.Rotate,
				slotUnit.Unit.HP,
				slotUnit.Unit.OnMap,
				slotUnit.Unit.Power,
				slotUnit.Unit.ActionPoints,
				slotUnit.Unit.Defend,
				squad.ID,
				slot,
				slotUnit.Unit.Move,
				slotUnit.Unit.GameID,
				slotUnit.Unit.BodyColor1,
				slotUnit.Unit.BodyColor2,
				slotUnit.Unit.WeaponColor1,
				slotUnit.Unit.WeaponColor2,
				slotUnit.Unit.MapID,
			)

			if err != nil {
				log.Fatal("update unit squad" + err.Error())
			}

			UpdateBody(units[slot].Unit, squad.ID, tx)
		}
	}
}
