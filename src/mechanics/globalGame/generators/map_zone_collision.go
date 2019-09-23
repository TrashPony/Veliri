package generators

import (
	"github.com/TrashPony/Veliri/src/mechanics/factories/gameTypes"
	"github.com/TrashPony/Veliri/src/mechanics/factories/maps"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/coordinate"
	"github.com/TrashPony/Veliri/src/mechanics/gameObjects/map"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/collisions"
	"github.com/TrashPony/Veliri/src/mechanics/globalGame/game_math"
)

// создаем сектора с данными о проходимости для быстродействия методов движения
func UpdateMapZoneCollision() {
	for _, mp := range maps.Maps.GetAllMap() {
		FillMapZone(mp)
		go BodyCollisionCash(mp)
	}
}

func FillMapZone(mp *_map.Map) {
	mp.GeoZones = make([][]*_map.Zone, 100)

	for x := 0; x < mp.QSize*game_math.HexagonWidth; x += game_math.DiscreteSize {

		mp.GeoZones[x/game_math.DiscreteSize] = make([]*_map.Zone, 100)

		for y := 0; y < mp.RSize*game_math.HexagonHeight; y += game_math.DiscreteSize {

			mp.GeoZones[x/game_math.DiscreteSize][y/game_math.DiscreteSize] = &_map.Zone{
				Size:      game_math.DiscreteSize,
				DiscreteX: x / game_math.DiscreteSize,
				DiscreteY: y / game_math.DiscreteSize,
			}

			collisions.FillMapZone(x, y, mp.GeoZones[x/game_math.DiscreteSize][y/game_math.DiscreteSize], mp)
		}
	}

	// когда у нас есть все зоны с заполненыеми регионами, то необзодимо найти все переходы из зоны в другие зоны
	// 1 из зоны можно перейти только в соседние зоны
	// 2 однако внутри зоны добавляет новая координата регион и из одной зоны можно войти в неоограниченое количество регионов
	// 3 тоесть перезод долежн выглядеть как FROM [zoneID][RegionID] TO [zoneID][RegionID]
	// 4 задача найти для каждого региона выходы в прилежащие регионы

	// TODO надо будет ввести понятие вместимость региона если корпус по ширине
	//  если машинка 30рх то ему нужна всего одна клетка для прохода если 50 то 2 (но тогда юнит стоит не в клетке :\ )
	for _, x := range mp.GeoZones {
		for _, zone := range x {
			if zone != nil {

				for _, region := range zone.Regions {
					if region != nil && region.Index != 0 {
						SearchEntries(region, mp, zone)
					}
				}
			}
		}
	}
}

func SearchEntries(region *_map.Region, mp *_map.Map, parentZone *_map.Zone) {

	zones := parentZone.GetNeighboursZone(mp)
	region.Links = make([]*_map.Link, 0)
	region.GlobalLinks = make(map[string]*_map.Link)

	checkCoordinate := func(xCell, yCell int, cellEntry *coordinate.Coordinate) (*_map.Zone, *_map.Region, *coordinate.Coordinate, *coordinate.Coordinate) {
		// координата НЕ должна принадлежать родительскому региону и региону 0
		// она должна быть проходимой
		// должна находится в другой зону

		for _, zone := range zones {

			if zone != nil && (zone.DiscreteX != parentZone.DiscreteX || zone.DiscreteY != parentZone.DiscreteY) {

				for _, region := range zone.Regions {

					if region != nil && region.Index != 0 {

						for _, x := range region.Cells {
							for _, cell := range x {
								if xCell == cell.X && yCell == cell.Y {
									return zone, region, cellEntry, cell
								}
							}
						}
					}
				}
			}
		}

		return nil, nil, nil, nil
	}

	createLink := func(toZone *_map.Zone, toRegion *_map.Region, cellEntry, cellOut *coordinate.Coordinate) {
		if toZone != nil && toRegion != nil {

			region.Links = append(region.Links, &_map.Link{
				Zone:   toZone,
				Region: toRegion,
				FromX:  cellEntry.X,
				FromY:  cellEntry.Y,
				ToX:    cellOut.X,
				ToY:    cellOut.Y,
			})

			globalLink := _map.Link{
				Zone:   toZone,
				Region: toRegion,
			}

			region.GlobalLinks[globalLink.GetGlobalKey()] = &globalLink
		}
	}

	for _, x := range region.Cells {
		for _, cell := range x {
			//строго лево
			createLink(checkCoordinate(cell.X-game_math.CellSize, cell.Y, cell))

			//строго право
			createLink(checkCoordinate(cell.X+game_math.CellSize, cell.Y, cell))

			//верх центр
			createLink(checkCoordinate(cell.X, cell.Y-game_math.CellSize, cell))

			//низ центр
			createLink(checkCoordinate(cell.X, cell.Y+game_math.CellSize, cell))

			//верх лево
			createLink(checkCoordinate(cell.X-game_math.CellSize, cell.Y-game_math.CellSize, cell))

			//верх право
			createLink(checkCoordinate(cell.X+game_math.CellSize, cell.Y-game_math.CellSize, cell))

			//низ лево
			createLink(checkCoordinate(cell.X-game_math.CellSize, cell.Y+game_math.CellSize, cell))

			//низ право
			createLink(checkCoordinate(cell.X+game_math.CellSize, cell.Y+game_math.CellSize, cell))
		}
	}
}

func BodyCollisionCash(mp *_map.Map) {
	bodies := gameTypes.Bodies.GetAllType()

	for _, body := range bodies {
		for _, x := range mp.GeoZones {
			for _, zone := range x {
				if zone != nil {
					for _, cell := range zone.Cells {
						collisions.CheckCollisionsOnStaticMap(cell.X, cell.Y, 0, mp, &body, false, true)
					}
				}
			}
		}
	}
}
