package attack

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func CollisionDamage(typeTarget string, idTarget int, bullet *unit.Bullet, mp *_map.Map) {
	// TODO вовзращать список убитых обьектов

	if bullet.Ammo.AreaCovers > 0 {
		Explosion(bullet, mp)
	} else {
		damage, equipDamage := bullet.Unit.GetDamage()
		DamageTarget(typeTarget, idTarget, mp, damage, equipDamage)
	}
}

func Explosion(bullet *unit.Bullet, mp *_map.Map) {
	// TODO вовзращать список убитых обьектов

	damage, equipDamage := bullet.Unit.GetDamage()
	// тип снаряда имеет зону поражения то дамаг по всему что вокруг

	explosionDamage := func(x, y, areaCovers int) int {
		// чем ближе к эпицентру тем более полный урон

		dist := game_math.GetBetweenDist(bullet.X, bullet.Y, x, y)

		if int(dist) < areaCovers {
			percentRange := (dist * 100) / float64(bullet.Ammo.AreaCovers)
			return damage/2 + int(float64(damage/2)*(percentRange/100))
		} else {
			// до обьекта снаряд не дастрелил, а получил он в кусочек геодаты поэтому получает половину урона
			return damage / 2
		}
	}

	for _, x := range mp.GetCopyMapDynamicObjects() {
		for _, obj := range x {
			if collisions.CircleDynamicObj(bullet.X, bullet.Y, bullet.Ammo.AreaCovers, obj) {
				DamageTarget("object", obj.ID, mp, explosionDamage(obj.X, obj.Y, bullet.Ammo.AreaCovers), equipDamage)
			}
		}
	}
}

func DamageTarget(typeTarget string, idTarget int, mp *_map.Map, damage, equipDamage int) {
	if typeTarget == "object" {
		obj := mp.GetDynamicObjectsByID(idTarget)
		obj.HP -= damage
		if obj.HP < 0 {
			mp.RemoveDynamicObject(obj)
		}
	}
}
