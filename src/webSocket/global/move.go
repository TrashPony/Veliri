package global

import (
	"github.com/TrashPony/Veliri/src/mechanics/db/squad/update"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
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
	// игнорировать юнитов в формирование при поиске пути т.к. они не могут пересечься

	if user.GetSquad() != nil && msg.UnitsID != nil {

		formationMove := false
		if FormationInit(user, msg.UnitsID) {
			formationMove = true
		}

		var toPos []*coordinate.Coordinate

		if len(msg.UnitsID) > 1 {
			toPos = move.GetUnitPos(msg.UnitsID, user, msg.ToX, msg.ToY)
		} else {
			toPos = make([]*coordinate.Coordinate, 0)
			toPos = append(toPos, &coordinate.Coordinate{X: int(msg.ToX), Y: int(msg.ToY)})
		}

		for i, unitID := range msg.UnitsID {

			moveUnit := user.GetSquad().GetUnitByID(unitID)

			moveUUID := uuid.NewV1().String()
			moveUnit.MoveUUID = moveUUID

			if !formationMove && !moveUnit.Body.MotherShip {
				moveUnit.Formation = false
			}

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
							float64(moveUnit.LastPathCell.Y), moveUnit.LastPathCell.Rotate, moveUUID, units, msg.UnitsID)
					} else {
						path, err = move.Unit(moveUnit, float64(toPos[i].X), float64(toPos[i].Y), float64(moveUnit.X),
							float64(moveUnit.Y), moveUnit.Rotate, moveUUID, units, msg.UnitsID)
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

func FormationInit(user *player.Player, unitsID []int) bool {
	for _, id := range unitsID {

		if user.GetSquad().MatherShip.ID == id && !user.GetSquad().MatherShip.Formation {
			user.GetSquad().MatherShip.Formation = true
			go FormationMove(user)
			return true
		} else {
			if user.GetSquad().MatherShip.Formation {
				return true
			}
		}

	}

	return false
}

func FormationMove(user *player.Player) {
	for {
		for _, unitSlot := range user.GetSquad().MatherShip.Units {

			if unitSlot.Unit != nil && unitSlot.Unit.OnMap && unitSlot.Unit.Formation {

				x, y := user.GetSquad().GetFormationCoordinate(unitSlot.Unit.FormationPos.X, unitSlot.Unit.FormationPos.Y)

				msg := Message{}
				msg.ToX, msg.ToY = float64(x), float64(y)
				msg.UnitsID = []int{unitSlot.Unit.ID}

				Move(user, msg, true)
			}
		}

		time.Sleep(100 * time.Millisecond)
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
		moveUnit.ActualPathCell = nil

		stopMove(moveUnit, false)

		moveUnit.LastPathCell = nil

		if moveUnit != nil {
			moveUnit.MoveChecker = false
		}

		if moveRepeat {
			msg.UnitsID = []int{moveUnit.ID}
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

		moveUnit.ActualPathCell = pathUnit

		newGravity := move.GetGravity(moveUnit.X, moveUnit.Y, user.GetSquad().MatherShip.MapID)
		if moveUnit.HighGravity != newGravity {
			moveUnit.HighGravity = newGravity
			go SendMessage(Message{Event: "ChangeGravity", IDUserSend: user.GetID(), ShortUnit: moveUnit.GetShortInfo(),
				IDMap: moveUnit.MapID, HighGravity: newGravity, Bot: user.Bot})
		}

		// колизии юнит - юнит
		noCollision, collisionUnit := collisionUnitToUnit(moveUnit, pathUnit, path, i)

		if !noCollision && collisionUnit != nil {

			timeCount := 0

			go SendMessage(Message{
				Event:     "MoveStop",
				ShortUnit: moveUnit.GetShortInfo(),
				PathUnit: &unit.PathUnit{
					Speed: 0,
				},
				IDMap: moveUnit.MapID,
			})

			for collisionUnit != nil && collisionUnit.MoveChecker && timeCount < 10 && !noCollision {

				noCollision, collisionUnit = collisionUnitToUnit(moveUnit, pathUnit, path, i)

				if collisionUnit != nil && !collisionUnit.MoveChecker {
					moveRepeat = true
					return
				}

				timeCount++
				time.Sleep(100 * time.Millisecond)
			}

			if collisionUnit != nil {
				moveRepeat = true
				return
			}

			//pathUnit.X, pathUnit.Y = x, y
			//go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID})
			//time.Sleep(time.Duration((pathUnit.Millisecond*percent)/100) * time.Millisecond)
			//
			//moveUnit.X, moveUnit.Y, moveUnit.Rotate = x, y, pathUnit.Rotate
			//
			//unitPath, toUnitPath := collisions.UnitToUnitCollisionReaction(moveUnit, globalGame.Clients.GetUnitByID(collisionUnit.ID))
			//
			//go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: unitPath, IDMap: moveUnit.MapID})
			//go SendMessage(Message{Event: "MoveTo", ShortUnit: collisionUnit, PathUnit: toUnitPath, IDMap: moveUnit.MapID})
			//time.Sleep(200 * time.Millisecond)
			//
			//moveRepeat = true
			//return
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
		mapBox, x, y, percent := collisions.BodyCheckCollisionBoxes(moveUnit, moveUnit.Body, pathUnit)
		if mapBox != nil {

			//доходим до куда можем
			pathUnit.X, pathUnit.Y, pathUnit.Millisecond = x, y, (pathUnit.Millisecond*percent)/100
			SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID})

			//// ждем события столкновения
			time.Sleep(time.Duration(pathUnit.Millisecond) * time.Millisecond)
			moveUnit.X, moveUnit.Y, moveUnit.Rotate = pathUnit.X, pathUnit.Y, pathUnit.Rotate

			// обрабатываем столкновение
			unitPos, boxPos := collisions.UnitToBoxCollisionReaction(moveUnit, mapBox)
			SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: unitPos, IDMap: moveUnit.MapID})
			SendMessage(Message{Event: "BoxTo", PathUnit: boxPos, IDMap: moveUnit.MapID, BoxID: mapBox.ID})

			// TODO тнимаение зп ящика и если 0 то уничтожать
			//go SendMessage(Message{Event: "DestroyBox", BoxID: mapBox.ID, IDMap: moveUnit.MapID})
			//boxes.Boxes.DestroyBox(mapBox)
			//moveUnit.CurrentSpeed -= float64(moveUnit.Speed)
			//moveRepeat = true
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
		moveUnit.WorkOutMovePower()
		go SendMessage(Message{Event: "WorkOutThorium", IDUserSend: user.GetID(), Unit: moveUnit,
			ThoriumSlots: moveUnit.Body.ThoriumSlots, IDMap: moveUnit.MapID, Bot: user.Bot})

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
			go update.Squad(user.GetSquad(), true)
		}
	}
}

func collisionUnitToUnit(moveUnit *unit.Unit, pathUnit *unit.PathUnit, path *[]*unit.PathUnit, i int) (bool, *unit.ShortUnitInfo) {
	if pathUnit.Speed == 0 {
		return true, nil
	}

	// колизии юнит - юнит
	noCollision, collisionUnit, _, _, _ := collisions.InitCheckCollision(moveUnit, pathUnit)

	if !noCollision {
		return noCollision, collisionUnit
	}

	if len(*path)-1 > i+1 && pathUnit.Speed > 0 {
		noCollision, collisionUnit, _, _, _ = collisions.InitCheckCollision(moveUnit, (*path)[i+1])
	}

	return noCollision, collisionUnit
}
