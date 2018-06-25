package get

import (
	"../../../dbConnect"
	"../../gameObjects/detail"
	"log"
)

func Bodies() (bodies []detail.Body) {
	bodies = make([]detail.Body, 0)

	rows, err := dbConnect.GetDBConnect().Query("select * from body_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var body detail.Body

	for rows.Next() {
		err := rows.Scan(&body.Id, &body.Name, &body.Type, &body.Weight, &body.HP, &body.MaxTowerWeight, &body.Armor, &body.VulToKinetics, &body.VulToThermo, &body.VulToEM, &body.VulToExplosion)
		if err != nil {
			log.Fatal("get bodies" + err.Error())
		}
		bodies = append(bodies, body)
	}

	return bodies
}

func Body(id int) (body *detail.Body) {

	rows, err := dbConnect.GetDBConnect().Query("select * from body_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	body = &detail.Body{}

	for rows.Next() {
		err := rows.Scan(&body.Id, &body.Name, &body.Type, &body.Weight, &body.HP, &body.MaxTowerWeight, &body.Armor, &body.VulToKinetics, &body.VulToThermo, &body.VulToEM, &body.VulToExplosion)
		if err != nil {
			log.Fatal("get body" + err.Error())
		}
	}

	return body
}
