package field

import (
	"../../game/objects"
)

type ActiveGame struct {
	Map *objects.Map
	Stat *objects.Game
	Players []*objects.UserStat
	Units map[int]map[int]*objects.Unit
	Coordinate map[int]map[int]*objects.Coordinate
}

func (game *ActiveGame) addPlayers(players []*objects.UserStat)  {
		game.Players = players
}

func (game *ActiveGame) addMap(mp *objects.Map)  {
	game.Map = mp
}

func (game *ActiveGame) addStat(stat *objects.Game)  {
	game.Stat = stat
}

func (game *ActiveGame) addUnits(unit map[int]map[int]*objects.Unit) {
	game.Units = unit
}

func (game *ActiveGame) addUnit(unit *objects.Unit) {
	if game.Units[unit.X] != nil {
		game.Units[unit.X][unit.Y] = unit
	} else {
		game.Units[unit.X] = make(map[int]*objects.Unit)
		game.Units[unit.X][unit.Y] = unit
	}
}

func (game *ActiveGame) delUnit(unit *objects.Unit) {
	delete(game.Units[unit.X], unit.Y)
}

func (game *ActiveGame) getMap()(mp *objects.Map)  {
	return game.Map
}

func (game *ActiveGame) getUnits() (units map[int]map[int]*objects.Unit) {
	return game.Units
}

func (game *ActiveGame) getPlayers() (Players []*objects.UserStat) {
	return game.Players
}