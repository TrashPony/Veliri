package detail

type Body struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Weight         int    `json:"weight"`
	Speed          int    `json:"speed"`
	HP             int    `json:"hp"`
	MaxTowerWeight int    `json:"max_tower_weight"`
	Armor          int    `json:"armor"`
	VulToKinetics  int    `json:"vul_to_kinetics"`
	VulToThermo    int    `json:"vul_to_thermo"`
	VulToEM        int    `json:"vul_to_em"`
	VulToExplosion int    `json:"vul_to_explosion"`
	RangeView      int    `json:"range_view"`
}
