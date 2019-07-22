package targetPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/hexLineDraw"
	"strconv"
)

func filter(targetUnit *unit.Unit, coordinates []*coordinate.Coordinate, game *localGame.Game, artillery bool, client *player.Player) (targets map[string]*coordinate.Coordinate) {

	targets = make(map[string]*coordinate.Coordinate)

	targetUnitCoordinate, _ := game.GetMap().GetCoordinate(targetUnit.GetQ(), targetUnit.GetR())

	targets[strconv.Itoa(targetUnitCoordinate.Q)+":"+strconv.Itoa(targetUnitCoordinate.R)] = targetUnitCoordinate

	for _, gameCoordinate := range coordinates {
		watchCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.Q, gameCoordinate.R)
		if find {
			pathLine := hexLineDraw.Draw(targetUnitCoordinate, watchCoordinate, game)

			// до !каждой! координаты мы строим линию, если мы достигаем по этой линии координату то добавляем ее если нет то не добавляем
			for i, pathCell := range pathLine {
				pastCoordinate := &coordinate.Coordinate{} // предыдущая координата

				if i > 0 {
					pastCoordinate = pathLine[i-1]
				} else {
					pastCoordinate = pathCell
				}

				if artillery {
					// если стреляет арта то она игнорирует все препятвия
					lastCoordinate := pathLine[len(pathLine)-1]
					targets[strconv.Itoa(lastCoordinate.Q)+":"+strconv.Itoa(lastCoordinate.R)] = lastCoordinate
					break
				}

				// оружие без параметра артилерии не может стрелять сквозь других юнитов
				if !(pathCell.Q == targetUnit.GetQ() && pathCell.R == targetUnit.GetR()) {

					standingUnit, findUnit := client.GetUnit(pathCell.Q, pathCell.R)
					if !findUnit {
						standingUnit, findUnit = client.GetHostileUnit(pathCell.Q, pathCell.R)
					}

					// игнорируем "себя" как перпятвие, и игнорируем юнитов которых не видит игрок
					if findUnit && standingUnit.ID != targetUnit.ID {
						//добавляем самого юнита и выходим
						targets[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
						break
					}
				}

				if !pathCell.Attack || checkLevelViewCoordinate(pathCell, pastCoordinate) ||
					checkLevelViewCoordinate(pathCell, targetUnitCoordinate) {
					// 1) смотрим что черезхх координату можно смотреть
					// 2) сравниваем высоту новой координаты с предыдущей, если высота больше или рано 2м, то через нее нельзя смотреть
					// 3) сравниваем высоты новой и стартовой координаты, опять же если новая выше чем на 2 то она непроглядная
					if checkLevelViewCoordinate(pathCell, pastCoordinate) || checkLevelViewCoordinate(pathCell, targetUnitCoordinate) {
						// если координата не проглядная из за высот то мы не можем видеть координату которая выше
						break
					} else {
						// иначе можем видеть но дальше нет
						targets[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
						break
					}
				} else {
					if len(pathLine) == i+1 {
						targets[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
					}
				}
			}
		}
	}
	return
}

func checkLevelViewCoordinate(one, past *coordinate.Coordinate) bool {
	if one.Level > past.Level {
		diffLevel := one.Level - past.Level
		if diffLevel < 2 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
