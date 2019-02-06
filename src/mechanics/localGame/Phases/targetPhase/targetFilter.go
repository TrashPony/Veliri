package targetPhase

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/localGame"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/hexLineDraw"
	"github.com/TrashPony/Veliri/src/mechanics/localGame/map/watchZone"
	"strconv"
)

func filter(gameObject watchZone.Watcher, coordinates []*coordinate.Coordinate, game *localGame.Game, artillery bool) (targets map[string]*coordinate.Coordinate) {

	targets = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetQ(), gameObject.GetR())

	targets[strconv.Itoa(watcherCoordinate.Q)+":"+strconv.Itoa(watcherCoordinate.R)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {
		watchCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.Q, gameCoordinate.R)
		if find {
			pathLine := hexLineDraw.Draw(watcherCoordinate, watchCoordinate, game)

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

				// оружие без параметры артилерии не может стрелять сквозь других юнитов
				if !(pathCell.Q == gameObject.GetQ() && pathCell.R == gameObject.GetR()) {
					_, findUnit := game.GetUnit(pathCell.Q, pathCell.R)
					if findUnit {
						//добавляем самого юнита и выходим
						targets[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
						break
					}
				}

				if !pathCell.Attack || checkLevelViewCoordinate(pathCell, pastCoordinate) ||
					checkLevelViewCoordinate(pathCell, watcherCoordinate) {
					// 1) смотрим что черезхх координату можно смотреть
					// 2) сравниваем высоту новой координаты с предыдущей, если высота больше или рано 2м, то через нее нельзя смотреть
					// 3) сравниваем высоты новой и стартовой координаты, опять же если новая выше чем на 2 то она непроглядная
					if checkLevelViewCoordinate(pathCell, pastCoordinate) || checkLevelViewCoordinate(pathCell, watcherCoordinate) {
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
