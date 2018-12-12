package gameTypes

import (
	"../../db/get"
	"../../gameObjects/detail"
)

type WeaponsStore struct {
	weapons map[int]detail.Weapon
}

var Weapons = NewWeaponsStore()

func NewWeaponsStore() *WeaponsStore {
	return &WeaponsStore{weapons: get.WeaponsType()}
}

func (w *WeaponsStore) GetByID(id int) (*detail.Weapon, bool) {
	var newWeapon detail.Weapon
	newWeapon, ok := w.weapons[id]
	return &newWeapon, ok
}

func (w *WeaponsStore) GetAllType() (map[int]detail.Weapon) {
	return w.weapons
}
