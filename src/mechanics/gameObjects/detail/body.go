package detail

import (
	"../equip"
	"../ammo"
)

type Body struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	MotherShip      bool   `json:"mother_ship"`
	Speed           int    `json:"speed"`
	Initiative      int    `json:"initiative"`
	MaxHP           int    `json:"max_hp"`
	Armor           int    `json:"armor"`
	EvasionCritical int    `json:"evasion_critical"`
	VulToKinetics   int    `json:"vul_to_kinetics"`
	VulToThermo     int    `json:"vul_to_thermo"`
	VulToEM         int    `json:"vul_to_em"`
	VulToExplosion  int    `json:"vul_to_explosion"`
	RangeView       int    `json:"range_view"`
	Accuracy        int    `json:"accuracy"`
	MaxPower        int    `json:"max_power"`
	RecoveryPower   int    `json:"recovery_power"`
	WallHack        bool   `json:"wall_hack"`

	EquippingI   map[int]*BodyEquipSlot `json:"equippingI"`
	EquippingII  map[int]*BodyEquipSlot `json:"equippingII"`
	EquippingIII map[int]*BodyEquipSlot `json:"equippingIII"`
	EquippingIV  map[int]*BodyEquipSlot `json:"equippingIV"`
	EquippingV   map[int]*BodyEquipSlot `json:"equippingV"`

	Weapons map[int]*BodyWeaponSlot `json:"weapons"`
}

func (body *Body) GetUsePower() int {
	var allPower int

	var counter = func(equip map[int]*BodyEquipSlot) int {
		var power int
		for _, slot := range equip {
			if slot.Equip != nil {
				power = power + slot.Equip.Power
			}
		}
		return power
	}

	allPower = allPower + counter(body.EquippingI)
	allPower = allPower + counter(body.EquippingII)
	allPower = allPower + counter(body.EquippingIII)
	allPower = allPower + counter(body.EquippingIV)
	allPower = allPower + counter(body.EquippingV)

	for _, slot := range body.Weapons {
		if slot.Weapon != nil {
			allPower = allPower + slot.Weapon.Power
		}
	}

	return allPower
}

type BodyEquipSlot struct {
	Type           int          `json:"type_slot"`
	Number         int          `json:"number_slot"`
	Equip          *equip.Equip `json:"equip"`
	InsertToDB     bool         `json:"insert_to_db"`
	Used           bool         `json:"used"` /* использовано или нет */
	StepsForReload int          `json:"steps_for_reload"`
	HP             int          `json:"hp"`
}

type BodyWeaponSlot struct {
	Type         int        `json:"type_slot"`
	Number       int        `json:"number_slot"`
	Weapon       *Weapon    `json:"weapon"`
	WeaponType   string     `json:"weapon_type"`
	Ammo         *ammo.Ammo `json:"ammo"`
	AmmoQuantity int        `json:"ammo_quantity"`
	InsertToDB   bool       `json:"insert_to_db"`
	HP           int        `json:"hp"`
}
