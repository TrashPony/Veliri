package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"math"
	"time"
)

func Move(user *player.Player, msg Message) {
	if user.GetSquad() != nil && msg.UnitsID != nil {
		for _, unitID := range msg.UnitsID {

			moveUnit := user.GetSquad().GetUnitByID(unitID)

			if moveUnit != nil {

				// обнуляем маршрут что бы игрок больше не двигался
				stopMove(moveUnit, false)

				mp, find := maps.Maps.GetByID(user.GetSquad().MapID)
				if find && user.InBaseID == 0 && !user.GetSquad().Evacuation {

					for user.GetSquad().MoveChecker {
						time.Sleep(10 * time.Millisecond) // без этого будет блокировка
						// Ожидаем пока не завершится текущая клетка хода
						// иначе будут рывки в игре из за того что пока путь просчитывается х у отряда будет
						// менятся и когда начнется движение то отряд телепортирует обратно
					}

					// todo переделать мове метод
					// todo если ходит сразу много юнитов изменять ToX и ToY так что бы они занимали центры гейсов
					path, err := globalGame.MoveUnit(moveUnit, msg.ToX, msg.ToY, mp)
					moveUnit.ActualPath = &path

					// todo паника когда игрок умер но его цикл движения не прекратился, из за чего происходят проверки тела которого уже нет
					go MoveGlobalUnit(msg, user, &path, moveUnit)

					if len(path) > 1 {
						moveUnit.ToX = float64(path[len(path)-1].X)
						moveUnit.ToY = float64(path[len(path)-1].Y)
					} else {
						moveUnit.ToX = float64(moveUnit.X)
						moveUnit.ToY = float64(moveUnit.Y)
					}

					if err != nil && len(path) == 0 {
						go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
					}
					go SendMessage(Message{Event: "PreviewPath", Path: path, IDUserSend: user.GetID(),
						IDMap: user.GetSquad().MapID, Bot: user.Bot, ShortUnit: moveUnit.GetShortInfo()})
				}
			}
		}
	}
}

func stopMove(userUnit *unit.Unit, resetSpeed bool) {
	if userUnit != nil {
		userUnit.ActualPath = nil // останавливаем прошлое движение
		if resetSpeed {
			userUnit.CurrentSpeed = 0
		}
	}
}

func MoveGlobalUnit(msg Message, user *player.Player, path *[]unit.PathUnit, moveUnit *unit.Unit) {
	moveRepeat := false
	user.GetSquad().MoveChecker = true

	defer func() {
		stopMove(moveUnit, false)
		if user.GetSquad() != nil {
			user.GetSquad().MoveChecker = false
		}
		if moveRepeat {
			Move(user, msg)
		}
	}()

	for i, pathUnit := range *path {
		// юнит или отряд умер или путь изменился
		if user.GetSquad() == nil || moveUnit == nil || moveUnit.HP <= 0 || moveUnit.ActualPath == nil || moveUnit.ActualPath != path {
			return
		}

		newGravity := globalGame.GetGravity(moveUnit.X, moveUnit.Y, user.GetSquad().MapID)
		if moveUnit.HighGravity != newGravity {
			moveUnit.HighGravity = newGravity
			go SendMessage(Message{Event: "ChangeGravity", IDUserSend: user.GetID(), ShortUnit: moveUnit.GetShortInfo(),
				IDMap: user.GetSquad().MapID, Bot: user.Bot})
			moveRepeat = true
			return
		}

		globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
		if moveUnit.Afterburner {
			SquadDamage(user, 1, moveUnit)
		}

		//// колизии игрок-игрок // todo юнит - юнит
		//noCollision, collisionUser := initCheckCollision(user, &pathUnit)
		//if !noCollision && collisionUser != nil {
		//
		//	playerToPlayerCollisionReaction(user, collisionUser)
		//	return
		//}

		if moveUnit.Body.MotherShip {
			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: user.GetSquad().MapID, Bot: user.Bot})
			}
		}

		// если на пути встречается ящик то мы его давим и падает скорость todo
		//mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, user.GetSquad().MapID, user.GetSquad().MatherShip.Body)
		//if mapBox != nil {
		//	go SendMessage(Message{Event: "DestroyBox", BoxID: mapBox.ID, IDMap: user.GetSquad().MapID})
		//	boxes.Boxes.DestroyBox(mapBox)
		//	user.GetSquad().CurrentSpeed -= float64(user.GetSquad().MatherShip.Body.Speed)
		//	moveRepeat = true
		//	return
		//}

		// если клиент отключился то останавливаем его
		if globalGame.Clients.GetById(user.GetID()) == nil {
			return
		}

		// todo
		//coor := globalGame.HandlerDetect(user)
		//if coor != nil && coor.HandlerOpen {
		//	HandlerParse(user, ws, coor)
		//	return
		//}

		// говорим юзеру как расходуется его топливо
		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), ShortUnit: moveUnit.GetShortInfo(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MapID, Bot: user.Bot})

		// оповещаем мир как двигается отряд
		go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: user.GetSquad().MapID})

		if i+1 != len(*path) { // бeз этого ифа канал будет ловить деад лок
			time.Sleep(100 * time.Millisecond)
			moveUnit.CurrentSpeed = pathUnit.Speed
		} else {
			moveUnit.CurrentSpeed = 0
		}

		moveUnit.Rotate = pathUnit.Rotate
		moveUnit.X = int(pathUnit.X)
		moveUnit.Y = int(pathUnit.Y)

		if ((pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != moveUnit.Q && pathUnit.R != moveUnit.R)) ||
			i+1 == len(*path) {

			moveUnit.Q = pathUnit.Q
			moveUnit.R = pathUnit.R

			if !user.Bot {
				go update.Squad(user.GetSquad(), false)
			}
		}
	}
}

