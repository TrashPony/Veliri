package ammo

type Ammo struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	TypeAttack   string  `json:"type_attack"`
	MinDamage    int     `json:"min_damage"`
	MaxDamage    int     `json:"max_damage"`
	AreaCovers   int     `json:"area_covers"`
	StandardSize int     `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	Size         float32 `json:"size"`
}
