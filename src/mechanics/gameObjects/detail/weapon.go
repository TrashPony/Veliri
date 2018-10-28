package detail

type Weapon struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	MinAttackRange int     `json:"min_attack_range"`
	Range          int     `json:"range"`
	Accuracy       int     `json:"accuracy"`
	AmmoCapacity   int     `json:"ammo_capacity"`
	Artillery      bool    `json:"artillery"`
	Power          int     `json:"power"`
	MaxHP          int     `json:"max_hp"`
	Type           string  `json:"type"`          /* firearms, missile_weapon, laser_weapon */
	StandardSize   int     `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	Size           float32 `json:"size"`          /* занимаемый обьем в кубо метрах */
}
