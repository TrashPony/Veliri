package mission

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
)

type Mission struct {
	ID             int                  `json:"id"`
	StartDialogID  int                  `json:"start_dialog_id"`
	Name           string               `json:"name"`
	Actions        []*Action            `json:"actions"`
	RewardCr       int                  `json:"reward_cr"`
	RewardItems    *inventory.Inventory `json:"reward_items"`
	EndDialogID    int                  `json:"end_dialog_id"`
	EndBaseID      int                  `json:"end_base_id"`
	Fraction       string               `json:"fraction"`
	StartBaseID    int                  `json:"start_base_id"`
	DeliveryItemId int                  `json:"delivery_item_id"`

	StartDialog *dialog.Dialog `json:"start_dialog"`
	EndDialog   *dialog.Dialog `json:"end_dialog_id"`
	ToBase      *base.Base     `json:"to_base"`
}

type Action struct {
	ID               int                  `json:"id"`
	TypeFuncMonitor  string               `json:"type_func_monitor"`
	Complete         bool                 `json:"complete"`
	Description      string               `json:"description"`
	ShortDescription string               `json:"short_description"`
	BaseID           int                  `json:"base_id"`
	Q                int                  `json:"q"`
	R                int                  `json:"r"`
	Count            int                  `json:"count"`
	CurrentCount     int                  `json:"current_count"`
	PlayerID         int                  `json:"player_id"`
	BotID            int                  `json:"bot_id"`
	DialogID         int                  `json:"dialog_id"`
	NeedItems        *inventory.Inventory `json:"need_items"`
	Number           int                  `json:"number"`
	Async            bool                 `json:"async"`
}
