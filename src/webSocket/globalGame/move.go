package globalGame

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/squad"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
	"github.com/gorilla/websocket"
	"math"
	"time"
)

func Move(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	if user != nil && user.GetSquad() != nil {
		// обнуляем маршрут что бы игрок больше не двигался
		stopMove(user, false)

		mp, find := maps.Maps.GetByID(user.GetSquad().MapID)
		if find && user.InBaseID == 0 && !user.GetSquad().Evacuation {

			for user.GetSquad().MoveChecker {
				time.Sleep(10 * time.Millisecond) // без этого будет блокировка
				// Ожидаем пока не завершится текущая клетка хода
				// иначе будут рывки в игре из за того что пока путь просчитывается х у отряда будет
				// менятся и когда начнется движение то отряд телепортирует обратно
			}

			path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)
			user.GetSquad().ActualPath = &path

			go MoveUserMS(ws, msg, user, &path)

			if len(path) > 1 {
				user.GetSquad().ToX = float64(path[len(path)-1].X)
				user.GetSquad().ToY = float64(path[len(path)-1].Y)
			} else {
				user.GetSquad().ToX = float64(user.GetSquad().GlobalX)
				user.GetSquad().ToY = float64(user.GetSquad().GlobalY)
			}

			if err != nil && len(path) == 0 {
				go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
			}
			go SendMessage(Message{Event: "PreviewPath", Path: path, IDUserSend: user.GetID(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
		}
	}
}

func stopMove(user *player.Player, resetSpeed bool) {
	if user.GetSquad() != nil {
		user.GetSquad().ActualPath = nil // останавливаем прошлое движение
		if resetSpeed {
			user.GetSquad().CurrentSpeed = 0
		}
	}
}

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path *[]squad.PathUnit) {
	moveRepeat := false
	user.GetSquad().MoveChecker = true

	defer func() {
		stopMove(user, false)
		if user.GetSquad() != nil {
			user.GetSquad().MoveChecker = false
		}
		if moveRepeat {
			Move(ws, msg)
		}
	}()

	for i, pathUnit := range *path {
		// TODO приложение падает если во время движения он погибает
		if user.GetSquad() == nil || user.GetSquad().ActualPath == nil || user.GetSquad().ActualPath != path {
			return
		}

		newGravity := globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)
		if user.GetSquad().HighGravity != newGravity {
			user.GetSquad().HighGravity = newGravity
			go SendMessage(Message{Event: "ChangeGravity", IDUserSend: user.GetID(), Squad: user.GetSquad(), IDMap: user.GetSquad().MapID, Bot: user.Bot})
			moveRepeat = true
			return
		}

		globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)
		if user.GetSquad().Afterburner {
			SquadDamage(user, 1, ws)
		}

		// колизии игрок-игрок
		noCollision, collisionUser := initCheckCollision(user, &pathUnit)
		if !noCollision && collisionUser != nil {

			playerToPlayerCollisionReaction(user, collisionUser)
			return
		}

		// находим аномалии
		equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
		anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
		if err == nil {
			go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: user.GetSquad().MapID, Bot: user.Bot})
		}

		// если на пути встречается ящик то мы его давим и падает скорость
		mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, user.GetSquad().MapID, user.GetSquad().MatherShip.Body)
		if mapBox != nil {
			go SendMessage(Message{Event: "DestroyBox", BoxID: mapBox.ID, IDMap: user.GetSquad().MapID})
			boxes.Boxes.DestroyBox(mapBox)
			user.GetSquad().CurrentSpeed -= float64(user.GetSquad().MatherShip.Body.Speed)
			moveRepeat = true
			return
		}

		// если клиент отключился то останавливаем его
		if ws == nil || globalGame.Clients.GetByWs(ws) == nil {
			return
		}

		coor := globalGame.HandlerDetect(user)
		if coor != nil && coor.HandlerOpen {
			HandlerParse(user, ws, coor)
			return
		}

		// говорим юзеру как расходуется его топливо
		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(),
			ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, IDMap: user.GetSquad().MapID, Bot: user.Bot})

		// оповещаем мир как двигается отряд
		go SendMessage(Message{Event: "MoveTo", OtherUser: user.GetShortUserInfo(true), PathUnit: pathUnit, IDMap: user.GetSquad().MapID})

		if i+1 != len(*path) { // бeз этого ифа канал будет ловить деад лок
			time.Sleep(100 * time.Millisecond)
			user.GetSquad().CurrentSpeed = pathUnit.Speed
		} else {
			user.GetSquad().CurrentSpeed = 0
		}

		user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
		user.GetSquad().GlobalX = int(pathUnit.X)
		user.GetSquad().GlobalY = int(pathUnit.Y)

		if ((pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != user.GetSquad().Q && pathUnit.R != user.GetSquad().R)) ||
			i+1 == len(*path) {
			user.GetSquad().Q = pathUnit.Q
			user.GetSquad().R = pathUnit.R

			if !user.Bot {
				go update.Squad(user.GetSquad(), false)
			}
		}
	}
}

