package ammo

type Ammo struct {
	// если тут появяться ссылочные типы данных включая срезы, карты и тд, надо будет делать глубокое копирование в
	// factory/gameTypes/ammo
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Type                string  `json:"type"`
	TypeAttack          string  `json:"type_attack"`
	MinDamage           int     `json:"min_damage"`
	MaxDamage           int     `json:"max_damage"`
	AreaCovers          int     `json:"area_covers"`
	StandardSize        int     `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	EquipDamage         int     `json:"equip_damage"`
	EquipCriticalDamage int     `json:"equip_critical_damage"`
	Size                float32 `json:"size"`
	ChaseTarget         bool    `json:"chase_target"`
	BulletSpeed         int     `json:"bullet_speed"`
}
