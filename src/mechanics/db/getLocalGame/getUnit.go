package getLocalGame

import (
	"database/sql"
	"strings"
	"strconv"
	"../../localGame/map/coordinate"
	"../../gameObjects/unit"
	//"../../dbConnect"
)

func AllUnits(idGame int) (map[int]map[int]*unit.Unit, []*unit.Unit ){

	/*rows, err := dbConnect.GetDBConnect().Query("Select ag.id, users.name, ag.x, ag.y, ag.rotate, ag.action, ag.target, ag.queue_attack, "+
		"ag.weight, ag.speed, ag.initiative, ag.damage, ag.range_attack, ag.min_attack_range, ag.area_attack, "+
		"ag.type_attack, ag.hp, ag.max_hp, ag.armor, ag.evasion_critical, ag.vul_kinetics, ag.vul_thermal, ag.vul_em, "+
		"ag.vul_explosive, ag.range_view, ag.accuracy, ag.wall_hack, "+
		"ag.id_chassis, ag.id_weapons, ag.id_tower, ag.id_body, ag.id_radar, ag.on_map "+
		"FROM action_game_unit as ag, users WHERE ag.id_game=$1 AND ag.id_user=users.id", idGame)
	if err != nil {
		println("get game unit")
		log.Fatal(err)
	}
	defer rows.Close()


*/ // todo переделать под новые запросы
	var units = make(map[int]map[int]*unit.Unit)
	var unitStorage = make([]*unit.Unit, 0)

	/*var targetKey string

	chassisID := sql.NullInt64{}
	weaponID := sql.NullInt64{}
	towerID := sql.NullInt64{}
	bodyID := sql.NullInt64{}
	radarID := sql.NullInt64{}

	for rows.Next() {
		var gameUnit unit.Unit
		err := rows.Scan(&gameUnit.Id, &gameUnit.Owner, &gameUnit.X, &gameUnit.Y, &gameUnit.Rotate, &gameUnit.Action, &targetKey, &gameUnit.Queue,
			&gameUnit.Weight, &gameUnit.MoveSpeed, &gameUnit.Initiative, &gameUnit.Damage, &gameUnit.RangeAttack, &gameUnit.MinAttackRange, &gameUnit.AreaAttack,
			&gameUnit.TypeAttack, &gameUnit.HP, &gameUnit.MaxHP, &gameUnit.Armor, &gameUnit.EvasionCritical, &gameUnit.VulKinetics, &gameUnit.VulThermal, &gameUnit.VulEM,
			&gameUnit.VulExplosive, &gameUnit.RangeView, &gameUnit.Accuracy, &gameUnit.WallHack, &chassisID, &weaponID, &towerID, &bodyID, &radarID, &gameUnit.OnMap)
		if err != nil {
			println("scan game unit")
			log.Fatal(err)
		}

		gameUnit.Target = ParseUnitTarget(targetKey)

		SetDetails(&gameUnit, chassisID, weaponID, towerID, bodyID, radarID)
		GetUnitEffects(&gameUnit)

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
	}*/

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

func SetDetails(unit *unit.Unit, weaponID, bodyID sql.NullInt64)  {
	if weaponID.Valid {
		//weapon := detailUnit.GetWeapon(int(weaponID.Int64))
		//unit.SetWeapon(weapon)
	}
	if bodyID.Valid {
		//body := detailUnit.GetBody(int(bodyID.Int64))
		//unit.SetBody(body)
	}
}
