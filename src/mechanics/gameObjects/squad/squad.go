package squad

import (
	"../../gameObjects/matherShip"
)

type Squad struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	Active     bool                   `json:"active"`
	MatherShip *matherShip.MatherShip `json:"mather_ship"`
	Inventory  map[int]*InventorySlot `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame     bool                   `json:"in_game"`
}

type InventorySlot struct {
	Item     interface{} `json:"item"`
	Quantity int         `json:"quantity"`
	Type     string      `json:"type"`
	ItemID   int         `json:"item_id"`
}
