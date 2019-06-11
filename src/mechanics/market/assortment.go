package market

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/ammo"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/blueprints"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/box"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/resource"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/trashItem"
)

type Assortment struct {
	Bodies     map[int]detail.Body               `json:"bodies"`
	Weapons    map[int]detail.Weapon             `json:"weapons"`
	Ammo       map[int]ammo.Ammo                 `json:"ammo"`
	Equips     map[int]equip.Equip               `json:"equips"`
	Resources  map[int]resource.Resource         `json:"resources"`
	Recycles   map[int]resource.RecycledResource `json:"recycles"`
	Details    map[int]resource.CraftDetail      `json:"details"`
	Blueprints map[int]blueprints.Blueprint      `json:"blueprints"`
	Boxes      map[int]box.Box                   `json:"boxes"`
	Trash      map[int]trashItem.TrashItem       `json:"trash"`
}

func GetAssortment() *Assortment {
	return &Assortment{
		Bodies:     gameTypes.Bodies.GetAllType(),
		Weapons:    gameTypes.Weapons.GetAllType(),
		Ammo:       gameTypes.Ammo.GetAllType(),
		Equips:     gameTypes.Equips.GetAllType(),
		Resources:  gameTypes.Resource.GetAllBaseResource(),
		Recycles:   gameTypes.Resource.GetAllRecycled(),
		Details:    gameTypes.Resource.GetAllDetails(),
		Blueprints: gameTypes.BluePrints.GetAllType(),
		Boxes:      gameTypes.Boxes.GetAllType(),
		Trash:      gameTypes.TrashItems.GetAllType(),
	}
}
