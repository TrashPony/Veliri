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

	for _, oldPage := range d.dialogs[newDialogType.ID].Pages {
		for _, newPage := range newDialogType.Pages {

			// перетаскиваем старые картинки в новый диалог
			if oldPage.ID == newPage.ID {
				newPage.SetPictures(
					oldPage.GetPicture("main"),
					oldPage.GetPicture("Replics"),
					oldPage.GetPicture("Explores"),
					oldPage.GetPicture("Reverses"),
				)
			}

			// это новые страницы их тоже надо перенести в новый диалог
			if oldPage.ID == 0 {
				newDialogType.Pages[oldPage.Number] = oldPage
			}
		}
	}

	// при обновление ид самого диалога остается без изменений
	dialogEditor.UpdateDialog(newDialogType)
	d.dialogs[newDialogType.ID] = *newDialogType
}

func (d *dialogStore) SetPicture(dialogID, pageID int, PicType, picture string) {
	for _, page := range d.dialogs[dialogID].Pages {
		if page.ID == pageID {
			page.SetPicture(picture, PicType)
		}
	}
}

func (d *dialogStore) AddPage(dialogID int) {
	maxNumber := 0
	for _, page := range d.dialogs[dialogID].Pages {
		if page.Number > maxNumber {
			maxNumber = page.Number
		}
	}
	maxNumber++
	d.dialogs[dialogID].Pages[maxNumber] = &dialog.Page{Number: maxNumber}
}

func (d *dialogStore) AddNewDialog(name string) {
	newDialog := dialog.Dialog{
		Name:  name,
		Pages: make(map[int]*dialog.Page),
	}

	newDialog.Pages[1] = &dialog.Page{Number: 1}

	dialogEditor.AddDialog(&newDialog)
	d.dialogs[newDialog.ID] = newDialog
}

func (d *dialogStore) DeleteDialog(id int) {

	deleteDialog, ok := d.dialogs[id]
	if ok {
		dialogEditor.DeleteDialog(&deleteDialog)
		delete(d.dialogs, id)
	}
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
