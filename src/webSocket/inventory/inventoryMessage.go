package inventory

import (
	"../../mechanics/gameObjects/squad"
)

type Message struct {
	Event string `json:"event"`

	BodyID        int `json:"id_body"`
	InventorySlot int `json:"inventory_slot"`
	EquipSlot	  int `json:"equip_slot"`
}

type Response struct {
	Event string       `json:"event"`
	Squad *squad.Squad `json:"squad"`
}
