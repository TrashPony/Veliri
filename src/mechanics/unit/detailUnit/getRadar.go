package detailUnit

import (
	"log"
	"../../../dbConnect"
)

type Radar struct {
	Id       int    `json:"id"`
	Name	 string	`json:"name"`
	Type     string `json:"type"`
	Weight   int    `json:"weight"`
	Power    int    `json:"power"`
	Through  bool   `json:"through"`
	Analysis int    `json:"analysis"`
}

func GetRadars() (radars []Radar) {
	radars = make([]Radar, 0)

	rows, err := dbConnect.GetDBConnect().Query("select * from radar_type")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var radar Radar

	for rows.Next() {
		err := rows.Scan(&radar.Id, &radar.Name, &radar.Type, &radar.Weight, &radar.Power, &radar.Through, &radar.Analysis)
		if err != nil {
			log.Fatal("get radars" + err.Error())
		}
		radars = append(radars, radar)
	}

	return radars
}

func GetRadar(id int) (radar *Radar) {

	rows, err := dbConnect.GetDBConnect().Query("select * from radar_type where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	radar = &Radar{}

	for rows.Next() {
		err := rows.Scan(&radar.Id, &radar.Name, &radar.Type, &radar.Weight, &radar.Power, &radar.Through, &radar.Analysis)
		if err != nil {
			log.Fatal("get radar" + err.Error())
		}
	}

	return radar
}
