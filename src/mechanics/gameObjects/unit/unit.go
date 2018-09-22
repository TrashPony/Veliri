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
