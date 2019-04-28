package lobby

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/dialog"
	inv "github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/skill"
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

	RecycleSlots        map[int]*lobby.RecycleItem `json:"recycle_slots"`
	PreviewRecycleSlots []*inv.Slot                `json:"preview_recycle_slots"`

	Storage *inv.Inventory `json:"storage"`

	BluePrint *blueprints.Blueprint        `json:"blue_print"`
	BlueWorks map[int]*blueprints.BlueWork `json:"blue_works"`
	BlueWork  *blueprints.BlueWork         `json:"blue_work"`
	BPItem    interface{}                  `json:"bp_item"`
	Count     int                          `json:"count"`
	MaxCount  int                          `json:"max_count"`

	StartTime     int64 `json:"start_time"`
	ToTime        int64 `json:"to_time"`
	MineralSaving int   `json:"mineral_saving"`
	TimeSaving    int   `json:"time_saving"`
	BluePrintID   int   `json:"blue_print_id"`

	Player *player.Player `json:"player"`

	DialogPage   *dialog.Page `json:"dialog_page"`
	DialogAction string       `json:"dialog_action"`
	ToPage       int          `json:"to_page"`
	AskID        int          `json:"ask_id"`

	Fraction string `json:"fraction"`

	File      string      `json:"file"`
	Biography string      `json:"biography"`
	Skill     skill.Skill `json:"skill"`

	Maps       map[int]*_map.ShortInfoMap `json:"maps"`
	SearchMaps []*maps.SearchMap          `json:"search_maps"`
}
