package dialog

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/missions"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func GetBaseGreeting(client *player.Player, base *base.Base) (*dialog.Page, *dialog.Dialog, *mission.Mission) {

	// TODO проверка на нападение на сектор если да то "На сектор напали! Почему ты все еще на базе!? Иди обороняй!"
	// todo если игрок на базе не своей фракции

	// сначала ищем диалоги которые положены по мисии, и доступны
	for _, userMission := range client.Missions {
		for _, action := range userMission.Actions {
			if action.BaseID == base.ID && action.Dialog != nil && userMission.CheckAvailableActionByIndex(action.Number) {
				client.SetOpenDialog(action.Dialog)
				return action.Dialog.Pages[1], action.Dialog, nil
			}
		}
	}

	// потом смотрим взял ли он уже миссию на этой базе, но еще не выполнил
	for _, userMission := range client.Missions {
		if userMission.StartBase.ID == base.ID {
			greeting := gameTypes.Dialogs.GetByID(userMission.NotFinishedDialogId)
			greeting.ProcessingDialogText(client.Login, base.Name, "", "", client.Fraction)
			client.SetOpenDialog(greeting)
			return greeting.Pages[1], greeting, nil
		}
	}

	// если игрок прошел обучение выдаем ему фракционный сюжет
	if client.Training >= 9 && base.ID == bases.Bases.GetCapital(client.Fraction).ID {

		storyMiss := missions.Missions.GetStory(client.StoryEpisode, client.Fraction)

		if storyMiss != nil {
			userMission := missions.Missions.GenerateMissionForUser(client, storyMiss)
			client.SetOpenDialog(userMission.StartDialog)
			return userMission.StartDialog.Pages[1], userMission.StartDialog, nil
		}
	}

	// нормальное привествие
	greeting := gameTypes.Dialogs.GetTypeGreeting(base.Fraction, "greeting")
	greeting.ProcessingDialogText(client.Login, base.Name, "", "", client.Fraction)
	client.SetOpenDialog(greeting)

	return greeting.Pages[1], greeting, nil
}
