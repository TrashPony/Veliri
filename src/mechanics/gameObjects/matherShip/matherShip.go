package matherShip

import (
	"../unit"
	"../detail"
	"../coordinate"
	"../effect"
)

type MatherShip struct {
	ID      int    `json:"id"`
	SquadID int    `json:"squad_id"`
	Owner   string `json:"owner"`

	Body   *detail.Body   `json:"body"`

	Units      map[int]*unit.Unit     `json:"units"`     // в роли ключей карты выступают

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

func (matherShip *MatherShip) GetID() int  {
	return matherShip.ID
}

func (matherShip *MatherShip) GetBody() *detail.Body  {
	return matherShip.Body
}

func (matherShip *MatherShip) GetX() int {
	return matherShip.X
}

func (matherShip *MatherShip) GetY() int {
	return matherShip.Y
}

func (matherShip *MatherShip) GetWatchZone() int {
	return matherShip.Body.RangeView
}

func (matherShip *MatherShip) GetOwnerUser() string {
	return matherShip.Owner
}

func (matherShip *MatherShip) GetTarget() *coordinate.Coordinate  {
	return matherShip.Target
}