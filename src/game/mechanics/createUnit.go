package mechanics

import (
	"log"
	"../objects"
	"errors"
	"strconv"
)

func CreateUnit(idGame int, idPlayer string, unitType string, x int, y int)(*objects.Unit, int, error) {

	var unit objects.Unit
	checkPlace := CheckPlace(idGame, x, y)
	if checkPlace { // если место не занято то дидем дальше
		unitType := objects.GetUnitType(unitType)
		success, price := Price(unitType.Price, idGame, idPlayer)
		if success { // если хватило денег то вносим изменения , вероятно тут надо применить транзацию
			rows, err := db.Query("INSERT INTO action_game_unit (id_game, id_type, id_user, hp, action, target, x, y, queue_attack) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
				idGame, unitType.Id, idPlayer, unitType.Hp, true, "", x, y, 0)
			defer rows.Close()
			if err != nil {
				log.Fatal(err)
			}
			unit, errFound := objects.GetXYUnits(idGame, x, y)
			if errFound != nil {
				return &unit, 0, errFound
			}
			return &unit, price, nil
		} else {
			return &unit, 0, errors.New("no many")
		}
	} else {
		return &unit, 0, errors.New("busy")
	}
	return &unit, 0, errors.New("unknown error")
}

func Price(cost int, idGame int, idPlayer string) (bool, int) {
	var price int

	rows, err := db.Query("Select price FROM action_game_user WHERE id_game=" + strconv.Itoa(idGame) + " AND id_user=" + idPlayer)
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

func CheckPlace(idGame int, x int, y int)(bool)  {
	var id int

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