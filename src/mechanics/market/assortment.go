package market

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/ammo"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
)

type Assortment struct {
	Bodies  map[int]detail.Body   `json:"bodies"`
	Weapons map[int]detail.Weapon `json:"weapons"`
	Ammo    map[int]ammo.Ammo     `json:"ammo"`
	Equips  map[int]equip.Equip   `json:"equips"`
}

func GetAssortment() *Assortment {
	return &Assortment{Bodies: gameTypes.Bodies.GetAllType(), Weapons: gameTypes.Weapons.GetAllType(),
		Ammo: gameTypes.Ammo.GetAllType(), Equips: gameTypes.Equips.GetAllType()}
}
