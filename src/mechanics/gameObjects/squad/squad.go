package squad

import (
	"../../gameObjects/unit"
	"math"
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
	Size       float32     `json:"size"`
}

func (squad *Squad) GetUseAllInventorySize() float32 {

	var inventorySquadSize float32
	for _, slot := range squad.Inventory {
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
