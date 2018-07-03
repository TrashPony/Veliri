package detail

import "../ammo"

type Weapon struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	MinAttackRange int        `json:"min_attack_range"`
	Range          int        `json:"range"`
	Accuracy       int        `json:"accuracy"`
	AmmoCapacity   int        `json:"ammo_capacity"`
	Artillery      bool       `json:"artillery"`
	Power          int        `json:"power"`
	Ammo           *ammo.Ammo `json:"ammo"`
}
