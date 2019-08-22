package missions

import (
	"database/sql"
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"log"
)

func AddMission(newMission *mission.Mission) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	err = tx.QueryRow("INSERT INTO missions (name, start_dialog_id, reward_cr, fraction, start_base_id, type) "+
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		newMission.Name, newMission.StartDialogID, newMission.RewardCr, newMission.Fraction, newMission.StartBaseID, newMission.Type).Scan(&newMission.ID)
	if err != nil {
		log.Fatal("add new mission :" + err.Error())
	}

	AddActions(newMission, tx)
	AddRewardItems(newMission, tx)

	err = tx.Commit()
	if err != nil {
		log.Fatal("add new mission : " + err.Error())
	}
}

func UpdateMission(updateMission, oldMission *mission.Mission) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE missions SET name = $2, start_dialog_id = $3, reward_cr = $4, fraction = $5, start_base_id = $6, type = $7 WHERE id = $1",
		updateMission.ID, updateMission.Name, updateMission.StartDialogID, updateMission.RewardCr, updateMission.Fraction, updateMission.StartBaseID, updateMission.Type)
	if err != nil {
		log.Fatal("update mission main info" + err.Error())
	}

	DeleteOldInfo(oldMission, tx)
	AddActions(updateMission, tx)
	AddRewardItems(updateMission, tx)
	for _, action := range updateMission.Actions {
		AddNeedItems(action, tx)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("update mission: " + err.Error())
	}
}

func DeleteMission(deleteMission *mission.Mission) {
	tx, err := dbConnect.GetDBConnect().Begin()
	defer tx.Rollback()

	DeleteOldInfo(deleteMission, tx)
	_, err = tx.Exec("DELETE FROM missions WHERE id=$1",
		deleteMission.ID)
	if err != nil {
		log.Fatal("delete mission" + err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("delete mission: " + err.Error())
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
	for i, action := range updateMission.Actions {

		err := tx.QueryRow("INSERT INTO actions (id_mission, type_monitor, description, short_description, "+
			"base_id, Q, R, radius, sec, count, dialog_id, number, async, alternative_dialog_id) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id",
			updateMission.ID, action.TypeFuncMonitor, action.Description, action.ShortDescription, action.BaseID, action.Q,
			action.R, action.Radius, action.Sec, action.Count, action.DialogID, action.Number, action.Async,
			action.AlternativeDialogId).Scan(&updateMission.Actions[i].ID)
		if err != nil {
			log.Fatal("add new action in mission " + err.Error())
		}
	}
}

func AddNeedItems(action *mission.Action, tx *sql.Tx) {
	// todo повторяющийся код
	if action.NeedItems == nil {
		return
	}

	for i, itemSlot := range action.NeedItems.Slots {
		_, err := tx.Exec("INSERT INTO need_action_items (id_actions, slot, item_type, item_id, quantity, hp) "+
			"VALUES ($1, $2, $3, $4, $5, $6)",
			action.ID, i, itemSlot.Type, itemSlot.ItemID, itemSlot.Quantity, itemSlot.HP)
		if err != nil {
			println("update need items" + err.Error())
		}
	}
}

func AddRewardItems(updateMission *mission.Mission, tx *sql.Tx) {

	if updateMission.RewardItems == nil {
		return
	}

	for i, itemSlot := range updateMission.RewardItems.Slots {
		_, err := tx.Exec("INSERT INTO reward_items (id_mission, slot, item_type, item_id, quantity, hp) "+
			"VALUES ($1, $2, $3, $4, $5, $6)",
			updateMission.ID, i, itemSlot.Type, itemSlot.ItemID, itemSlot.Quantity, itemSlot.HP)

		if err != nil {
			println("update reward items" + err.Error())
		}
	}
}
