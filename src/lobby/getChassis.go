package lobby

import "log"

type Chassis struct {
	Id            int    `json:"id"`
	Type          string `json:"type"`
	HP            int    `json:"hp"`
	MoveSpeed     int    `json:"move_speed"`
	Initiative    int    `json:"initiative"`
	Size          int    `json:"size"`
	MaxWeaponSize int    `json:"max_weapon_size"`
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
		err := rows.Scan(&chassis.Id, &chassis.Type, &chassis.HP, &chassis.MoveSpeed, &chassis.Initiative, &chassis.Size, &chassis.MaxWeaponSize)
		if err != nil {
			log.Fatal(err)
		}
		chassiss = append(chassiss, chassis)
	}

	return chassiss
}
