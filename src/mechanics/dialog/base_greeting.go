package dialog

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func GetBaseGreeting(client *player.Player, base *base.Base) (*dialog.Page, *dialog.Dialog) {

	// TODO проверка на нападение на сектор если да то "На сектор напали! Почему ты все еще на базе!? Иди обороняй!"

	// сначала ищем диалоги которые положены по мисии
	for _, mission := range client.Missions {
		for _, action := range mission.Actions {
			if action.BaseID == base.ID && action.Dialog != nil {
				client.SetOpenDialog(action.Dialog)
				return action.Dialog.Pages[1], action.Dialog
			}
		}
	}

	// потом смотрим взял ли он уже миссию на этой базе
	for _, mission := range client.Missions {
		if mission.StartBase.ID == base.ID {
			greeting := gameTypes.Dialogs.GetTypeGreeting(client.Fraction, "greeting_before_mission_not_complete")
			greeting.ProcessingDialogText(client.Login, base.Name, "", "")
			client.SetOpenDialog(greeting)
			return greeting.Pages[1], greeting
		}
	}

	// нормальное привествие
	greeting := gameTypes.Dialogs.GetTypeGreeting(client.Fraction, "greeting")
	greeting.ProcessingDialogText(client.Login, base.Name, "", "")
	client.SetOpenDialog(greeting)
	return greeting.Pages[1], greeting
}
