package dialog

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/missions"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func actionDialog(client *player.Player, ask *dialog.Ask) (string, error, *dialog.Dialog, *mission.Mission) {

	if ask.TypeAction == "get_base_mission" {
		newMission := missions.Missions.GenerateMissionForUser()
		newMission.StartDialog.ProcessingDialogText(client.GetLogin(), "", newMission.ToBase.Name)
		return "new_dialog", nil, newMission.StartDialog, newMission
	}

	if ask.TypeAction == "start_training" {
		client.Training = 1
		return "start_training", nil, nil, nil
	}

	if ask.TypeAction == "miss_training" {
		client.Training = 999
		return "miss_training", nil, nil, nil
	}

	if ask.TypeAction == "close" {
		client.SetOpenDialog(nil)
		return "close", nil, nil, nil
	}

	if ask.TypeAction == "accept_mission" {

	}

	if ask.TypeAction == "get_reward" {

	}

	if ask.TypeAction == "get_base_greeting" {
		userBase, _ := bases.Bases.Get(client.InBaseID)
		_, greeting := GetBaseGreeting(client, userBase)
		return "get_base_greeting", nil, greeting, nil
	}

	dbPlayer.UpdateUser(client)
	return "", nil, nil, nil
}
