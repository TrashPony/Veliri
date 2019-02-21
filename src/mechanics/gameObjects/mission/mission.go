package mission

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"

type Mission struct {
	StartDialog   *dialog.Dialog `json:"start_dialog"`
	Actions       []*Action      `json:"actions"`
	CurrentAction *Action        `json:"current_action"`
	
}

type Action struct {
}
