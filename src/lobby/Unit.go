package lobby

type Unit struct {
	Chassis *Chassis
	Weapon  *Weapon
	Tower   *Tower
	Body    *Body
	Radar   *Radar

	// Движение
	Speed      int
	Initiative int

	// Атака
	Damage         int
	RangeAttack    int
	MinAttackRange int
	AreaAttack     int
	TypeAttack     string

	// Выживаемость
	HP              int
	Armor           int
	EvasionCritical int
	VulKinetics     int // Уязвимости
	VulThermal      int
	VulEM           int
	VulExplosive    int

	// Навигация
	RangeView       int
	Accuracy		int
	WallHack		bool
}

func (unit *Unit) CalculateParametersUnit() {
	weight := WeightUnit(unit)

	// расчет шасси
	if unit.Chassis != nil {
		unit.EvasionCritical = unit.Chassis.Maneuverability * 2
		unit.Initiative = unit.Initiative + unit.Chassis.Maneuverability

		percentWeight := unit.Chassis.Carrying * (weight/100)

		if percentWeight > 75 {
			percentWeight = percentWeight - 75
			fine := (percentWeight/5) + 1
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

func WeightUnit(unit *Unit) (weight int)  {

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
