package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
)

func RotateGunToTarget(rotateGunUnit *unit.Unit, targetX, targetY int, time int) (*unit.PathUnit, int) {

	realRotateGun := int(float64(rotateGunUnit.GetWeaponSlot().Weapon.RotateSpeed) / float64(1000/time))

	xWeapon, yWeapon := rotateGunUnit.GetWeaponPos()
	needRotate := game_math.GetBetweenAngle(float64(targetX), float64(targetY), float64(xWeapon), float64(yWeapon))

	rotate := rotateGunUnit.GunRotate

	countRotateAngle := move.RotateUnit(&rotate, &needRotate, realRotateGun)
	return &unit.PathUnit{RotateGun: rotate, Millisecond: time, Animate: true}, countRotateAngle
}

func RotateGunToBody(rotateGunUnit *unit.Unit, time int) (*unit.PathUnit, int) {

	realRotateGun := int(float64(rotateGunUnit.GetWeaponSlot().Weapon.RotateSpeed) / float64(1000/time))

	needRotate := rotateGunUnit.Rotate
	rotate := rotateGunUnit.GunRotate

	countRotateAngle := move.RotateUnit(&rotate, &needRotate, realRotateGun)
	return &unit.PathUnit{RotateGun: rotate, Millisecond: time, Animate: true}, countRotateAngle
}
