package gameTypes

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/get"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"math/rand"
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

func (w *weaponsStore) GetRandom() *detail.Weapon {

	count := 0
	count2 := rand.Intn(len(w.weapons))
	for id := range w.weapons {
		if count == count2 {
			weapon, _ := w.GetByID(id)
			return weapon
		}
		count++
	}
	return nil

}

func (w *weaponsStore) GetAllType() map[int]detail.Weapon {
	return w.weapons
}
