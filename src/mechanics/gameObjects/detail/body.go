package detail

import (
	"../equip"
	"../ammo"
)

type Body struct {
	ID              int                     `json:"id"`
	Name            string                  `json:"name"`
	MotherShip      bool                    `json:"mother_ship"`
	Speed           int                     `json:"speed"`
	Initiative      int                     `json:"initiative"`
	MaxHP           int                     `json:"max_hp"`
	Armor           int                     `json:"armor"`
	EvasionCritical int                     `json:"evasion_critical"`
	VulToKinetics   int                     `json:"vul_to_kinetics"`
	VulToThermo     int                     `json:"vul_to_thermo"`
	VulToEM         int                     `json:"vul_to_em"`
	VulToExplosion  int                     `json:"vul_to_explosion"`
	RangeView       int                     `json:"range_view"`
	Accuracy        int                     `json:"accuracy"`
	MaxPower        int                     `json:"max_power"`
	RecoveryPower   int                     `json:"recovery_power"`
	WallHack        bool                    `json:"wall_hack"`
	Equipping       map[int]*BodyEquipSlot  `json:"equipping"`
	Weapons         map[int]*BodyWeaponSlot `json:"weapons"`
}

type BodyEquipSlot struct {
	Type       int          `json:"type_slot"`
	Number     int          `json:"number_slot"`
	Equip      *equip.Equip `json:"equip"`
	InsertToDB bool         `json:"insert_to_db"`
}

type BodyWeaponSlot struct {
	Type       int        `json:"type_slot"`
	Number     int        `json:"number_slot"`
	Weapon     *Weapon    `json:"weapon"`
	WeaponType string     `json:"weapon_type"`
	Ammo       *ammo.Ammo `json:"ammo"`
	InsertToDB bool       `json:"insert_to_db"`
}
