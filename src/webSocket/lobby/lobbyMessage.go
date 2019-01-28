package lobby

type Message struct {
	Event     string `json:"event"`
	UserID    int    `json:"user_id"`

	StorageSlot  int   `json:"storage_slot"`
	StorageSlots []int `json:"storage_slots"`
}
