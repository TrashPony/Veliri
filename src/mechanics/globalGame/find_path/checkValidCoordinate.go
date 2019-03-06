package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/player"
)

func checkValidForMoveCoordinate(client *player.Player, gameMap *_map.Map, x int, y int, gameUnit *unit.Unit, scaleMap int) (*coordinate.Coordinate, bool) {

	fakeBody, _ := gameTypes.Bodies.GetByID(gameUnit.Body.ID)
	fakeBody.SideRadius = fakeBody.FrontRadius

	possible, _, _, _ := globalGame.CheckCollisionsOnStaticMap(x*scaleMap, y*scaleMap, gameUnit.Rotate, gameMap, fakeBody)

	if possible {
		return &coordinate.Coordinate{X: x, Y: y}, true
	}
	return nil, false
}
