package targetPhase

import (
	"../../map/watchZone"
	"../../../localGame"
	"../../map/bresenhamLineFilter"
	"strconv"
	"../../../gameObjects/coordinate"
)

func filter(gameObject watchZone.Watcher, coordinates []*coordinate.Coordinate, game *localGame.Game) (watch map[string]*coordinate.Coordinate) {

	watch = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetQ(), gameObject.GetR())
	watch[strconv.Itoa(watcherCoordinate.X)+":"+strconv.Itoa(watcherCoordinate.Y)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {

		targetCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.X, gameCoordinate.Y)
		//todo костыль связаный с тем что getRadius берет координаты не из игры
		if find {
			passedCoordinate := bresenhamLineFilter.Draw(gameObject.GetQ(), gameObject.GetR(), targetCoordinate, game, "Target")
			if passedCoordinate != nil {
				watch[strconv.Itoa(passedCoordinate.X)+":"+strconv.Itoa(passedCoordinate.Y)] = passedCoordinate
			}
		}
	}

	return
}
