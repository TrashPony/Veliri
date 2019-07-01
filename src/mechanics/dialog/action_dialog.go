package dialog

import (
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/missions"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	mission2 "github.com/TrashPony/Veliri/src/mechanics/mission"
)

func actionDialog(client *player.Player, ask *dialog.Ask) (string, error, *dialog.Page, *mission.Mission) {

	if ask.TypeAction == "get_base_mission" {
		newMission := missions.Missions.GenerateMissionForUser(client)

		// базы фракций могут выдавать только свои квесты
		userBase, _ := bases.Bases.Get(client.InBaseID)
		if newMission.Fraction == userBase.Fraction {
			client.SetOpenDialog(newMission.StartDialog)
			return "new_dialog", nil, newMission.StartDialog.Pages[1], newMission
		} else {
			return "", nil, client.GetOpenDialog().Pages[ask.ToPage], nil
		}
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
		missions.Missions.AcceptMission(client, client.GetOpenDialog().Mission)
	}

	if ask.TypeAction == "get_reward" {
		// завершение мисии
		_, page := mission2.Complete(client, client.GetOpenDialog().Mission)
		return "get_reward", nil, page, nil
	}

	if ask.TypeAction == "get_base_greeting" {
		userBase, _ := bases.Bases.Get(client.InBaseID)
		_, greeting := GetBaseGreeting(client, userBase)
		client.SetOpenDialog(greeting)
		return "get_base_greeting", nil, greeting.Pages[1], nil
	}

	dbPlayer.UpdateUser(client)
	return "", nil, nil, nil
}
