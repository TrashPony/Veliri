package blueWorks

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"log"
)

func InsertDW(work *blueprints.BlueWork) {
	err := dbConnect.GetDBConnect().QueryRow("INSERT INTO "+
		"created_blueprint "+
		"(id_blueprint, id_base, id_user, start_time, finish_time, mineral_tax_percentage, time_tax_percentage) "+
		"VALUES "+
		"($1, $2, $3, $4, $5, $6, $7) "+
		"RETURNING id",
		work.BlueprintID, work.BaseID, work.UserID, work.StartTime.UTC(), work.FinishTime.UTC(), work.MineralTaxPercentage, work.TimeTaxPercentage).Scan(&work.ID)
	if err != nil {
		log.Fatal("add new blue work " + err.Error())
	}
}
