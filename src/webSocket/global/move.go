package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/boxes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/detail"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"github.com/satori/go.uuid"
	"time"
)

func Move(user *player.Player, msg Message, newAction bool) {

	// TODO движение отряда
	// отряд двигается с минимальной скорость из всех юнитов
	// при старте движения все юниты сначало должны принять необходимое положение, угол
	// и когда все будут готовы то начинать двигатся в составе формирования

	if user.GetSquad() != nil && msg.UnitsID != nil {

		var toPos []*coordinate.Coordinate

		if len(msg.UnitsID) > 1 {
			// TODO toPos = move.GetUnitPos(msg.UnitsID, user.GetSquad().MatherShip.MapID, int(msg.ToX), int(msg.ToY))
		} else {
			toPos = make([]*coordinate.Coordinate, 0)
			toPos = append(toPos, &coordinate.Coordinate{X: int(msg.ToX), Y: int(msg.ToY)})
		}

		for i, unitID := range msg.UnitsID {

			moveUnit := user.GetSquad().GetUnitByID(unitID)

			moveUUID := uuid.NewV1().String()
			moveUnit.MoveUUID = moveUUID

			if moveUnit != nil && moveUnit.OnMap {

				if newAction {
					moveUnit.FollowUnitID = 0
					moveUnit.Return = false
				}

				// обнуляем маршрут что бы игрок больше не двигался
				stopMove(moveUnit, false)

				mp, find := maps.Maps.GetByID(moveUnit.MapID)
				if find && user.InBaseID == 0 && !moveUnit.Evacuation {

					for moveUnit.LastPathCell == nil && moveUnit.MoveChecker {
						// ждем пока юнит начнет инициализацию выхода из текущего метода движения если оно есть
						// пока он доходит последнюю клетку LastPathCell мы начинаем формировать путь
						// это уменьшаяет лаги на клиента
						time.Sleep(time.Millisecond)
					}

					var path []*unit.PathUnit
					var err error

					units := globalGame.Clients.GetAllShortUnits(moveUnit.MapID, true)

					if moveUnit.LastPathCell != nil {
						path, err = move.Unit(moveUnit, float64(toPos[i].X), float64(toPos[i].Y), float64(moveUnit.LastPathCell.X),
							float64(moveUnit.LastPathCell.Y), moveUnit.LastPathCell.Rotate, moveUUID, units)
					} else {
						path, err = move.Unit(moveUnit, float64(toPos[i].X), float64(toPos[i].Y), float64(moveUnit.X),
							float64(moveUnit.Y), moveUnit.Rotate, moveUUID, units)
					}

					if err != nil && len(path) == 0 {
						println(err.Error())
						go SendMessage(Message{Event: "Error", Error: err.Error(), IDUserSend: user.GetID(), IDMap: moveUnit.MapID, Bot: user.Bot})
					}

					for moveUnit.MoveChecker && moveUnit.OnMap {
						if moveUnit.MoveUUID != moveUUID {
							return
						}

						time.Sleep(time.Millisecond)
						// Ожидаем пока не завершится текущая клетка хода
						// иначе будут рывки в игре из за того что пока путь просчитывается х у отряда будет
						// менятся и когда начнется движение то отряд телепортирует обратно
					}

					if moveUnit.MoveUUID == moveUUID {
						go MoveGlobalUnit(msg, user, &path, moveUnit, mp)
						go FollowUnit(user, moveUnit, msg)

						go SendMessage(Message{Event: "PreviewPath", Path: path, IDUserSend: user.GetID(),
							IDMap: moveUnit.MapID, Bot: user.Bot, ShortUnit: moveUnit.GetShortInfo()})
					}
				}
			}
		}
	}
}

func stopMove(moveUnit *unit.Unit, resetSpeed bool) {
	if moveUnit != nil {
		moveUnit.ActualPath = nil // останавливаем прошлое движение

		if resetSpeed {
			moveUnit.CurrentSpeed = 0
		}
	}
}

