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

func move(ws *websocket.Conn, msg Message) {
	user := globalGame.Clients.GetByWs(ws)

	if user != nil && user.GetSquad() != nil {
		// обнуляем маршрут что бы игрок больше не двигался
		stopMove(ws, false)

		mp, find := maps.Maps.GetByID(user.GetSquad().MapID)
		if find && user.InBaseID == 0 && !user.GetSquad().Evacuation {

			path, err := globalGame.MoveSquad(user, msg.ToX, msg.ToY, mp)
			user.GetSquad().ActualPath = &path

			if len(path) > 1 {
				user.GetSquad().ToX = float64(path[len(path)-1].X)
				user.GetSquad().ToY = float64(path[len(path)-1].Y)
			} else {
				user.GetSquad().ToX = float64(user.GetSquad().GlobalX)
				user.GetSquad().ToY = float64(user.GetSquad().GlobalY)
			}

			if err != nil && len(path) == 0 {
				globalPipe <- Message{Event: "Error", Error: err.Error(), idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
			}
			globalPipe <- Message{Event: "PreviewPath", Path: path, idUserSend: user.GetID(), idMap: user.GetSquad().MapID, Bot: user.Bot}
		}
	}
}

func stopMove(ws *websocket.Conn, reserSpeed bool) {
	user := globalGame.Clients.GetByWs(ws)
	if user.GetSquad() != nil && user.GetSquad().MoveChecker {
		user.GetSquad().MoveChecker = false
		user.GetSquad().ActualPath = nil // останавливаем прошлое движение
		if reserSpeed {
			user.GetSquad().CurrentSpeed = 0
		}
	}
}

func MoveUserMS(ws *websocket.Conn, user *player.Player) {
	for {
		user.GetSquad().MoveChecker = false

		for user.GetSquad().ActualPath == nil {
			// ожидает пока появится актуальный путь
			time.Sleep(100 * time.Millisecond)
		}

		// если мы получили путь то сохраняем на него сылку что бы мы могли всегда сравнить его с аутальным
		currentPath := user.GetSquad().ActualPath

		for i, pathUnit := range *currentPath {
			user.GetSquad().MoveChecker = true

			if (*currentPath)[i].Traversed {
				continue
			}
			(*currentPath)[i].Traversed = true

			if user.GetSquad().ActualPath == nil || user.GetSquad().ActualPath != currentPath {
				break
			}

			newGravity := globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)
			if user.GetSquad().HighGravity != newGravity {
				user.GetSquad().HighGravity = newGravity
				globalPipe <- Message{Event: "ChangeGravity", idUserSend: user.GetID(), Squad: user.GetSquad(), idMap: user.GetSquad().MapID, Bot: user.Bot}
				go move(ws, Message{ToX: user.GetSquad().ToX, ToY: user.GetSquad().ToY})
				break
			}

			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)
			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			// колизии игрок-игрок // TODO столкновения,  урон
			noCollision, collisionUser := globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, globalGame.Clients.GetAll())
			if !noCollision && collisionUser != nil {
				playerToPlayerCollisionReaction(user, collisionUser)
				if !user.Bot {
					go move(ws, Message{ToX: user.GetSquad().ToX, ToY: user.GetSquad().ToY})
				}
				break
			}

			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				globalPipe <- Message{Event: "AnomalySignal", idUserSend: user.GetID(), Anomalies: anomalies, idMap: user.GetSquad().MapID, Bot: user.Bot}
			}

			// если на пути встречается ящик то мы его давим и падает скорость
			mapBox := globalGame.CheckCollisionsBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, user.GetSquad().MapID, user.GetSquad().MatherShip.Body)
			if mapBox != nil {
				globalPipe <- Message{Event: "DestroyBox", BoxID: mapBox.ID, idMap: user.GetSquad().MapID}
				boxes.Boxes.DestroyBox(mapBox)
				user.GetSquad().CurrentSpeed -= float64(user.GetSquad().MatherShip.Body.Speed)
				go move(ws, Message{ToX: user.GetSquad().ToX, ToY: user.GetSquad().ToY})
				break
			}

			// если клиент отключился то останавливаем его
			if ws == nil || globalGame.Clients.GetByWs(ws) == nil {
				break
			}

			coor := globalGame.HandlerDetect(user)
			if coor != nil && coor.HandlerOpen {
				HandlerParse(user, coor)
				break
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID, Bot: user.Bot}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit, idMap: user.GetSquad().MapID}

			time.Sleep(100 * time.Millisecond)
			user.GetSquad().CurrentSpeed = pathUnit.Speed

			user.GetSquad().MatherShip.Rotate = pathUnit.Rotate
			user.GetSquad().GlobalX = int(pathUnit.X)
			user.GetSquad().GlobalY = int(pathUnit.Y)

			// помечаем что ячейка пройдена

			if (pathUnit.Q != 0 && pathUnit.R != 0) && (pathUnit.Q != user.GetSquad().Q && pathUnit.R != user.GetSquad().R) {
				user.GetSquad().Q = pathUnit.Q
				user.GetSquad().R = pathUnit.R

				if !user.Bot {
					go update.Squad(user.GetSquad(), false)
				}
			}
		}
	}
}

func playerToPlayerCollisionReaction(user, toUser *player.Player) {
	// задаем переменные массы шаров
	mass1 := 1
	mass2 := 1

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

	// TODO проверка нового места на колизию (уперся в стенку, уперся в другово игрока)
	// TODO если нет разгона то надо отьехать назад
	speed1 := math.Sqrt((xVel1 * xVel1) + (yVel1 * yVel1))
	speed2 := math.Sqrt((xVel2 * xVel2) + (yVel2 * yVel2))

	if user.GetSquad().CurrentSpeed > float64(user.GetSquad().MatherShip.Speed)*1.2 {
		user.GetSquad().CurrentSpeed = speed1
		user.GetSquad().GlobalX += int(float64(-speed1) * math.Cos(needRad))
		user.GetSquad().GlobalY += int(float64(-speed1) * math.Sin(needRad))
	} else {
		// если скорость пинимальна то мы тащим а не оталкиваемся
		// todo всеравно остаются рывки
	}

	toUser.GetSquad().GlobalX += int(float64(speed2) * math.Cos(needRad))
	toUser.GetSquad().GlobalY += int(float64(speed2) * math.Sin(needRad))

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

	globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: userPath, idMap: user.GetSquad().MapID}
	globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(toUser), PathUnit: toUserPath, idMap: user.GetSquad().MapID}
	time.Sleep(200 * time.Millisecond)
}
