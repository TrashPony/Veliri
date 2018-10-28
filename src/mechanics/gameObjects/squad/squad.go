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
	Size       float32     `json:"size"`
}

func (squad *Squad) GetUseAllInventorySize() float32 {

	var unitSquadSize float32
	var msBodySize float32

	if squad.MatherShip.Body != nil {
		msBodySize = squad.MatherShip.Body.GetUseCapacitySize()

		for _, squadUnit := range squad.MatherShip.Units {
			if squadUnit.Unit != nil {
				unitSquadSize = unitSquadSize + squadUnit.Unit.Body.GetUseCapacitySize()
				unitSquadSize = unitSquadSize + squadUnit.Unit.Body.CapacitySize
			}
		}
	}

	var inventorySquadSize float32
	for _, slot := range squad.Inventory {
		if slot.Item != nil {
			inventorySquadSize = inventorySquadSize + slot.Size
		}
	}

	return msBodySize + unitSquadSize + inventorySquadSize
}
