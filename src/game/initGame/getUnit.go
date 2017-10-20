package initGame

import (
	"database/sql"
	"log"
	"errors"
)

func GetUnit(query string) ([]Unit)  {
	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select ag.id, ag.id_game, t.damage, t.movespeed, t.init, t.rangeattack, t.rangeview, t.areaattack, t.typeattack, t.price, t.type, u.name, ag.hp, ag.action, ag.target, ag.x, ag.y FROM action_game_unit as ag, unittype as t, users as u WHERE " + query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make([]Unit, 0)
	var unit Unit

	for rows.Next() {
		err := rows.Scan(&unit.Id, &unit.IdGame, &unit.Damage, &unit.MoveSpeed, &unit.Init, &unit.RangeAttack, &unit.WatchZone, &unit.AreaAttack,
			&unit.TypeAttack, &unit.Price, &unit.NameType, &unit.NameUser, &unit.Hp, &unit.Action, &unit.Target, &unit.X, &unit.Y)
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

func GetAllUnits(idGame string)([]Unit)  {
	units := GetUnit(" ag.id_game=" + idGame + " AND ag.id_type=t.id AND ag.id_user=u.id")
	return units
}

func GetXYUnits(idGame string, x string, y string)(Unit, error)  {
	units := GetUnit(" ag.id_game=" + idGame + " AND ag.id_type=t.id AND ag.id_user=u.id AND ag.x=" + x + "AND ag.y=" + y)
	if len(units) > 0 {
		return units[0], nil
	} else {
		var unit Unit
		units = append(units, unit)
		return units[0], errors.New("unit not found")
	}
}

func GetUnitType(nameType string) (UnitType) {
	var unitType UnitType

	db, err := sql.Open("postgres", "postgres://postgres:yxHie25@192.168.101.95:5432/game") // подключаемся к нашей бд
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("Select * From unittype WHERE type=$1", nameType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&unitType.Id, &unitType.Type, &unitType.Damage, &unitType.Hp,
			&unitType.MoveSpeed, &unitType.Init, &unitType.RangeAttack, &unitType.WatchZone,
			&unitType.AreaAttack, &unitType.TypeAttack, &unitType.Price)
		if err != nil {
			log.Fatal(err)
		}
	}
	return unitType
}

func GetUnitsCoordinate(units []Unit)([]Coordinate)  {
	var coordinates []Coordinate
	for i := 0; i < len(units); i++ {
		var coordinate Coordinate
		coordinate.X = units[i].X
		coordinate.Y = units[i].Y
		coordinates = append(coordinates, coordinate)
	}
	return coordinates
}