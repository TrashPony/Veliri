package gameTypes

import (
	"../../db/get"
	"../../gameObjects/detail"
)

type weaponsStore struct {
	weapons map[int]detail.Weapon
}

var Weapons = newWeaponsStore()

func newWeaponsStore() *weaponsStore {
	return &weaponsStore{weapons: get.WeaponsType()}
}

func (w *weaponsStore) GetByID(id int) (*detail.Weapon, bool) {
	var newWeapon detail.Weapon
	newWeapon, ok := w.weapons[id]
	return &newWeapon, ok
}

func (w *weaponsStore) GetAllType() (map[int]detail.Weapon) {
	return w.weapons
}
