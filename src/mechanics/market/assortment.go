package market

import (
	"../gameObjects/detail"
	"../gameObjects/ammo"
	"../gameObjects/equip"
	"../db/get"
)

type Assortment struct {
	Bodies  []*detail.Body   `json:"bodies"`
	Weapons []*detail.Weapon `json:"weapons"`
	Ammo    []*ammo.Ammo     `json:"ammo"`
	Equips  []*equip.Equip   `json:"equips"`
}

func GetAssortment() *Assortment{
	return &Assortment{Bodies: get.GetBodiesType(), Weapons:get.GetWeaponsType(), Ammo: get.GetAmmoType(), Equips: get.GetEquipsType()}
}
