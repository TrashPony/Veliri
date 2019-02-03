package blueWorks

import (
	"../../../dbConnect"
	"../../gameObjects/blueprints"
	"log"
)

func BlueWorks() map[int]map[int]map[int]*blueprints.BlueWork {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" id_blueprint," +
		" id_base," +
		" id_user," +
		" finish_time " +
		" " +
		"FROM created_blueprint")
	if err != nil {
		log.Fatal("get all blueWork " + err.Error())
	}
	defer rows.Close()

	blueWorks := make(map[int]map[int]map[int]*blueprints.BlueWork) // [user_ID, base_ID, []works]

	for rows.Next() {
		var blueWork blueprints.BlueWork
		err := rows.Scan(&blueWork.ID, &blueWork.BlueprintID, &blueWork.BaseID, &blueWork.UserID, &blueWork.FinishTime)
		if err != nil {
			log.Fatal("get scan all blueWork " + err.Error())
		}

		// D8
		if blueWorks[blueWork.UserID] != nil {
			if blueWorks[blueWork.UserID][blueWork.BaseID] != nil {
				blueWorks[blueWork.UserID][blueWork.BaseID][blueWork.ID] = &blueWork
			} else {
				blueWorks[blueWork.UserID][blueWork.BaseID] = make(map[int]*blueprints.BlueWork, 0)
				blueWorks[blueWork.UserID][blueWork.BaseID][blueWork.ID] = &blueWork
			}
		} else {
			blueWorks[blueWork.UserID] = make(map[int]map[int]*blueprints.BlueWork)
			blueWorks[blueWork.UserID][blueWork.BaseID] = make(map[int]*blueprints.BlueWork, 0)
			blueWorks[blueWork.UserID][blueWork.BaseID][blueWork.ID] = &blueWork
		}
	}

	return blueWorks
}
