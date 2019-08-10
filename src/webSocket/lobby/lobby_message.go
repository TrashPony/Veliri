package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/base"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/lobby"
)

type Message struct {
	Event  string `json:"event"`
	UserID int    `json:"user_id"`
	Error  string `json:"error"`

	ID int `json:"id"`

	StorageSlot  int   `json:"storage_slot"`
	RecyclerSlot int   `json:"recycler_slot"`
	StorageSlots []int `json:"storage_slots"`

	RecycleSlots        map[string]map[int]*lobby.RecycleItem `json:"recycle_slots"`
	PreviewRecycleSlots []*inv.Slot                           `json:"preview_recycle_slots"`
	UserRecycleSkill    int                                   `json:"user_recycle_skill"`
	Storage             *inv.Inventory                        `json:"storage"`

	BluePrint *blueprints.Blueprint        `json:"blue_print"`
	BlueWorks map[int]*blueprints.BlueWork `json:"blue_works"`
	BlueWork  *blueprints.BlueWork         `json:"blue_work"`
	BPItem    interface{}                  `json:"bp_item"`
	Count     int                          `json:"count"`
	MaxCount  int                          `json:"max_count"`

	StartTime     int64  `json:"start_time"`
	ToTime        int64  `json:"to_time"`
	MineralSaving int    `json:"mineral_saving"`
	TimeSaving    int    `json:"time_saving"`
	BluePrintID   int    `json:"blue_print_id"`
	ItemSource    string `json:"item_source"`

	Fraction string `json:"fraction"`

	DialogPage *dialog.Page `json:"dialog_page"`

	UserWorkSkillDetailPercent int `json:"user_work_skill_detail_percent"`
	UserWorkSkillTimePercent   int `json:"user_work_skill_time_percent"`

	Base       *base.Base `json:"base"`
	Efficiency int        `json:"efficiency"`
}
