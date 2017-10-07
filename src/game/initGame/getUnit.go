package initGame

import (
	"database/sql"
	"log"
	"strconv"
)

func GetUnit(idGame string) ([]Unit)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM action_game_unit WHERE id_game=$1", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make([]Unit, 0)
	var unit Unit

	for rows.Next() {
		err := rows.Scan(&unit.id, &unit.id_game, &unit.id_type, &unit.id_user, &unit.hp, &unit.action, &unit.target, &unit.x, &unit.y,)
		if err != nil {
			log.Fatal(err)
		}
		units = append(units, unit)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return units
}

func GetUnitList(idGame string) ([]string)  {
	units := GetUnit(idGame)
	unitList := make([]string, 0)

	var idType = ""
	var idUser = ""
	var hp = ""
	var action = ""
	var target = ""
	var x = ""
	var y = ""
	// ндеееее :\
	for i := 0; i < len(units); i++ {
		idType = idType + GetUnitType(units[i].id_type) + ":"
		idUser = idUser + strconv.Itoa(units[i].id_user) + ":"
		hp = hp + strconv.Itoa(units[i].hp) + ":"
		action = action + strconv.FormatBool(units[i].action) + ":"
		target = target + strconv.Itoa(units[i].target) + ":"
		x = x + strconv.Itoa(units[i].x) + ":"
		y = y + strconv.Itoa(units[i].y) + ":"
	}
	// ндооооо :\
	unitList = append(unitList, idType)
	unitList = append(unitList, idUser)
	unitList = append(unitList, hp)
	unitList = append(unitList, action)
	unitList = append(unitList, target)
	unitList = append(unitList, x)
	unitList = append(unitList, y)
	// мдеееее :\

	return unitList
}

func GetUnitType(idType int) (string) {
	var unitType string

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select type From unittype WHERE id=$1", idType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&unitType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return unitType
}


type Unit struct {
	id int
	id_game int
	id_type int
	id_user int
	hp int
	action bool
	target int
	x int
	y int
}