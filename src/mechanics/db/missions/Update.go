package missions

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"log"
)

func UpdateMission(updateMission *mission.Mission) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE missions SET name = $2, start_dialog_id = $3, reward_cr = $4, fraction = $5, start_base_id = $6, type = $7 WHERE id = $1",
		updateMission.ID, updateMission.Name, updateMission.StartDialogID, updateMission.RewardCr, updateMission.Fraction, updateMission.StartBaseID, updateMission.Type)
	if err != nil {
		log.Fatal("update mission main info" + err.Error())
	}

	DeleteOldInfo(updateMission, tx)

	err = tx.Commit()
	if err != nil {
		log.Fatal("update mission: " + err.Error())
	}
}

func DeleteOldInfo(updateMission *mission.Mission, tx *sql.Tx) {

	_, err := tx.Exec("DELETE FROM reward_items WHERE id=$1",
		updateMission.ID)
	if err != nil {
		log.Fatal("delete reward_items in mission" + err.Error())
	}

	for _, action := range updateMission.Actions {
		_, err := tx.Exec("DELETE FROM need_action_items WHERE id_actions=$1",
			action.ID)
		if err != nil {
			log.Fatal("delete old need_items action" + err.Error())
		}
	}

	_, err = tx.Exec("DELETE FROM actions WHERE dialog_id=$1",
		updateMission.ID)
	if err != nil {
		log.Fatal("delete old actions" + err.Error())
	}
}

func AddActions(updateMission *mission.Mission, tx *sql.Tx) {

}
