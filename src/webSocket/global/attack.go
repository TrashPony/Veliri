package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/attack"
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

	fireGun := func(attackUnit *unit.Unit) {
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

			if needRotate == attackUnit.GunRotate && dist <= weaponSlot.Weapon.Range {

				bullets, startAttack := attack.Fire(attackUnit)
				if startAttack {
					for _, bullet := range bullets {
						// todo отправка сообщения что бы проигралась анимация выстрела на клиенте
						go FlyBullet(bullet, attackUnit.MapID)
					}
				}

				if attackUnit.Target.Type == "map" {
					// если это атака тупо в карту то происхдит только 1 выстрел
					attackUnit.Target = nil
				}
			}
		}
	}

	for {
		if user != nil && user.GetSquad() != nil && user.GetSquad().MatherShip != nil {

			if user.GetSquad().MatherShip.GetWeaponSlot() != nil && user.GetSquad().MatherShip.GetWeaponSlot().Weapon != nil {
				fireGun(user.GetSquad().MatherShip)
				rotateGun(user.GetSquad().MatherShip)
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap &&
					unitSlot.Unit.GetWeaponSlot() != nil && unitSlot.Unit.GetWeaponSlot().Weapon != nil {

					fireGun(unitSlot.Unit)
					rotateGun(unitSlot.Unit)
				}
			}
		}

		time.Sleep(time.Duration(tickTime) * time.Millisecond) // время 1 такта поворота
	}
}

// функция которая заставляет лететь снаряды)
func FlyBullet(bullet *unit.Bullet, idMap int) {
	tickTime := 250

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	for {

		// TODO детальная проверка что бы пуля долетала прям до пикселя
		dist := game_math.GetBetweenDist(int(bullet.X), int(bullet.Y), int(bullet.Target.X), int(bullet.Target.Y))
		if dist < realSpeed+5 {
			break
		}

		radRotate := float64(bullet.Rotate) * math.Pi / 180
		stopX := realSpeed * math.Cos(radRotate) // идем по вектору движения выстрела
		stopY := realSpeed * math.Sin(radRotate)

		go SendMessage(Message{
			Event:    "FlyBullet",
			Bullet:   bullet,
			PathUnit: &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X + int(stopX), Y: bullet.Y + int(stopY), Millisecond: tickTime},
			IDMap:    idMap},
		)

		bullet.X, bullet.Y = bullet.X+int(stopX), bullet.Y+int(stopY)

		time.Sleep(time.Duration(tickTime) * time.Millisecond)
	}
}
