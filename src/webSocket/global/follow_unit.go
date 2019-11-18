package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"time"
)

func FollowUnit(user *player.Player, moveUnit *unit.Unit, msg Message) {
	// если юнит преследует другово юнита, то достаем его и мониторим его положение
	// если по какойто причине (столкновение, гравитация и тд) надо перестроить маршрут то сохраняем FollowUnitID
	// однако если сам игрок сгенерил событие движения то мы не сохраняем параметр FollowUnitID

	var followUnit *unit.Unit
	if moveUnit.FollowUnitID != 0 {
		followUnit = globalGame.Clients.GetUnitByID(moveUnit.FollowUnitID)
	} else {
		return
	}

	if followUnit != nil {
		for {

			if moveUnit.FollowUnitID == 0 || !moveUnit.OnMap || !followUnit.OnMap || moveUnit.MapID != followUnit.MapID {
				moveUnit.FollowUnitID = 0
				moveUnit.Return = false
				return
			}

			dist := game_math.GetBetweenDist(followUnit.X, followUnit.Y, int(moveUnit.X), int(moveUnit.Y))
			if dist < 90 {

				stopMove(moveUnit, true)

				if moveUnit.Return {
					go ReturnUnit(user, moveUnit)
					return
				}

				time.Sleep(100 * time.Millisecond)
				continue
			}

			dist = game_math.GetBetweenDist(followUnit.X, followUnit.Y, int(moveUnit.ToX), int(moveUnit.ToY))
			if dist > 90 || moveUnit.ActualPath == nil {
				msg.ToX = float64(followUnit.X)
				msg.ToY = float64(followUnit.Y)
				Move(user, msg, false)
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}

func FollowTarget(user *player.Player, followUnit *unit.Unit, mp *_map.Map) {

	for {

		target := followUnit.GetTarget()
		if target == nil || !target.Follow {
			// юнит перестал преследовать цель
			return
		}

		if target.Type == "object" {
			obj := user.GetMapDynamicObjectByID(followUnit.MapID, target.ID)
			if obj == nil {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				followUnit.SetTarget(nil)
			} else {
				target.X, target.Y = obj.X, obj.Y
			}
		}

		if target.Type == "box" {
			mapBox, mx := boxes.Boxes.Get(target.ID)
			mx.Unlock()

			if mapBox == nil || mapBox.MapID != followUnit.MapID {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				followUnit.SetTarget(nil)
			} else {
				target.X, target.Y = mapBox.X, mapBox.Y
			}
		}

		if target.Type == "unit" {
			targetUnit := globalGame.Clients.GetUnitByID(target.ID)
			if targetUnit == nil || targetUnit.MapID != followUnit.MapID {
				// по той или иной причине юнит перестал видит цель и больше не знает существует оно или нет
				followUnit.SetTarget(nil)
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

		// преследовать если оружия не достает (-50 что бы не рыпатся при любом движение цели) или если не прострельнут до цели
		if followUnit.GetDistWeaponToTarget() < followUnit.GetWeaponRange()-50 && !collisionWeaponRangeCollision(followUnit, mp, target) {
			// иначе стоим стреляем до отмены приказа или пока цель не пропадет
			stopMove(followUnit, true)
		} else {
			// что бы не генерить всегда новые события проверяем, может юнит уже на пути к цели
			dist := int(game_math.GetBetweenDist(target.X, target.Y, int(followUnit.ToX), int(followUnit.ToY)))
			if !followUnit.MoveChecker || dist > followUnit.GetWeaponRange()-50 {
				Move(user, Message{ToX: float64(target.X), ToY: float64(target.Y), UnitsID: []int{followUnit.ID}}, false)
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}
