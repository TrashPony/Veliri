package matherShip

import (
	"../unit"
	"../equip"
	"../detail"
	"../ammo"
	"../coordinate"
	"../effect"
)

type MatherShip struct {
	ID      int    `json:"id"`
	SquadID int    `json:"squad_id"`
	Owner   string `json:"owner"`

	Weapon *detail.Weapon `json:"weapon"`
	Body   *detail.Body   `json:"body"`
	Ammo   *ammo.Ammo     `json:"ammo"`

	Units      map[int]*unit.Unit     `json:"units"`     // в роли ключей карты выступают
	Equip      map[int]*equip.Equip   `json:"equip"`     // номера слотов

	MotherShipSlot int `json:"mother_ship_slot"`

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
