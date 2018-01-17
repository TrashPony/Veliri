package game

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func GetUnit(query string) map[int]map[int]*Unit {

	rows, err := db.Query("Select ag.id, ag.id_game, t.damage, t.move_speed, t.initiative, t.range_attack, t.range_view, t.area_attack, t.type_attack, t.price, t.type, u.name, ag.hp, ag.action, ag.target, ag.x, ag.y, ag.rotate, ag.queue_attack FROM action_game_unit as ag, unit_type as t, users as u WHERE " + query)
	if err != nil { // TODO неправильный запрос
		log.Fatal(err)
	}
	defer rows.Close()

	var units = make(map[int]map[int]*Unit)
	var targetKey string

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.Id, &unit.IdGame, &unit.Damage, &unit.MoveSpeed, &unit.Initiative, &unit.RangeAttack, &unit.WatchZone, &unit.AreaAttack,
			&unit.TypeAttack, &unit.Price, &unit.ChassisType, &unit.WeaponType, &unit.Owner, &unit.Hp, &unit.Action, &targetKey, &unit.X, &unit.Y, &unit.Queue)
		if err != nil {
			log.Fatal(err)
		}

		targ := strings.Split(targetKey, ":")

		if len(targ) > 1 {
			x, ok := strconv.Atoi(targ[0])
			y, ok := strconv.Atoi(targ[1])
			if ok == nil {
				target := Coordinate{X: x, Y: y}
				unit.Target = &target
			}
		}

		if units[unit.X] != nil {
			units[unit.X][unit.Y] = &unit
		} else {
			units[unit.X] = make(map[int]*Unit)
			units[unit.X][unit.Y] = &unit
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return units
}

func GetAllUnits(idGame int) map[int]map[int]*Unit {
	units := GetUnit(" ag.id_game=" + strconv.Itoa(idGame) + " AND ag.id_type=t.id AND ag.id_user=u.id")
	return units
}

func GetXYUnits(idGame int, x int, y int) (Unit, error) {
	units := GetUnit(" ag.id_game=" + strconv.Itoa(idGame) + " AND ag.id_type=t.id AND ag.id_user=u.id AND ag.x=" + strconv.Itoa(x) + "AND ag.y=" + strconv.Itoa(y))
	if len(units) > 0 {
		unit, ok := units[x][y]
		if ok {
			return *unit, nil
		} else {
			return *unit, errors.New("unit not found")
		}
	} else {
		var unit Unit
		return unit, errors.New("unit not found")
	}
}