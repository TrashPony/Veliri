package update

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"log"
)

func MotherShip(squad *squad.Squad, tx *sql.Tx) {

	ship := squad.MatherShip

	var bodyID sql.NullInt64

	if ship.Body == nil {
		bodyID = sql.NullInt64{Int64: 0, Valid: false}
	} else {
		bodyID = sql.NullInt64{Int64: int64(ship.Body.ID), Valid: true}
	}

	if ship != nil && ship.ID != 0 {

		if ship.Body == nil {
			_, err := tx.Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1 AND id_squad_unit=$2", squad.ID, ship.ID)
			if err != nil {
				log.Fatal("delete all unit equip " + err.Error())
			}

			_, err = tx.Exec("DELETE FROM squad_units_inventory WHERE id_unit=$1", ship.ID)
			if err != nil {
				log.Fatal("delete inventory ms " + err.Error())
			}

		} else {
			UpdateBody(ship, squad.ID, tx)
		}

		_, err := tx.Exec(
			"UPDATE squad_units SET "+
				"id_body = $1, "+
				"x = $2, "+
				"y = $3, "+
				"rotate = $4, "+
				"hp = $5, "+
				"power = $6, "+
				"action_point = $7, "+
				"defend = $8,"+
				"move = $11,"+
				"body_color_1 = $12,"+
				"body_color_2 = $13,"+
				"weapon_color_1 = $14,"+
				"weapon_color_2 = $15,"+
				"id_map = $16 "+
				"WHERE id_squad = $9 AND mother_ship = $10",
			bodyID,
			ship.X,
			ship.Y,
			ship.Rotate,
			ship.HP,
			ship.Power,
			ship.ActionPoints,
			ship.Defend,
			squad.ID,
			true, // mother_ship = $10
			ship.Move,
			ship.BodyColor1,
			ship.BodyColor2,
			ship.WeaponColor1,
			ship.WeaponColor2,
			ship.MapID,
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
				"x, "+
				"y, "+
				"rotate, "+
				"hp, "+
				"power, "+
				"mother_ship, "+
				"on_map, "+
				"action_point, "+
				"defend, "+
				"move,"+
				"id_map "+
				") "+
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id",
				squad.ID,
				bodyID,
				ship.X,
				ship.Y,
				ship.Rotate,
				ship.HP,
				ship.Power,
				true, // mother_ship
				true, // on_map
				ship.Speed,
				ship.Defend,
				ship.Move,
				ship.MapID,
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
