package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
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

				mp, find := maps.Maps.GetByID(moveUnit.MapID)
				if find && user.InBaseID == 0 && !moveUnit.Evacuation {

					for moveUnit.MoveChecker { // TODO чекер для юнита
						time.Sleep(10 * time.Millisecond) // без этого будет блокировка
						// Ожидаем пока не завершится текущая клетка хода
						// иначе будут рывки в игре из за того что пока путь просчитывается х у отряда будет
						// менятся и когда начнется движение то отряд телепортирует обратно
					}

					// todo если ходит сразу много юнитов изменять ToX и ToY так что бы они занимали центры гейсов
					path, err := globalGame.MoveUnit(moveUnit, msg.ToX, msg.ToY, mp)
					moveUnit.ActualPath = &path

					go MoveGlobalUnit(msg, user, &path, moveUnit)

					if len(path) > 1 {
						moveUnit.ToX = float64(path[len(path)-1].X)
						moveUnit.ToY = float64(path[len(path)-1].Y)
					} else {
						moveUnit.ToX = float64(moveUnit.X)
						moveUnit.ToY = float64(moveUnit.Y)
					}

					if err != nil && len(path) == 0 {
						go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: moveUnit.MapID, Bot: user.Bot})
					}
					go SendMessage(Message{Event: "PreviewPath", Path: path, IDUserSend: user.GetID(),
						IDMap: moveUnit.MapID, Bot: user.Bot, ShortUnit: moveUnit.GetShortInfo()})
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

		go SendMessage(Message{
			Event:     "MoveTo",
			ShortUnit: userUnit.GetShortInfo(),
			PathUnit: unit.PathUnit{
				X:           userUnit.X,
				Y:           userUnit.Y,
				Rotate:      userUnit.Rotate,
				Millisecond: 100,
				Speed:       userUnit.CurrentSpeed,
			},
			IDMap: userUnit.MapID,
		})
	}
}

func MoveGlobalUnit(msg Message, user *player.Player, path *[]unit.PathUnit, moveUnit *unit.Unit) {
	moveRepeat := false
	moveUnit.MoveChecker = true

	defer func() {
		stopMove(moveUnit, false)
		if user.GetSquad() != nil {
			moveUnit.MoveChecker = false
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

		newGravity := globalGame.GetGravity(moveUnit.X, moveUnit.Y, user.GetSquad().MatherShip.MapID)
		if moveUnit.HighGravity != newGravity {
			moveUnit.HighGravity = newGravity
			go SendMessage(Message{Event: "ChangeGravity", IDUserSend: user.GetID(), ShortUnit: moveUnit.GetShortInfo(),
				IDMap: moveUnit.MapID, Bot: user.Bot})
			moveRepeat = true
			return
		}

		globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
		if moveUnit.Afterburner {
			SquadDamage(user, 1, moveUnit)
		}

		// колизии юнит - юнит
		noCollision, collisionUnit := initCheckCollision(moveUnit, &pathUnit)
		if !noCollision && collisionUnit != nil {

			playerToPlayerCollisionReaction(moveUnit, globalGame.Clients.GetUnitByID(collisionUnit.ID))
			return
		}

		if moveUnit.Body.MotherShip && moveUnit.HP > 0 {
			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: moveUnit.MapID, Bot: user.Bot})
			}
		}

		// если на пути встречается ящик то мы его давим и падает скорость
		mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, moveUnit.MapID, moveUnit.Body)
		if mapBox != nil {
			go SendMessage(Message{Event: "DestroyBox", BoxID: mapBox.ID, IDMap: moveUnit.MapID})
			boxes.Boxes.DestroyBox(mapBox)
			moveUnit.CurrentSpeed -= float64(moveUnit.Speed)
			moveRepeat = true
			return
		}

		// если клиент отключился то останавливаем его
		if globalGame.Clients.GetById(user.GetID()) == nil {
			return
		}

		// входит на базы и телепортироваться в другие сектор могут ток мп (мобильные платформы)
		coor := globalGame.HandlerDetect(moveUnit)
		if moveUnit.Body.MotherShip && coor != nil && coor.HandlerOpen {
			HandlerParse(user, coor)
			return
		}

		// говорим юзеру как расходуется его топливо
		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), Unit: moveUnit,
			ThoriumSlots: moveUnit.Body.ThoriumSlots, IDMap: moveUnit.MapID, Bot: user.Bot})

		// оповещаем мир как двигается отряд
		go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID})

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

