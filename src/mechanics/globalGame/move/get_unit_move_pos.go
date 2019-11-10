package move

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/player"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/find_path"
)

func GetUnitPos(unitsID []int, user *player.Player, toX, toY float64) []*coordinate.Coordinate {

	toPos := make([]*coordinate.Coordinate, 0)
	mp, _ := maps.Maps.GetByID(user.GetSquad().MatherShip.MapID)
	units := globalGame.Clients.GetAllShortUnits(mp.Id)

	for _, id := range unitsID {
		moveUnit := user.GetSquad().GetUnitByID(id)
		if moveUnit != nil {

			x, y, _ := find_path.SearchEndPoint(toX, toY, toX, toY, moveUnit, mp, units)
			toPos = append(toPos, &coordinate.Coordinate{X: int(x), Y: int(y)})

			// переназначаем юнита что бы учитывать его будущие положение при поиске позиции
			shortUnitInfo := moveUnit.GetShortInfo()
			shortUnitInfo.X = int(x)
			shortUnitInfo.Y = int(y)

			units[moveUnit.ID] = shortUnitInfo
		}
	}

	return toPos
}
