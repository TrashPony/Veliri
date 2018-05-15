package Squad

import "../../detailUnit"

type Unit struct {
	Chassis *detailUnit.Chassis `json:"chassis"`
	Weapon  *detailUnit.Weapon  `json:"weapon"`
	Tower   *detailUnit.Tower   `json:"tower"`
	Body    *detailUnit.Body    `json:"body"`
	Radar   *detailUnit.Radar   `json:"radar"`

	Weight int `json:"weight"`

	// Движение
	Speed      int `json:"speed"`
	Initiative int `json:"initiative"`

	// Атака
	Damage         int    `json:"damage"`
	RangeAttack    int    `json:"range_attack"`
	MinAttackRange int    `json:"min_attack_range"`
	AreaAttack     int    `json:"area_attack"`
	TypeAttack     string `json:"type_attack"`

	// Выживаемость
	HP              int `json:"hp"`
	Armor           int `json:"armor"`
	EvasionCritical int `json:"evasion_critical"`
	VulKinetics     int `json:"vul_kinetics"` // Уязвимости
	VulThermal      int `json:"vul_thermal"`
	VulEM           int `json:"vul_em"`
	VulExplosive    int `json:"vul_explosive"`

	// Навигация
	RangeView int  `json:"range_view"`
	Accuracy  int  `json:"accuracy"`
	WallHack  bool `json:"wall_hack"`
}

func (unit *Unit) CalculateParametersUnit() {
	unit.ZeroingStats() // обнуляем стату юнита

	weight := WeightUnit(unit)
	unit.Weight = weight

	// расчет шасси
	if unit.Chassis != nil {
		unit.EvasionCritical = unit.Chassis.Maneuverability * 2
		unit.Initiative = unit.Initiative + unit.Chassis.Maneuverability

		percentWeight := unit.Chassis.Carrying * (weight / 100)

		if percentWeight > 75 {
			percentWeight = percentWeight - 75
			fine := (percentWeight / 5) + 1
			unit.Speed = unit.Chassis.Speed - fine
		} else {
			unit.Speed = unit.Chassis.Speed
		}
	}

	// расчет оружия
	if unit.Weapon != nil {
		unit.TypeAttack = unit.Weapon.Type
		unit.Damage = unit.Weapon.Damage
		unit.MinAttackRange = unit.Weapon.MinAttackRange
		unit.RangeAttack = unit.Weapon.Range
		unit.AreaAttack = unit.Weapon.AreaCovers

		unit.Accuracy = unit.Accuracy + unit.Weapon.Accuracy
	}

	// расчет башни
	if unit.Tower != nil {
		unit.HP = unit.HP + unit.Tower.HP
		unit.RangeView = unit.RangeView + unit.Tower.PowerRadar

		unit.VulEM = unit.VulEM + unit.Tower.VulToEM
		unit.VulExplosive = unit.VulExplosive + unit.Tower.VulToExplosion
		unit.VulKinetics = unit.VulKinetics + unit.Tower.VulToKinetics
		unit.VulThermal = unit.VulThermal + unit.Tower.VulToThermo

		unit.Armor = unit.Armor + unit.Tower.Armor
	}

	// Расчет корпуса
	if unit.Body != nil {
		unit.HP = unit.HP + unit.Body.HP

		unit.VulEM = unit.VulEM + unit.Body.VulToEM
		unit.VulExplosive = unit.VulExplosive + unit.Body.VulToExplosion
		unit.VulKinetics = unit.VulKinetics + unit.Body.VulToKinetics
		unit.VulThermal = unit.VulThermal + unit.Body.VulToThermo

		unit.Armor = unit.Armor + unit.Body.Armor
	}

	// расчет навигации
	if unit.Radar != nil {
		unit.RangeView = unit.RangeView + unit.Radar.Power
		unit.WallHack = unit.Radar.Through

		unit.Initiative = unit.Initiative + unit.Radar.Analysis
		unit.Accuracy = unit.Accuracy + unit.Radar.Analysis
	}
}

