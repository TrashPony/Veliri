package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"math"
	"time"
)

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

	if bullet == nil {
		return
	}

	tickTime := 100

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	radRotate := float64(bullet.Rotate) * math.Pi / 180

	// пуля летит по параболе, поэтому до половины пути она немного приподнимается по Z,
	// а после половины пути стремительно идет к 0, это сказывается на анимации фронта, но не на логике
	startDist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
	minDist := startDist

	SendMessage(Message{
		Event:         "FlyBullet",
		Bullet:        bullet,
		PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: tickTime},
		IDMap:         idMap,
		NeedCheckView: true},
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
			Event:         "FlyBullet",
			Bullet:        bullet,
			PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: ms},
			IDMap:         idMap,
			NeedCheckView: true},
		)

		if end || minDist < currentDist {
			// для отыгрыша анимации взрыва
			// TODO появление динамического обьекта кратера
			SendMessage(Message{
				Event:         "ExplosionBullet",
				Bullet:        bullet,
				IDMap:         idMap,
				NeedCheckView: true,
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

		// TODO колизии с обьектами, юнитами и геодатой по курсу движения
		// урон по ящикам и транспортам проходит только если стрелять по ним целенаправленно

		dist = game_math.GetBetweenDist(int(x), int(y), int(toX), int(toY))
		if dist <= 3 || minDist < dist {
			bullet.X, bullet.Y = int(x), int(y)
			return percentPath, 25 > game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		}

		minDist = dist
	}
}
