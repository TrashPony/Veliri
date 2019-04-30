package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/getlantern/deepcopy"
)

type dialogStore struct {
	dialogs map[int]dialog.Dialog
}

var Dialogs = newDialogStore()

func newDialogStore() *dialogStore {
	return &dialogStore{dialogs: get.Dialogs()}
}

func (d *dialogStore) GetByID(id int) *dialog.Dialog {

	var newDialog dialog.Dialog
	deepcopy.Copy(&newDialog, d.dialogs[id])

	return &newDialog
}

func (d *dialogStore) GetBaseGreeting(fraction string) *dialog.Dialog {
	for _, gameDialog := range d.dialogs {
		if gameDialog.Fraction == fraction && gameDialog.AccessType == "base" && gameDialog.Name == "Приветствие" {
			return d.GetByID(gameDialog.ID)
		}
	}
	return nil
}