func (unit *Unit) ZeroingStats() {
	unit.Weight = 0
	// todo это немного странно но другово решения я не нашел
	// Движение
	unit.Speed = 0
	unit.Initiative = 0

	// Атака
	unit.Damage = 0
	unit.RangeAttack = 0
	unit.MinAttackRange = 0
	unit.AreaAttack = 0
	unit.TypeAttack = ""

	// Выживаемость
	unit.HP = 0
	unit.Armor = 0
	unit.EvasionCritical = 0
	unit.VulKinetics = 0
	unit.VulThermal = 0
	unit.VulEM = 0
	unit.VulExplosive = 0

	// Навигация
	unit.RangeView = 0
	unit.Accuracy = 0
	unit.WallHack = false
}

func (unit *Unit) DelChassis() {
	if unit.Chassis != nil {
		unit.EvasionCritical = 0
		unit.Initiative = unit.Initiative - unit.Chassis.Maneuverability
		unit.Speed = 0

		unit.Chassis = nil
	}
}

func (unit *Unit) DelWeapon() {
	if unit.Weapon != nil {
		unit.TypeAttack = ""
		unit.Damage = 0
		unit.MinAttackRange = 0
		unit.RangeAttack = 0
		unit.AreaAttack = 0

		unit.Accuracy = unit.Accuracy - unit.Weapon.Accuracy

		unit.Weapon = nil
	}
}

func (unit *Unit) DelTower() {
	if unit.Tower != nil {
		unit.HP = unit.HP - unit.Tower.HP
		unit.RangeView = unit.RangeView - unit.Tower.PowerRadar

		unit.VulEM = unit.VulEM - unit.Tower.VulToEM
		unit.VulExplosive = unit.VulExplosive - unit.Tower.VulToExplosion
		unit.VulKinetics = unit.VulKinetics - unit.Tower.VulToKinetics
		unit.VulThermal = unit.VulThermal - unit.Tower.VulToThermo

		unit.Armor = unit.Armor - unit.Tower.Armor

		unit.Tower = nil
	}
}

func (unit *Unit) DelBody() {
	if unit.Body != nil {
		unit.HP = unit.HP - unit.Body.HP

		unit.VulEM = unit.VulEM - unit.Body.VulToEM
		unit.VulExplosive = unit.VulExplosive - unit.Body.VulToExplosion
		unit.VulKinetics = unit.VulKinetics - unit.Body.VulToKinetics
		unit.VulThermal = unit.VulThermal - unit.Body.VulToThermo

		unit.Armor = unit.Armor - unit.Body.Armor

		unit.Body = nil
	}
}

func (unit *Unit) DelRadar() {
	if unit.Radar != nil {
		unit.RangeView = unit.RangeView - unit.Radar.Power
		unit.WallHack = false

		unit.Initiative = unit.Initiative - unit.Radar.Analysis
		unit.Accuracy = unit.Accuracy - unit.Radar.Analysis

		unit.Radar = nil
	}
}

func (unit *Unit) SetChassis(chassis *detailUnit.Chassis) {
	unit.Chassis = chassis
}

func (unit *Unit) SetWeapon(weapon *detailUnit.Weapon) {
	unit.Weapon = weapon
}

func (unit *Unit) SetTower(tower *detailUnit.Tower) {
	unit.Tower = tower
}

func (unit *Unit) SetBody(body *detailUnit.Body) {
	unit.Body = body
}

func (unit *Unit) SetRadar(radar *detailUnit.Radar) {
	unit.Radar = radar
}

func WeightUnit(unit *Unit) (weight int) {

	if unit.Weapon != nil {
		weight = weight + unit.Weapon.Weight
	}

	if unit.Tower != nil {
		weight = weight + unit.Tower.Weight
	}

	if unit.Body != nil {
		weight = weight + unit.Body.Weight
	}

	if unit.Radar != nil {
		weight = weight + unit.Radar.Weight
	}

	return
}
