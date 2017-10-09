package initGame

import (
	"database/sql"
	"log"
	"strconv"
)

func GetUnits(query string) ([]Unit)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * FROM action_game_unit " + query)
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

func GetUnit(idGame string, x string, y string)(bool, []string)  {
	unit := GetUnits("WHERE id_game=" + idGame + " AND x="+ x + " AND y=" + y)
	unitParam := make([]string, 0)

	if len(unit) > 0 {
		typeUnit := GetUnitType(unit[0].id_type)
		unitParam = append(unitParam, typeUnit.Type)                      // 0
		unitParam = append(unitParam, GetUserName(unit[0].id_user))       // 1
		unitParam = append(unitParam, strconv.Itoa(unit[0].hp))           // 2
		unitParam = append(unitParam, strconv.FormatBool(unit[0].action)) // 3
		unitParam = append(unitParam, strconv.Itoa(unit[0].target))       // 4
		unitParam = append(unitParam, strconv.Itoa(typeUnit.damage))      // 5
		unitParam = append(unitParam, strconv.Itoa(typeUnit.movespeed))   // 6
		unitParam = append(unitParam, strconv.Itoa(typeUnit.init))        // 7
		unitParam = append(unitParam, strconv.Itoa(typeUnit.rangeattack)) // 8
		unitParam = append(unitParam, strconv.Itoa(typeUnit.rangeview))   // 9
		unitParam = append(unitParam, strconv.Itoa(typeUnit.areaattack))  // 10
		unitParam = append(unitParam, typeUnit.typeattack)                // 11

		return true, unitParam
	}
	return false, unitParam
}

func GetUnitList(idGame string) ([]string)  {
	units := GetUnits("WHERE id_game=" + idGame)
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
		unitType := GetUnitType(units[i].id_type)
		idType = idType + unitType.Type + ":"
		idUser = idUser + GetUserName(units[i].id_user) + ":"
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

func GetUserName(idUser int) (string) {
	var UserName string

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select name From users WHERE id=$1", idUser)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&UserName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return UserName
}

func GetUnitType(idType int) (UnitType) {
	var unitType UnitType

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * From unittype WHERE id=$1", idType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&unitType.id, &unitType.Type, &unitType.damage, &unitType.hp,
			&unitType.movespeed, &unitType.init, &unitType.rangeattack, &unitType.rangeview,
			&unitType.areaattack, &unitType.typeattack, &unitType.price)
		if err != nil {
			log.Fatal(err)
		}
	}

	return unitType
}

type UnitType struct {
	id int
	Type string
	damage int
	hp int
	movespeed int
	init int
	rangeattack int
	rangeview int
	areaattack int
	typeattack string
	price int
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