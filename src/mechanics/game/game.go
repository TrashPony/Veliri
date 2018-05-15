package game

import (
	"../gameMap"
	"../player"
	"../matherShip"
	"../unit"
)

type Game struct {
	Id          int
	Name        string
	MapID       int
	Step        int
	Phase       string
	Winner      string
	Map         *gameMap.Map
	players     []*player.Player
	unitStorage []*unit.Unit
	units       map[int]map[int]*unit.Unit
	MatherShips map[int]map[int]*matherShip.MatherShip
}

func (game *Game) SetMatherShips(matherShips map[int]map[int]*matherShip.MatherShip) {
	game.MatherShips = matherShips
}

func (game *Game) SetPlayers(players []*player.Player) {
	game.players = players
}

func (game *Game) SetMap(gameMap *gameMap.Map) {
	game.Map = gameMap
}

func (game *Game) SetUnits(unit map[int]map[int]*unit.Unit) {
	game.units = unit
}

func (game *Game) SetUnitsStorage(unit []*unit.Unit) {
	game.unitStorage = unit
}

func (game *Game) SetUnit(gameUnit *unit.Unit) {
	if game.units[gameUnit.X] != nil {
		game.units[gameUnit.X][gameUnit.Y] = gameUnit
	} else {
		game.units[gameUnit.X] = make(map[int]*unit.Unit)
		game.units[gameUnit.X][gameUnit.Y] = gameUnit
	}
}

func (game *Game) DelUnit(unit *unit.Unit) {
	delete(game.units[unit.X], unit.Y)
}

func (game *Game) GetMap() (mp *gameMap.Map) {
	return game.Map
}

func (game *Game) GetUnits() (units map[int]map[int]*unit.Unit) {
	return game.units
}

func (game *Game) GetUnitsStorage() (units []*unit.Unit) {
	return game.unitStorage
}

func (game *Game) GetUnit(x, y int) (unit *unit.Unit, find bool) {
	unit, find = game.units[x][y]
	return
}

func (game *Game) GetPlayers() (Players []*player.Player) {
	return game.players
}

func (game *Game) GetMatherShips() (matherShips map[int]map[int]*matherShip.MatherShip) {
	return game.MatherShips
}

func (game *Game) GetMatherShip(x, y int) (matherShip *matherShip.MatherShip, find bool) {
	matherShip, find = game.MatherShips[x][y]
	return
}

func (game *Game) GetStep() int {
	return game.Step
}

func (game *Game) GetPhase() string {
	return game.Phase
}
