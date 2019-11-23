package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/getlantern/deepcopy"
	"github.com/satori/go.uuid"
)

// метод определяет начально положение пули скорость выстрела и направление полета
func Fire(user *player.Player, attackUnit *unit.Unit) ([]*unit.Bullet, bool) {

	bullets := make([]*unit.Bullet, 0)
	firePos := attackUnit.GetWeaponFirePos()

	target := attackUnit.GetTarget()
	GetXYTarget(user, attackUnit, target)

	bulletTarget := &unit.Target{}
	err := deepcopy.Copy(bulletTarget, target)
	if err != nil {
		return nil, false
	}

	//  создаем обьект пули, дать ему направление и начальную позицию
	for i := 0; i < attackUnit.GetWeaponSlot().Weapon.CountFireBullet; i++ {

		bullet := &unit.Bullet{
			UUID:      uuid.NewV1().String(),
			X:         firePos[i].X,
			Y:         firePos[i].Y,
			Z:         1,
			Weapon:    attackUnit.GetWeaponSlot().Weapon,
			Ammo:      attackUnit.GetWeaponSlot().Ammo,
			Rotate:    attackUnit.GunRotate,
			Speed:     attackUnit.GetWeaponSlot().Weapon.BulletSpeed + attackUnit.GetWeaponSlot().Ammo.BulletSpeed,
			Target:    bulletTarget,
			OwnerID:   attackUnit.OwnerID,
			UnitID:    attackUnit.ID,
			MaxRange:  attackUnit.GetWeaponRange(),
			Artillery: attackUnit.GetWeaponSlot().Weapon.Artillery,
		}

		bullets = append(bullets, bullet)
	}

	return bullets, true
}
