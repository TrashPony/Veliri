package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func Attack(user *player.Player, msg Message) {

	mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)

	if msg.Type == "map" {
		// - атаковать в землю
		for _, unitID := range msg.UnitsID {
			fireUnit := user.GetSquad().GetUnitByID(unitID)
			if fireUnit.OnMap && fireUnit.GetWeaponSlot() != nil && fireUnit.GetWeaponSlot().Weapon != nil {
				attackToMap(user, fireUnit, mp, msg.X, msg.Y)
			}
		}
	}

	if msg.Type == "object" {
		// - атаковать обьекты на карте
		for _, unitID := range msg.UnitsID {
			fireUnit := user.GetSquad().GetUnitByID(unitID)
			if fireUnit.OnMap && fireUnit.GetWeaponSlot() != nil && fireUnit.GetWeaponSlot().Weapon != nil {
				attackToObject(user, fireUnit, mp, msg.X, msg.Y)
			}
		}
	}

	if msg.Type == "box" {
		// - ящик
		for _, unitID := range msg.UnitsID {
			fireUnit := user.GetSquad().GetUnitByID(unitID)
			if fireUnit.OnMap && fireUnit.GetWeaponSlot() != nil && fireUnit.GetWeaponSlot().Weapon != nil {
				attackToBox(user, fireUnit, mp, msg.BoxID)
			}
		}
	}

	if msg.Type == "unit" {
		// - игроков/нпс
		for _, unitID := range msg.UnitsID {
			fireUnit := user.GetSquad().GetUnitByID(unitID)
			if fireUnit.OnMap && fireUnit.GetWeaponSlot() != nil && fireUnit.GetWeaponSlot().Weapon != nil {
				attackToUnit(user, fireUnit, mp, msg.UnitID)
			}
		}
	}

	if msg.Type == "reservoir" {
		// todo атаковать руду, почему бы и нет? :D
	}

	if msg.Type == "transport" {
		// todo защитников баз
	}
}

func attackToMap(user *player.Player, fireUnit *unit.Unit, mp *_map.Map, toX, toY int) {
	xWeapon, yWeapon := fireUnit.GetWeaponPos()
	firePos := fireUnit.GetWeaponFirePos()

	dist := int(game_math.GetBetweenDist(toX, toY, xWeapon, yWeapon))

	units := globalGame.Clients.GetAllShortUnits(mp.Id)
	boxs := boxes.Boxes.GetAllBoxByMapID(mp.Id)

	delete(units, fireUnit.ID) // удаляем из карты что бы не обрабатывать в колизиях

	collisionInLine := collisions.SearchCircleCollisionInLine(float64(firePos[0].X), float64(firePos[0].Y),
		float64(toX), float64(toY), mp, 3, units, boxs, nil)

	if dist <= fireUnit.GetWeaponSlot().Weapon.Range && !collisionInLine {
		// — если цель в пределах достягаемости то атакуем но не отменяем прошлой команды
		fireUnit.SetTarget(&unit.Target{Type: "map", X: toX, Y: toY, Follow: false})
	} else {
		// — если цель вне досягаемости отменяем прошлую команду идем до цели
		fireUnit.FollowUnitID = 0
		fireUnit.SetTarget(&unit.Target{Type: "map", X: toX, Y: toY, Follow: true})
		Move(user, Message{ToX: float64(toX), ToY: float64(toY), UnitsID: []int{fireUnit.ID}}, false)
	}
}

func attackToObject(user *player.Player, fireUnit *unit.Unit, mp *_map.Map, xObj, yObj int) {
	obj := mp.GetDynamicObjects(xObj, yObj)
	if obj != nil {
		fireUnit.SetTarget(&unit.Target{Type: "object", ID: obj.ID, Follow: true})
		go FollowTarget(user, fireUnit, mp)
	}
}

func attackToBox(user *player.Player, fireUnit *unit.Unit, mp *_map.Map, id int) {
	mapBox, mx := boxes.Boxes.Get(id)
	mx.Unlock()

	if mapBox != nil && mapBox.MapID == fireUnit.MapID {
		fireUnit.SetTarget(&unit.Target{Type: "box", ID: mapBox.ID, Follow: true})
		go FollowTarget(user, fireUnit, mp)
	}
}

func attackToUnit(user *player.Player, fireUnit *unit.Unit, mp *_map.Map, id int) {
	targetUnit := globalGame.Clients.GetUnitByID(id)
	if targetUnit != nil && targetUnit.MapID == fireUnit.MapID {
		fireUnit.SetTarget(&unit.Target{Type: "unit", ID: targetUnit.ID, Follow: true})
		go FollowTarget(user, fireUnit, mp)
	}
}
