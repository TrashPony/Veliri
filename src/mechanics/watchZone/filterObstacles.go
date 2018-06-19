package watchZone

import (
	"strconv"
	"../game"
	"../coordinate"
	"../bresenhamLineFilter"
)

func filter(gameObject Watcher, coordinates []*coordinate.Coordinate, game *game.Game) (watch map[string]*coordinate.Coordinate) {

	watch = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetX(), gameObject.GetY())
	watch[strconv.Itoa(watcherCoordinate.X)+":"+strconv.Itoa(watcherCoordinate.Y)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {

		watchCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.X, gameCoordinate.Y)
		//todo костыль связаный с тем что getRadius берет координаты не из игры
		if find {
			passedCoordinate := bresenhamLineFilter.Draw(gameObject.GetX(), gameObject.GetY(), watchCoordinate, game, "View")
			if passedCoordinate != nil {
				watch[strconv.Itoa(passedCoordinate.X)+":"+strconv.Itoa(passedCoordinate.Y)] = passedCoordinate
			}
		}
	}

	return
}
