package detailUnit

import (
	"log"
	"../dbConnect"
)

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

	rows, err := dbConnect.GetDBConnect().Query("select * from body_type")
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

func GetBody(id int) (body *Body) {

	rows, err := dbConnect.GetDBConnect().Query("select * from body_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	body = &Body{}

	for rows.Next() {
		err := rows.Scan(&body.Id, &body.Name, &body.Type, &body.Weight, &body.HP, &body.MaxTowerWeight, &body.Armor, &body.VulToKinetics, &body.VulToThermo, &body.VulToEM, &body.VulToExplosion)
		if err != nil {
			log.Fatal("get body" + err.Error())
		}
	}

	return body
}
