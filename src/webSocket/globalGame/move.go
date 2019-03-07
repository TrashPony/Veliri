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
	user := Clients.GetByWs(ws)

	if user != nil && user.GetSquad() != nil {
		mp, find := maps.Maps.GetByID(user.GetSquad().MapID)

		if find && user.InBaseID == 0 && !user.GetSquad().Evacuation {

			stopMove(ws, false)

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
			if err != nil {
				DisconnectUser(user)
			}

			go MoveUserMS(ws, msg, user, &path)
			user.GetSquad().MoveChecker = true
		}
	}
}

func stopMove(ws *websocket.Conn, reserSpeed bool) {
	user := Clients.GetByWs(ws)
	if user.GetSquad() != nil && user.GetSquad().MoveChecker && user.GetSquad().GetMove() != nil {
		user.GetSquad().GetMove() <- true // останавливаем прошлое движение
		if reserSpeed {
			user.GetSquad().CurrentSpeed = 0
		}
	}
}

func MoveUserMS(ws *websocket.Conn, msg Message, user *player.Player, path *[]squad.PathUnit) {

	moveRepeat := false

	defer func() {
		user.GetSquad().MoveChecker = false
		stopMove(ws, false)
		if moveRepeat {
			move(ws, msg)
		}
	}()

	for i, pathUnit := range *path {
		select {
		case exitNow := <-user.GetSquad().GetMove():
			if exitNow {
				return
			}
		default:

			if user.GetSquad().ActualPath != path {
				return
			}

			newGravity := globalGame.GetGravity(user.GetSquad().GlobalX, user.GetSquad().GlobalY, user.GetSquad().MapID)
			if user.GetSquad().HighGravity != newGravity {
				user.GetSquad().HighGravity = newGravity
				globalPipe <- Message{Event: "ChangeGravity", idUserSend: user.GetID(), Squad: user.GetSquad(), idMap: user.GetSquad().MapID, Bot: user.Bot}
				moveRepeat = true
				return
			}

			globalGame.WorkOutThorium(user.GetSquad().MatherShip.Body.ThoriumSlots, user.GetSquad().Afterburner, user.GetSquad().HighGravity)
			if user.GetSquad().Afterburner {
				// TODO ломание корпуса
			}

			// колизии игрок-игрок // TODO столкновения,  урон
			noCollision, collisionUser := globalGame.CheckCollisionsPlayers(user, pathUnit.X, pathUnit.Y, pathUnit.Rotate, user.GetSquad().MapID, Clients.GetAll())
			if !noCollision && collisionUser != nil {
				playerToPlayerCollisionReaction(user, collisionUser)
				moveRepeat = true
				return
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
				moveRepeat = true
				return
			}

			// если клиент отключился то останавливаем его
			if ws == nil || Clients.GetByWs(ws) == nil {
				return
			}

			coor := globalGame.HandlerDetect(user)
			if coor != nil && coor.HandlerOpen {
				HandlerParse(user, coor)
				return
			}

			// говорим юзеру как расходуется его топливо
			globalPipe <- Message{Event: "WorkOutThorium", idUserSend: user.GetID(),
				ThoriumSlots: user.GetSquad().MatherShip.Body.ThoriumSlots, idMap: user.GetSquad().MapID, Bot: user.Bot}

			// оповещаем мир как двигается отряд
			globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: pathUnit, idMap: user.GetSquad().MapID}

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
}

func playerToPlayerCollisionReaction(user, toUser *player.Player) {
	// задаем переменные массы шаров
	mass1 := 1
	mass2 := 1

	startDist := globalGame.GetBetweenDist(user.GetSquad().GlobalX, user.GetSquad().GlobalY, toUser.GetSquad().GlobalX, toUser.GetSquad().GlobalY)

	// задаем переменные скорости
	// расчет для первой машины
	radRotate := float64(user.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel1 := float64(user.GetSquad().CurrentSpeed) * math.Cos(radRotate) // идем по вектору движения корпуса
	yVel1 := float64(user.GetSquad().CurrentSpeed) * math.Sin(radRotate)

	// расчет для второй машины
	radRotate = float64(toUser.GetSquad().MatherShip.Rotate) * math.Pi / 180
	xVel2 := float64(toUser.GetSquad().CurrentSpeed) * math.Cos(radRotate) // идем по вектору движения корпуса
	yVel2 := float64(toUser.GetSquad().CurrentSpeed) * math.Sin(radRotate)

	run := user.GetSquad().GlobalX - toUser.GetSquad().GlobalX
	rise := user.GetSquad().GlobalY - toUser.GetSquad().GlobalY

	//Угол между осью х и линией действия
	Alfa := math.Atan2(float64(rise), float64(run))
	cosAlfa := math.Cos(Alfa)
	sinAlfa := math.Sin(Alfa)

	// находим скорости вдоль линии действия
	xVel1prime := float64(xVel1)*cosAlfa + float64(yVel1)*sinAlfa
	xVel2prime := float64(xVel2)*cosAlfa + float64(yVel2)*sinAlfa

	// находим скорости перпендикулярные линии действия
	yVel1prime := float64(yVel1)*cosAlfa - float64(xVel1)*sinAlfa
	yVel2prime := (yVel2)*cosAlfa - float64(xVel2)*sinAlfa

	// применяем законы сохранения
	P := float64(mass1)*xVel1prime + float64(mass2)*xVel2prime
	V := xVel1prime - xVel2prime
	v2f := (P + float64(mass1)*V) / (float64(mass1) + float64(mass2))
	v1f := v2f - xVel1prime + xVel2prime
	xVel1prime = v1f
	xVel2prime = v2f

	// Проецируем обратно на оси Х и У.
	xVel1 = xVel1prime*cosAlfa - yVel1prime*sinAlfa
	xVel2 = xVel2prime*cosAlfa - yVel2prime*sinAlfa
	yVel1 = yVel1prime*cosAlfa + xVel1prime*sinAlfa
	yVel2 = yVel2prime*cosAlfa + xVel2prime*sinAlfa

	// TODO проверка нового места на колизию (уперся в стенку, уперся в другово игрока)
	// TODO если игрок заденет другово игрока при повороте жопой с отрицательной скоростью то происходит обратное слипание

	endDist := globalGame.GetBetweenDist(user.GetSquad().GlobalX+int(xVel1), user.GetSquad().GlobalY+int(yVel1),
		toUser.GetSquad().GlobalX+int(xVel2), toUser.GetSquad().GlobalY+int(xVel2))

	//println(int(startDist), int(endDist))
	//println(startDist > endDist)

	if startDist > endDist {
		user.GetSquad().GlobalX -= int(xVel1)
		user.GetSquad().GlobalY -= int(yVel1)
		toUser.GetSquad().GlobalX -= int(xVel2)
		toUser.GetSquad().GlobalY -= int(yVel2)
	} else {
		user.GetSquad().GlobalX += int(xVel1)
		user.GetSquad().GlobalY += int(yVel1)
		toUser.GetSquad().GlobalX += int(xVel2)
		toUser.GetSquad().GlobalY += int(yVel2)
	}

	//меняем скорость у того кто врезался на минимальную
	user.GetSquad().CurrentSpeed = float64(user.GetSquad().MatherShip.Body.Speed)

	userPath := squad.PathUnit{
		X:           user.GetSquad().GlobalX,
		Y:           user.GetSquad().GlobalY,
		Rotate:      user.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       3,
	}

	toUserPath := squad.PathUnit{
		X:           toUser.GetSquad().GlobalX,
		Y:           toUser.GetSquad().GlobalY,
		Rotate:      toUser.GetSquad().MatherShip.Rotate,
		Millisecond: 200,
		Speed:       3,
	}

	globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(user), PathUnit: userPath, idMap: user.GetSquad().MapID}
	globalPipe <- Message{Event: "MoveTo", OtherUser: GetShortUserInfo(toUser), PathUnit: toUserPath, idMap: user.GetSquad().MapID}
	time.Sleep(100 * time.Millisecond)
}
