package dialog

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/bases"
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/factories/missions"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func GetBaseGreeting(client *player.Player, base *base.Base) (*dialog.Page, *dialog.Dialog, *mission.Mission) {

	// TODO проверка на нападение на сектор если да то "На сектор напали! Почему ты все еще на базе!? Иди обороняй!"

	// сначала ищем диалоги которые положены по мисии
	for _, userMission := range client.Missions {
		for _, action := range userMission.Actions {
			if action.BaseID == base.ID && action.Dialog != nil {
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
		userMission := missions.Missions.GetStory(client.StoryEpisode, client.Fraction)
		greeting := gameTypes.Dialogs.GetByID(userMission.StartDialogID)

		// ToSectorName берется из первого экшона где он указан, в сюжетных заданиях рандомно нечего не генерится
		toSectorName := ""
		for _, action := range userMission.Actions {
			if action.MapID != 0 {
				mp, _ := maps.Maps.GetByID(action.MapID)
				toSectorName = mp.Name
			}
		}

		greeting.ProcessingDialogText(client.Login, base.Name, "", toSectorName, client.Fraction)
		client.SetOpenDialog(greeting)
		return greeting.Pages[1], greeting, nil
	}

	// нормальное привествие
	greeting := gameTypes.Dialogs.GetTypeGreeting(base.Fraction, "greeting")
	greeting.ProcessingDialogText(client.Login, base.Name, "", "", client.Fraction)
	client.SetOpenDialog(greeting)

	// todo если игрок на базе не своей фракции

	// todo invalid memory address or nil pointer dereference т.к. нет диалога для другой фракции
	return greeting.Pages[1], greeting, nil
}
