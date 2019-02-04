package lobby

import (
	"../../mechanics/gameObjects/blueprints"
	inv "../../mechanics/gameObjects/inventory"
	"../../mechanics/lobby"
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
	BPItem    interface{}                  `json:"bp_item"`
	Count     int                          `json:"count"`
	MaxCount  int                          `json:"max_count"`
}
