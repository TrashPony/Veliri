package detail

import "github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"

type Weapon struct {
	// если тут появяться ссылочные типы данных включая срезы, карты и тд, надо будет делать глубокое копирование в
	// factory/gameTypes/weapons
	ID                  int                      `json:"id"`
	Name                string                   `json:"name"`
	MinAttackRange      int                      `json:"min_attack_range"`
	Range               int                      `json:"range"`
	Accuracy            int                      `json:"accuracy"`
	AmmoCapacity        int                      `json:"ammo_capacity"`
	Artillery           bool                     `json:"artillery"`
	Power               int                      `json:"power"`
	MaxHP               int                      `json:"max_hp"`
	Type                string                   `json:"type"`          /* firearms, missile_weapon, laser_weapon */
	StandardSize        int                      `json:"standard_size"` /* small - 1, medium - 2, big - 3 */
	Size                float32                  `json:"size"`          /* занимаемый обьем в кубо метрах */
	EquipDamage         int                      `json:"equip_damage"`
	EquipCriticalDamage int                      `json:"equip_critical_damage"`
	XAttach             int                      `json:"x_attach"`
	YAttach             int                      `json:"y_attach"`
	RotateSpeed         int                      `json:"rotate_speed"`
	CountFireBullet     int                      `json:"count_fire_bullet"`
	BulletSpeed         int                      `json:"bullet_speed"`
	ReloadTime          int                      `json:"reload_time"`
	DelayFollowingFire  int                      `json:"delay_following_fire"`
	FirePositions       []*coordinate.Coordinate `json:"fire_positions"`
}
