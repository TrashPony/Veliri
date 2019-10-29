package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/attack"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"time"
)

func Attack(user *player.Player, msg Message) {
	//todo методы атаки,
	// просто землю
	// обьекты на карте
	// ящик
	// игроков/нпс

	// TODO адаптировать под всех юнитов
	if msg.Type == "map" {

		for _, unitID := range msg.UnitsID {
			fireUnit := user.GetSquad().GetUnitByID(unitID)
			fireUnit.Target = &unit.Target{Type: msg.Type, X: msg.X, Y: msg.Y}
		}

		// todo если есть колизия между целью и юнитом или цель за пределами радиуса
		//  то ищем путь до цели пока нет будет доступна атака
		//  по карте мы стреляем 1 раз
	}

	// если цель юнит/ящик/обьект то берется Х:У по вектору проходящий через юнита в конце дальности атаки оружия
	// тогда оружие промахивается и оружие бьет дальше - это относится ко всему кроме артилерии и тактических ракет
	if msg.Type == "unit" {

	}

	if msg.Type == "box" {

	}

	if msg.Type == "object" {

	}
}

// функция следит за орудиями всех юнитов
// если есть цель то орудие всегда повернуто к цели, и функция следит что бы оно так и было независимо от движения
// если цели нет то орудие всегда идет по курсу тела
func GunWorker(user *player.Player) {

	tickTime := 250

	sendRotate := func(rotateUnit *unit.Unit, pathUnit *unit.PathUnit) {
		go SendMessage(Message{
			Event:     "RotateGun",
			ShortUnit: rotateUnit.GetShortInfo(),
			PathUnit:  pathUnit,
			IDMap:     rotateUnit.MapID},
		)
	}

	rotateGun := func(rotateUnit *unit.Unit) {

		if rotateUnit.Target != nil {

			pathUnit, diffAngle := attack.RotateGunToTarget(
				rotateUnit,
				rotateUnit.Target.X,
				rotateUnit.Target.Y,
				tickTime,
			)

			if diffAngle > 0 {
				sendRotate(rotateUnit, pathUnit)
				rotateUnit.GunRotate = pathUnit.RotateGun
			}
		} else {
			pathUnit, diffAngle := attack.RotateGunToBody(rotateUnit, tickTime)
			if diffAngle > 0 {
				sendRotate(rotateUnit, pathUnit)
				rotateUnit.GunRotate = pathUnit.RotateGun
			}
		}
	}

	fireGun := func(attackUnit *unit.Unit) bool {
		if attackUnit.Target != nil {

			// TODO проверка перезарядки
			// смотрим что бы оружие было повернуто в необходимом положение
			xWeapon, yWeapon := attackUnit.GetWeaponPos()
			needRotate := game_math.GetBetweenAngle(float64(attackUnit.Target.X), float64(attackUnit.Target.Y), float64(xWeapon), float64(yWeapon))
			if needRotate < 0 {
				needRotate += 360
			}

			// и дистанция оружия доставала до цели
			weaponSlot := attackUnit.GetWeaponSlot()
			dist := int(game_math.GetBetweenDist(attackUnit.Target.X, attackUnit.Target.Y, xWeapon, yWeapon))

			if debug.Store.WeaponFirePos {
				debug.Store.AddMessage("CreateLine", "orange", attackUnit.Target.X,
					attackUnit.Target.Y, xWeapon, yWeapon, 0, attackUnit.MapID, 0)
			}

			if needRotate == attackUnit.GunRotate && dist <= weaponSlot.Weapon.Range {

				bullets, startAttack := attack.Fire(attackUnit)
				if startAttack {
					for _, bullet := range bullets {

						// для отыгрыша анимации выстрела
						SendMessage(Message{
							Event:  "FireWeapon",
							X:      bullet.X,
							Y:      bullet.Y,
							Weapon: weaponSlot.Weapon,
							IDMap:  attackUnit.MapID,
						})

						go FlyBullet(bullet, attackUnit.MapID)
						time.Sleep(time.Duration(weaponSlot.Weapon.DelayFollowingFire) * time.Millisecond)
					}

					if attackUnit.Target.Type == "map" {
						// если это атака тупо в карту то происхдит только 1 выстрел
						attackUnit.Target = nil
					}

					return true
				}
			}
		}

		return false
	}

	for {
		if user != nil && user.GetSquad() != nil && user.GetSquad().MatherShip != nil {

			if user.GetSquad().MatherShip.GetWeaponSlot() != nil && user.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {

				if !fireGun(user.GetSquad().MatherShip) {
					rotateGun(user.GetSquad().MatherShip)
				}
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap &&
					unitSlot.Unit.GetWeaponSlot() != nil && unitSlot.Unit.GetWeaponSlot().Weapon != nil {

					if !fireGun(unitSlot.Unit) {
						rotateGun(unitSlot.Unit)
					}
				}
			}
		}

		time.Sleep(time.Duration(tickTime) * time.Millisecond) // время 1 такта поворота
	}
}

