package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/equip"
	"github.com/getlantern/deepcopy"
)

type equipStore struct {
	equips map[int]equip.Equip
}

var Equips = newEquipStore()

func newEquipStore() *equipStore {
	return &equipStore{equips: get.EquipsType()}
}

func (e *equipStore) GetByID(id int) (*equip.Equip, bool) {

	var newEquip equip.Equip
	factoryEquip, ok := e.equips[id]

	err := deepcopy.Copy(&newEquip, &factoryEquip) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &newEquip, ok
}

func (e *equipStore) GetAllType() map[int]equip.Equip {
	return e.equips
}
