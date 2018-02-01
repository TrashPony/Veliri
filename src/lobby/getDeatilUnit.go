package lobby

import (
	"log"
)

type UnitPrototype struct {
	chassis Chassis
	weapon Weapon
	tower Tower
	body Body
	radar Radar
}

type Chassis struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Carrying        int    `json:"carrying"`
	Maneuverability int    `json:"maneuverability"`
	Speed           int    `json:"max_speed"`
}

func GetChassis() (chassiss []Chassis) {
	chassiss = make([]Chassis, 0)

	rows, err := db.Query("select * from chassis_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var chassis Chassis

	for rows.Next() {
		err := rows.Scan(&chassis.Id, &chassis.Name, &chassis.Type, &chassis.Carrying, &chassis.Maneuverability, &chassis.Speed)
		if err != nil {
			log.Fatal("get chassiss" + err.Error())
		}
		chassiss = append(chassiss, chassis)
	}

	return chassiss
}

type Weapon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	Damage         int    `json:"damage"`
	MinAttackRange int    `json:"min_attack_range"`
	Accuracy       int    `json:"accuracy"`
	AreaCovers     int    `json:"area_covers"`
}

func GetWeapons() (weapons []Weapon) {
	weapons = make([]Weapon, 0)

	rows, err := db.Query("select * from weapon_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weapon Weapon

	for rows.Next() {
		err := rows.Scan(&weapon.Id, &weapon.Name, &weapon.Type, &weapon.Weight, &weapon.Damage, &weapon.MinAttackRange, &weapon.Accuracy, &weapon.AreaCovers)
		if err != nil {
			log.Fatal("get weapon" + err.Error())
		}
		weapons = append(weapons, weapon)
	}

	return weapons
}

type Tower struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	HP             int    `json:"hp"`
	PowerRadar     int    `json:"power_radar"`
	Armor          int    `json:"armor"`
	VulToKinetics  int    `json:"vul_to_kinetics"`
	VulToThermo    int    `json:"vul_to_thermo"`
	VulToEM        int    `json:"vul_to_em"`
	VulToExplosion int    `json:"vul_to_explosion"`
}

func GetTowers() (towers []Tower) {
	towers = make([]Tower, 0)

	rows, err := db.Query("select * from tower_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tower Tower

	for rows.Next() {
		err := rows.Scan(&tower.Id, &tower.Name, &tower.Type, &tower.Weight, &tower.HP, &tower.PowerRadar, &tower.Armor, &tower.VulToKinetics, &tower.VulToThermo, &tower.VulToEM, &tower.VulToExplosion)
		if err != nil {
			log.Fatal("get towers" + err.Error())
		}
		towers = append(towers, tower)
	}

	return towers
}

type Body struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	HP             int    `json:"hp"`
	MaxTowerWeight int    `json:"max_tower_weight"`
	Armor          int    `json:"armor"`
	VulToKinetics  int    `json:"vul_to_kinetics"`
	VulToThermo    int    `json:"vul_to_thermo"`
	VulToEM        int    `json:"vul_to_em"`
	VulToExplosion int    `json:"vul_to_explosion"`
}

func GetBodies() (bodies []Body) {
	bodies = make([]Body, 0)

	rows, err := db.Query("select * from body_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var body Body

	for rows.Next() {
		err := rows.Scan(&body.Id, &body.Name, &body.Type, &body.Weight, &body.HP, &body.MaxTowerWeight, &body.Armor, &body.VulToKinetics, &body.VulToThermo, &body.VulToEM, &body.VulToExplosion)
		if err != nil {
			log.Fatal("get bodies" + err.Error())
		}
		bodies = append(bodies, body)
	}

	return bodies
}

type Radar struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Weight   int    `json:"weight"`
	Power    int    `json:"power"`
	Through  bool   `json:"through"`
	Analysis int    `json:"analysis"`
}

func GetRadars() (radars []Radar) {
	radars = make([]Radar, 0)

	rows, err := db.Query("select * from radar_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var radar Radar

	for rows.Next() {
		err := rows.Scan(&radar.Id, &radar.Name, &radar.Type, &radar.Weight, &radar.Power, &radar.Through, &radar.Analysis)
		if err != nil {
			log.Fatal("get radar" + err.Error())
		}
		radars = append(radars, radar)
	}

	return radars
}
