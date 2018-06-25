package detail

type Weapon struct {
	Id             int    `json:"id"`
	Name		   string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	Damage         int    `json:"damage"`
	MinAttackRange int    `json:"min_attack_range"`
	Range		   int    `json:"range"`
	Accuracy       int    `json:"accuracy"`
	AreaCovers     int    `json:"area_covers"`
}
