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

			if debug.Store.WeaponFirePos {
				debug.Store.AddMessage("CreateLine", "orange", attackUnit.Target.X,
					attackUnit.Target.Y, xWeapon, yWeapon, 0, attackUnit.MapID, 0)
			}

			if needRotate == attackUnit.GunRotate && dist <= weaponSlot.Weapon.Range {

				bullets, startAttack := attack.Fire(attackUnit)
				if startAttack {
					for _, bullet := range bullets {
						// todo отправка сообщения что бы проигралась анимация выстрела на клиенте
						// задержка перед выстрелом
						time.Sleep(250 * time.Millisecond)
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
				rotateGun(user.GetSquad().MatherShip)
				fireGun(user.GetSquad().MatherShip)
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap &&
					unitSlot.Unit.GetWeaponSlot() != nil && unitSlot.Unit.GetWeaponSlot().Weapon != nil {

					rotateGun(unitSlot.Unit)
					fireGun(unitSlot.Unit)
				}
			}
		}

		time.Sleep(time.Duration(tickTime) * time.Millisecond) // время 1 такта поворота
	}
}

// функция которая заставляет лететь снаряды)
func FlyBullet(bullet *unit.Bullet, idMap int) {
	tickTime := 100

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	radRotate := float64(bullet.Rotate) * math.Pi / 180

	SendMessage(Message{
		Event:    "FlyBullet",
		Bullet:   bullet,
		PathUnit: &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: tickTime},
		IDMap:    idMap},
	)
	time.Sleep(time.Duration(tickTime) * time.Millisecond)

	for {

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

		if end {
			break
		}

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

		stopX := float64(2) * math.Cos(radRotate) // идем по вектору движения корпуса
		stopY := float64(2) * math.Sin(radRotate)

		x += stopX
		y += stopY

		dist = game_math.GetBetweenDist(int(x), int(y), int(toX), int(toY))
		if dist <= 1 || minDist < dist {
			bullet.X, bullet.Y = int(x), int(y)
			return percentPath, 15 > game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		}

		minDist = dist
	}
}