func initCheckCollision(user *player.Player, pathUnit *unit.PathUnit) (bool, *player.Player) {
	// вынесено в отдельную функцию что бы можно было беспробленнмно сделать defer rLock.Unlock()
	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()
	return globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, users)
}

func playerToPlayerCollisionReaction(user, toUser *player.Player) {
	// задаем переменные массы шаров
	mass1 := user.GetSquad().MatherShip.Body.CapacitySize
	mass2 := toUser.GetSquad().MatherShip.Body.CapacitySize

	if user.GetSquad().MatherShip.CurrentSpeed < 2 {
		user.GetSquad().MatherShip.CurrentSpeed = 2
	}

	// задаем переменные скорости
	// расчет для первой машины
	radRotate1 := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel1 := float64(user.GetSquad().MatherShip.CurrentSpeed) * math.Cos(radRotate1) // идем по вектору движения корпуса
	yVel1 := float64(user.GetSquad().MatherShip.CurrentSpeed) * math.Sin(radRotate1)

	// расчет для второй машины
	radRotate2 := float64(toUser.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel2 := float64(toUser.GetSquad().MatherShip.CurrentSpeed) * math.Cos(radRotate2) // идем по вектору движения корпуса
	yVel2 := float64(toUser.GetSquad().MatherShip.CurrentSpeed) * math.Sin(radRotate2)

	//Угол между осью х и линией действия
	needRad := math.Atan2(float64(toUser.GetSquad().MatherShip.Y-user.GetSquad().MatherShip.Y), float64(toUser.GetSquad().MatherShip.X-user.GetSquad().MatherShip.X))
	cosAlfa := math.Cos(needRad)
	sinAlfa := math.Sin(needRad)

	// находим скорости вдоль линии действия
	xVel1prime := xVel1*cosAlfa + yVel1*sinAlfa
	xVel2prime := xVel2*cosAlfa + yVel2*sinAlfa

	// находим скорости перпендикулярные линии действия
	yVel1prime := yVel1*cosAlfa - xVel1*sinAlfa
	yVel2prime := yVel2*cosAlfa - xVel2*sinAlfa

	//// применяем законы сохранения
	P := float64(mass1)*xVel1prime + float64(mass2)*xVel2prime
	V := xVel1prime - xVel2prime
	v2f := (P + float64(mass1)*V) / (float64(mass1) + float64(mass2))
	v1f := v2f - xVel1prime + xVel2prime
	xVel1prime = v1f
	xVel2prime = v2f

	// Проецируем обратно на оси Х и У.
	xVel1 = xVel1prime*cosAlfa - yVel1prime*sinAlfa
	yVel1 = yVel1prime*cosAlfa + xVel1prime*sinAlfa

	xVel2 = xVel2prime*cosAlfa - yVel2prime*sinAlfa
	yVel2 = yVel2prime*cosAlfa + xVel2prime*sinAlfa

	speed1 := math.Sqrt((xVel1 * xVel1) + (yVel1 * yVel1))
	speed2 := math.Sqrt((xVel2 * xVel2) + (yVel2 * yVel2))

	user.GetSquad().MatherShip.CurrentSpeed = speed1
	user.GetSquad().MatherShip.X += int(float64(-speed1) * math.Cos(needRad))
	user.GetSquad().MatherShip.Y += int(float64(-speed1) * math.Sin(needRad))

	// проверка нового места толкаемого юзера на колизию в статичной карте
	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	possibleMove, _, _, _ := globalGame.CheckCollisionsOnStaticMap(
		int(toUser.GetSquad().MatherShip.X+int(float64(speed2)*math.Cos(needRad))),
		int(toUser.GetSquad().MatherShip.Y+int(float64(speed2)*math.Sin(needRad))),
		toUser.GetSquad().MatherShip.Rotate,
		mp,
		toUser.GetSquad().MatherShip.Body,
		false,
	)

	// проверка нового места толкаемого юзера на колизию с другими юзерами // TODO не отдебажено
	noCollision, _ := initCheckCollision(toUser, &unit.PathUnit{
		X:      int(toUser.GetSquad().MatherShip.X + int(float64(speed2)*math.Cos(needRad))),
		Y:      int(toUser.GetSquad().MatherShip.Y + int(float64(speed2)*math.Sin(needRad))),
		Rotate: toUser.GetSquad().MatherShip.Rotate,
	})

	if possibleMove && noCollision {
		toUser.GetSquad().MatherShip.X += int(float64(speed2) * math.Cos(needRad))
		toUser.GetSquad().MatherShip.Y += int(float64(speed2) * math.Sin(needRad))
	} else {
		// оталкиваем игрока инициализирующего столкновение иначе они застрянут
		user.GetSquad().MatherShip.X += int(float64(-speed2) * math.Cos(needRad))
		user.GetSquad().MatherShip.Y += int(float64(-speed2) * math.Sin(needRad))
	}

	userPath := unit.PathUnit{
		X:           user.GetSquad().MatherShip.X,
		Y:           user.GetSquad().MatherShip.Y,
		Rotate:      user.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       user.GetSquad().MatherShip.CurrentSpeed,
	}

	toUserPath := unit.PathUnit{
		X:           toUser.GetSquad().MatherShip.X,
		Y:           toUser.GetSquad().MatherShip.Y,
		Rotate:      toUser.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       speed2,
	}

	go SendMessage(Message{Event: "MoveTo", OtherUser: user.GetShortUserInfo(true), PathUnit: userPath, IDMap: user.GetSquad().MapID})
	go SendMessage(Message{Event: "MoveTo", OtherUser: toUser.GetShortUserInfo(true), PathUnit: toUserPath, IDMap: user.GetSquad().MapID})
	time.Sleep(200 * time.Millisecond)
}
