package inventory

import (
	"../../mechanics/gameObjects/squad"
)

type Message struct {
	Event string `json:"event"`

	BodyID   int `json:"id_body"`
	WeaponID int `json:"weapon_id"`
	EquipID  int `json:"equip_id"`
	AmmoID   int `json:"ammo_id"`

	InventorySlot int `json:"inventory_slot"`
	EquipSlot     int `json:"equip_slot"`
	EquipSlotType int `json:"equip_slot_type"`

	UnitSlot int `json:"unit_slot"`
}

type Response struct {
	Event string       `json:"event"`
	Squad *squad.Squad `json:"squad"`
	Error string       `json:"error"`
}
