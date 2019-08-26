package boxInMap

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/inventory"
	"time"
)

type Box struct {
	ID           int     `json:"id"`
	TypeID       int     `json:"type_id"`
	MapID        int     `json:"map_id"`
	Type         string  `json:"type"`
	CapacitySize float32 `json:"capacity_size"`
	FoldSize     float32 `json:"fold_size"`
	Protect      bool    `json:"protect"`
	ProtectLvl   int     `json:"protect_lvl"`
	password     int
	DestroyTime  time.Time            `json:"destroy_time"`
	Underground  bool                 `json:"underground"`
	Q            int                  `json:"q"`
	R            int                  `json:"r"`
	Rotate       int                  `json:"rotate"`
	storage      *inventory.Inventory // содержимое не публично т.к. что бы узнать содержимое надо его открыть или просканирова
	HP           int                  `json:"hp"`
}

func (b *Box) SetPassword(password int) {
	b.password = password
}

func (b *Box) GetPassword() int {
	return b.password
}

func (b *Box) CreateStorage() {

}

func (b *Box) GetStorage() *inventory.Inventory {
	if b.storage == nil {
		b.storage = &inventory.Inventory{}
	}
	return b.storage
}
