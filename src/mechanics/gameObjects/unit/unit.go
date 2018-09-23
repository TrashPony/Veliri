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

// // // //

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
	return unit.Body.RangeView
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
	return unit.Body.WallHack
}

func (unit *Unit) CalculateParams() {
	// начальные параметры тела
	unit.Speed = unit.Body.Speed
	unit.Initiative = unit.Body.Initiative
	unit.MaxHP = unit.Body.MaxHP
	unit.Armor = unit.Body.Armor
	unit.EvasionCritical = unit.Body.EvasionCritical
	unit.VulToKinetics = unit.Body.VulToKinetics
	unit.VulToThermo = unit.Body.VulToThermo
	unit.VulToEM = unit.Body.VulToEM
	unit.VulToExplosion = unit.Body.VulToExplosion
	unit.RangeView = unit.Body.RangeView
	unit.Accuracy = unit.Body.Accuracy
	unit.MaxPower = unit.Body.MaxPower
	unit.RecoveryPower = unit.Body.RecoveryPower
	unit.WallHack = unit.Body.WallHack

	// смотрим пасивное обородование
	// todo

	// смотрим эффекты которые весят на юните
	// todo

	// высчитывает повер рековери
	unit.RecoveryPower = unit.Body.RecoveryPower - (unit.Body.GetUsePower() / 4)
	// востанавление энергии зависит от используемой энергии, чем больше обородования тем выше штраф
	// TODO штраф так же должен зависеть от скила пользвотеля
}
