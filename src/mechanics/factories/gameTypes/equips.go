package gameTypes

import (
	"../../db/get"
	"../../gameObjects/equip"
	"github.com/getlantern/deepcopy"
)

type EquipStore struct {
	equips map[int]equip.Equip
}

var Equips = NewEquipStore()

func NewEquipStore() *EquipStore {
	return &EquipStore{equips: get.EquipsType()}
}

func (e *EquipStore) GetByID(id int) (*equip.Equip, bool) {

	var newEquip equip.Equip
	factoryEquip, ok := e.equips[id]

	err := deepcopy.Copy(&newEquip, &factoryEquip) // функция глубокого копировния (very slow, but work)
	if err != nil {
		println(err.Error())
	}

	return &newEquip, ok
}

func (e *EquipStore) GetAllType() (map[int]equip.Equip) {
	return e.equips
}