package detailUnit

import (
	"log"
	"../../../dbConnect"
)

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

	rows, err := dbConnect.GetDBConnect().Query("select * from chassis_type")
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

func GetChass(id int) (chassis *Chassis) {

	rows, err := dbConnect.GetDBConnect().Query("select * from chassis_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	chassis = &Chassis{}

	for rows.Next() {
		err := rows.Scan(&chassis.Id, &chassis.Name, &chassis.Type, &chassis.Carrying, &chassis.Maneuverability, &chassis.Speed)
		if err != nil {
			log.Fatal("get chass" + err.Error())
		}
	}

	return chassis
}