func FollowUnit(user *player.Player, moveUnit *unit.Unit, msg Message) {
	// если юнит преследует другово юнита, то достаем его и мониторим его положение
	// если по какойто причине (столкновение, гравитация и тд) надо перестроить маршрут то сохраняем FollowUnitID
	// однако если сам игрок сгенерил событие движения то мы не сохраняем параметр FollowUnitID

	var followUnit *unit.Unit
	if moveUnit.FollowUnitID != 0 {
		followUnit = globalGame.Clients.GetUnitByID(moveUnit.FollowUnitID)
	} else {
		return
	}

	if followUnit != nil {
		for {

			if moveUnit.FollowUnitID == 0 || !moveUnit.OnMap || !followUnit.OnMap || moveUnit.MapID != followUnit.MapID {
				moveUnit.FollowUnitID = 0
				moveUnit.Return = false
				return
			}

			dist := game_math.GetBetweenDist(followUnit.X, followUnit.Y, int(moveUnit.X), int(moveUnit.Y))
			if dist < 90 {

				stopMove(moveUnit, true)

				if moveUnit.Return {
					go ReturnUnit(user, moveUnit)
					return
				}

				time.Sleep(100 * time.Millisecond)
				continue
			}

			dist = game_math.GetBetweenDist(followUnit.X, followUnit.Y, int(moveUnit.ToX), int(moveUnit.ToY))
			if dist > 90 || moveUnit.ActualPath == nil {
				msg.ToX = float64(followUnit.X)
				msg.ToY = float64(followUnit.Y)
				Move(user, msg, false)
				return
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
}

func MoveGlobalUnit(msg Message, user *player.Player, path *[]*unit.PathUnit, moveUnit *unit.Unit, mp *_map.Map) {
	moveRepeat := false
	moveUnit.MoveChecker = true

	defer func() {
		stopMove(moveUnit, false)

		moveUnit.LastPathCell = nil

		if moveUnit != nil {
			moveUnit.MoveChecker = false
		}

		if moveRepeat {
			Move(user, msg, false)
		}

		go SendMessage(Message{
			Event:     "MoveStop",
			ShortUnit: moveUnit.GetShortInfo(),
			PathUnit: &unit.PathUnit{
				Speed: moveUnit.CurrentSpeed,
			},
			IDMap: moveUnit.MapID,
		})
	}()

	countExit := 1
	moveUnit.ActualPath = path

	for i, pathUnit := range *path {
		// юнит или отряд умер
		if user.GetSquad() == nil || moveUnit == nil || !moveUnit.OnMap || moveUnit.HP <= 0 {
			return
		}

		if moveUnit.ActualPath == nil || moveUnit.ActualPath != path {
			// если актуальный путь сменился то выполняем еще 1 итерацию из старого пути дабы дать время сгенерить новый путь
			if countExit <= 0 {
				return
			}
			countExit--

			if moveUnit.LastPathCell == nil && len(*path)-1 >= i+countExit {
				moveUnit.LastPathCell = (*path)[i+countExit]
			} else {
				if len(*path)-1 < i+countExit {
					moveUnit.LastPathCell = (*path)[len(*path)-1]
				}
			}
		}

		newGravity := move.GetGravity(moveUnit.X, moveUnit.Y, user.GetSquad().MatherShip.MapID)
		if moveUnit.HighGravity != newGravity {
			moveUnit.HighGravity = newGravity
			go SendMessage(Message{Event: "ChangeGravity", IDUserSend: user.GetID(), ShortUnit: moveUnit.GetShortInfo(),
				IDMap: moveUnit.MapID, HighGravity: newGravity, Bot: user.Bot})
		}

		// колизии юнит - юнит
		noCollision, collisionUnit, x, y, percent := collisions.InitCheckCollision(moveUnit, pathUnit)
		if !noCollision && collisionUnit != nil {

			pathUnit.X, pathUnit.Y = x, y
			go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID})
			time.Sleep(time.Duration((pathUnit.Millisecond*percent)/100) * time.Millisecond)

			moveUnit.X, moveUnit.Y, moveUnit.Rotate = x, y, pathUnit.Rotate

			unitPath, toUnitPath := collisions.UnitToUnitCollisionReaction(moveUnit, globalGame.Clients.GetUnitByID(collisionUnit.ID))

			go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: unitPath, IDMap: moveUnit.MapID})
			go SendMessage(Message{Event: "MoveTo", ShortUnit: collisionUnit, PathUnit: toUnitPath, IDMap: moveUnit.MapID})
			time.Sleep(200 * time.Millisecond)

			moveRepeat = true
			return
		}

		//possibleMove, _, _, _ := collisions.CheckCollisionsOnStaticMap(pathUnit.X, pathUnit.Y, pathUnit.Rotate, mp, moveUnit.Body, false)
		//if !possibleMove {
		//	println("collision")
		//	return
		//}

		if moveUnit.Body.MotherShip && moveUnit.HP > 0 {
			// находим аномалии
			equipSlot := user.GetSquad().MatherShip.Body.FindApplicableEquip("geo_scan")
			anomalies, err := globalGame.GetVisibleAnomaly(user, equipSlot)
			if err == nil {
				go SendMessage(Message{Event: "AnomalySignal", IDUserSend: user.GetID(), Anomalies: anomalies, IDMap: moveUnit.MapID, Bot: user.Bot})
			}
		}

		// если на пути встречается ящик то мы его давим и падает скорость
		mapBox := collisions.BodyCheckCollisionBoxes(int(pathUnit.X), int(pathUnit.Y), pathUnit.Rotate, moveUnit.MapID, moveUnit.Body)
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

		// расход топлива
		// TODO если расход топлима изменился или топливо кончилось то останавливатся или перерасчитывать путь
		if !moveUnit.Body.MotherShip {

			fakeThoriumSlots := make(map[int]*detail.ThoriumSlot)
			fakeThoriumSlots[1] = &detail.ThoriumSlot{Number: 1, WorkedOut: float32(moveUnit.Power), Inversion: true, Count: 1}

			move.WorkOutThorium(fakeThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
			if moveUnit.Afterburner {
				SquadDamage(user, 1, moveUnit)
			}

			moveUnit.Power = int(fakeThoriumSlots[1].WorkedOut)

			go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), Unit: moveUnit,
				ThoriumSlots: fakeThoriumSlots, IDMap: moveUnit.MapID, Bot: user.Bot})
		} else {

			move.WorkOutThorium(moveUnit.Body.ThoriumSlots, moveUnit.Afterburner, moveUnit.HighGravity)
			if moveUnit.Afterburner {
				SquadDamage(user, 1, moveUnit)
			}
			go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), Unit: moveUnit,
				ThoriumSlots: moveUnit.Body.ThoriumSlots, IDMap: moveUnit.MapID, Bot: user.Bot})
		}

		// оповещаем мир как двигается отряд
		go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID})

		if i+1 != len(*path) { // бeз этого ифа канал будет ловить деад лок
			time.Sleep(time.Duration(pathUnit.Millisecond) * time.Millisecond)
			moveUnit.CurrentSpeed = pathUnit.Speed
		} else {
			moveUnit.CurrentSpeed = 0
		}

		moveUnit.Rotate = pathUnit.Rotate
		moveUnit.X = int(pathUnit.X)
		moveUnit.Y = int(pathUnit.Y)

		// todo оптимизировать а то каждый раз обновлять это ужасно
		if !user.Bot { // TODO апдейт юнитов
			go update.Squad(user.GetSquad(), false)
		}
	}
}