func initCheckCollision(moveUnit *unit.Unit, pathUnit *unit.PathUnit) (bool, *unit.ShortUnitInfo) {
	// вынесено в отдельную функцию что бы можно было беспробленнмно сделать defer rLock.Unlock()
	units := globalGame.Clients.GetAllShortUnits(moveUnit.MapID)
	return globalGame.CheckCollisionsPlayers(moveUnit, pathUnit.X, pathUnit.Y, pathUnit.Rotate, units)
}

func playerToPlayerCollisionReaction(takeUnit, toUnit *unit.Unit) {
	// задаем переменные массы шаров
	mass1 := takeUnit.Body.CapacitySize
	mass2 := toUnit.Body.CapacitySize

	if takeUnit.CurrentSpeed < 2 {
		takeUnit.CurrentSpeed = 2
	}

	// задаем переменные скорости
	// расчет для первой машины
	radRotate1 := float64(takeUnit.Rotate) * math.Pi / 180
	xVel1 := float64(takeUnit.CurrentSpeed) * math.Cos(radRotate1) // идем по вектору движения корпуса
	yVel1 := float64(takeUnit.CurrentSpeed) * math.Sin(radRotate1)

	// расчет для второй машины
	radRotate2 := float64(toUnit.Rotate) * math.Pi / 180
	xVel2 := float64(toUnit.CurrentSpeed) * math.Cos(radRotate2) // идем по вектору движения корпуса
	yVel2 := float64(toUnit.CurrentSpeed) * math.Sin(radRotate2)

	//Угол между осью х и линией действия
	needRad := math.Atan2(float64(toUnit.Y-takeUnit.Y), float64(toUnit.X-takeUnit.X))
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

	takeUnit.CurrentSpeed = speed1
	takeUnit.X += int(float64(-speed1) * math.Cos(needRad))
	takeUnit.Y += int(float64(-speed1) * math.Sin(needRad))

	// проверка нового места толкаемого юзера на колизию в статичной карте
	mp, _ := maps.Maps.GetByID(takeUnit.MapID)

	possibleMove, _, _, _ := globalGame.CheckCollisionsOnStaticMap(
		int(toUnit.X+int(float64(speed2)*math.Cos(needRad))),
		int(toUnit.Y+int(float64(speed2)*math.Sin(needRad))),
		toUnit.Rotate,
		mp,
		toUnit.Body,
		false,
	)

	// проверка нового места толкаемого юзера на колизию с другими юзерами // TODO не отдебажено
	noCollision, _ := initCheckCollision(toUnit, &unit.PathUnit{
		X:      int(toUnit.X + int(float64(speed2)*math.Cos(needRad))),
		Y:      int(toUnit.Y + int(float64(speed2)*math.Sin(needRad))),
		Rotate: toUnit.Rotate,
	})

	if possibleMove && noCollision {
		toUnit.X += int(float64(speed2) * math.Cos(needRad))
		toUnit.Y += int(float64(speed2) * math.Sin(needRad))
	} else {
		// оталкиваем игрока инициализирующего столкновение иначе они застрянут
		takeUnit.X += int(float64(-speed2) * math.Cos(needRad))
		takeUnit.Y += int(float64(-speed2) * math.Sin(needRad))
	}

	userPath := unit.PathUnit{
		X:           takeUnit.X,
		Y:           takeUnit.Y,
		Rotate:      takeUnit.Rotate,
		Millisecond: 200,
		Speed:       takeUnit.CurrentSpeed,
	}

	toUserPath := unit.PathUnit{
		X:           toUnit.X,
		Y:           toUnit.Y,
		Rotate:      toUnit.Rotate,
		Millisecond: 200,
		Speed:       speed2,
	}

	go SendMessage(Message{Event: "MoveTo", ShortUnit: takeUnit.GetShortInfo(), PathUnit: userPath, IDMap: takeUnit.MapID})
	go SendMessage(Message{Event: "MoveTo", ShortUnit: toUnit.GetShortInfo(), PathUnit: toUserPath, IDMap: toUnit.MapID})
	time.Sleep(200 * time.Millisecond)
}
