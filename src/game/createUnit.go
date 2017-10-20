package game

import (
	"database/sql"
	"log"
	"./initGame"
	"errors"
)

func CreateUnit(idGame string, idPlayer string, unitType string, x string, y string)(initGame.Unit, int, error) {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд

	if err != nil {
		log.Fatal(err)
	}
	var unit initGame.Unit
	checkPlace := CheckPlace(idGame, x, y)
	if checkPlace { // если место не занято то дидем дальше
		unitType := initGame.GetUnitType(unitType)
		success, price := Price(unitType.Price, idGame, idPlayer)
		if success { // если хватило денег то вносим изменения , вероятно тут надо применить транзацию
			rows, err := db.Query("INSERT INTO action_game_unit (id_game, id_type, id_user, hp, action, target, x, y) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				idGame, unitType.Id, idPlayer, unitType.Hp, true, 0, x, y)
			defer rows.Close()
			if err != nil {
				log.Fatal(err)
			}
			unit, errFound := initGame.GetXYUnits(idGame, x, y)
			if errFound != nil {
				return unit, 0, errFound
			}
			return unit, price, nil
		} else {
			return unit, 0, errors.New("no many")
		}
	} else {
		return unit, 0, errors.New("busy")
	}
	return unit, 0, errors.New("unknown error")
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

	if cost <= price {
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

func CheckPlace(idGame string, x string, y string)(bool)  {
	var id int

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select id FROM action_game_unit WHERE x = $1 AND y = $2 AND id_game= $3", x, y, idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}

	if id == 0 {
		return true
	} else {
		return false
	}
}