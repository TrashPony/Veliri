package blueWorks

import (
	"../../../dbConnect"
	"../../gameObjects/blueprints"
	"log"
)

func InsertDW(work *blueprints.BlueWork) {
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
		"created_blueprint "+
		"(id_blueprint, id_base, id_user, finish_time) "+
		"VALUES "+
		"($1, $2, $3, $4) "+
		"RETURNING id",
		work.BlueprintID, work.BaseID, work.UserID, work.FinishTime).Scan(&work.ID)
	if err != nil {
		log.Fatal("add new blue work " + err.Error())
	}
}
