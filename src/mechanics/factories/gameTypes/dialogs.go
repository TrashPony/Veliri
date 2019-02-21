package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
)

type dialogStore struct {
	dialogs map[int]dialog.Dialog
}

var Dialogs = newDialogStore()

func newDialogStore() *dialogStore {
	return &dialogStore{dialogs: get.Dialogs()}
}

func (d *dialogStore) GetByID(id int) dialog.Dialog {
	return d.dialogs[id]
}
