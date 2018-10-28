package unit

import (
	"../coordinate"
	"../detail"
	"../effect"
)

type Unit struct {
	ID      int    `json:"id"`
	SquadID int    `json:"squad_id"`
	Owner   string `json:"owner"`

	Body *detail.Body `json:"body"`

	Q int `json:"q"`
	R int `json:"r"`

	Rotate int  `json:"rotate"`
	OnMap  bool `json:"on_map"`

	Target      *coordinate.Coordinate `json:"target"`
	QueueAttack int                    `json:"queue_attack"`
	Defend      bool                   `json:"defend"`

	HP           int `json:"hp"`
	Power        int `json:"power"`
	ActionPoints int `json:"action_points"`

	Speed           int  `json:"speed"`
	Initiative      int  `json:"initiative"`
	MaxHP           int  `json:"max_hp"`
	Armor           int  `json:"armor"`
	EvasionCritical int  `json:"evasion_critical"`
	VulToKinetics   int  `json:"vul_to_kinetics"`
	VulToThermo     int  `json:"vul_to_thermo"`
	VulToEM         int  `json:"vul_to_em"`
	VulToExplosion  int  `json:"vul_to_explosion"`
	RangeView       int  `json:"range_view"`
	Accuracy        int  `json:"accuracy"`
	MaxPower        int  `json:"max_power"`
	RecoveryPower   int  `json:"recovery_power"`
	RecoveryHP      int  `json:"recovery_HP"`
	WallHack        bool `json:"wall_hack"`

	Effects []*effect.Effect `json:"effects"`
	MS      bool             `json:"ms"`
	Units   map[int]*Slot    `json:"units"` // в роли ключей карты выступают
}

type Slot struct {
	Unit       *Unit `json:"unit"`
	NumberSlot int   `json:"number_slot"`
}

func (unit *Unit) GetID() int {
	return unit.ID
}

func (unit *Unit) GetBody() *detail.Body {
	return unit.Body
}

func (unit *Unit) GetTarget() *coordinate.Coordinate {
	return unit.Target
}

func (unit *Unit) DelBody() {
	if unit.Body != nil {
		unit.Body = nil
	}
}

func (unit *Unit) DelEquip() {

}

func (unit *Unit) DelAmmo() {

}

func (unit *Unit) SetBody(body *detail.Body) {
	unit.Body = body
}

func (unit *Unit) SetEquip() {

}

func (unit *Unit) SetAmmo() {

}

func (unit *Unit) SetQ(q int) {
	unit.Q = q
}

func (unit *Unit) GetQ() int {
	return unit.Q
}

func (unit *Unit) SetR(y int) {
	unit.R = y
}

func (unit *Unit) GetR() int {
	return unit.R
}

func (unit *Unit) GetY() int {
	return -unit.Q - unit.R
}

func (unit *Unit) GetWatchZone() int {
	return unit.RangeView
}

func (unit *Unit) GetOwnerUser() string {
	return unit.Owner
}

func (unit *Unit) GetOnMap() bool {
	return unit.OnMap
}

func (unit *Unit) SetOnMap(bool bool) {
	unit.OnMap = bool
}

func (unit *Unit) GetWallHack() bool {
	return unit.WallHack
}

func (unit *Unit) GetAmmoCount() int { // по диз доку оружие в юните может быть только одно
	for _, weaponSlot := range unit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			return weaponSlot.AmmoQuantity
		}
	}

	return 0
}

func (unit *Unit) GetWeaponSlot() *detail.BodyWeaponSlot { // по диз доку оружие в юните может быть только одно
	for _, weaponSlot := range unit.Body.Weapons {
		if weaponSlot.Weapon != nil {
			return weaponSlot
		}
	}

	return nil
}

