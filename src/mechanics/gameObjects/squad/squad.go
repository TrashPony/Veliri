package squad

import (
	"../../gameObjects/matherShip"
)

type Squad struct {
	ID         int                    `json:"id"`
	Name       string                 `json:"name"`
	MatherShip *matherShip.MatherShip `json:"mather_ship"`
	Inventory  map[int]interface{}    `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame     bool                   `json:"in_game"`
}
