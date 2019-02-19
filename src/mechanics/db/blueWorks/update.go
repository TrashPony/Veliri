package blueWorks

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"log"
)

func UpdateBW(work *blueprints.BlueWork) {
	_, err := dbConnect.DB.Exec(""+
		"UPDATE created_blueprint "+
		"SET finish_time = $2"+
		"WHERE id=$1",
		work.ID, work.FinishTime.UTC())
	if err != nil {
		log.Fatal("update storage item" + err.Error())
	}
}
