package DetailUnit

import "log"

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

func GetTower(id int) (tower *Tower) {

	rows, err := db.Query("select * from tower_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tower = &Tower{}

	for rows.Next() {
		err := rows.Scan(&tower.Id, &tower.Name, &tower.Type, &tower.Weight, &tower.HP, &tower.PowerRadar, &tower.Armor, &tower.VulToKinetics, &tower.VulToThermo, &tower.VulToEM, &tower.VulToExplosion)
		if err != nil {
			log.Fatal("get tower" + err.Error())
		}
	}

	return tower
}
