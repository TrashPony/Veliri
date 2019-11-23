package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/getlantern/deepcopy"
	"github.com/satori/go.uuid"
	"math/rand"
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
		// если количество пулей больше чем точек то пули вылетают по кругу

		bullet := &unit.Bullet{
			UUID:      uuid.NewV1().String(),
			FirePos:   i % len(firePos),
			Z:         1,
			Weapon:    attackUnit.GetWeaponSlot().Weapon,
			Ammo:      attackUnit.GetWeaponSlot().Ammo,
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

func AccuracyWeapon(attackUnit *unit.Unit, bullet *unit.Bullet) {

	weaponSlot := attackUnit.GetWeaponSlot()

	if bullet.Target.Type == "map" || bullet.Artillery {
		// если эта атака по карте или артилерия то точность выражена в площади поражения

		// чем дальше атака тем она менее точнее от 0 до 10ти
		dist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		percentRange := (dist * 100) / float64(attackUnit.GetWeaponRange())
		randomK := percentRange / 10

		getRandomOffsetZone := func() int {
			return rand.Intn(int(float64(weaponSlot.Weapon.Accuracy)*float64(randomK)*2)) - int(float64(weaponSlot.Weapon.Accuracy)*float64(randomK))
		}

		bullet.Target.X = bullet.Target.X + getRandomOffsetZone()
		bullet.Target.Y = bullet.Target.Y + getRandomOffsetZone()

		bullet.Rotate = game_math.GetBetweenAngle(float64(bullet.Target.X), float64(bullet.Target.Y), float64(bullet.X), float64(bullet.Y))

	} else {
		// если атака прямой атакой, то точно выражена в угле атаки
		bullet.Rotate = attackUnit.GunRotate + rand.Intn(weaponSlot.Weapon.Accuracy*2) - weaponSlot.Weapon.Accuracy
	}
}
