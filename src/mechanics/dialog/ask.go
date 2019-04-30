package dialog

import (
	"errors"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/mission"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
)

func Ask(client *player.Player, gameDialog *dialog.Dialog, place string, toPage, askID int) (*dialog.Page, error, string, *mission.Mission) {
	if gameDialog == nil {
		return nil, errors.New("no dialog"), "", nil
	}

	var ask *dialog.Ask
	for _, page := range gameDialog.Pages {
		for _, gameAsk := range page.Asc {
			if gameAsk.ID == askID {
				ask = &gameAsk
				break
			}
		}
	}

	if ask == nil {
		return nil, errors.New("no ask"), "", nil
	}

	// переменная будет говорить соденение например обнови инвентарь или обнови сторедж и тд
	actionInfo := ""

	if (gameDialog.AccessType == "world" || gameDialog.AccessType == "base") && place == "base" {

		if page, ok := gameDialog.Pages[toPage]; ok || toPage == 0 {

			if ask.TypeAction != "" {
				var err error
				var newDialog *dialog.Dialog
				var newMission *mission.Mission

				actionInfo, err, newDialog, newMission = actionDialog(client, ask)

				if err != nil {
					// TODO влияние ошибки на диалог, например обмен предметов но нет нужного количества
				}

				if newDialog != nil {
					client.SetOpenDialog(newDialog)
					return newDialog.Pages[1], nil, actionInfo, newMission
				}
			}

			if toPage == 0 {
				client.SetOpenDialog(nil)
				return nil, nil, actionInfo, nil
			}

			return page, nil, actionInfo, nil

		} else {
			return nil, errors.New("no page"), actionInfo, nil
		}
	}
	return nil, errors.New("unknown error"), "", nil
}
