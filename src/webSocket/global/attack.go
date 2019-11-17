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
		// todo обьекты на карте
	}

	if msg.Type == "unit" {
		// todo игроков/нпс
	}

	if msg.Type == "box" {
		// todo ящик
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
		float64(toX), float64(toY), mp, 3, units, boxs)

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
