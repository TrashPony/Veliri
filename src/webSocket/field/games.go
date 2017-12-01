package field

import (
	"../../game"
)

type ActiveGame struct {
	mapInfo    *game.Map
	stat       *game.Game
	players    []*game.UserStat
	units      map[int]map[int]*game.Unit
	structure  map[int]map[int]*game.Structure
	coordinate map[int]map[int]*game.Coordinate
}

func (activeGame *ActiveGame) setStructure(structure map[int]map[int]*game.Structure) {
	activeGame.structure = structure
}

func (activeGame *ActiveGame) setPlayers(players []*game.UserStat) {
	activeGame.players = players
}

func (activeGame *ActiveGame) setInfoMap(mapInfo *game.Map) {
	activeGame.mapInfo = mapInfo
}

func (activeGame *ActiveGame) setStat(stat *game.Game) {
	activeGame.stat = stat
}

func (activeGame *ActiveGame) setUnits(unit map[int]map[int]*game.Unit) {
	activeGame.units = unit
}

func (activeGame *ActiveGame) setUnit(unit *game.Unit) {
	if activeGame.units[unit.X] != nil {
		activeGame.units[unit.X][unit.Y] = unit
	} else {
		activeGame.units[unit.X] = make(map[int]*game.Unit)
		activeGame.units[unit.X][unit.Y] = unit
	}
}

func (activeGame *ActiveGame) setMap(coordinate map[int]map[int]*game.Coordinate)  {
	activeGame.coordinate = coordinate
}

func (activeGame *ActiveGame) delUnit(unit *game.Unit) {
	delete(activeGame.units[unit.X], unit.Y)
}

func (activeGame *ActiveGame) getMap() (mp *game.Map) {
	return activeGame.mapInfo
}

func (activeGame *ActiveGame) getUnits() (units map[int]map[int]*game.Unit) {
	return activeGame.units
}

func (activeGame *ActiveGame) getPlayers() (Players []*game.UserStat) {
	return activeGame.players
}

func (activeGame *ActiveGame) getStructure() (Structure map[int]map[int]*game.Structure) {
	return activeGame.structure
}
