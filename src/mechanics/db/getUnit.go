package db

import (
	"log"
	"database/sql"
	"strings"
	"strconv"
	"../../detailUnit"
	"../coordinate"
	"../unit"
)

func GetAllUnits(idGame int) (map[int]map[int]*unit.Unit, []*unit.Unit ){

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



	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)

	var targetKey string

	chassisID := sql.NullInt64{}
	weaponID := sql.NullInt64{}
	towerID := sql.NullInt64{}
	bodyID := sql.NullInt64{}
	radarID := sql.NullInt64{}

	for rows.Next() {
		var gameUnit unit.Unit
		err := rows.Scan(&gameUnit.Id, &gameUnit.Owner, &gameUnit.X, &gameUnit.Y, &gameUnit.Rotate, &gameUnit.Action, &targetKey, &gameUnit.Queue,
			&gameUnit.Weight, &gameUnit.MoveSpeed, &gameUnit.Initiative, &gameUnit.Damage, &gameUnit.RangeAttack, &gameUnit.MinAttackRange, &gameUnit.AreaAttack,
			&gameUnit.TypeAttack, &gameUnit.HP, &gameUnit.Armor, &gameUnit.EvasionCritical, &gameUnit.VulKinetics, &gameUnit.VulThermal, &gameUnit.VulEM,
			&gameUnit.VulExplosive, &gameUnit.RangeView, &gameUnit.Accuracy, &gameUnit.WallHack, &chassisID, &weaponID, &towerID, &bodyID, &radarID, &gameUnit.OnMap)
		if err != nil {
			println("scan game unit")
			log.Fatal(err)
		}

		gameUnit.Target = ParseUnitTarget(targetKey)
		SetDetails(&gameUnit, chassisID, weaponID, towerID, bodyID, radarID)

		if gameUnit.OnMap {
			if units[gameUnit.X] != nil { // кладем юнита в матрицу
				units[gameUnit.X][gameUnit.Y] = &gameUnit
			} else {
				units[gameUnit.X] = make(map[int]*unit.Unit)
				units[gameUnit.X][gameUnit.Y] = &gameUnit
			}
		} else {
			unitStorage = append(unitStorage, &gameUnit)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return units, unitStorage
}

func ParseUnitTarget(targetKey string) *coordinate.Coordinate {
	targetCell := strings.Split(targetKey, ":")

	if len(targetCell) > 1 { // устанавливаем таргет если он есть
		x, ok := strconv.Atoi(targetCell[0])
		y, ok := strconv.Atoi(targetCell[1])
		if ok == nil {
			target := coordinate.Coordinate{X: x, Y: y}
			return &target
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func SetDetails(unit *unit.Unit, chassisID, weaponID, towerID, bodyID, radarID sql.NullInt64)  {
	if chassisID.Valid {
		chassis := detailUnit.GetChass(int(chassisID.Int64))
		unit.SetChassis(chassis)
	}
	if weaponID.Valid {
		weapon := detailUnit.GetWeapon(int(weaponID.Int64))
		unit.SetWeapon(weapon)
	}
	if towerID.Valid {
		tower := detailUnit.GetTower(int(towerID.Int64))
		unit.SetTower(tower)
	}
	if bodyID.Valid {
		body := detailUnit.GetBody(int(bodyID.Int64))
		unit.SetBody(body)
	}
	if radarID.Valid {
		radar := detailUnit.GetRadar(int(radarID.Int64))
		unit.SetRadar(radar)
	}
}
