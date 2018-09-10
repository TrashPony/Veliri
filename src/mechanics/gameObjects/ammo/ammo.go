package ammo

type Ammo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	TypeAttack string `json:"type_attack"`
	MinDamage  int    `json:"min_damage"`
	MaxDamage  int    `json:"max_damage"`
	AreaCovers int    `json:"area_covers"`
}
