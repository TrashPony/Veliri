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

	InventorySlot  int   `json:"inventory_slot"`
	InventorySlots []int `json:"inventory_slots"`

	StorageSlot  int   `json:"storage_slot"`
	StorageSlots []int `json:"storage_slots"`

	EquipSlot     int `json:"equip_slot"`
	EquipSlotType int `json:"equip_slot_type"`

	UnitSlot    int `json:"unit_slot"`
	ThoriumSlot int `json:"thorium_slot"`

	Source      string `json:"source"`
	Destination string `json:"destination"`
}

type Response struct {
	Event         string       `json:"event"`
	Squad         *squad.Squad `json:"squad"`
	InventorySize float32      `json:"inventory_size"`
	InBase        bool         `json:"in_base"`
	Error         string       `json:"error"`
}
