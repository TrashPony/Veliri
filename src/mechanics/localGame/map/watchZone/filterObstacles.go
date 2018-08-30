package watchZone

import (
	"strconv"
	"../../../localGame"
	"../../../gameObjects/coordinate"
)

func filter(gameObject Watcher, coordinates []*coordinate.Coordinate, game *localGame.Game) (watch map[string]*coordinate.Coordinate) {

	watch = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetQ(), gameObject.GetR())

	watch[strconv.Itoa(watcherCoordinate.X)+":"+strconv.Itoa(watcherCoordinate.Z)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {

		watchCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.X, gameCoordinate.GetZ())
		//todo костыль связаный с тем что getRadius берет координаты не из игры
		if find {
			//passedCoordinate := bresenhamLineFilter.Draw(gameObject.GetQ(), gameObject.GetY(), watchCoordinate, game, "View")
			//if passedCoordinate != nil {
				watch[strconv.Itoa(watchCoordinate.X)+":"+strconv.Itoa(watchCoordinate.Y)] = watchCoordinate
			//}
		}
	}

	return
}