func (unit *Unit) CalculateParams() {

	if unit.Body == nil {
		unit.Speed = 0
		unit.Initiative = 0
		unit.MaxHP = 0
		unit.Armor = 0
		unit.EvasionCritical = 0
		unit.VulToKinetics = 0
		unit.VulToThermo = 0
		unit.VulToEM = 0
		unit.VulToExplosion = 0
		unit.RangeView = 0
		unit.Accuracy = 0
		unit.MaxPower = 0
		unit.RecoveryPower = 0
		unit.RecoveryHP = 0
		unit.WallHack = false

		return
	}

	// начальные параметры тела
	unit.Speed = unit.Body.Speed
	unit.Initiative = unit.Body.Initiative
	unit.MaxHP = unit.Body.MaxHP
	unit.Armor = unit.Body.Armor
	unit.EvasionCritical = unit.Body.EvasionCritical
	unit.VulToKinetics = unit.Body.VulToKinetics
	unit.VulToThermo = unit.Body.VulToThermo
	unit.VulToExplosion = unit.Body.VulToExplosion
	unit.RangeView = unit.Body.RangeView
	unit.Accuracy = unit.Body.Accuracy
	unit.MaxPower = unit.Body.MaxPower
	unit.RecoveryPower = unit.Body.RecoveryPower
	unit.RecoveryHP = unit.Body.RecoveryHP
	unit.WallHack = unit.Body.WallHack

	// смотрим пасивное обородование
	var checkEffect = func(equipEffect *effect.Effect, parameter *int) {
		if equipEffect.Type == "enhances" {
			if equipEffect.Percentages {
				*parameter += *parameter / 100 * equipEffect.Quantity
			} else {
				*parameter += equipEffect.Quantity
			}
		}

		if equipEffect.Type == "reduced" {
			if equipEffect.Percentages {
				*parameter += *parameter / 100 * equipEffect.Quantity
			} else {
				*parameter += equipEffect.Quantity
			}
		}
	}

	var checkParams = func(equipEffect *effect.Effect) {
		if equipEffect.Parameter == "speed" {
			checkEffect(equipEffect, &unit.Speed)
		}

		if equipEffect.Parameter == "initiative" {
			checkEffect(equipEffect, &unit.Initiative)
		}

		if equipEffect.Parameter == "max_hp" {
			checkEffect(equipEffect, &unit.MaxHP)
		}

		if equipEffect.Parameter == "armor" {
			checkEffect(equipEffect, &unit.Armor)
		}

		if equipEffect.Parameter == "evasion_critical" {
			checkEffect(equipEffect, &unit.EvasionCritical)
		}

		if equipEffect.Parameter == "vulnerability_to_kinetics" {
			checkEffect(equipEffect, &unit.VulToKinetics)
		}

		if equipEffect.Parameter == "vulnerability_to_thermo" {
			checkEffect(equipEffect, &unit.VulToThermo)
		}

		if equipEffect.Parameter == "vulnerability_to_em" {
			checkEffect(equipEffect, &unit.VulToEM)
		}

		if equipEffect.Parameter == "vulnerability_to_explosion" {
			checkEffect(equipEffect, &unit.VulToExplosion)
		}

		if equipEffect.Parameter == "range_view" {
			checkEffect(equipEffect, &unit.RangeView)
		}

		if equipEffect.Parameter == "accuracy" {
			checkEffect(equipEffect, &unit.Accuracy)
		}

		if equipEffect.Parameter == "max_power" {
			checkEffect(equipEffect, &unit.MaxPower)
		}

		if equipEffect.Parameter == "recovery_power" {
			checkEffect(equipEffect, &unit.RecoveryPower)
		}

		if equipEffect.Parameter == "recovery_hp" {
			checkEffect(equipEffect, &unit.RecoveryHP)
		}
	}

	var checkPassiveEquip = func(equip map[int]*detail.BodyEquipSlot, gameUnit *Unit) {
		for _, slot := range equip {
			if slot.Equip != nil && !slot.Equip.Active {
				for _, equipEffect := range slot.Equip.Effects {
					checkParams(equipEffect)
				}
			}
		}
	}

	checkPassiveEquip(unit.Body.EquippingI, unit)
	checkPassiveEquip(unit.Body.EquippingII, unit)
	checkPassiveEquip(unit.Body.EquippingIII, unit)
	checkPassiveEquip(unit.Body.EquippingIV, unit)
	checkPassiveEquip(unit.Body.EquippingV, unit)

	// смотрим наложеные в игре эфекты
	for _, unitEffect := range unit.Effects {
		checkParams(unitEffect)
	}

	// высчитывает повер рековери
	unit.RecoveryPower = unit.Body.RecoveryPower - (unit.Body.GetUsePower() / 4)
	// востанавление энергии зависит от используемой энергии, чем больше обородования тем выше штраф
	// штраф так же должен зависеть от скила пользвотеля
}
