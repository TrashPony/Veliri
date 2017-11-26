package field

import (
	"../../game/objects"
)

type ActiveGame struct {
	mapInfo    *objects.Map
	stat       *objects.Game
	players    []*objects.UserStat
	units      map[int]map[int]*objects.Unit
	structure  map[int]map[int]*objects.Structure
	coordinate map[int]map[int]*objects.Coordinate
}

func (game *ActiveGame) setStructure(structure map[int]map[int]*objects.Structure) {
	game.structure = structure
}

func (game *ActiveGame) setPlayers(players []*objects.UserStat) {
	game.players = players
}

func (game *ActiveGame) setInfoMap(mapInfo *objects.Map) {
	game.mapInfo = mapInfo
}

func (game *ActiveGame) setStat(stat *objects.Game) {
	game.stat = stat
}

func (game *ActiveGame) setUnits(unit map[int]map[int]*objects.Unit) {
	game.units = unit
}

func (game *ActiveGame) setUnit(unit *objects.Unit) {
	if game.units[unit.X] != nil {
		game.units[unit.X][unit.Y] = unit
	} else {
		game.units[unit.X] = make(map[int]*objects.Unit)
		game.units[unit.X][unit.Y] = unit
	}
}

func (game *ActiveGame) setMap(coordinate map[int]map[int]*objects.Coordinate)  {
	game.coordinate = coordinate
}

func (game *ActiveGame) delUnit(unit *objects.Unit) {
	delete(game.units[unit.X], unit.Y)
}

func (game *ActiveGame) getMap() (mp *objects.Map) {
	return game.mapInfo
}

func (game *ActiveGame) getUnits() (units map[int]map[int]*objects.Unit) {
	return game.units
}

func (game *ActiveGame) getPlayers() (Players []*objects.UserStat) {
	return game.players
}

func (game *ActiveGame) getStructure() (Structure map[int]map[int]*objects.Structure) {
	return game.structure
}
