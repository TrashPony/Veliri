package watchZone

import (
	"../../../gameObjects/coordinate"
	"../../../localGame"
	"../hexLineDraw"
	"strconv"
)

func filter(gameObject Watcher, coordinates []*coordinate.Coordinate, game *localGame.Game) (watch map[string]*coordinate.Coordinate) {
	// todo вохможно этот код можно легко обьеденить с localGame/Phases/targetPhase/targetFilter.go
	watch = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetQ(), gameObject.GetR())

	watch[strconv.Itoa(watcherCoordinate.Q)+":"+strconv.Itoa(watcherCoordinate.R)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {
		watchCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.Q, gameCoordinate.R)
		if find {
			pathLine := hexLineDraw.Draw(watcherCoordinate, watchCoordinate, game)
			// линия пути 0 элемент - начало
			for i, pathCell := range pathLine {
				pastCoordinate := &coordinate.Coordinate{} // предыдущая координата

				if i > 0 {
					pastCoordinate = pathLine[i-1]
				} else {
					pastCoordinate = pathCell
				}

				if !pathCell.View || checkLevelViewCoordinate(pathCell, pastCoordinate) ||
					checkLevelViewCoordinate(pathCell, watcherCoordinate) {
					// 1) смотрим что черезхх координату можно смотреть
					// 2) сравниваем высоту новой координаты с предыдущей, если высота больше или рано 2м, то через нее нельзя смотреть
					// 3) сравниваем высоты новой и стартовой координаты, опять же если новая выше чем на 2 то она непроглядная
					if checkLevelViewCoordinate(pathCell, pastCoordinate) || checkLevelViewCoordinate(pathCell, watcherCoordinate) {
						// если координата не проглядная из за высот то мы не можем видеть координату которая выше
						break
					} else {
						// иначе можем видеть но дальше нет
						watch[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
						break
					}
				} else {
					if len(pathLine) == i+1 {
						watch[strconv.Itoa(pathCell.Q)+":"+strconv.Itoa(pathCell.R)] = pathCell
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
