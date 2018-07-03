package updateSquad

import (
	"../../gameObjects/squad"
	"../../../dbConnect"
	"log"
	"strconv"
)

func MotherShip(squad *squad.Squad) {

	ship := squad.MatherShip

	if ship != nil && ship.ID != 0 && ship.Body != nil {

		var target string

		if ship.Target != nil {
			target = strconv.Itoa(ship.Target.X) + ":" + strconv.Itoa(ship.Target.Y)
		}

		_, err := dbConnect.GetDBConnect().Exec(
			"UPDATE squad_mother_ship "+
				"SET id_body = $1, x = $2, y = $3, rotate = $4, action = $5, target = $6, queue_attack = $7, hp = $8 "+
				"WHERE id_squad = $9",
			ship.Body.ID, ship.X, ship.Y, ship.Rotate, ship.Action, target, ship.QueueAttack, ship.HP, squad.ID)

		if err != nil {
			log.Fatal("update motherShip squad" + err.Error())
		}

		//todo обновление туши
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
