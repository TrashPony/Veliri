package gameTypes

import (
	"../../db/get"
	"../../gameObjects/ammo"
)

type AmmoStore struct {
	ammo map[int]ammo.Ammo
}

var Ammo = NewAmmoStore()

func NewAmmoStore() *AmmoStore {
	return &AmmoStore{ammo: get.AmmoType()}
}

func (a *AmmoStore) GetByID(id int) (*ammo.Ammo, bool) {
	var newAmmo ammo.Ammo
	newAmmo, ok := a.ammo[id]
	return &newAmmo, ok
}
