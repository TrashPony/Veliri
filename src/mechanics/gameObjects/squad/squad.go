package squad

import (
	"../../gameObjects/unit"
)

type Squad struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	Active     bool                   `json:"active"`
	MatherShip *unit.Unit             `json:"mather_ship"`
	Inventory  map[int]*InventorySlot `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame     bool                   `json:"in_game"`
}

type InventorySlot struct {
	Item       interface{} `json:"item"`
	Quantity   int         `json:"quantity"`
	Type       string      `json:"type"`
	ItemID     int         `json:"item_id"`
	InsertToDB bool        `json:"insert_to_db"`
	HP         int         `json:"hp"`
}