func initCheckCollision(user *player.Player, pathUnit *squad.PathUnit) (bool, *player.Player) {
	// вынесено в отдельную функцию что бы можно было беспробленнмно сделать defer rLock.Unlock()
	users, rLock := globalGame.Clients.GetAll()
	defer rLock.Unlock()
	return globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, users)
}

func playerToPlayerCollisionReaction(user, toUser *player.Player) {
	// задаем переменные массы шаров
	mass1 := user.GetSquad().MatherShip.Body.CapacitySize
	mass2 := toUser.GetSquad().MatherShip.Body.CapacitySize

	if user.GetSquad().CurrentSpeed < 2 {
		user.GetSquad().CurrentSpeed = 2
	}

	// задаем переменные скорости
	// расчет для первой машины
	radRotate1 := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel1 := float64(user.GetSquad().CurrentSpeed) * math.Cos(radRotate1) // идем по вектору движения корпуса
	yVel1 := float64(user.GetSquad().CurrentSpeed) * math.Sin(radRotate1)

	// расчет для второй машины
	radRotate2 := float64(toUser.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel2 := float64(toUser.GetSquad().CurrentSpeed) * math.Cos(radRotate2) // идем по вектору движения корпуса
	yVel2 := float64(toUser.GetSquad().CurrentSpeed) * math.Sin(radRotate2)

	//Угол между осью х и линией действия
	needRad := math.Atan2(float64(toUser.GetSquad().GlobalY-user.GetSquad().GlobalY), float64(toUser.GetSquad().GlobalX-user.GetSquad().GlobalX))
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

	user.GetSquad().CurrentSpeed = speed1
	user.GetSquad().GlobalX += int(float64(-speed1) * math.Cos(needRad))
	user.GetSquad().GlobalY += int(float64(-speed1) * math.Sin(needRad))

	// проверка нового места толкаемого юзера на колизию в статичной карте
	mp, _ := maps.Maps.GetByID(user.GetSquad().MapID)
	possibleMove, _, _, _ := globalGame.CheckCollisionsOnStaticMap(
		int(toUser.GetSquad().GlobalX+int(float64(speed2)*math.Cos(needRad))),
		int(toUser.GetSquad().GlobalY+int(float64(speed2)*math.Sin(needRad))),
		toUser.GetSquad().MatherShip.Rotate,
		mp,
		toUser.GetSquad().MatherShip.Body,
		false,
	)

	// проверка нового места толкаемого юзера на колизию с другими юзерами // TODO не отдебажено
	noCollision, _ := initCheckCollision(toUser, &squad.PathUnit{
		X:      int(toUser.GetSquad().GlobalX + int(float64(speed2)*math.Cos(needRad))),
		Y:      int(toUser.GetSquad().GlobalY + int(float64(speed2)*math.Sin(needRad))),
		Rotate: toUser.GetSquad().MatherShip.Rotate,
	})

	if possibleMove && noCollision {
		toUser.GetSquad().GlobalX += int(float64(speed2) * math.Cos(needRad))
		toUser.GetSquad().GlobalY += int(float64(speed2) * math.Sin(needRad))
	} else {
		// оталкиваем игрока инициализирующего столкновение иначе они застрянут
		user.GetSquad().GlobalX += int(float64(-speed2) * math.Cos(needRad))
		user.GetSquad().GlobalY += int(float64(-speed2) * math.Sin(needRad))
	}

	userPath := squad.PathUnit{
		X:           user.GetSquad().GlobalX,
		Y:           user.GetSquad().GlobalY,
		Rotate:      user.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       user.GetSquad().CurrentSpeed,
	}

	toUserPath := squad.PathUnit{
		X:           toUser.GetSquad().GlobalX,
		Y:           toUser.GetSquad().GlobalY,
		Rotate:      toUser.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       speed2,
	}

	go SendMessage(Message{Event: "MoveTo", OtherUser: user.GetShortUserInfo(true), PathUnit: userPath, IDMap: user.GetSquad().MapID})
	go SendMessage(Message{Event: "MoveTo", OtherUser: toUser.GetShortUserInfo(true), PathUnit: toUserPath, IDMap: user.GetSquad().MapID})
	time.Sleep(200 * time.Millisecond)
}
