package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/satori/go.uuid"
)

// метод определяет начально положение пули скорость выстрела и направление полета
func Fire(attackUnit *unit.Unit) ([]*unit.Bullet, bool) {

	bullets := make([]*unit.Bullet, 0)
	firePos := attackUnit.GetWeaponFirePos()

	//  создаем обьект пули, дать ему направление и начальную позицию
	for i := 0; i < attackUnit.GetWeaponSlot().Weapon.CountFireBullet; i++ {
		bullet := &unit.Bullet{
			UUID:    uuid.NewV1().String(),
			X:       firePos[i].X,
			Y:       firePos[i].Y,
			Z:       1,
			Weapon:  attackUnit.GetWeaponSlot().Weapon,
			Ammo:    attackUnit.GetWeaponSlot().Ammo,
			Rotate:  attackUnit.GunRotate,
			Speed:   attackUnit.GetWeaponSlot().Weapon.BulletSpeed + attackUnit.GetWeaponSlot().Ammo.BulletSpeed,
			Target:  attackUnit.GetTarget(),
			OwnerID: attackUnit.OwnerID,
		}

		bullets = append(bullets, bullet)
	}

	return bullets, true
}
