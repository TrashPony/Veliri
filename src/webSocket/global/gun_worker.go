package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/attack"
	"time"
)

// функция следит за орудиями всех юнитов
// если есть цель то орудие всегда повернуто к цели, и функция следит что бы оно так и было независимо от движения
// если цели нет то орудие всегда идет по курсу тела
// todo расчет угла по движещейся цели
func GunWorker(user *player.Player) {

	user.GetSquad().GunWorkerWork = true
	defer func() {
		user.GetSquad().GunWorkerWork = false
	}()

	mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	tickTime := 100

	for {

		if user.GetSquad().GunWorkerExit {
			user.GetSquad().GunWorkerExit = false
			return
		}

		if user != nil && user.GetSquad() != nil && user.GetSquad().MatherShip != nil {

			if user.GetSquad().MatherShip.GetWeaponSlot() != nil && user.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {

				go FireGun(user, user.GetSquad().MatherShip, mp)
				RotateGun(user, user.GetSquad().MatherShip, tickTime)
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap &&
					unitSlot.Unit.GetWeaponSlot() != nil && unitSlot.Unit.GetWeaponSlot().Weapon != nil {

					go FireGun(user, unitSlot.Unit, mp)
					RotateGun(user, unitSlot.Unit, tickTime)
				}
			}
		}

		time.Sleep(time.Duration(tickTime) * time.Millisecond) // время 1 такта поворота
	}
}

func RotateGun(user *player.Player, rotateUnit *unit.Unit, tickTime int) {
	sendRotate := func(rotateUnit *unit.Unit, pathUnit *unit.PathUnit) {
		go SendMessage(Message{
			Event:         "RotateGun",
			ShortUnit:     rotateUnit.GetShortInfo(),
			PathUnit:      pathUnit,
			IDMap:         rotateUnit.MapID,
			NeedCheckView: true},
		)

		rotateUnit.GunRotate = pathUnit.RotateGun
	}

	target := rotateUnit.GetTarget()
	if attack.GetXYTarget(user, rotateUnit, target) {

		pathUnit, diffAngle := attack.RotateGunToTarget(
			rotateUnit,
			target.X,
			target.Y,
			tickTime,
		)

		if diffAngle > 0 {
			go sendRotate(rotateUnit, pathUnit)
		}
	} else {
		pathUnit, diffAngle := attack.RotateGunToBody(rotateUnit, tickTime)
		if diffAngle > 0 {
			go sendRotate(rotateUnit, pathUnit)
		}
	}
}

func FireGun(user *player.Player, attackUnit *unit.Unit, mp *_map.Map) {

	if attackUnit.FreezeGun {
		return
	}

	defer func() {
		attackUnit.FreezeGun = false
	}()
	attackUnit.FreezeGun = true

	target := attackUnit.GetTarget()
	weaponSlot := attackUnit.GetWeaponSlot()

	if target != nil && attack.CheckFireToTarget(attackUnit, mp, target) {

		bullets, startAttack := attack.Fire(user, attackUnit)
		if startAttack {

			go attack.ReloadGun(attackUnit)

			for _, bullet := range bullets {

				// если юнит находится в движение то из за DelayFollowingFire, позиция сместиться
				realPos := attackUnit.GetWeaponFirePos()[bullet.FirePos]
				bullet.X, bullet.Y = realPos.X, realPos.Y

				// влияние точности оружия на выстрел
				attack.AccuracyWeapon(attackUnit, bullet)

				// для отыгрыша анимации выстрела
				SendMessage(Message{
					Event:         "FireWeapon",
					X:             bullet.X,
					Y:             bullet.Y,
					Weapon:        weaponSlot.Weapon,
					IDMap:         attackUnit.MapID,
					NeedCheckView: true,
				})

				// лазеры особенные :)
				if weaponSlot.Weapon.Type == "laser" {
					go FlyLaser(bullet, mp, attackUnit)
				} else {
					go FlyBullet(user, bullet, mp, attackUnit)
				}

				// задержка орудия после выстрела, если с 1 ордуия летит много снарядов то они вылетят не одновременно
				time.Sleep(time.Duration(weaponSlot.Weapon.DelayFollowingFire) * time.Millisecond)
			}

			if target.Type == "map" {
				// если это атака тупо в карту то происхдит только 1 выстрел
				attackUnit.SetTarget(nil)
			}
		}
	}
}
