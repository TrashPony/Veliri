package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"time"
)

const tickTime = 100

// полет лазера
func FlyLaser(bullet *unit.Bullet, gameMap *_map.Map) {
	// лазер летит со скоростью света, поэтому что все нам надо это отдать стартовое ХУ и конечную ХУ
	// конечная ХУ это координата колизии или карты куда стреляет игрок

	radRotate := float64(bullet.Rotate) * math.Pi / 180

	startX, startY := bullet.X, bullet.Y

	if bullet.Target.Type != "map" {
		bullet.Target.X = bullet.X + int(float64(bullet.MaxRange)*math.Cos(radRotate))
		bullet.Target.Y = bullet.Y + int(float64(bullet.MaxRange)*math.Sin(radRotate))
	}

	_, _, deltaTime := detailFlyBullet(bullet, float64(bullet.Target.X), float64(bullet.Target.Y), radRotate, gameMap)

	go SendMessage(Message{
		Event:         "FlyLaser",
		Bullet:        bullet,
		PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: startX, Y: startY, Millisecond: tickTime - int(deltaTime)},
		IDMap:         gameMap.Id,
		NeedCheckView: true,
	})
}

// функция которая заставляет лететь снаряды летящие по прямой
func FlyBullet(bullet *unit.Bullet, gameMap *_map.Map) {

	if bullet == nil {
		return
	}

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	radRotate := float64(bullet.Rotate) * math.Pi / 180

	// пуля летит по параболе, поэтому до половины пути она немного приподнимается по Z,
	// а после половины пути стремительно идет к 0, это сказывается на анимации фронта, но не на логике
	startDist := 0.0
	if bullet.Target.Type == "map" {
		startDist = game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
	} else {
		startDist = float64(bullet.MaxRange)
		bullet.Target.X = bullet.X + int(float64(bullet.MaxRange)*math.Cos(radRotate))
		bullet.Target.Y = bullet.Y + int(float64(bullet.MaxRange)*math.Sin(radRotate))
	}

	minDist := startDist

	SendMessage(Message{
		Event:         "FlyBullet",
		Bullet:        bullet,
		PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: tickTime},
		IDMap:         gameMap.Id,
		NeedCheckView: true},
	)
	time.Sleep(time.Duration(tickTime) * time.Millisecond)

	for {

		currentDist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		// высота пульки
		bullet.Z = 1 - ((1.0 / float64(startDist)) * (startDist - currentDist))

		stopX := realSpeed * math.Cos(radRotate) // идем по вектору движения выстрела
		stopY := realSpeed * math.Sin(radRotate)

		// deltaTime - время затрачено на проверку колизий, оно существенно поэтому надо учитывать
		percent, end, deltaTime := detailFlyBullet(bullet, float64(bullet.X)+stopX, float64(bullet.Y)+stopY, radRotate, gameMap)

		ms := (tickTime * percent) / 100

		go SendMessage(Message{
			Event:         "FlyBullet",
			Bullet:        bullet,
			PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: ms - int(deltaTime)},
			IDMap:         gameMap.Id,
			NeedCheckView: true},
		)

		minDist = currentDist
		time.Sleep(time.Duration(ms-int(deltaTime)) * time.Millisecond)

		if end || minDist < currentDist {
			// для отыгрыша анимации взрыва
			// TODO появление динамического обьекта кратера
			SendMessage(Message{
				Event:         "ExplosionBullet",
				Bullet:        bullet,
				IDMap:         gameMap.Id,
				NeedCheckView: true,
			})

			break
		}
	}
}

// самоводящиеся ракеты
func FlyChaseRocket() {
	// TODO
}

// артилерийские установки
func FlyArtillery() {
	// TODO
}

func detailFlyBullet(bullet *unit.Bullet, toX, toY, radRotate float64, gameMap *_map.Map) (int, bool, int64) {

	startTime := time.Now()

	startDist := game_math.GetBetweenDist(bullet.X, bullet.Y, int(toX), int(toY))
	minDist := startDist
	dist := startDist

	x, y := float64(bullet.X), float64(bullet.Y)

	for {
		percentPath := 100 - (int((dist * 100) / startDist))

		stopX := float64(10) * math.Cos(radRotate) // идем по вектору движения пуди
		stopY := float64(10) * math.Sin(radRotate)

		x += stopX
		y += stopY

		// колизии с обьектами, юнитами и геодатой по курсу движения
		// урон по ящикам и транспортам проходит только если стрелять по ним целенаправленно
		units := globalGame.Clients.GetAllShortUnits(gameMap.Id)
		delete(units, bullet.UnitID) // стреляющий не может быть колизие

		collision, typeCollision, id := collisions.CircleAllCollisionCheck(int(x), int(y), 3, gameMap, units, nil)
		if typeCollision != "" {
			println(typeCollision, id)
		}

		dist = game_math.GetBetweenDist(int(x), int(y), int(toX), int(toY))
		distToTarget := game_math.GetBetweenDist(int(x), int(y), bullet.Target.X, bullet.Target.Y)

		if dist <= 3 || minDist < dist || distToTarget < 12 || collision {

			bullet.X, bullet.Y = int(x), int(y)

			if collision {
				bullet.Target.X, bullet.Target.Y = int(x), int(y)
			}

			return percentPath, distToTarget < 12 || collision, time.Since(startTime).Nanoseconds() / int64(time.Millisecond)
		}

		minDist = dist
	}
}
