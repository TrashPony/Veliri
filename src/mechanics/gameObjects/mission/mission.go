package mission

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
)

type Mission struct {
	ID                  int                  `json:"id"`
	UUID                string               `json:"uuid"`
	StartDialogID       int                  `json:"start_dialog_id"`
	NotFinishedDialogId int                  `json:"not_finished_dialog_id"`
	Name                string               `json:"name"`
	Actions             []*Action            `json:"actions"`
	RewardCr            int                  `json:"reward_cr"`
	RewardItems         *inventory.Inventory `json:"reward_items"`
	Fraction            string               `json:"fraction"`
	StartBaseID         int                  `json:"start_base_id"`
	Type                string               `json:"type"`
	MainStory           bool                 `json:"main_story"`
	Story               int                  `json:"story"`

	// методанные необходмые для правильной работы квеста
	StartDialog *dialog.Dialog     `json:"start_dialog"`
	StartBase   *base.Base         `json:"start_base"`
	StartMap    *_map.ShortInfoMap `json:"start_map"`
}

func (m *Mission) CheckAvailableActionByIndex(number int) bool {
	// функция проверяет дост к действию т.к. они выполняются последовательно надо что бы были выполнены все последующие
	for _, action := range m.Actions {
		// любое не выполненое действие снизу сразу не дает доступ к действию number,
		// однако исключение являются асинхронные действия (action.Async), т.к. они не имеют порядка выполнения
		if action.Number < number && !action.Complete && !action.Async {
			return false
		}
	}

	return true
}

type Action struct {
	ID              int    `json:"id"`
	TypeFuncMonitor string `json:"type_func_monitor"`
	Complete        bool   `json:"complete"`

	//мета информация
	Description         string               `json:"description"`
	ShortDescription    string               `json:"short_description"`
	BaseID              int                  `json:"base_id"`
	X                   int                  `json:"x"`
	Y                   int                  `json:"y"`
	MapID               int                  `json:"map_id"`
	Radius              int                  `json:"radius"`
	Sec                 int                  `json:"sec"`
	Count               int                  `json:"count"`
	CurrentCount        int                  `json:"current_count"`
	PlayerID            int                  `json:"player_id"`
	BotID               int                  `json:"bot_id"`
	DialogID            int                  `json:"dialog_id"`
	AlternativeDialogId int                  `json:"alternative_dialog_id"`
	NeedItems           *inventory.Inventory `json:"need_items"`
	Number              int                  `json:"number"`
	Async               bool                 `json:"async"`
	Dialog              *dialog.Dialog       `json:"dialog"`
	OwnerPlace          bool                 `json:"owner_place"`
	EndText             string               `json:"end_text"`
}
