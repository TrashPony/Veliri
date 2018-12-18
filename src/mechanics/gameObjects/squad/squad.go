package squad

import (
	"../../gameObjects/inventory"
	"../../gameObjects/unit"
	"math"
)

type Squad struct {
	ID         int                 `json:"id"`
	Name       string              `json:"name"`
	Active     bool                `json:"active"`
	MatherShip *unit.Unit          `json:"mather_ship"`
	Inventory  inventory.Inventory `json:"inventory"` // в роли ключей карты выступают номера слотов где содержиться итем
	InGame     bool                `json:"in_game"`
	Q          int                 `json:"q"`
	R          int                 `json:"r"`
	MapID      int                 `json:"map_id"`
}

func (squad *Squad) GetUseAllInventorySize() float32 {

	var inventorySquadSize float32
	for _, slot := range squad.Inventory.Slots {
		if slot.Item != nil {
			inventorySquadSize = inventorySquadSize + slot.Size
		}
	}

	return float32(round(float64(inventorySquadSize), 1))
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
