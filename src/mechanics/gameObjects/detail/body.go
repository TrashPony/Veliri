package detail

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/ammo"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"math/rand"
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

	VulToKinetics  int `json:"vul_to_kinetics"`
	VulToThermo    int `json:"vul_to_thermo"`
	VulToExplosion int `json:"vul_to_explosion"`

	RangeView  int `json:"range_view"`
	RangeRadar int `json:"range_radar"`

	Accuracy      int  `json:"accuracy"`
	MaxPower      int  `json:"max_power"`
	RecoveryPower int  `json:"recovery_power"`
	RecoveryHP    int  `json:"recovery_hp"`
	WallHack      bool `json:"wall_hack"`

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

	Height int `json:"height"`
	Width  int `json:"width"`
}

func (body *Body) FindApplicableEquip(applicable string) *BodyEquipSlot {

	var findEquip = func(equip map[int]*BodyEquipSlot) *BodyEquipSlot {
		for _, slot := range equip {
			if slot.Equip != nil && slot.Equip.Applicable == applicable {
				return slot
			}
		}
		return nil
	}

	bodyEquip := findEquip(body.EquippingI)
	if bodyEquip != nil {
		return bodyEquip
	}

	bodyEquip = findEquip(body.EquippingII)
	if bodyEquip != nil {
		return bodyEquip
	}

	bodyEquip = findEquip(body.EquippingIII)
	if bodyEquip != nil {
		return bodyEquip
	}

	bodyEquip = findEquip(body.EquippingIV)
	if bodyEquip != nil {
		return bodyEquip
	}

	bodyEquip = findEquip(body.EquippingV)
	if bodyEquip != nil {
		return bodyEquip
	}

	return nil
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

func (body *Body) GetRandomEquip() *BodyEquipSlot {

	typeSlot := rand.Intn(5) + 1

	randEquip := func(equips map[int]*BodyEquipSlot) *BodyEquipSlot {
		count := 0

		if equips == nil || len(equips) == 0 {
			return body.GetRandomEquip()
		}

		count2 := rand.Intn(len(equips))
		for _, slot := range equips {
			if count == count2 {
				return slot
			}
			count++
		}
		return body.GetRandomEquip()
	}

	if typeSlot == 1 {
		return randEquip(body.EquippingI)
	}
	if typeSlot == 2 {
		return randEquip(body.EquippingII)
	}
	if typeSlot == 3 {
		return randEquip(body.EquippingIII)
	}
	if typeSlot == 4 {
		return randEquip(body.EquippingIV)
	}
	if typeSlot == 5 {
		return randEquip(body.EquippingV)
	}

	return body.GetRandomEquip()
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
	XAttach        int                    `json:"x_attach"`
	YAttach        int                    `json:"y_attach"`
	Mining         bool                   `json:"mining"`
}

type ThoriumSlot struct {
	Number    int     `json:"number_slot"`
	Count     int     `json:"count"`
	MaxCount  int     `json:"max_count"`
	WorkedOut float32 `json:"worked_out"` /* параметр показывает что топливо вырабатано на сколь-ко то процентов */
	Inversion bool    `json:"inversion"`
}

type BodyWeaponSlot struct {
	Type         int        `json:"type_slot"` // по диз доку он может быть только 3
	Number       int        `json:"number_slot"`
	Weapon       *Weapon    `json:"weapon"`
	WeaponType   string     `json:"weapon_type"`
	Ammo         *ammo.Ammo `json:"ammo"`
	AmmoQuantity int        `json:"ammo_quantity"`
	InsertToDB   bool       `json:"insert_to_db"`
	HP           int        `json:"hp"`
	XAttach      int        `json:"x_attach"`
	YAttach      int        `json:"y_attach"`
	Reload       bool       `json:"reload"`
}
