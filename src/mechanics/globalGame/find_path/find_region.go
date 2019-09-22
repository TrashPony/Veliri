package find_path

import (
	"errors"
	"fmt"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/unit"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/debug"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
	"strconv"
	"time"
)

// функция находит путь из 1го региона в другой
// работает почти как А* но взятие соседий происходит по загатовленым связям между зонами по регионам
// на выходе мы должны получить точки регионов пути

type RegionParent struct {
	Region *_map.Region
	Parent *RegionParent
	Wave   int
}

func FindRegionPath(mp *_map.Map, start, end *coordinate.Coordinate, gameUnit *unit.Unit, uuid string) (error, []*_map.Region) {

	startTime := time.Now()
	defer func() {
		if debug.Store.Move {
			elapsed := time.Since(startTime)
			fmt.Println("time region path: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 6, 64))
		}
	}()

	// TODO проверить доступность этих регионов
	startZone := mp.GeoZones[start.X/game_math.DiscreteSize][start.Y/game_math.DiscreteSize]
	endZone := mp.GeoZones[end.X/game_math.DiscreteSize][end.Y/game_math.DiscreteSize]
	if startZone == nil || endZone == nil {
		return errors.New("no zone"), nil
	}

	if debug.Store.RegionFindDebug {
		debug.Store.AddMessage("CreateRect", "blue", startZone.DiscreteX*game_math.DiscreteSize,
			startZone.DiscreteY*game_math.DiscreteSize, 0, 0, game_math.DiscreteSize, mp.Id, 20)
		debug.Store.AddMessage("CreateRect", "blue", endZone.DiscreteX*game_math.DiscreteSize,
			endZone.DiscreteY*game_math.DiscreteSize, 0, 0, game_math.DiscreteSize, mp.Id, 20)
	}

	// TODO иногда не удается найти регион из за того что туша попала в регион 0
	startRegion := startZone.GetRegionsByXY(start.X, start.Y)
	endRegion := endZone.GetRegionsByXY(end.X, end.Y)

	if endRegion == nil || startRegion == nil {
		return errors.New("no region"), nil
	}

	path := make([]*_map.Region, 0)
	path = append(path, startRegion)

	if startRegion.GetKey() == endRegion.GetKey() {
		// в поиске пути участвует только 1 регион, не зона а именно регион
		return nil, path
	}

	if debug.Store.RegionFindDebug {
		drawRegion(startRegion, mp.Id, "green")
		drawRegion(endRegion, mp.Id, "green")
	}

	// key := ZoneID, RegionID
	openRegions := make(map[string]*RegionParent)
	openRegions[startRegion.GetKey()] = &RegionParent{Region: startRegion}

	closeRegion := make(map[string]*RegionParent)

	wave := 0
	for uuid == gameUnit.MoveUUID {

		if len(openRegions) <= 0 {
			return errors.New("region path no find"), nil
		}

		wave++
		curr := GetFistRegion(&openRegions, &closeRegion, wave)
		if curr.Region.GetKey() == endRegion.GetKey() {

			for !(curr.Region.GetKey() == startRegion.GetKey()) { // идем обратно до тех пока пока не дойдем до стартовой точки
				curr = curr.Parent

				if !(curr.Region.GetKey() == startRegion.GetKey()) {
					path = append(path, curr.Region)
					if debug.Store.RegionFindDebug {
						drawRegion(curr.Region, mp.Id, "orange")
					}
				}
			}
			break
		}

		GetNeighboursRegion(curr, &openRegions, &closeRegion, mp, wave)
	}

	path = append(path, endRegion)
	return nil, path
}

func drawRegion(region *_map.Region, mpID int, color string) {
	for _, x := range region.Cells {
		for _, cell := range x {
			debug.Store.AddMessage("CreateRect", color, cell.X, cell.Y, 0, 0, game_math.CellSize, mpID, 0)
		}
	}
}

func GetNeighboursRegion(currRegion *RegionParent, openRegions, closeRegion *map[string]*RegionParent, mp *_map.Map, wave int) {
	for _, link := range currRegion.Region.GlobalLinks {
		// если мы еще не ходили сюда то добавляем
		if (*openRegions)[link.Region.GetKey()] == nil && (*closeRegion)[link.Region.GetKey()] == nil {

			(*openRegions)[link.Region.GetKey()] = &RegionParent{link.Region, currRegion, wave}

			if debug.Store.RegionFindDebug {
				drawRegion(link.Region, mp.Id, "green")
			}
		}
	}
}

func GetFistRegion(openRegions, closeRegion *map[string]*RegionParent, maxWave int) *RegionParent {

	// выбираем зону с минимальной волной

	var minRegion *RegionParent

	for _, region := range *openRegions {
		if region.Wave < maxWave {
			minRegion = region
			maxWave = region.Wave
		}
	}

	(*closeRegion)[minRegion.Region.GetKey()] = minRegion
	delete(*openRegions, minRegion.Region.GetKey())

	return minRegion
}
