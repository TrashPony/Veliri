package dialog

import (
	"errors"
	dbPlayer "github.com/TrashPony/Veliri/src/mechanics/db/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func Ask(client *player.Player, gameDialog *dialog.Dialog, place string, toPage, askID int) (*dialog.Page, error, string) {
	if gameDialog == nil {
		return nil, errors.New("no dialog"), ""
	}

	var ask *dialog.Ask

	for _, page := range gameDialog.Pages {
		for _, gameAsk := range page.Asc {
			if gameAsk.ID == askID {
				ask = &gameAsk
			}
		}
	}

	if ask == nil {
		return nil, errors.New("no ask"), ""
	}

	// переменная будет говорить соденение например обнови инвентарь или обнови сторедж и тд
	actionInfo := ""

	if (gameDialog.AccessType == "world" || gameDialog.AccessType == "base") && place == "base" {
		if page, ok := gameDialog.Pages[toPage]; ok || toPage == 0 {

			if ask.GetAction() != "" {
				var err error

				actionInfo, err = actionDialog(client, ask)
				if err != nil {
					// TODO влияние ошибки на диалог, например обмен предметов но нет нужного количества
				}
			}

			if toPage == 0 {
				client.SetOpenDialog(nil)
				return nil, nil, actionInfo
			}

			return &page, nil, actionInfo

		} else {
			return nil, errors.New("no page"), actionInfo
		}
	}

	return nil, errors.New("unknown error"), ""
}

func actionDialog(client *player.Player, ask *dialog.Ask) (string, error) {
	if ask.GetAction() == "start_training" {
		client.Training = 1
		return "updateTraining", nil
	}

	if ask.GetAction() == "miss_training" {
		client.Training = 999
		return "updateTraining", nil
	}

	dbPlayer.UpdateUser(client)
	return "", nil
}
