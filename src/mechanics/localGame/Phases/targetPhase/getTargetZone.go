package targetPhase

import (
	"../../../gameObjects/coordinate"
	"../../../gameObjects/unit"
	"../../Phases"
	"../../map/watchZone"
	"../../../localGame"
	"../../map/bresenhamLineFilter"
	"strconv"
)

func GetTargetCoordinate(gameUnit *unit.Unit, activeGame *localGame.Game) map[string]map[string]*coordinate.Coordinate {

	openCoordinate := make(map[string]map[string]*coordinate.Coordinate)

	RadiusCoordinates := coordinate.GetCoordinatesRadius(gameUnit.GetX(), gameUnit.GetY(), gameUnit.Weapon.Range)
	zone := filter(gameUnit, RadiusCoordinates, activeGame)

	for _, gameCoordinate := range zone {
		if !(gameCoordinate.X == gameUnit.X && gameCoordinate.Y == gameUnit.Y) {
			Phases.AddCoordinate(openCoordinate, gameCoordinate)
		}
	}

	return openCoordinate
}

func filter(gameObject watchZone.Watcher, coordinates []*coordinate.Coordinate, game *localGame.Game) (watch map[string]*coordinate.Coordinate) {

	watch = make(map[string]*coordinate.Coordinate)

	watcherCoordinate, _ := game.GetMap().GetCoordinate(gameObject.GetX(), gameObject.GetY())
	watch[strconv.Itoa(watcherCoordinate.X)+":"+strconv.Itoa(watcherCoordinate.Y)] = watcherCoordinate

	for _, gameCoordinate := range coordinates {

		targetCoordinate, find := game.GetMap().GetCoordinate(gameCoordinate.X, gameCoordinate.Y)
		//todo костыль связаный с тем что getRadius берет координаты не из игры
		if find {
			passedCoordinate := bresenhamLineFilter.Draw(gameObject.GetX(), gameObject.GetY(), targetCoordinate, game, "Target")
			if passedCoordinate != nil {
				watch[strconv.Itoa(passedCoordinate.X)+":"+strconv.Itoa(passedCoordinate.Y)] = passedCoordinate
			}
		}
	}

	return
}
