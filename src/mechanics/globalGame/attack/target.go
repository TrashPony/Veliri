package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

// TODO
func GetXYTarget(user *player.Player, tUnit *unit.Unit, target *unit.Target) bool {

	if target == nil {
		return false
	}

	if target.Type == "object" {
		obj := user.GetMapDynamicObjectByID(tUnit.MapID, target.ID)
		if obj == nil {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			tUnit.SetTarget(nil)
			return false
		} else {
			target.X, target.Y = obj.X, obj.Y
		}
	}

	if target.Type == "box" {
		mapBox, mx := boxes.Boxes.Get(target.ID)
		mx.Unlock()

		if mapBox == nil || mapBox.MapID != tUnit.MapID {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			tUnit.SetTarget(nil)
			return false
		} else {
			target.X, target.Y = mapBox.X, mapBox.Y
		}
	}

	if target.Type == "unit" {
		targetUnit := globalGame.Clients.GetUnitByID(target.ID)
		if targetUnit == nil || targetUnit.MapID != tUnit.MapID {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			tUnit.SetTarget(nil)
			return false
		} else {
			target.X, target.Y = targetUnit.X, targetUnit.Y
		}
	}

	if target.Type == "reservoir" {
		// todo атаковать руду, почему бы и нет? :D
	}

	if target.Type == "transport" {
		// todo защитников баз
	}

	return true
}

// TODO
func GetXYTarget2(user *player.Player, target *unit.Target, mp *_map.Map) bool {

	if target == nil {
		return false
	}

	if target.Type == "object" {
		obj := user.GetMapDynamicObjectByID(mp.Id, target.ID)
		if obj == nil {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			return false
		} else {
			target.X, target.Y = obj.X, obj.Y
		}
	}

	if target.Type == "box" {
		mapBox, mx := boxes.Boxes.Get(target.ID)
		mx.Unlock()

		if mapBox == nil || mapBox.MapID != mp.Id {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			return false
		} else {
			target.X, target.Y = mapBox.X, mapBox.Y
		}
	}

	if target.Type == "unit" {
		targetUnit := globalGame.Clients.GetUnitByID(target.ID)
		if targetUnit == nil || targetUnit.MapID != mp.Id {
			// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
			return false
		} else {
			target.X, target.Y = targetUnit.X, targetUnit.Y
		}
	}

	if target.Type == "reservoir" {
		// todo атаковать руду, почему бы и нет? :D
	}

	if target.Type == "transport" {
		// todo защитников баз
	}

	return true
}

func CheckFireToTarget(attackUnit *unit.Unit, mp *_map.Map, target *unit.Target) bool {

	if attackUnit.GetWeaponSlot().Reload {
		// оружие перезаряжается
		return false
	}

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
		if CollisionWeaponRangeCollision(attackUnit, mp, target) {
			return false
		}
	} else {
		return false
	}

	return true
}

func CollisionWeaponRangeCollision(attackUnit *unit.Unit, mp *_map.Map, target *unit.Target) bool {
	units := globalGame.Clients.GetAllShortUnits(mp.Id)
	boxs := boxes.Boxes.GetAllBoxByMapID(mp.Id)
	firePos := attackUnit.GetWeaponFirePos()

	delete(units, attackUnit.ID) // удаляем из карты что бы не обрабатывать в колизиях

	return collisions.SearchCircleCollisionInLine(float64(firePos[0].X), float64(firePos[0].Y),
		float64(target.X), float64(target.Y), mp, 2, units, boxs, target)
}
