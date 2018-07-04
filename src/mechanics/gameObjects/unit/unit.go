package unit

import (
	"../detail"
	"../effect"
	"../coordinate"
)

type Unit struct {
	ID      int    `json:"id"`
	SquadID int    `json:"squad_id"`
	Owner   string `json:"owner"`

	Body   *detail.Body   `json:"body"`

	X      int  `json:"x"`
	Y      int  `json:"y"`
	Rotate int  `json:"rotate"`
	OnMap  bool `json:"on_map"`

	Action      bool                   `json:"action"`
	Target      *coordinate.Coordinate `json:"target"`
	QueueAttack int                    `json:"queue_attack"`

	HP int `json:"hp"`

	Effects []*effect.Effect `json:"effects"`
}

func (unit *Unit) GetID() int  {
	return unit.ID
}

func (unit *Unit) GetBody() *detail.Body  {
	return unit.Body
}

func (unit *Unit) GetTarget() *coordinate.Coordinate  {
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

func (unit *Unit) SetX(x int) {
	unit.X = x
}

func (unit *Unit) GetX() int {
	return unit.X
}

func (unit *Unit) SetY(y int) {
	unit.Y = y
}

func (unit *Unit) GetY() int {
	return unit.Y
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