package dialog

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func GetBaseGreeting(client *player.Player, base *base.Base) (*dialog.Page, *dialog.Dialog) {
	greeting := gameTypes.Dialogs.GetBaseGreeting(client.Fraction)

	// TODO проверка на то есть ли задания у игрока для этой базы, если есть то надо выводить ченить "хорош прохлаждатся! иди работай"
	// TODO проверка на нападение на сектор если да то "На сектор напали! Почему ты все еще на базе!? Иди обороняй!"

	greeting.ProcessingDialogText(client.Login, base.Name, "")

	client.SetOpenDialog(greeting)
	return greeting.Pages[1], greeting
}
