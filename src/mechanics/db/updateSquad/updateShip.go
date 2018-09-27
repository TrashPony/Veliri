package updateSquad

import (
	"../../gameObjects/squad"
	"database/sql"
	"log"
)

func MotherShip(squad *squad.Squad, tx *sql.Tx) {

	ship := squad.MatherShip

	if ship != nil && ship.ID != 0 {

		var bodyID sql.NullInt64

		if ship.Body == nil {
			bodyID = sql.NullInt64{Int64: 0, Valid: false}

			_, err := tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1 AND id_squad_unit=$2",
				squad.ID, ship.ID)
			if err != nil {
				log.Fatal("delete all unit equip " + err.Error())
			}

		} else {
			bodyID = sql.NullInt64{Int64: int64(ship.Body.ID), Valid: true}
			UpdateBody(ship, squad.ID, tx)
		}

		_, err := tx.Exec(
			"UPDATE squad_units SET "+
				"id_body = $1, "+
				"q = $2, "+
				"r = $3, "+
				"rotate = $4, "+
				"target = $5, "+
				"queue_attack = $6, "+
				"hp = $7, "+
				"power = $9, "+
				"action_point = $11, "+
				"defend = $12 "+
				"WHERE id_squad = $8 AND mother_ship = $10",
			bodyID,
			ship.Q,
			ship.R,
			ship.Rotate,
			parseTarget(ship),
			ship.QueueAttack,
			ship.HP,
			squad.ID,
			ship.Power,
			true, // mother_ship = $10
			ship.ActionPoints,
			ship.Defend,
		)

		if err != nil {
			log.Fatal("update motherShip squad " + err.Error())
		}

	} else {
		if ship.ID == 0 || ship.Body != nil {
			id := 0
			err := tx.QueryRow("INSERT INTO squad_units ("+
				"id_squad, "+
				"id_body, "+
				"q, "+
				"r, "+
				"rotate, "+
				"target, "+
				"queue_attack, "+
				"hp, "+
				"power, "+
				"mother_ship, "+
				"on_map, "+
				"action_point, "+
				"defend"+
				") "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
				squad.ID,
				ship.Body.ID,
				ship.Q,
				ship.R,
				ship.Rotate,
				parseTarget(ship),
				ship.QueueAttack,
				ship.HP,
				ship.Power,
				true, // mother_ship
				true, // on_map
				ship.Speed,
				ship.Defend,
			).Scan(&id)
			if err != nil {
				log.Fatal("add new ship to squad " + err.Error())
			}

			ship.ID = id

			if ship.Body != nil {
				UpdateBody(ship, squad.ID, tx)
			}

		} else {
			_, err := tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete all unit equip " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete all units " + err.Error())
			}
		}
	}
}