// полет лазера
func FlyLaser() {
	// TODO
}

// самоводящиеся ракеты
func FlyChaseRocket() {
	// TODO
}

// артилерийские установки
func FlyArtillery() {
	// TODO
}

// функция которая заставляет лететь снаряды летящие по прямой
func FlyBullet(bullet *unit.Bullet, idMap int) {

	tickTime := 100

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	radRotate := float64(bullet.Rotate) * math.Pi / 180

	// пуля летит по параболе, поэтому до половины пути она немного приподнимается по Z,
	// а после половины пути стремительно идет к 0, это сказывается на анимации фронта, но не на логике
	startDist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
	minDist := startDist

	SendMessage(Message{
		Event:    "FlyBullet",
		Bullet:   bullet,
		PathUnit: &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: tickTime},
		IDMap:    idMap},
	)
	time.Sleep(time.Duration(tickTime) * time.Millisecond)

	for {

		currentDist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		percentPath := 100 - (currentDist*100)/startDist

		if percentPath > 50 {
			// после 50% до цели пуля снижается и в конце пути удаляется о землю
			bullet.Z = 1 - (((percentPath - 50) * 2) / 100)
			if bullet.Z < 0 {
				bullet.Z = 0
			}
		}

		stopX := realSpeed * math.Cos(radRotate) // идем по вектору движения выстрела
		stopY := realSpeed * math.Sin(radRotate)

		percent, end := detailFlyBullet(bullet, float64(bullet.X)+stopX, float64(bullet.Y)+stopY, radRotate)

		ms := (tickTime * percent) / 100

		go SendMessage(Message{
			Event:    "FlyBullet",
			Bullet:   bullet,
			PathUnit: &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: ms},
			IDMap:    idMap},
		)

		if end || minDist < currentDist {
			// для отыгрыша анимации взрыва
			// TODO появление динамического обьекта кратера
			SendMessage(Message{
				Event:  "ExplosionBullet",
				Bullet: bullet,
				IDMap:  idMap,
			})

			break
		}

		minDist = currentDist
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

func detailFlyBullet(bullet *unit.Bullet, toX, toY, radRotate float64) (int, bool) {

	startDist := game_math.GetBetweenDist(bullet.X, bullet.Y, int(toX), int(toY))
	minDist := startDist
	dist := startDist

	x, y := float64(bullet.X), float64(bullet.Y)

	for {
		percentPath := 100 - (int((dist * 100) / startDist))

		stopX := float64(1) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(1) * math.Sin(radRotate)

		x += stopX
		y += stopY

		dist = game_math.GetBetweenDist(int(x), int(y), int(toX), int(toY))
		if dist <= 3 || minDist < dist {
			bullet.X, bullet.Y = int(x), int(y)
			return percentPath, 25 > game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		}

		minDist = dist
	}
}
