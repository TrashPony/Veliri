package blueWorks

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"log"
)

func BlueWorks() map[int]*blueprints.BlueWork {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" id_blueprint," +
		" id_base," +
		" id_user," +
		" start_time," +
		" finish_time, " +
		" mineral_tax_percentage," +
		" time_tax_percentage" +
		" " +
		"FROM created_blueprint")
	if err != nil {
		log.Fatal("get all blueWork " + err.Error())
	}
	defer rows.Close()

	blueWorks := make(map[int]*blueprints.BlueWork) // [user_ID, base_ID, []works]

	for rows.Next() {
		var blueWork blueprints.BlueWork
		err := rows.Scan(
			&blueWork.ID,
			&blueWork.BlueprintID,
			&blueWork.BaseID,
			&blueWork.UserID,
			&blueWork.StartTime,
			&blueWork.FinishTime,
			&blueWork.MineralTaxPercentage,
			&blueWork.TimeTaxPercentage,
		)

		if err != nil {
			log.Fatal("get scan all blueWork " + err.Error())
		}

		blueWorks[blueWork.ID] = &blueWork
	}

	return blueWorks
}
