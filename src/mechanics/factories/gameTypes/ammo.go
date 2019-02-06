package gameTypes

import (
	"../../db/get"
	"../../gameObjects/ammo"
)

type ammoStore struct {
	ammo map[int]ammo.Ammo
}

var Ammo = newAmmoStore()

func newAmmoStore() *ammoStore {
	return &ammoStore{ammo: get.AmmoType()}
}

func (a *ammoStore) GetByID(id int) (*ammo.Ammo, bool) {
	var newAmmo ammo.Ammo
	newAmmo, ok := a.ammo[id]
	return &newAmmo, ok
}

func (a *ammoStore) GetAllType() (map[int]ammo.Ammo) {
	return a.ammo
}