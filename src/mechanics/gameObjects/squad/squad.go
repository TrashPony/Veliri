package squad

import (
	"../../gameObjects/inventory"
	"../../gameObjects/unit"
)

type Squad struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Active       bool                `json:"active"`
	MatherShip   *unit.Unit          `json:"mather_ship"`
	Inventory    inventory.Inventory `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame       bool                `json:"in_game"`
	Q            int                 `json:"q"`
	R            int                 `json:"r"`
	GlobalX      int                 `json:"global_x"` /* вычасляема координата на пиксельной сетке */
	GlobalY      int                 `json:"global_y"` /* вычасляема координата на пиксельной сетке */
	MapID        int                 `json:"map_id"`
	CurrentSpeed float64             `json:"current_speed"`
}