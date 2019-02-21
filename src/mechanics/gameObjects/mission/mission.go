package mission

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"

type Mission struct {
	StartDialog   dialog.Dialog `json:"start_dialog"`
	Actions       []*Action     `json:"actions"`
	CurrentAction *Action       `json:"current_action"`
	RewardCr      int           `json:"reward_cr"`
	//todo награда []итемс
}

type Action struct {
	TypeFuncMonitor  string `json:"type_func_monitor"`
	TypeFuncComplete string `json:"type_func_complete"`
	Description      string `json:"description"`
	BaseID           int    `json:"base_id"`
	Q                int    `json:"q"`
	R                int    `json:"r"`
	Count            int    `json:"count"`
	CurrentCount     int    `json:"current_count"`
	PlayerID         int    `json:"player_id"`
	BotID            int    `json:"bot_id"`
	DialogID         int    `json:"dialog_id"`
}
