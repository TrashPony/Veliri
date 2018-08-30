package updateSquad

import (
	"../../gameObjects/squad"
	"log"
	"database/sql"
)

func MotherShip(squad *squad.Squad, tx *sql.Tx) {

	ship := squad.MatherShip

	if ship != nil && ship.ID != 0 {

		var bodyID sql.NullInt64

		if ship.Body == nil {
			bodyID = sql.NullInt64{Int64:0, Valid: false}

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
			"UPDATE squad_units "+
				"SET id_body = $1, q = $2, r = $3, rotate = $4, action = $5, target = $6, queue_attack = $7, hp = $8, use_equip = $10, power = $11 "+
				"WHERE id_squad = $9 AND mother_ship = $12",
			bodyID, ship.Q, ship.R, ship.Rotate, ship.Action, parseTarget(ship), ship.QueueAttack, ship.HP, squad.ID, ship.UseEquip, ship.Power, true)

		if err != nil {
			log.Fatal("update motherShip squad" + err.Error())
		}

	} else {
		if ship.ID == 0 || ship.Body != nil {
			id := 0
			err := tx.QueryRow("INSERT INTO squad_units (id_squad, id_body, q, r, rotate, action, target, queue_attack, hp, use_equip, power, mother_ship, on_map) " +
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
				squad.ID, ship.Body.ID, ship.Q, ship.R, ship.Rotate, ship.Action,
				parseTarget(ship), ship.QueueAttack, ship.HP, ship.UseEquip, ship.Power, true, true).Scan(&id)
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
