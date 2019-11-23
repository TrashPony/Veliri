package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"time"
)

func ReloadGun(reloadUnit *unit.Unit) {
	reloadUnit.GetWeaponSlot().Reload = true
	time.Sleep(time.Duration(reloadUnit.GetWeaponSlot().Weapon.ReloadTime) * time.Millisecond)
	reloadUnit.GetWeaponSlot().Reload = false
}
