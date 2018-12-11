package gameTypes

import (
	"../../db/get"
	"../../gameObjects/equip"
)

type EquipStore struct {
	equips map[int]equip.Equip
}

var Equips = NewEquipStore()

func NewEquipStore() *EquipStore {
	return &EquipStore{equips: get.EquipsType()}
}

func (e *EquipStore) GetByID(id int) (*equip.Equip, bool) {
	// TODO копировать эффекты из обьекта в карте в новый эквип т.к. это сылочный тип данных
	var newEquip equip.Equip
	newEquip, ok := e.equips[id]
	return &newEquip, ok
}
