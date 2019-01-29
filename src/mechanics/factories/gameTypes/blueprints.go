package gameTypes

import (
	"../../db/get"
	"../../gameObjects/blueprints"
	"github.com/getlantern/deepcopy"
)

type bluePrintsStore struct {
	bluePrints map[int]blueprints.Blueprint
}

var BluePrints = NewBluePrints()

func NewBluePrints() *bluePrintsStore {
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
