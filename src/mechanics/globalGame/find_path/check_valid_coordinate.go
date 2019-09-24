package find_path

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
)

func checkValidForMoveCoordinate(gameMap *_map.Map, x, y, xSize, ySize int, gameUnit *unit.Unit,
	scaleMap int, regions []*_map.Region, units map[int]*unit.ShortUnitInfo) (*coordinate.Coordinate, bool) {

	// за пределами карты
	if x > xSize || y > ySize || x < 0 || y < 0 {
		return nil, false
	}

	if units != nil {
		free, _ := collisions.CheckCollisionsPlayers(gameUnit, x*scaleMap+scaleMap/2, y*scaleMap+scaleMap/2,
			0, units, false, true, true)
		if !free {
			return nil, false
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
		return nil, false
	} else {
		c := &coordinate.Coordinate{X: x, Y: y} //, Rotate: game_math.GetBetweenAngle(float64(x), float64(y), float64(pX), float64(pY))}
		//+scaleMap/2 потомучто юнит находится в центре клетки
		possible, _ := collisions.CheckCollisionsOnStaticMap(x*scaleMap+scaleMap/2, y*scaleMap+scaleMap/2, c.Rotate, gameMap, gameUnit.Body, false, true)

		if possible {
			return c, true
		}

		return nil, false
	}
}
