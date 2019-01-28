package lobby

import (
	"../../mechanics/gameObjects/inventory"
)

type Message struct {
	Event  string `json:"event"`
	UserID int    `json:"user_id"`
	Error  string `json:"error"`

	StorageSlot  int   `json:"storage_slot"`
	RecyclerSlot int   `json:"recycler_slot"`
	StorageSlots []int `json:"storage_slots"`

	RecycleSlots        map[int]*inventory.Slot `json:"recycle_slots"`
	PreviewRecycleSlots []*inventory.Slot       `json:"preview_recycle_slots"`
}
