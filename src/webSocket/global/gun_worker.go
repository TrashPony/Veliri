package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/attack"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
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

				if !FireGun(user.GetSquad().MatherShip, mp) {
					RotateGun(user, user.GetSquad().MatherShip, tickTime)
				}
			}

			for _, unitSlot := range user.GetSquad().MatherShip.Units {
				if unitSlot != nil && unitSlot.Unit != nil && unitSlot.Unit.OnMap &&
					unitSlot.Unit.GetWeaponSlot() != nil && unitSlot.Unit.GetWeaponSlot().Weapon != nil {

					if !FireGun(unitSlot.Unit, mp) {
						RotateGun(user, unitSlot.Unit, tickTime)
					}
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
	if target != nil {

		// TODO обьеденить эти блоки с блоками из follow_unit и attack они одинаковы
		if target.Type == "object" {
			obj := user.GetMapDynamicObjectByID(rotateUnit.MapID, target.ID)
			if obj == nil {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				rotateUnit.SetTarget(nil)
				return
			} else {
				target.X, target.Y = obj.X, obj.Y
			}
		}

		if target.Type == "box" {
			mapBox, mx := boxes.Boxes.Get(target.ID)
			mx.Unlock()

			if mapBox == nil || mapBox.MapID != rotateUnit.MapID {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				rotateUnit.SetTarget(nil)
				return
			} else {
				target.X, target.Y = mapBox.X, mapBox.Y
			}
		}

		if target.Type == "unit" {
			targetUnit := globalGame.Clients.GetUnitByID(target.ID)
			if targetUnit == nil || targetUnit.MapID != rotateUnit.MapID {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				rotateUnit.SetTarget(nil)
			} else {
				target.X, target.Y = targetUnit.X, targetUnit.Y
			}
		}

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

func FireGun(attackUnit *unit.Unit, mp *_map.Map) bool {

	target := attackUnit.GetTarget()
	weaponSlot := attackUnit.GetWeaponSlot()

	if target != nil && CheckFireToTarget(attackUnit, mp, target) {

		bullets, startAttack := attack.Fire(attackUnit)
		if startAttack {

			for _, bullet := range bullets {

				// для отыгрыша анимации выстрела
				SendMessage(Message{
					Event:         "FireWeapon",
					X:             bullet.X,
					Y:             bullet.Y,
					Weapon:        weaponSlot.Weapon,
					IDMap:         attackUnit.MapID,
					NeedCheckView: true,
				})

				go FlyBullet(bullet, attackUnit.MapID)
				time.Sleep(time.Duration(weaponSlot.Weapon.DelayFollowingFire) * time.Millisecond)
			}

			if target.Type == "map" {
				// если это атака тупо в карту то происхдит только 1 выстрел
				attackUnit.SetTarget(nil)
			}

			return true
		}
	}

	return false
}

func CheckFireToTarget(attackUnit *unit.Unit, mp *_map.Map, target *unit.Target) bool {
	// TODO проверка перезарядки

	// смотрим что бы оружие было повернуто в необходимом положение
	xWeapon, yWeapon := attackUnit.GetWeaponPos()
	needRotate := game_math.GetBetweenAngle(float64(target.X), float64(target.Y), float64(xWeapon), float64(yWeapon))
	if needRotate < 0 {
		needRotate += 360
	}

	if debug.Store.WeaponFirePos {
		debug.Store.AddMessage("CreateLine", "orange", target.X,
			target.Y, xWeapon, yWeapon, 0, attackUnit.MapID, 0)
	}

	if needRotate == attackUnit.GunRotate && attackUnit.GetDistWeaponToTarget() <= attackUnit.GetWeaponRange() {
		// и между оружием и целью нет колизий
		if collisionWeaponRangeCollision(attackUnit, mp, target) {
			return false
		}
	} else {
		return false
	}

	return true
}

func collisionWeaponRangeCollision(attackUnit *unit.Unit, mp *_map.Map, target *unit.Target) bool {
	units := globalGame.Clients.GetAllShortUnits(mp.Id)
	boxs := boxes.Boxes.GetAllBoxByMapID(mp.Id)
	firePos := attackUnit.GetWeaponFirePos()

	delete(units, attackUnit.ID) // удаляем из карты что бы не обрабатывать в колизиях

	return collisions.SearchCircleCollisionInLine(float64(firePos[0].X), float64(firePos[0].Y),
		float64(target.X), float64(target.Y), mp, 3, units, boxs, target)
}
