package update

import (
	"../../../gameObjects/unit"
	//"strconv"
	//"../../dbConnect"
)

func Unit(unit *unit.Unit) error {

	//var target string
	// todo переделать под новую структуру бд
	if unit.Target != nil {
		//target = strconv.Itoa(unit.Target.X) + ":" + strconv.Itoa(unit.Target.Y)
	}

	/*_, err := dbConnect.GetDBConnect().Exec("UPDATE action_game_unit "+
		"SET x=$2, y=$3, rotate=$4, on_map=$5, "+
		"action=$6, target=$7, queue_attack=$8, "+
		"weight=$9, speed=$10, initiative=$11, damage=$12, range_attack=$13, min_attack_range=$14, area_attack=$15, type_attack=$16, "+
		"hp=$17, max_hp=$18, armor=$19, evasion_critical=$20, vul_kinetics=$21, vul_thermal=$22, vul_em=$23, vul_explosive=$24, range_view=$25, "+
		"accuracy=$26, wall_hack=$27 "+

		"WHERE id=$1", unit.Id,

		unit.X, unit.Y, unit.Rotate, unit.OnMap,
		unit.Action, target, unit.Queue,
		unit.Weight, unit.MoveSpeed, unit.Initiative, unit.Damage, unit.RangeAttack, unit.MinAttackRange, unit.AreaAttack, unit.TypeAttack,
		unit.HP, unit.MaxHP, unit.Armor, unit.EvasionCritical, unit.VulKinetics, unit.VulThermal, unit.VulEM, unit.VulExplosive, unit.RangeView,
		unit.Accuracy, unit.WallHack)

	if err != nil {
		println("Error update unit params")
		return err
	}

	err = UpdateUnitEffects(unit)
	if err != nil {
		return err
	}*/

	return nil
}