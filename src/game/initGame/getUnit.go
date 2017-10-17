package initGame

import (
	"database/sql"
	"log"
)

func GetUnits(idGame string) ([]Unit)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select ag.id, ag.id_game, t.type, u.name, ag.hp, ag.action, ag.target, ag.x, ag.y FROM action_game_unit as ag, unittype as t, users as u WHERE ag.id_game=$1 AND ag.id_type=t.id AND ag.id_user=u.id", idGame)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make([]Unit, 0)
	var unit Unit

	for rows.Next() {
		err := rows.Scan(&unit.Id, &unit.Id_game, &unit.NameType, &unit.NameUser, &unit.Hp, &unit.Action, &unit.Target, &unit.X, &unit.Y,)
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

/*func GetUnit(idGame string, x string, y string)(bool, []string)  {
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
*/
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
	Id int
	Id_game int
	NameType string
	NameUser string
	Hp int
	Action bool
	Target int
	X int
	Y int
}