package market

import (
	"../gameObjects/ammo"
	"../gameObjects/detail"
	"../gameObjects/equip"
)

type Assortment struct {
	Bodies  []*detail.Body   `json:"bodies"`
	Weapons []*detail.Weapon `json:"weapons"`
	Ammo    []*ammo.Ammo     `json:"ammo"`
	Equips  []*equip.Equip   `json:"equips"`
}

func GetAssortment() *Assortment {
	return nil //&Assortment{Bodies: get.BodiesType(), Weapons:get.WeaponsType(), Ammo: get.AmmoType(), Equips: get.EquipsType()}
}
