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
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/move"
	"github.com/satori/go.uuid"
	"time"
)

func Move(user *player.Player, msg Message, newAction bool) {
	// TODO рефакторинг
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

					target := moveUnit.GetTarget()
					if target != nil {
						if target.Type == "map" {
							moveUnit.SetTarget(nil)
						} else {
							moveUnit.SetFollowTarget(false)
						}
					}
				}

				// обнуляем маршрут что бы игрок больше не двигался
				move.StopMove(moveUnit, false)

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

					units := globalGame.Clients.GetAllShortUnits(moveUnit.MapID)

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

func MoveGlobalUnit(msg Message, user *player.Player, path *[]*unit.PathUnit, moveUnit *unit.Unit, mp *_map.Map) {
	moveRepeat := false
	moveUnit.MoveChecker = true

	defer func() {
		// без обертки не сработает moveRepeat
		MoveGlobalUnitEnd(user, moveUnit, moveRepeat, msg)
	}()

	countExit := 1
	moveUnit.ActualPath = path

	// если юнит вошел в метод с целью и флагов приследования, то это движения для атаки
	followTarget := false
	if moveUnit.GetTarget() != nil && moveUnit.GetTarget().Follow {
		followTarget = true
	}

	for i, pathUnit := range *path {

		startTime := time.Now()

		if move.CheckExit(user, moveUnit, path, &i, &countExit, followTarget) {
			return
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
				IDMap:         moveUnit.MapID,
				NeedCheckView: true,
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
			SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID, NeedCheckView: true})

			//// ждем события столкновения
			time.Sleep(time.Duration(pathUnit.Millisecond) * time.Millisecond)
			moveUnit.X, moveUnit.Y, moveUnit.Rotate = pathUnit.X, pathUnit.Y, pathUnit.Rotate

			// обрабатываем столкновение
			unitPos, boxPos := collisions.UnitToBoxCollisionReaction(moveUnit, mapBox)
			SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: unitPos, IDMap: moveUnit.MapID, NeedCheckView: true})
			SendMessage(Message{Event: "BoxTo", PathUnit: boxPos, IDMap: moveUnit.MapID, BoxID: mapBox.ID, NeedCheckView: true})

			// TODO отнимаение хп ящика и если 0 то уничтожать
			//go SendMessage(Message{Event: "DestroyBox", BoxID: mapBox.ID, IDMap: moveUnit.MapID})
			//boxes.Boxes.DestroyBox(mapBox)
			//moveUnit.CurrentSpeed -= float64(moveUnit.Speed)
			//moveRepeat = true
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
		go SendMessage(Message{Event: "MoveTo", ShortUnit: moveUnit.GetShortInfo(), PathUnit: pathUnit, IDMap: moveUnit.MapID, NeedCheckView: true})
		move.SetPosition(moveUnit, pathUnit, time.Since(startTime).Nanoseconds()/int64(time.Millisecond))

		if i+1 != len(*path) { // бeз этого ифа канал будет ловить деад лок
			moveUnit.CurrentSpeed = pathUnit.Speed
		} else {
			moveUnit.CurrentSpeed = 0
		}

		// todo оптимизировать а то каждый раз обновлять это ужасно
		if !user.Bot { // TODO апдейт юнитов
			go update.Squad(user.GetSquad(), true)
		}
	}
}

func MoveGlobalUnitEnd(user *player.Player, moveUnit *unit.Unit, moveRepeat bool, msg Message) {
	if moveUnit != nil {

		moveUnit.ActualPathCell = nil
		move.StopMove(moveUnit, false)
		moveUnit.LastPathCell = nil

		if moveRepeat {
			msg.UnitsID = []int{moveUnit.ID}
			go Move(user, msg, false)
		} else {
			moveUnit.CurrentSpeed = 0
		}

		go SendMessage(Message{
			Event:     "MoveStop",
			ShortUnit: moveUnit.GetShortInfo(),
			PathUnit: &unit.PathUnit{
				Speed: moveUnit.CurrentSpeed,
			},
			IDMap: moveUnit.MapID,
		})

		moveUnit.MoveChecker = false
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
