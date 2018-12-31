package detail

import (
	"../ammo"
	"../coordinate"
	"../equip"
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
	VulToExplosion  int    `json:"vul_to_explosion"`
	RangeView       int    `json:"range_view"`
	Accuracy        int    `json:"accuracy"`
	MaxPower        int    `json:"max_power"`
	RecoveryPower   int    `json:"recovery_power"`
	RecoveryHP      int    `json:"recovery_hp"`
	WallHack        bool   `json:"wall_hack"`

	CapacitySize       float32 `json:"capacity_size"`        /* вместимость корпуса к кубо-метрах */
	StandardSize       int     `json:"standard_size"`        /* small - 1, medium - 2, big - 3, размер корпуса (если корпус мс то неучитывается)*/
	StandardSizeSmall  bool    `json:"standard_size_small"`  /* оружие которое может использовать корпус small, medium, big */
	StandardSizeMedium bool    `json:"standard_size_medium"` /* оружие которое может использовать корпус small, medium, big */
	StandardSizeBig    bool    `json:"standard_size_big"`    /* оружие которое может использовать корпус small, medium, big */

	EquippingI   map[int]*BodyEquipSlot `json:"equippingI"`
	EquippingII  map[int]*BodyEquipSlot `json:"equippingII"`
	EquippingIII map[int]*BodyEquipSlot `json:"equippingIII"`
	EquippingIV  map[int]*BodyEquipSlot `json:"equippingIV"`
	EquippingV   map[int]*BodyEquipSlot `json:"equippingV"`

	ThoriumSlots map[int]*ThoriumSlot `json:"thorium_slots"` /* слоты в которых хранится топливо */

	Weapons map[int]*BodyWeaponSlot `json:"weapons"`
}

func (body *Body) GetEquip(typeSlot, numberSlot int) *BodyEquipSlot {
	if typeSlot == 1 {
		return body.EquippingI[numberSlot]
	}
	if typeSlot == 2 {
		return body.EquippingII[numberSlot]
	}
	if typeSlot == 3 {
		return body.EquippingIII[numberSlot]
	}
	if typeSlot == 4 {
		return body.EquippingIV[numberSlot]
	}
	if typeSlot == 5 {
		return body.EquippingV[numberSlot]
	}

	return nil
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

func (body *Body) GetUseCapacitySize() float32 {
	var allSize float32

	var counter = func(equip map[int]*BodyEquipSlot) float32 {
		var size float32
		for _, slot := range equip {
			if slot.Equip != nil {
				size = size + slot.Equip.Size
			}
		}
		return size
	}

	allSize = allSize + counter(body.EquippingI)
	allSize = allSize + counter(body.EquippingII)
	allSize = allSize + counter(body.EquippingIII)
	allSize = allSize + counter(body.EquippingIV)
	allSize = allSize + counter(body.EquippingV)

	for _, slot := range body.Weapons {
		if slot.Weapon != nil {
			allSize = allSize + slot.Weapon.Size

			if slot.Ammo != nil {
				allSize = allSize + slot.Ammo.Size*float32(slot.AmmoQuantity)
			}
		}
	}

	return allSize
}

type BodyEquipSlot struct {
	Type           int                    `json:"type_slot"`
	Number         int                    `json:"number_slot"`
	Equip          *equip.Equip           `json:"equip"`
	InsertToDB     bool                   `json:"insert_to_db"`
	Used           bool                   `json:"used"` /* использовано или нет */
	StepsForReload int                    `json:"steps_for_reload"`
	HP             int                    `json:"hp"`
	Target         *coordinate.Coordinate `json:"target"`
	StandardSize   int                    `json:"standard_size"` /* определяет тип вмещаемого юнита если это ангар */
}

type ThoriumSlot struct {
	Number    int     `json:"number_slot"`
	Count     int     `json:"count"`
	MaxCount  int     `json:"max_count"`
	WorkedOut float32 `json:"worked_out"` /* параметр показывает что топливо вырабатано на сколь-ко то процентов */
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
