package gameTypes

import (
	"../../db/get"
	"../../gameObjects/blueprints"
	"github.com/getlantern/deepcopy"
)

type bluePrintsStore struct {
	bluePrints map[int]blueprints.Blueprint
}

var BluePrints = newBluePrints()

func newBluePrints() *bluePrintsStore {
	return &bluePrintsStore{bluePrints: get.BlueprintsType()}
}

func (b *bluePrintsStore) GetAllType() map[int]blueprints.Blueprint {
	return b.bluePrints
}

func (b *bluePrintsStore) GetByID(id int) (*blueprints.Blueprint, bool) {
	var newBluePrint blueprints.Blueprint
	newBluePrint, ok := b.bluePrints[id]
	return &newBluePrint, ok
}

func (b *bluePrintsStore) GetByItemTypeAndID(itemID int, itemType string) *blueprints.Blueprint {
	var newBluePrint blueprints.Blueprint

	for _, bluePrint := range b.bluePrints {
		if bluePrint.ItemId == itemID && bluePrint.ItemType == itemType {
			err := deepcopy.Copy(&newBluePrint, &bluePrint) // функция глубокого копировния (very slow, but work)
			if err != nil {
				println(err.Error())
			} else {
				return &newBluePrint
			}
		}
	}

	return nil
}

func (b *bluePrintsStore) GetItems(id int) interface{} {
	bp, ok := b.bluePrints[id]

	if ok {
		if bp.ItemType == "weapon" {
			weapon, _ := Weapons.GetByID(bp.ItemId)
			return weapon
		}
		if bp.ItemType == "equip" {
			equip, _ := Equips.GetByID(bp.ItemId)
			return equip
		}
		if bp.ItemType == "detail" {
			detail, _ := Resource.GetDetailByID(bp.ItemId)
			return detail
		}
		if bp.ItemType == "ammo" {
			ammo, _ := Ammo.GetByID(bp.ItemId)
			return ammo
		}
		if bp.ItemType == "body" {
			body, _ := Bodies.GetByID(bp.ItemId)
			return body
		}
		if bp.ItemType == "boxes" {
			box, _ := Boxes.GetByID(bp.ItemId)
			return box
		}
	}

	return nil
}
