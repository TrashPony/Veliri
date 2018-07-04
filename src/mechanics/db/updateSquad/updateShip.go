package updateSquad

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
)

func MotherShip(squad *squad.Squad) {

	ship := squad.MatherShip

	if ship != nil && ship.ID != 0 && ship.Body != nil {

		_, err := dbConnect.GetDBConnect().Exec(
			"UPDATE squad_mother_ship "+
				"SET id_body = $1, x = $2, y = $3, rotate = $4, action = $5, target = $6, queue_attack = $7, hp = $8 "+
				"WHERE id_squad = $9",
			ship.Body.ID, ship.X, ship.Y, ship.Rotate, ship.Action, parseTarget(ship), ship.QueueAttack, ship.HP, squad.ID)

		if err != nil {
			log.Fatal("update motherShip squad" + err.Error())
		}

		UpdateBody(ship, squad.ID, "squad_mother_ship_equipping")
		//todo обновление эфектов

	} else {
		if ship.ID == 0 && ship.Body != nil {
			id := 0
			err := dbConnect.GetDBConnect().QueryRow("INSERT INTO squad_mother_ship (id_squad, id_body, x, y, rotate, on_map, action, target, queue_attack, hp) " +
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
				squad.ID, ship.Body.ID, ship.X, ship.Y, ship.Rotate, ship.OnMap, ship.Action,
				parseTarget(ship), ship.QueueAttack, ship.HP).Scan(&id)
			if err != nil {
				log.Fatal("add new unit to squad " + err.Error())
			}

			ship.ID = id

			UpdateBody(ship, squad.ID, "squad_mother_ship_equipping")
			//todo обновление эфектов

		} else {
			_, err := dbConnect.GetDBConnect().Exec("DELETE FROM squad_mother_ship_equipping WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete all ship equip " + err.Error())
			}

			_, err = dbConnect.GetDBConnect().Exec("DELETE FROM squad_mother_ship WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete ship" + err.Error())
			}

			_, err = dbConnect.GetDBConnect().Exec("DELETE FROM squad_units_equipping WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete all unit equip " + err.Error())
			}

			_, err = dbConnect.GetDBConnect().Exec("DELETE FROM squad_units WHERE id_squad=$1",
				squad.ID)
			if err != nil {
				log.Fatal("delete all units " + err.Error())
			}
		}
	}
}
