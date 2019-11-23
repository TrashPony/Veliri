package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/attack"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"math"
	"time"
)

const tickTime = 100

// полет лазера
func FlyLaser(bullet *unit.Bullet, gameMap *_map.Map) {
	// лазер летит со скоростью света, поэтому что все нам надо это отдать стартовое ХУ и конечную ХУ
	// конечная ХУ это координата колизии или карты куда стреляет игрок

	startTime := time.Now()
	radRotate := float64(bullet.Rotate) * math.Pi / 180
	startX, startY := bullet.X, bullet.Y

	if bullet.Target.Type != "map" {
		bullet.Target.X = bullet.X + int(float64(bullet.MaxRange)*math.Cos(radRotate))
		bullet.Target.Y = bullet.Y + int(float64(bullet.MaxRange)*math.Sin(radRotate))
	}

	detailFlyBullet(bullet, float64(bullet.Target.X), float64(bullet.Target.Y), radRotate, gameMap, 10, nil, &startTime)
	go SendMessage(Message{
		Event:         "FlyLaser",
		Bullet:        bullet,
		PathUnit:      &unit.PathUnit{Rotate: int(bullet.Rotate), X: startX, Y: startY, Millisecond: 50},
		IDMap:         gameMap.Id,
		NeedCheckView: true,
	})
}

// функция которая заставляет лететь снаряды летящие по прямой
func FlyBullet(user *player.Player, bullet *unit.Bullet, gameMap *_map.Map) {

	if bullet == nil {
		return
	}

	realSpeed := float64(bullet.Speed / (1000 / tickTime))
	startSpeed := float64(bullet.Speed / (1000 / tickTime))
	radRotate := float64(bullet.Rotate) * math.Pi / 180

	startDist := 0.0

	if bullet.Target.Type == "map" || bullet.Artillery {
		// если цель карта или стреляет арта то юнит целинаправленно стреляет в х,у и снаряд не может пролететь мимо
		// так же меняется максимальное растояние выстрела по нему будет считаться путь)
		startDist = game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)
		bullet.MaxRange = int(startDist)
	} else {
		// в остальных случаях если это танк ган, лазер или ракеты то они могут пролететь мимо и останоятся только на макс дальности выстрела
		startDist = float64(bullet.MaxRange)
		bullet.Target.X = bullet.X + int(float64(bullet.MaxRange)*math.Cos(radRotate))
		bullet.Target.Y = bullet.Y + int(float64(bullet.MaxRange)*math.Sin(radRotate))
	}

	// первое появление пули из ствола
	sendFlyBullet("FlyBullet", bullet, gameMap.Id, tickTime) // тут еще и слип

	// пройденное растояние, что бы наприм ракеты не могли пролететь больше своей дальности
	distanceTraveled := 0.0

	for {
		startTime := time.Now()

		if bullet.Ammo.ChaseTarget && attack.GetXYTarget2(user, bullet.Target, gameMap) {
			// самоводящиеся ракеты преследуют цель
			needRotate := game_math.GetBetweenAngle(float64(bullet.Target.X), float64(bullet.Target.Y), float64(bullet.X), float64(bullet.Y))
			move.RotateUnit(&bullet.Rotate, &needRotate, 3)
			radRotate = float64(bullet.Rotate) * math.Pi / 180
		}

		currentDist := game_math.GetBetweenDist(bullet.X, bullet.Y, bullet.Target.X, bullet.Target.Y)

		realSpeed = attack.GetZAndSpeedBullet(bullet, startDist, currentDist, startSpeed)

		// идем по вектору движения выстрела
		stopX := realSpeed * math.Cos(radRotate)
		stopY := realSpeed * math.Sin(radRotate)

		// deltaTime - время затрачено на проверку колизий, оно существенно поэтому надо учитывать
		percent, end, deltaTime := detailFlyBullet(bullet, float64(bullet.X)+stopX, float64(bullet.Y)+stopY, radRotate, gameMap, realSpeed, &distanceTraveled, &startTime)

		ms := tickTime
		if end {
			ms = (tickTime * percent) / 100
		}

		sendFlyBullet("FlyBullet", bullet, gameMap.Id, ms-int(deltaTime)) // тут еще и слип

		if end {
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

// артилерийские ракеты
func FlyArtilleryRocket(bullet *unit.Bullet, gameMap *_map.Map) {
	// TODO
}

func detailFlyBullet(bullet *unit.Bullet, toX, toY, radRotate float64, gameMap *_map.Map, speed float64,
	distanceTraveled *float64, startTime *time.Time) (int, bool, int64) {

	startDist := game_math.GetBetweenDist(bullet.X, bullet.Y, int(toX), int(toY))
	minDist := startDist
	dist := startDist

	x, y := float64(bullet.X), float64(bullet.Y)

	if speed > 10 {
		speed = speed / 10
	}

	for {

		if distanceTraveled != nil {
			*distanceTraveled += speed
		}

		percentPath := 100 - (int((dist * 100) / startDist))

		x += speed * math.Cos(radRotate) // идем по вектору движения пуди
		y += speed * math.Sin(radRotate)

		dist = game_math.GetBetweenDist(int(x), int(y), int(toX), int(toY))

		var collision bool
		var typeCollision string
		var id int

		if !bullet.Artillery || bullet.Z <= 1.1 {
			// колизии с обьектами, юнитами и геодатой по курсу движения
			// проверяем колизии ток для прямо летящих снарядов или низко летящей арты

			units := globalGame.Clients.GetAllShortUnits(gameMap.Id)
			delete(units, bullet.UnitID) // стреляющий не может быть в колизии

			// урон по ящикам и транспортам проходит только если стрелять по ним целенаправленно
			collision, typeCollision, id = collisions.CircleAllCollisionCheck(int(x), int(y), 2, gameMap, units, nil)

			if typeCollision != "" {
				// todo обработка колизий
				println(typeCollision, id)
			}
		}

		end := checkEndPath(bullet, distanceTraveled) || collision

		if dist <= 3 || minDist < dist || end {

			bullet.X, bullet.Y = int(x), int(y)

			if collision {
				bullet.Target.X, bullet.Target.Y = int(x), int(y)
			}

			return percentPath, end, time.Since(*startTime).Nanoseconds() / int64(time.Millisecond)
		}

		minDist = dist
	}
}

func checkEndPath(bullet *unit.Bullet, distanceTraveled *float64) bool {

	if distanceTraveled == nil {
		return false
	}

	// конец пути не самоводящегося снаряда кончается когда он достигает цель
	return *distanceTraveled >= float64(bullet.MaxRange)
}

func sendFlyBullet(event string, bullet *unit.Bullet, mapID int, ms int) {
	go SendMessage(Message{
		Event:         event,
		Bullet:        bullet,
		PathUnit:      &unit.PathUnit{Rotate: bullet.Rotate, X: bullet.X, Y: bullet.Y, Millisecond: ms},
		IDMap:         mapID,
		NeedCheckView: true},
	)

	time.Sleep(time.Duration(ms) * time.Millisecond)
}
