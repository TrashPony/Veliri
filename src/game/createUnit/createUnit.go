package createUnit

import (
	"database/sql"
	"log"
)

func CreateUnit(idGame string, idPlayer string, unitType string, x string, y string)(bool, int) {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд

	if err != nil {
		log.Fatal(err)
	}

	unit := GetUnitType(unitType, idGame, idPlayer)
	success, price := Price(unit.price, idGame, idPlayer)

	if (success) {
		rows, err := db.Query("INSERT INTO action_game_unit (id_game, id_type, id_user, hp, action, target, x, y) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			idGame, unit.id, idPlayer, unit.hp, true, 0, x, y)
		defer rows.Close()
		if err != nil {
			log.Fatal(err)
		}

		return success, price
	}
	return false, 0
}

func GetUnitType(unitType string, idGame string, idPlayer string) (UnitType)  {
	var unit UnitType

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id, hp, price From unittype WHERE type=" + "'" +unitType + "'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&unit.id, &unit.hp, &unit.price)
		if err != nil {
			log.Fatal(err)
		}
	}

	return unit
}

func Price(cost int, idGame string, idPlayer string) (bool, int) {
	var price int

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select price FROM action_game_user WHERE id_game=" + idGame + " AND id_user=" + idPlayer)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&price)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cost < price {
		price = price - cost
		_ , err := db.Exec("UPDATE action_game_user SET price = $1 WHERE id_game = $2 AND id_user = $3", price, idGame, idPlayer)
		if err != nil {
			log.Fatal(err)
		} else {
			return true, price
		}

	}
	return false, 0
}

type UnitType struct {
	id int
	hp int
	price int
}