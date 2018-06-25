package ammo

type Ammo struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	TypeAttack string `json:"type_attack"`
	Damage     int    `json:"damage"`
	AreaCovers int    `json:"area_covers"`
}
