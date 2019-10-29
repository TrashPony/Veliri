package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func checkValidForMoveCoordinate(gameMap *_map.Map, x, y, xSize, ySize int, gameUnit *unit.Unit,
	scaleMap int, regions []*_map.Region, units map[int]*unit.ShortUnitInfo, unitsID []int) (*coordinate.Coordinate, bool, bool) {

	// за пределами карты
	if x > xSize || y > ySize || x < 0 || y < 0 {
		return nil, false, false
	}

	if units != nil {
		free, collisionUnit := collisions.CheckCollisionsPlayers(gameUnit, x*scaleMap+scaleMap/2, y*scaleMap+scaleMap/2,
			0, units, false, true, true, false, true, unitsID)

		if !free {
			dist := game_math.GetBetweenDist(x*scaleMap+scaleMap/2, y*scaleMap+scaleMap/2, collisionUnit.X, collisionUnit.Y)
			if !collisionUnit.MoveChecker || dist < 350 {
				return nil, false, true
			}
		}
	}

	find := true

	if regions != nil {
		find = false
		for _, region := range regions {
			_, find = region.Cells[x][y]
			if find {
				break
			}
		}
	}

	if !find {
		return nil, false, false
	} else {
		c := &coordinate.Coordinate{X: x, Y: y} //, Rotate: game_math.GetBetweenAngle(float64(x), float64(y), float64(pX), float64(pY))}
		//+scaleMap/2 потомучто юнит находится в центре клетки
		possible, _ := collisions.BodyCheckCollisionsOnStaticMap(x*scaleMap+scaleMap/2, y*scaleMap+scaleMap/2, 0, gameMap, gameUnit.Body, false, true)

		if possible {
			return c, true, false
		}

		return nil, false, false
	}
}
