package collisions

import (
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

func FillMapZone(x, y int, zone *_map.Zone, mp *_map.Map) {
	// +50 что бы в область папали пограничные препятсвия
	zoneRect := GetRect(float64(x), float64(y), game_math.DiscreteSize+50, game_math.DiscreteSize+50)

	zone.Obstacle = make([]*_map.ObstaclePoint, 0)

	for i := 0; i < len(mp.GeoData); i++ {
		if zoneRect.detectCollisionRectToCircle(&point{x: float64(mp.GeoData[i].X), y: float64(mp.GeoData[i].Y)}, mp.GeoData[i].Radius) {
			zone.Obstacle = append(zone.Obstacle, mp.GeoData[i])
		}
	}

	zone.Cells = make([]*coordinate.Coordinate, 0)
	for xCell := x + game_math.DiscreteSize - game_math.CellSize; xCell >= x; xCell -= game_math.CellSize {
		for yCell := y + game_math.DiscreteSize - game_math.CellSize; yCell >= y; yCell -= game_math.CellSize {
			zone.Cells = append(zone.Cells, &coordinate.Coordinate{X: xCell, Y: yCell})
		}
	}

	zone.Regions = make([]*_map.Region, 100)
	index := 0

	// все не проходимые клетки это всегда регион 0
	obstacleRegion := _map.Region{Index: index, Cells: make(map[int]map[int]*coordinate.Coordinate, 0), Zone: zone}
	zone.Regions[index] = &obstacleRegion

	for _, cell := range zone.Cells {

		cellRect := GetRect(float64(cell.X), float64(cell.Y), game_math.CellSize, game_math.CellSize)

		for i := 0; i < len(mp.GeoData); i++ {

			dist := game_math.GetBetweenDist(cell.X, cell.Y, mp.GeoData[i].X, mp.GeoData[i].Y)

			if int(dist) < mp.GeoData[i].Radius+game_math.CellSize*2 {
				if cellRect.detectCollisionRectToCircle(&point{x: float64(mp.GeoData[i].X), y: float64(mp.GeoData[i].Y)}, mp.GeoData[i].Radius) {
					cell.Find = true

					if obstacleRegion.Cells[cell.X/game_math.CellSize] == nil {
						obstacleRegion.Cells[cell.X/game_math.CellSize] = make(map[int]*coordinate.Coordinate)
					}

					obstacleRegion.Cells[cell.X/game_math.CellSize][cell.Y/game_math.CellSize] = cell
					continue
				}
			}
		}
	}

	index++

	// ищем координу которую еще не искали, берем всех ее соседей то тех пока мы не упремя в стену или в предел зоны
	for _, cell := range zone.Cells {
		if !cell.Find {
			zone.Regions[index] = CreateRegion(zone, cell, index)
			index++
		}
	}
}

func CreateRegion(zone *_map.Zone, start *coordinate.Coordinate, index int) *_map.Region {

	region := _map.Region{Index: index, Cells: make(map[int]map[int]*coordinate.Coordinate, 0), Zone: zone}

	openPoints := make(map[string]*coordinate.Coordinate, 0)
	openPoints[start.Key()] = start

	getCurrPoint := func() *coordinate.Coordinate {
		for _, point := range openPoints {
			return point
		}

		return nil
	}

	checkCoordinate := func(x, y int) *coordinate.Coordinate {
		for _, zonePoint := range zone.Cells {
			if zonePoint.X == x && zonePoint.Y == y && !zonePoint.Find {
				return zonePoint
			}
		}

		return nil
	}

	appendOpen := func(newCoordinate *coordinate.Coordinate) {
		if newCoordinate != nil {
			openPoints[newCoordinate.Key()] = newCoordinate
		}
	}

	for len(openPoints) > 0 {

		curr := getCurrPoint()
		curr.Find = true

		if region.Cells[curr.X/game_math.CellSize] == nil {
			region.Cells[curr.X/game_math.CellSize] = make(map[int]*coordinate.Coordinate)
		}

		region.Cells[curr.X/game_math.CellSize][curr.Y/game_math.CellSize] = curr

		delete(openPoints, curr.Key())

		//строго лево
		appendOpen(checkCoordinate(curr.X-game_math.CellSize, curr.Y))

		//строго право
		appendOpen(checkCoordinate(curr.X+game_math.CellSize, curr.Y))

		//верх центр
		appendOpen(checkCoordinate(curr.X, curr.Y-game_math.CellSize))

		//низ центр
		appendOpen(checkCoordinate(curr.X, curr.Y+game_math.CellSize))

		//верх лево
		appendOpen(checkCoordinate(curr.X-game_math.CellSize, curr.Y-game_math.CellSize))

		//верх право
		appendOpen(checkCoordinate(curr.X+game_math.CellSize, curr.Y-game_math.CellSize))

		//низ лево
		appendOpen(checkCoordinate(curr.X-game_math.CellSize, curr.Y+game_math.CellSize))

		//низ право
		appendOpen(checkCoordinate(curr.X+game_math.CellSize, curr.Y+game_math.CellSize))
	}

	return &region
}
