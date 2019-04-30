package missions

import (
	"github.com/TrashPony/Veliri/src/dbConnect"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"log"
)

func Missions() map[int]*mission.Mission {
	rows, err := dbConnect.GetDBConnect().Query("" +
		"SELECT " +
		" id," +
		" name," +
		" start_dialog_id," +
		" reward_cr," +
		" end_dialog_id," +
		" end_base_id," +
		" fraction," +
		" start_base_id," +
		" delivery_item_id " +
		" " +
		"FROM missions")
	if err != nil {
		log.Fatal("get all missions " + err.Error())
	}
	defer rows.Close()

	allMissions := make(map[int]*mission.Mission)

	for rows.Next() {

		var gameMission mission.Mission

		err := rows.Scan(&gameMission.ID, &gameMission.Name, &gameMission.StartDialogID, &gameMission.RewardCr,
			&gameMission.EndDialogID, &gameMission.EndBaseID, &gameMission.Fraction, &gameMission.StartBaseID,
			&gameMission.DeliveryItemId)
		if err != nil {
			log.Fatal("scan all missions " + err.Error())
		}

		// всятие итемов в нагруду, экшинов и необходимых придметов для экшинов
		rewardItems(&gameMission)
		getMissionActions(&gameMission)

		allMissions[gameMission.ID] = &gameMission
	}

	return allMissions
}

func rewardItems(missionGame *mission.Mission) {
	missionGame.RewardItems = &inventory.Inventory{}
	missionGame.RewardItems.Slots = make(map[int]*inventory.Slot)
	missionGame.RewardItems.SetSlotsSize(999)

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
		"FROM reward_items "+
		"WHERE id_mission = $1", missionGame.ID)
	if err != nil {
		log.Fatal("rewardItems in missions" + err.Error())
	}
	defer rows.Close()

	missionGame.RewardItems.FillInventory(rows)
}

func getMissionActions(missionGame *mission.Mission) {
	missionGame.Actions = make([]*mission.Action, 0)

	rows, err := dbConnect.GetDBConnect().Query(""+
		"SELECT"+
		" id,"+
		" type_monitor,"+
		" complete,"+
		" description,"+
		" short_description,"+
		" base_id,"+
		" q,"+
		" r,"+
		" count,"+
		" current_count,"+
		" player_id,"+
		" bot_id,"+
		" dialog_id,"+
		" number,"+
		" async "+
		" "+
		"FROM actions "+
		"WHERE id_mission = $1", missionGame.ID)
	if err != nil {
		log.Fatal("actions in missions" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var actions mission.Action

		err := rows.Scan(&actions.ID, &actions.TypeFuncMonitor, &actions.Complete, &actions.Description,
			&actions.ShortDescription, &actions.BaseID, &actions.Q, &actions.R, &actions.Count, &actions.CurrentCount,
			&actions.PlayerID, &actions.BotID, &actions.DialogID, &actions.Number, &actions.Async)
		if err != nil {
			log.Fatal("scan actions in missions " + err.Error())
		}

		needActionItems(&actions)

		missionGame.Actions = append(missionGame.Actions, &actions)
	}
}

func needActionItems(action *mission.Action) {
	action.NeedItems = &inventory.Inventory{}
	action.NeedItems.Slots = make(map[int]*inventory.Slot)
	action.NeedItems.SetSlotsSize(999)

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp "+
		"FROM reward_items "+
		"WHERE id_actions = $1", action.ID)
	if err != nil {
		log.Fatal("need items in action" + err.Error())
	}
	defer rows.Close()

	action.NeedItems.FillInventory(rows)
}
