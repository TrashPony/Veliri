package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/dialogEditor"
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

func (d *dialogStore) GetAll() map[int]dialog.Dialog {
	return d.dialogs
}

func (d *dialogStore) GetByID(id int) *dialog.Dialog {

	var newDialog dialog.Dialog
	deepcopy.Copy(&newDialog, d.dialogs[id])

	return &newDialog
}

func (d *dialogStore) GetTypeGreeting(fraction, typeDialog string) *dialog.Dialog {
	for _, gameDialog := range d.dialogs {
		if gameDialog.Fraction == fraction && gameDialog.AccessType == "base" && gameDialog.Type == typeDialog {
			return d.GetByID(gameDialog.ID)
		}
	}
	return nil
}

func (d *dialogStore) UpdateTypeDialog(newDialogType *dialog.Dialog) {
	// при обновление ид самого диалога остается без изменений
	dialogEditor.UpdateDialog(newDialogType)
	d.dialogs[newDialogType.ID] = *newDialogType
}

func (d *dialogStore) AddNewDialog(newDialogType *dialog.Dialog) {
	dialogEditor.AddDialog(newDialogType)
	d.dialogs[newDialogType.ID] = *newDialogType
}

func (d *dialogStore) GetDialogPageByID(pageID int) *dialog.Page {
	for _, gameDialog := range d.dialogs {
		for _, page := range gameDialog.Pages {
			if page.ID == pageID {
				return page
			}
		}
	}

	return nil
}
