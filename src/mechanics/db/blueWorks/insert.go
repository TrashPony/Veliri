package blueWorks

import (
	"../../../dbConnect"
	"../../gameObjects/blueprints"
	"log"
)

func InsertDW(work *blueprints.BlueWork) {
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
		"created_blueprint "+
		"(id_blueprint, id_base, id_user, finish_time, mineral_saving_percentage, time_saving_percentage) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6) "+
		"RETURNING id",
		work.BlueprintID, work.BaseID, work.UserID, work.FinishTime.UTC(), work.MineralSavingPercentage, work.TimeSavingPercentage).Scan(&work.ID)
	if err != nil {
		log.Fatal("add new blue work " + err.Error())
	}
}
