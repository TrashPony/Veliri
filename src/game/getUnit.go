package game

import (
	"strconv"
	"log"
	"strings"
	"database/sql"
	"../DetailUnit"
)

func GetAllUnits(idGame int) (map[int]map[int]*Unit, []*Unit ){

	rows, err := db.Query("Select ag.id, users.name, ag.x, ag.y, ag.rotate, ag.action, ag.target, ag.queue_attack, "+
		"ag.Weight, ag.Speed, ag.Initiative, ag.Damage, ag.RangeAttack, ag.MinAttackRange, ag.AreaAttack, "+
		"ag.TypeAttack, ag.HP, ag.Armor, ag.EvasionCritical, ag.VulKinetics, ag.VulThermal, ag.VulEM, "+
		"ag.VulExplosive, ag.RangeView, ag.Accuracy, ag.WallHack, "+
		"ag.id_chassis, ag.id_weapons, ag.id_tower, ag.id_body, ag.id_radar, ag.on_map "+
		"FROM action_game_unit as ag, users WHERE ag.id_game=$1 AND ag.id_user=users.id", idGame)
	if err != nil {
		println("get game unit")
		log.Fatal(err)
	}
	defer rows.Close()

	/*

	Select ag.id, users.name, ag.x, ag.y, ag.rotate, ag.action, ag.target, ag.queue_attack, ag.Weight, ag.Speed FROM action_game_unit as ag, users WHERE ag.id_game=1 AND ag.id_user=users.id;
	 */

	var units = make(map[int]map[int]*Unit)
	var notGameUnits = make([]*Unit, 0)

	var targetKey string

	chassisID := sql.NullInt64{}
	weaponID := sql.NullInt64{}
	towerID := sql.NullInt64{}
	bodyID := sql.NullInt64{}
	radarID := sql.NullInt64{}

	for rows.Next() {
		var unit Unit
		err := rows.Scan(&unit.Id, &unit.Owner, &unit.X, &unit.Y, &unit.Rotate, &unit.Action, &targetKey, &unit.Queue,
			&unit.Weight, &unit.MoveSpeed, &unit.Initiative, &unit.Damage, &unit.RangeAttack, &unit.MinAttackRange, &unit.AreaAttack,
			&unit.TypeAttack, &unit.HP, &unit.Armor, &unit.EvasionCritical, &unit.VulKinetics, &unit.VulThermal, &unit.VulEM,
			&unit.VulExplosive, &unit.RangeView, &unit.Accuracy, &unit.WallHack, &chassisID, &weaponID, &towerID, &bodyID, &radarID, &unit.OnMap)
		if err != nil {
			println("scan game unit")
			log.Fatal(err)
		}

		unit.Target = ParseUnitTarget(targetKey)
		SetDetails(&unit, chassisID, weaponID, towerID, bodyID, radarID)

		if unit.OnMap {
			if units[unit.X] != nil { // кладем юнита в матрицу
				units[unit.X][unit.Y] = &unit
			} else {
				units[unit.X] = make(map[int]*Unit)
				units[unit.X][unit.Y] = &unit
			}
		} else {
			notGameUnits = append(notGameUnits, &unit)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return units, notGameUnits
}

func ParseUnitTarget(targetKey string) *Coordinate {
	targetCell := strings.Split(targetKey, ":")

	if len(targetCell) > 1 { // устанавливаем таргет если он есть
		x, ok := strconv.Atoi(targetCell[0])
		y, ok := strconv.Atoi(targetCell[1])
		if ok == nil {
			target := Coordinate{X: x, Y: y}
			return &target
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func SetDetails(unit *Unit, chassisID, weaponID, towerID, bodyID, radarID sql.NullInt64)  {
	if chassisID.Valid {
		chassis := DetailUnit.GetChass(int(chassisID.Int64))
		unit.SetChassis(chassis)
	}
	if weaponID.Valid {
		weapon := DetailUnit.GetWeapon(int(weaponID.Int64))
		unit.SetWeapon(weapon)
	}
	if towerID.Valid {
		tower := DetailUnit.GetTower(int(towerID.Int64))
		unit.SetTower(tower)
	}
	if bodyID.Valid {
		body := DetailUnit.GetBody(int(bodyID.Int64))
		unit.SetBody(body)
	}
	if radarID.Valid {
		radar := DetailUnit.GetRadar(int(radarID.Int64))
		unit.SetRadar(radar)
	}
}
