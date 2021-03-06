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
		" fraction," +
		" start_base_id," +
		" type," +
		" not_finished_dialog_id," +
		" main_story," +
		" story " +
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
			&gameMission.Fraction, &gameMission.StartBaseID, &gameMission.Type, &gameMission.NotFinishedDialogId,
			&gameMission.MainStory, &gameMission.Story)
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

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp, place_user_id "+
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
		" description,"+
		" short_description,"+
		" base_id,"+
		" x,"+
		" y,"+
		" count,"+
		" dialog_id,"+
		" number,"+
		" async,"+
		" radius,"+
		" sec,"+
		" alternative_dialog_id,"+
		" map_id,"+
		" owner_place,"+
		" end_text "+
		" "+
		"FROM actions "+
		"WHERE id_mission = $1", missionGame.ID)
	if err != nil {
		log.Fatal("actions in missions" + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var actions mission.Action

		err := rows.Scan(&actions.ID, &actions.TypeFuncMonitor, &actions.Description,
			&actions.ShortDescription, &actions.BaseID, &actions.X, &actions.Y, &actions.Count,
			&actions.DialogID, &actions.Number, &actions.Async, &actions.Radius, &actions.Sec,
			&actions.AlternativeDialogId, &actions.MapID, &actions.OwnerPlace, &actions.EndText)
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

	rows, err := dbConnect.GetDBConnect().Query("SELECT slot, item_type, item_id, quantity, hp, place_user_id "+
		"FROM need_action_items "+
		"WHERE id_actions = $1", action.ID)
	if err != nil {
		log.Fatal("need items in action" + err.Error())
	}
	defer rows.Close()

	action.NeedItems.FillInventory(rows)
}
